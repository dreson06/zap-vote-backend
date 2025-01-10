package v1

import (
	"github.com/labstack/echo/v4"
)

func (vc *VoteController) FacultyVotePOST(e echo.Context) error {
	return vc.process(e, VoteType{
		getVotes:    vc.electionStore.GetFacultyVotesTx,
		updateVotes: vc.electionStore.UpdateFacultyVote,
	})
}
