package v1

import (
	"github.com/labstack/echo/v4"
	"zapvote/internal/api/response"
	"zapvote/internal/model/vote"
)

func (vc *VoteController) FacultyVotePOST(e echo.Context) error {
	body := &requestBody{}
	if err := e.Bind(body); err != nil {
		return response.BadRequestError(e, "")
	}

	var final vote.Category
	switch body.Faculty {
	case 2:
		final = vote.Banking
		break
	case 3:
		final = vote.Finance
		break
	case 4:
		final = vote.Computer
		break
	default:
		return response.BadRequestError(e, "invalid input")
	}

	return vc.process(e, VoteType{
		getVotes:    vc.electionStore.GetFacultyVotesTx,
		updateVotes: vc.electionStore.UpdateFacultyVote,
		voteType:    final,
		body:        body,
	})
}
