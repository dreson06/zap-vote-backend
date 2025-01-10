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

func (ec *ElectionController) ElectionGET(e echo.Context) error {
	id := e.Param("id")
	el, err := ec.electionStore.GetElection(id)
	if err != nil {
		return response.ServerError(e, err, "")
	}
	return response.JSON(e, el)
}

func (ec *ElectionController) PresidentialCandidatesGET(e echo.Context) error {
	candidates, err := ec.electionStore.GetAllPresidentialCandidates()
	if err != nil {
		return response.ServerError(e, err, "")
	}
	return response.JSON(e, candidates)
}

func (ec *ElectionController) FacultyCandidatesGET(e echo.Context) error {
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

func (ec *ElectionController) ClassRepCandidatesGET(e echo.Context) error {
	course := e.Param("course")
	if course == "" {
		return response.BadRequestError(e, "course is required")
	}
	candidates, err := ec.electionStore.GetClassRepCandidates(course)
	if err != nil {
		return response.ServerError(e, err, "")
	}
	return response.JSON(e, candidates)
}

func (ec *ElectionController) FacultyResultsGET(e echo.Context) error {
	candidates, err := ec.electionStore.GetFacultyResults()
	if err != nil {
		return response.ServerError(e, err, "")
	}
	return response.JSON(e, candidates)
}

func (ec *ElectionController) ClassRepResultsGET(e echo.Context) error {
	code := e.Param("code")
	candidates, err := ec.electionStore.GetClassRepResults(code)
	if err != nil {
		return response.ServerError(e, err, "")
	}
	return response.JSON(e, candidates)
}

func (ec *ElectionController) TotalVotesGET(e echo.Context) error {
	id := e.Param("id")
	votes, err := ec.electionStore.GetTotalVotes(id)
	if err != nil {
		return response.ServerError(e, err, "")
	}
	v := map[string]int64{"total_votes": votes}
	return response.JSON(e, v)
}
