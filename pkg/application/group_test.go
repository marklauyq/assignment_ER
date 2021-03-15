package application

import (
	"fmt"
	"playground/pkg/domain"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGroupService(t *testing.T) {
	testSuite := []struct{
		name string
		scores domain.Scores
		expected interface{}
	}{
		{
			name:"simple test case",
			scores:map[string][]domain.User{
				"managers" : {
					{UserID: 1, Score: 5},
					{UserID: 2, Score: 5},
					{UserID: 3, Score: 5},
				},
			},
			expected: &domain.AverageScores{
				"managers":5.0,
			},
		},
		{
			name:"test decimal point average score",
			scores:map[string][]domain.User{
				"managers" : {
					{UserID: 1, Score: 1},
					{UserID: 2, Score: 5},
					{UserID: 3, Score: 3},
					{UserID: 4, Score: 1},
				},
			},
			expected: &domain.AverageScores{
				"managers":2.5,
			},
		},
		{
			name:"2 groups",
			scores:map[string][]domain.User{
				"managers" : {
					{UserID: 1, Score: 5},
					{UserID: 2, Score: 5},
					{UserID: 3, Score: 5},
				},
				"users" : {
					{UserID: 4, Score: 4},
					{UserID: 5, Score: 3},
					{UserID: 6, Score: 5},
				},
			},
			expected: &domain.AverageScores{
				"managers":5.0,
				"users":4.0,
			},
		},
		{
			name:"less than 2 users should return 0",
			scores:map[string][]domain.User{
				"managers" : {
					{UserID: 1, Score: 1},
					{UserID: 2, Score: 5},
				},
			},
			expected: &domain.AverageScores{
				"managers":0,
			},
		},

		{
			name:"no users in group should throw an error",
			scores:map[string][]domain.User{
				"managers" : {},
			},
			expected: fmt.Errorf(ErrorMinOneUser, "managers"),
		},
		{
			name:"User cannot exist in multiple groups",
			scores:map[string][]domain.User{
				"managers" : {
					{UserID: 1, Score: 1},
					{UserID: 2, Score: 5},
				},
				"random" : {
					{UserID: 1, Score: 1},
				},
			},
			expected: fmt.Errorf(ErrorOneUserExistOnce, 1),
		},
		{
			name:"Score cannot be greater than 5",
			scores:map[string][]domain.User{
				"managers" : {
					{UserID: 1, Score: 1},
					{UserID: 2, Score: 5},
				},
				"random" : {
					{UserID: 3, Score: 6},
				},
			},
			expected: fmt.Errorf(ErrorScoreOutOfRange, 3),
		},
		{
			name:"Score cannot be greater less than 0",
			scores:map[string][]domain.User{
				"managers" : {
					{UserID: 1, Score: 1},
					{UserID: 2, Score: 5},
				},
				"random" : {
					{UserID: 3, Score: -1},
				},
			},
			expected: fmt.Errorf(ErrorScoreOutOfRange, 3),
		},
	}

	for _, tt := range testSuite {
		t.Run(tt.name, func(t *testing.T){
			s := NewGroupService()
			result,err := s.GetAverageScore(tt.scores)

			switch tt.expected.(type) {
			case error :
				assert.Error(t,err)
				assert.Equal(t, err, tt.expected)
			case *domain.AverageScores:
				assert.Equal(t, tt.expected, result)
			}
		})
	}
}
