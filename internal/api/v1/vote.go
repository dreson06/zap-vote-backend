package v1

import (
	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo/v4"
	"zapvote/internal/api/middleware/auth"
	"zapvote/internal/api/response"
	"zapvote/internal/model/vote"
	"zapvote/internal/services/electionstore"
	"zapvote/internal/services/userstore"
	"zapvote/internal/services/votestore"
)

type VoteController struct {
	voteStore     votestore.Store
	userStore     userstore.Store
	electionStore electionstore.Store
	db            *sqlx.DB
}

type requestBody struct {
	ID         string `json:"id"`
	ElectionID string `json:"election_id"`
	DeviceID   string `json:"device_id"`
}

func NewVoteController(voteStore votestore.Store, userStore userstore.Store, electionStore electionstore.Store, db *sqlx.DB) *VoteController {
	return &VoteController{
		voteStore:     voteStore,
		userStore:     userStore,
		electionStore: electionStore,
		db:            db,
	}
}

type VoteType struct {
	getVotes    func(*sqlx.Tx, string) (int64, error)
	updateVotes func(*sqlx.Tx, string, int64) error
}

func (vc *VoteController) HasVotedGET(e echo.Context) error {
	userID := auth.GetID(e)
	electionID := e.Param("election-id")
	hasVoted, err := vc.voteStore.HasVoted(userID, electionID)
	if err != nil {
		return response.ServerError(e, err, "")
	}
	res := map[string]bool{"has_voted": hasVoted}
	return response.JSON(e, res)
}

func (vc *VoteController) process(e echo.Context, voteType VoteType) error {
	userID := auth.GetID(e)
	body := &requestBody{}
	if err := e.Bind(body); err != nil {
		return response.BadRequestError(e, "")
	}
	if body.ID == "" || body.DeviceID == "" {
		return response.BadRequestError(e, "information missing")
	}

	u, err := vc.userStore.GetUserSimple(userID)
	if err != nil {
		return response.ServerError(e, err, "")
	}

	//check if the device id is the same
	if u.DeviceID != body.DeviceID {
		return response.OtherErrors(e, response.StatusUnauthorizedVote, "vote from device you registered")
	}

	hasVoted, err := vc.voteStore.HasVoted(userID, body.ElectionID)
	if err != nil {
		return response.ServerError(e, err, "")
	}

	if hasVoted {
		return response.OtherErrors(e, response.StatusUnauthorizedVote, "you can vote only once")
	}

	tx, err := vc.db.Beginx()
	if err != nil {
		return response.ServerError(e, err, "")
	}

	votes, err := voteType.getVotes(tx, body.ID)
	if err != nil {
		_ = tx.Rollback()
		return response.ServerError(e, err, "")
	}
	newVote := votes + 1
	err = voteType.updateVotes(tx, body.ID, newVote)
	if err != nil {
		_ = tx.Rollback()
		return response.ServerError(e, err, "")
	}

	v := &vote.Vote{
		UserID:     userID,
		ElectionID: body.ElectionID,
	}

	err = vc.voteStore.CreateTx(tx, v)
	if err != nil {
		_ = tx.Rollback()
		return response.ServerError(e, err, "")
	}

	err = tx.Commit()
	if err != nil {
		return response.ServerError(e, err, "")
	}
	return response.Success(e)
}
