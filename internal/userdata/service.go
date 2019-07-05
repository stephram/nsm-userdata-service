package userdata

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/go-openapi/runtime/middleware"
	log "github.com/sirupsen/logrus"
	"github.com/stephram/nsm-userdata-service/internal/repo"
	"github.com/stephram/nsm-userdata-service/pkg/app"
	"github.com/stephram/nsm-userdata-service/restapi/models"
	op "github.com/stephram/nsm-userdata-service/restapi/operations"
)

const (
	GameOnUser  = "GameOnUser"
	ConsoleUser = "ConsoleUser"
	TokenID     = "tokenId"
)

type Service interface {
	GetHealth(params op.GetHealthParams) middleware.Responder
	GetInfo(params op.GetInfoParams) middleware.Responder

	GetUser(params op.GetUserParams) middleware.Responder
	PostUser(params op.PostUserParams) middleware.Responder

	GetGameOnResults(params op.GetGameOnResultsParams) middleware.Responder
	PostGameOnResults(params op.PostGameOnResultsParams) middleware.Responder

	GetInteractions(params op.GetInteractionsParams) middleware.Responder
	PostInteraction(params op.PostInteractionParams) middleware.Responder
}

type serviceImpl struct {
	repo repo.UserRepo
}

func createRepo(testMode bool) repo.UserRepo {
	return repo.New(nil)
}

func New() Service {
	return &serviceImpl{
		repo: createRepo(true),
	}
}

func (s *serviceImpl) GetHealth(params op.GetHealthParams) middleware.Responder {
	log.WithField("params", params).Infof("")
	return op.NewGetHealthOK().WithPayload(&models.HealthResponse{
		Status: aws.String("UP"),
	})
}

func (s *serviceImpl) GetInfo(params op.GetInfoParams) middleware.Responder {
	log.WithField("params", params).Infof("")
	return op.NewGetInfoOK().WithPayload(&models.InfoResponse{
		Name:       aws.String(app.Name),
		Author:     aws.String(app.Author),
		BranchName: aws.String(app.Branch),
		BuildDate:  aws.String(app.BuildDate),
		GitCommit:  aws.String(app.Commit),
		Version:    aws.String(app.Version),
	})
}

func (s *serviceImpl) GetUser(params op.GetUserParams) middleware.Responder {
	var user *models.User
	var err error

	if user, err = s.repo.FindUser(params.TokenID); err != nil {
		return op.NewGetUserNotFound().WithPayload(&models.APIError{
			Code:    aws.Int64(-9),
			Message: aws.String(err.Error()),
		})
	}
	return op.NewGetUserOK().WithPayload(user)
}

func (s *serviceImpl) PostUser(params op.PostUserParams) middleware.Responder {
	var user *models.User
	var err error

	if user, err = s.repo.StoreUser(params.User); err != nil {
		return op.NewPostUserInternalServerError().WithPayload(&models.APIError{
			Code:    aws.Int64(-10),
			Message: aws.String(err.Error()),
		})
	}

	if len(*params.User.TokenID) > 0 {
		// It's an update
		return op.NewPostUserOK().WithPayload(user)
	}
	// It's a new one
	return op.NewPostUserCreated().WithPayload(user)
}

func (s *serviceImpl) GetGameOnResults(params op.GetGameOnResultsParams) middleware.Responder {
	var gameOnResults *models.GameOnResults
	var err error

	if gameOnResults, err = s.repo.FindGameOnResults(params.TokenID); err != nil {
		return op.NewGetGameOnResultsNotFound().WithPayload(&models.APIError{
			Code:    aws.Int64(-12),
			Message: aws.String(err.Error()),
		})
	}
	return op.NewGetGameOnResultsOK().WithPayload(gameOnResults)
}

func (s *serviceImpl) PostGameOnResults(params op.PostGameOnResultsParams) middleware.Responder {
	var gor *models.GameOnResults
	var err error

	if gor, err = s.repo.StoreGameOnResults(params.TokenID, params.GameOnResults); err != nil {
		return op.NewPostGameOnResultsInternalServerError().WithPayload(&models.APIError{
			Code:    aws.Int64(-11),
			Message: aws.String(err.Error()),
		})
	}
	return op.NewPostGameOnResultsOK().WithPayload(gor)
}

func (s *serviceImpl) GetInteractions(params op.GetInteractionsParams) middleware.Responder {
	return op.NewGetInteractionsOK().WithPayload(models.ConsoleInteractions{})
}

func (s *serviceImpl) PostInteraction(params op.PostInteractionParams) middleware.Responder {
	return op.NewPostInteractionOK().WithPayload(&models.ConsoleInteraction{})
}
