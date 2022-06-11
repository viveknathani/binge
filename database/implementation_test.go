package database

import (
	"log"
	"os"
	"testing"

	"github.com/viveknathani/binge/entity"
)

const dsn = "postgres://viveknathani:root@localhost:5432/binge?sslmode=disable"

var db *Database

func TestMain(t *testing.M) {

	db = &Database{}
	err := db.Initialize(dsn)
	if err != nil {
		log.Fatal(err)
	}

	// create tables
	_, err = db.pool.Exec("create table if not exists videos(id uuid primary key,\"videoId\" varchar not null unique);")
	if err != nil {
		log.Fatal(err)
	}

	_, err = db.pool.Exec("create table if not exists shows(id uuid primary key, name varchar not null);")
	if err != nil {
		log.Fatal(err)
	}

	_, err = db.pool.Exec("create table if not exists episodes(id uuid primary key, \"episodeNumber\" int not null, name varchar not null, season int not null, \"showId\" uuid references shows(\"id\") on delete cascade, \"videoId\" varchar references videos(\"videoId\") on delete cascade);")
	if err != nil {
		log.Fatal(err)
	}

	_, err = db.pool.Exec("create table if not exists movies(id uuid primary key, name varchar not null, \"videoId\" varchar references videos(\"videoId\") on delete cascade);")
	if err != nil {
		log.Fatal(err)
	}

	code := t.Run()
	db.Close()
	os.Exit(code)
}

func TestCreateVideo(t *testing.T) {

	v := &entity.Video{
		VideoId: "abc2402",
	}

	err := db.CreateVideo(v)
	if err != nil {
		log.Fatal(err)
	}
}

func TestCreateAndGetMovie(t *testing.T) {

	v := &entity.Video{
		VideoId: "fancy2402",
	}

	m := &entity.Movie{
		Name:    "Begin Again",
		VideoId: v.VideoId,
	}

	err := db.CreateVideo(v)
	if err != nil {
		log.Fatal(err)
	}

	err = db.CreateMovie(m)
	if err != nil {
		log.Fatal(err)
	}

	got, err := db.GetMovies()
	if err != nil {
		log.Fatal(err)
	}
	if len(*got) != 1 {
		log.Fatal("Incorrect length for movies array.")
	}

	if (*got)[0] != *m {
		log.Println(*m)
		log.Println((*got)[0])
		log.Fatal("Inequality.")
	}
}

func TestCreateAndGetShow(t *testing.T) {

	s := &entity.Show{
		Name: "Friends",
	}

	err := db.CreateShow(s)
	if err != nil {
		log.Fatal(err)
	}

	got, err := db.GetShows()
	if err != nil {
		log.Fatal(err)
	}
	if len(*got) != 1 {
		log.Fatal("Incorrect length for movies array.")
	}

	if (*got)[0] != *s {
		log.Println(*s)
		log.Println((*got)[0])
		log.Fatal("Inequality.")
	}
}

func TestCreateAndGetEpisodes(t *testing.T) {

	s := &entity.Show{
		Name: "HIMYM",
	}

	err := db.CreateShow(s)
	if err != nil {
		log.Fatal(err)
	}

	v := &entity.Video{
		VideoId: "no2402",
	}

	err = db.CreateVideo(v)
	if err != nil {
		log.Fatal(err)
	}

	e := &entity.Episode{
		Name:          "Pilot",
		ShowId:        s.Id,
		VideoId:       v.VideoId,
		EpisodeNumber: 1,
		Season:        1,
	}

	err = db.CreateEpisode(e)
	if err != nil {
		log.Fatal(err)
	}

	got, err := db.GetEpisodes(s.Id)
	if err != nil {
		log.Fatal(err)
	}
	if len(*got) != 1 {
		log.Fatal("Incorrect length for movies array.")
	}

	if (*got)[0] != *e {
		log.Println(*e)
		log.Println((*got)[0])
		log.Fatal("Inequality.")
	}
}
