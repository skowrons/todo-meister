package github

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

const (
	gitHubBaseURL = "https://api.github.com"
)

type gitHubIssues struct {
	ID    int    `json:"id"`
	Title string `json:"title"`
	Body  string `json:"body"`
}

func getIssuesByLabel(github *github, label string) []gitHubIssues {
	url := fmt.Sprintf("%s/repos/%s/%s?labels=%s", gitHubBaseURL, github.Owner, github.Repo, label)
	log.Println(url)
	req, _ := http.NewRequest(http.MethodGet, url, nil)
	req.Header.Add("Authorization", github.AccessKey)

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		panic(err)
	}
	defer res.Body.Close()

	data, err := io.ReadAll(res.Body)
	if err != nil {
		panic(err)
	}

	log.Println(string(data))

	return []gitHubIssues{}
}
