package strawpoll

import (
	"testing"
)

func TestCreatePoll(t *testing.T) {
	input := PollRecord{
		Title: "(Test) ESS-1111",
		Description: "This is a test",
	}

	err := CreatePoll(input)

	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
}

func TestCreatePollTriggerValidation(t *testing.T) {
	input := PollRecord{
		Title: "",
		Description: "",
	}

	err := CreatePoll(input)

	if err == nil {
		t.Errorf("Expected validation errors, but got none. %v", err)
	}
}
