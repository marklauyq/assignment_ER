package main

import (
	"encoding/json"
	"net/http"
	"playground/pkg/application"
	"playground/pkg/domain"
)


type ResponseStruct struct {
	Success bool `json:"success"`
	Data    struct{
		Scores *domain.AverageScores `json:"scores,omitempty"`
	}`json:"data"`
	Errors []string `json:"errors"`
}


type RequestStruct struct {
	Scores domain.Scores `json:"scores"`
}

func GroupHandler(s domain.ScoreService) func (w http.ResponseWriter, req *http.Request) {
	return func (w http.ResponseWriter, req *http.Request) {
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		var request RequestStruct

		err := json.NewDecoder(req.Body).Decode(&request)
		if err != nil {
			writeError(w, err.Error(), http.StatusBadRequest)
			return
		}
		avgScore, err := s.GetAverageScore(request.Scores)
		if err != nil {
			writeError(w, err.Error(), http.StatusBadRequest)
			return
		}

		resp := ResponseStruct{
			Success: true,
			Data: struct {
				Scores *domain.AverageScores `json:"scores,omitempty"`
			}{
				Scores: avgScore,
			},
			Errors: []string{},
		}
		writeResponse(w, resp)
	}
}


func writeError(w http.ResponseWriter, s string, status int) {
	resp := ResponseStruct{
		Success: false,
		Errors: []string{
			s,
		},
	}
	w.WriteHeader(status)
	writeResponse(w, resp)
}

func writeResponse(w http.ResponseWriter, resp ResponseStruct) {
	data, err := json.Marshal(resp)
	if err != nil {
		w.Write([]byte("unable to marshal the response"))
		return
	}
	w.Write(data)
}

func main() {
	service := application.NewGroupService()
	http.HandleFunc("/", GroupHandler(service))

	http.ListenAndServe(":8080", nil)
}
