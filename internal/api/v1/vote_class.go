package v1

import (
	"github.com/labstack/echo/v4"
	"zapvote/internal/api/response"
	"zapvote/internal/model/vote"
)

func (vc *VoteController) VoteClassPOST(e echo.Context) error {
	body := &requestBody{}
	if err := e.Bind(body); err != nil {
		return response.BadRequestError(e, "")
	}
	return vc.process(e, VoteType{
		getVotes:    vc.electionStore.GetClassVotes,
		updateVotes: vc.electionStore.UpdateClassRepVote,
		voteType:    vote.Class,
		body:        body,
	})
}
