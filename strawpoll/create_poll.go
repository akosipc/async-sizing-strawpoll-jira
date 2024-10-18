package strawpoll

import (
	"fmt"
	"bytes"
	"net/http"
	"encoding/json"
)

type CreatePollInput struct {
	Title					string 					`json:"title"`
	InputType				string					`json:"type" default:"multiple_choice"`

	Meta					PollMetaInput			`json:"meta"`
	Config					PollConfigInput			`json:"poll_config"`
	Options					[]PollOptionInput		`json:"poll_options"`
}

type PollMetaInput struct {
	Location				string			`json:"location"`
	Description				string			`json:"description"`
}

type PollOptionInput struct {
	Value					string			`json:"value"`
	Position				int				`json:"position"`
	MaxVotes				int				`json:"max_votes" default:"0"`
	InputType				string			`json:"type" default:"text"`
	Description				string			`json:"description"`
}

type PollConfigInput struct {
	DeadlineAt				int64			`json:"deadline_at"`

	IsPrivate				bool			`json:"is_private" default:"true"`
	RequireVoterNames 		bool			`json:"require_voter_names" default:"true"`
	ResultsVisibility		string			`json:"results_visibility" default:"after_vote"`
}

func CreatePollRequest(requestBody CreatePollInput) (*http.Request, error) {
	jsonBody, err := json.Marshal(requestBody)
	if err != nil {
		return nil, fmt.Errorf("error during marshalling the request body: %v", err)
	}

	req, err := http.NewRequest("POST", "https://api.strawpoll.com/v3/polls", bytes.NewBuffer(jsonBody))
	if err != nil {
		return nil, fmt.Errorf("error creating the request: %v", err)
	}

	req.Header.Set("X-API-Key", " 73df0ed6-8cea-11ef-92ef-ba349c717338")
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Content-Type", "application/json")

	return req, nil
}
