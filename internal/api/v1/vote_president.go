package v1

import (
	"github.com/labstack/echo/v4"
	"zapvote/internal/api/response"
	"zapvote/internal/model/vote"
)

func (vc *VoteController) VotePresidentialPOST(e echo.Context) error {
	body := &requestBody{}
	if err := e.Bind(body); err != nil {
		return response.BadRequestError(e, "")
	}
	return vc.process(e, VoteType{
		getVotes:    vc.electionStore.GetPresidentVotesTx,
		updateVotes: vc.electionStore.UpdatePresidentVoteTx,
		voteType:    vote.Presidential,
		body:        body,
	})
}
