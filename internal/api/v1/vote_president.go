package v1

import (
	"github.com/labstack/echo/v4"
)

func (vc *VoteController) VotePresidentialPOST(e echo.Context) error {
	return vc.process(e, VoteType{
		getVotes:    vc.electionStore.GetPresidentVotesTx,
		updateVotes: vc.electionStore.UpdatePresidentVoteTx,
	})
}
