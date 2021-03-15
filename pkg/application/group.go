package application

import (
	"fmt"
	"playground/pkg/domain"
)

type scoreService struct {}


func (a scoreService) GetAverageScore(scores domain.Scores) (*domain.AverageScores, error) {
	if err := a.validateScore(scores); err != nil {
		return nil,err
	}

	avgScores := make(domain.AverageScores, len(scores))
	for gName, g := range scores {
		totalUsers := len(g)
		if totalUsers <= 2 {
			avgScores[gName] = 0
			continue
		}

		var totalScore float64
		for _, u := range g{
			totalScore += u.Score
		}
		avgScores[gName] = totalScore / float64(totalUsers)
	}

	return &avgScores,nil
}

const (
	ErrorMinOneUser       = `%s requires at least 1 user`
	ErrorOneUserExistOnce = `%d can only exist in 1 group at 1 time`
	ErrorScoreOutOfRange  = `user %d score should be between 0 and 5`
)
func (a scoreService) validateScore(scores domain.Scores) error {
	groupMap := make( map[int]bool)
	for gName, g := range scores {
		if len(g)== 0 {
			return fmt.Errorf(ErrorMinOneUser, gName)
		}

		for _, u := range g{
			_, ok := groupMap[u.UserID]
			if ok {
				return fmt.Errorf(ErrorOneUserExistOnce, u.UserID)
			}
			groupMap[u.UserID] = true

			if u.Score < 0  || u.Score > 5.0 {
				return fmt.Errorf(ErrorScoreOutOfRange, u.UserID)
			}
		}
	}
	return nil
}


func NewGroupService() domain.ScoreService {
	return &scoreService{}
}



