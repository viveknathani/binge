package database

import (
	"database/sql"

	"github.com/google/uuid"
	"github.com/viveknathani/binge/entity"
)

const (
	statementInsertVideo    = "insert into videos (id, \"videoId\") values ($1, $2);"
	statementInsertShow     = "insert into shows (id, name) values ($1, $2);"
	statementInsertEpisode  = "insert into episodes (id, \"episodeNumber\", name, season, \"videoId\", \"showId\") values ($1, $2, $3, $4, $5, $6);"
	statementInsertMovie    = "insert into movies (id, name, \"videoId\") values ($1, $2, $3);"
	statementSelectShows    = "select * from shows;"
	statementSelectMovies   = "select * from movies;"
	statementSelectEpisodes = "select * from episodes where text(\"showId\") = $1;"
)

func (db *Database) CreateVideo(v *entity.Video) error {

	v.Id = uuid.New().String()
	err := db.execWithTransaction(statementInsertVideo, v.Id, v.VideoId)
	return err
}

func (db *Database) CreateShow(s *entity.Show) error {

	s.Id = uuid.New().String()
	err := db.execWithTransaction(statementInsertShow, s.Id, s.Name)
	return err
}

func (db *Database) CreateMovie(m *entity.Movie) error {

	m.Id = uuid.New().String()
	err := db.execWithTransaction(statementInsertMovie, m.Id, m.Name, m.VideoId)
	return err
}

func (db *Database) CreateEpisode(e *entity.Episode) error {

	e.Id = uuid.New().String()
	err := db.execWithTransaction(statementInsertEpisode, e.Id, e.EpisodeNumber, e.Name, e.Season, e.VideoId, e.ShowId)
	return err
}

func (db *Database) GetMovies() (*[]entity.Movie, error) {

	result := make([]entity.Movie, 0)
	err := db.queryWithTransaction(statementSelectMovies, func(rows *sql.Rows) error {
		for rows.Next() {

			var m entity.Movie
			err := rows.Scan(&m.Id, &m.Name, &m.VideoId)
			if err != nil {
				return err
			}
			result = append(result, m)
		}
		return nil
	})

	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (db *Database) GetShows() (*[]entity.Show, error) {

	result := make([]entity.Show, 0)
	err := db.queryWithTransaction(statementSelectShows, func(rows *sql.Rows) error {
		for rows.Next() {

			var s entity.Show
			err := rows.Scan(&s.Id, &s.Name)
			if err != nil {
				return err
			}
			result = append(result, s)
		}
		return nil
	})

	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (db *Database) GetEpisodes(showId string) (*[]entity.Episode, error) {

	result := make([]entity.Episode, 0)
	err := db.queryWithTransaction(statementSelectEpisodes, func(rows *sql.Rows) error {
		for rows.Next() {

			var e entity.Episode
			err := rows.Scan(&e.Id, &e.EpisodeNumber, &e.Name, &e.Season, &e.ShowId, &e.VideoId)
			if err != nil {
				return err
			}
			result = append(result, e)
		}
		return nil
	}, showId)

	if err != nil {
		return nil, err
	}

	return &result, nil
}
