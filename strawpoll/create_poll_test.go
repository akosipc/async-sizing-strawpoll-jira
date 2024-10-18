package strawpoll

import (
	"testing"
)

func TestCreatePollRequest(t *testing.T) {
	input := CreatePollInput{
		Title: "Test Poll",
		Meta: PollMetaInput{
			Location: "Test Location",
			Description: "Test Description",
		},
		Options: []PollOptionInput{
			{
				InputType: "text",
				Value: "Option 1",
				Position: 0,
			},
			{
				InputType: "text",
				Value: "Option 2",
				Position: 1,
			},
		},
	}	

	req, err := CreatePollRequest(input)

	if err != nil {
		t.Fatalf("CreatePollRequest returned an error: %v", err)
	}

	if req.Method != "POST" {
		t.Errorf("Expected POST method, got %s", req.Method)
	}

	expectedURL := "https://api.strawpoll.com/v3/polls"
	if req.URL.String() != expectedURL {
		t.Errorf("Expected URL %s, got %s", expectedURL, req.URL.String())
	}
}


