package v1

import (
	"github.com/labstack/echo/v4"
)

func (vc *VoteController) VoteClassPOST(e echo.Context) error {
	return vc.process(e, VoteType{
		getVotes:    vc.electionStore.GetClassVotes,
		updateVotes: vc.electionStore.UpdateClassRepVote,
	})
}
