package domain

type ScoreService interface {
	GetAverageScore(Scores) (*AverageScores, error)
}

type AverageScores map[string]float64

type Scores map[string][]User

type Group struct {
	Users []User
}

type User struct {
	UserID int `json:"userId"`
	Score float64 `json:"score"`
}
