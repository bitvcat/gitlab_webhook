package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"
)

/*
[xxx:main] push 3 new commits by @xxxx
- commit_id1 name - commit_msg1
- commit_id2 name - commit_msg2
*/
type PushType struct {
	Ref      string
	User     string `json:"user_name"`
	UserName string `json:"user_username"`
	Commits  []CommitType
}

func handlePush(h *HeaderType, data []byte) {
	var pushData PushType
	err := json.Unmarshal(data, &pushData)
	if err != nil {
		log.Printf("Failed to parse request: %s", err)
		return
	}
	s := strings.Split(pushData.Ref, "/")
	branchName := s[len(s)-1]
	markdowns := make([]string, 0)

	// title
	branch := fmt.Sprintf("[[%s:%s](%s/-/tree/%s)]", h.Repository.Name, branchName, h.Repository.HomePage, branchName)
	commitNum := len(pushData.Commits)
	pushTitle := fmt.Sprintf("%s push %d new commits by @%s", branch, commitNum, pushData.UserName)
	markdowns = append(markdowns, pushTitle)

	// commits
	for _, commit := range pushData.Commits {
		commitStr := fmt.Sprintf("- [%s](%s) %s - %s", commit.Id[0:8], commit.Url, commit.Author.Name, commit.Message)
		markdowns = append(markdowns, commitStr)
	}
	//fmt.Println(strings.Join(markdowns, "\n"))

	// send to mm
	http.Post(FlagHook, "application/json", strings.NewReader("{\"text\":\""+strings.Join(markdowns, "\n")+"\"}"))
}
