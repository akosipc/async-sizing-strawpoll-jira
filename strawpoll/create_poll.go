package strawpoll

import (
	"fmt"
	"bytes"
	"net/http"
	"encoding/json"
)

type CreatePollInput struct {
	Title					string 			`json:"title"`
	InputType				string			`json:"type" default:"multiple_choice"`

	Meta					PollMetaInput
	Config					PollConfigInput
	Options					[]PollOptionInput
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
	AllowComments			bool			`json:"allow_comments" default:"false"`
	AllowIndeterminate		bool			`json:"allow_indeterminate" default:"false"`
	AllowOtherOption		bool			`json:"allow_other_option" default:"false"`
	AllowVPNUsers			bool			`json:"allow_vpn_users" default:"false"`
	IsMultipleChoice		bool			`json:"is_multiple_choice" default:"false"`
	IsPrivate				bool			`json:"is_private" default:"false"`
	HideParticipants		bool			`json:"hide_participants" default:"true"`
	RandomizeOptions		bool			`json:"randomize_options" default:"false"`
	RequireVoterNames 		bool			`json:"require_voter_names" default:"true"`
	UseCustomDesign			bool			`json:"use_custom_design" default:"false"`

	DeadlineAt				int				`json:"deadline_at"`
	NumberOfWinners			int				`json:"number_of_winners" default:"1"`
	MultipleChoiceMin		int				`json:"multiple_choice_min" default:"1"`
	MultipleChoiceMax		int				`json:"multiple_choice_max" default:"1"`

	DuplicationChecking		string			`json:"duplication_checking" default:"ip"`
	EditVotePermissions		string			`json:"edit_vote_permissions" default:"voter"`
	ResultsVisibility		string			`json:"results_visibility" default:"after_vote"`
	VoteType				string			`json:"vote_type" default:"default"`
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

	req.Header.Set("X-API-TOKEN", " 73df0ed6-8cea-11ef-92ef-ba349c717338")
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Content-Type", "application/json")

	return req, nil
}
