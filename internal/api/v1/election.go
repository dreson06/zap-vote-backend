package v1

import (
	"github.com/labstack/echo/v4"
	"zapvote/internal/api/response"
	"zapvote/internal/services/electionstore"
)

type ElectionController struct {
	electionStore electionstore.Store
}

func NewElectionController(electionStore electionstore.Store) *ElectionController {
	return &ElectionController{
		electionStore: electionStore,
	}
}

func (ec *ElectionController) GetPresidentialCandidates(e echo.Context) error {
	candidates, err := ec.electionStore.GetPresidentialCandidates()
	if err != nil {
		return response.ServerError(e, err, "")
	}
	return response.JSON(e, candidates)
}

func (ec *ElectionController) GetFacultyCandidates(e echo.Context) error {
	faculty := e.Param("faculty")
	if faculty == "" {
		return response.BadRequestError(e, "faculty is required")
	}
	candidates, err := ec.electionStore.GetFacultyCandidates(faculty)
	if err != nil {
		return response.ServerError(e, err, "")
	}
	return response.JSON(e, candidates)
}
