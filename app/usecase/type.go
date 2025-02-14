package usecase

import "github/revaldimijaya/lacak-api/app/repository"

type ResponseScoredCity struct {
	Name      string  `json:"name"`
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
	Score     float64 `json:"score"`
}

type RequestScoredCity struct {
	Query  string
	LatStr string
	LonStr string
}

type Usecase struct {
	Repository repository.Repository
}
