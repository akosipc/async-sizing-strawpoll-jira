package strawpoll

import (
	"io"
	"fmt"
	"time"
	"net/http"

	"github.com/go-playground/validator/v10"
)


type PollRecord struct {
	Title			string			`validate:"required"`
	Description		string
}

func CreatePoll(poll PollRecord) error {
	validate := validator.New()

	err := validate.Struct(poll)
	if err != nil {
		return fmt.Errorf("error with the struct: %v", err)
	}

	req, err := CreatePollRequest(
		CreatePollInput{
			Title: poll.Title,
			Meta: PollMetaInput {
				Location: "From async-sizing-bot",
				Description: poll.Description,
			},
			Config: PollConfigInput{ 
				DeadlineAt: time.Now().AddDate(0, 0, 3).Unix(),
			},
			Options: []PollOptionInput{
				{
					InputType: "text",
					Value: "1",
					Position: 0,
				},
				{
					InputType: "text",
					Value: "2",
					Position: 1,
				},
				{
					InputType: "text",
					Value: "3",
					Position: 2,
				},
				{
					InputType: "text",
					Value: "5",
					Position: 3,
				},
				{
					InputType: "text",
					Value: "8",
					Position: 4,
				},
				{
					InputType: "text",
					Value: "amigo session please",
					Position: 5,
				},
			},
		},
	)

	if err != nil {
		return fmt.Errorf("error building the request: %v", err)
	}

	client := &http.Client{}	
	response, err := client.Do(req)

	if err != nil {
		return fmt.Errorf("error sending request: %v", err)
	}
	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)
	if err != nil {
		return fmt.Errorf("error reading response body: %v", err)
	}

	if response.StatusCode != http.StatusCreated {
		return fmt.Errorf("unexpected status code: %d, body: %s", response.StatusCode, string(body))
	}

	fmt.Printf("Response: %s\n", string(body))
	return nil
}
