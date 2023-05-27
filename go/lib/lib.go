package lib

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
)

const (
	InputUrlFormat = "https://hackattic.com/challenges/%s/problem?access_token=%s"
	SolveUrlFormat = "https://hackattic.com/challenges/%s/solve?access_token=%s"
)

func GetChallengeInput(problem, accessToken string) []byte {
	inputUrl := fmt.Sprintf(InputUrlFormat, problem, accessToken)
	resp, err := http.Get(inputUrl)

	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()

	input, err := io.ReadAll(resp.Body)

	if err != nil {
		panic(err)
	}

	return input
}

func SubmitChallengeSolution(problem, accessToken, soln string) {
	submitUrl := fmt.Sprintf(SolveUrlFormat, problem, accessToken)

	solutionBuffer := bytes.NewBuffer([]byte(soln))
	resp, err := http.Post(submitUrl, "application/json", solutionBuffer)

	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()

	submissionResult, err := io.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	fmt.Println(string(submissionResult))
}
