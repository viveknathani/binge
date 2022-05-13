package service

import (
	"context"

	"github.com/viveknathani/binge/database"
	"github.com/viveknathani/binge/entity"
	"github.com/viveknathani/binge/shared"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type Service struct {
	Database *database.Database
	Logger   *zap.Logger
}

func zapReqID(ctx context.Context) zapcore.Field {

	return zapcore.Field{
		Key:    "requestID",
		String: shared.ExtractRequestID(ctx),
		Type:   zapcore.StringType,
	}
}

func (service *Service) GetShowsAndMovies(ctx context.Context) (*[]entity.Show, *[]entity.Movie, error) {

	shows, err := service.Database.GetShows()
	if err != nil {
		service.Logger.Error(err.Error(), zapReqID(ctx))
		return nil, nil, err
	}

	movies, err := service.Database.GetMovies()
	if err != nil {
		service.Logger.Error(err.Error(), zapReqID(ctx))
		return nil, nil, err
	}

	return shows, movies, nil
}

func (service *Service) GetEpisodes(ctx context.Context, showId string) (*[]entity.Episode, error) {

	episodes, err := service.Database.GetEpisodes(showId)
	if err != nil {
		service.Logger.Error(err.Error(), zapReqID(ctx))
		return nil, err
	}

	return episodes, nil
}

func (service *Service) AddVideo(ctx context.Context, v *entity.Video) error {

	err := service.Database.CreateVideo(v)
	if err != nil {
		service.Logger.Error(err.Error(), zapReqID(ctx))
		return err
	}

	return nil
}

func (service *Service) AddShow(ctx context.Context, s *entity.Show) error {

	err := service.Database.CreateShow(s)
	if err != nil {
		service.Logger.Error(err.Error(), zapReqID(ctx))
		return err
	}

	return nil
}

func (service *Service) AddMovie(ctx context.Context, m *entity.Movie) error {

	err := service.Database.CreateMovie(m)
	if err != nil {
		service.Logger.Error(err.Error(), zapReqID(ctx))
		return err
	}

	return nil
}

func (service *Service) AddEpisode(ctx context.Context, e *entity.Episode) error {

	err := service.Database.CreateEpisode(e)
	if err != nil {
		service.Logger.Error(err.Error(), zapReqID(ctx))
		return err
	}

	return nil
}
