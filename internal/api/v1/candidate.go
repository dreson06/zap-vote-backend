package v1

import (
	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
	"zapvote/internal/api/response"
	"zapvote/internal/model/candidate"
	"zapvote/internal/services/candidatestore"
)

type CandidateController struct {
	candidateStore candidatestore.Store
}

func NewCandidateController(candidateStore candidatestore.Store) *CandidateController {
	return &CandidateController{
		candidateStore: candidateStore,
	}
}

func (cc *CandidateController) AddPOST(e echo.Context) error {
	body := &candidate.Candidate{}
	if err := e.Bind(body); err != nil {
		return response.BadRequestError(e, "wrong body")
	}
	if body.Name == "" || body.CourseCode == "" || body.Position == "" || body.Type == "" {
		return response.BadRequestError(e, "information missing")
	}
	c := candidate.New()
	c.ID = primitive.NewObjectID().Hex()
	c.Name = body.Name
	c.Department = body.Department
	c.CourseCode = body.CourseCode
	c.Position = body.Position
	c.Type = body.Type
	c.CreatedAt = time.Now()
	c.UpdatedAt = time.Now()

	err := cc.candidateStore.Create(c)
	if err != nil {
		return response.ServerError(e, err, "")
	}
	return response.Success(e)
}

func (cc *CandidateController) GetFromDepartment(e echo.Context) error {
	department := e.QueryParam("department")
	if department == "" {
		return response.BadRequestError(e, "department required")
	}
	candidates, err := cc.candidateStore.GetCandidateByDepartment(department)
	if err != nil {
		return response.ServerError(e, err, "")
	}
	return response.JSON(e, candidates)
}

//func (cc *CandidateController) SpecificGET(e echo.Context) error {
//	course := e.QueryParam("course")
//	if course == "" {
//		return response.BadRequestError(e, "can not find course")
//	}
//	candidates, err := cc.candidateStore.GetSpecific(course)
//	if err != nil {
//		return response.ServerError(e, err, "")
//	}
//	return response.JSON(e, candidates)
//}
//
//func (cc *CandidateController) GeneralGET(e echo.Context) error {
//	candidates, err := cc.candidateStore.GetGeneral()
//	if err != nil {
//		return response.ServerError(e, err, "")
//	}
//	return response.JSON(e, candidates)
//}
