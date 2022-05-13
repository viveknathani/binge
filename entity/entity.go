package entity

type Video struct {
	Id      string `json:"id"`
	VideoId string `json:"videoId"`
}

type Movie struct {
	Id      string `json:"id"`
	Name    string `json:"name"`
	VideoId string `json:"videoId"`
}

type Show struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}

type Episode struct {
	Id            string `json:"id"`
	Name          string `json:"name"`
	ShowId        string `json:"showId"`
	VideoId       string `json:"videoId"`
	EpisodeNumber int    `json:"episodeNumber"`
	Season        int    `json:"season"`
}
