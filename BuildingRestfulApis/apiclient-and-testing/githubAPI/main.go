package main

import (
	"fmt"
	"github.com/levigross/grequests"
	"os"
)

var GITHUB_TOKEN = os.Getenv("GITHUB_TOKEN")
var ro = &grequests.RequestOptions{Auth: []string{GITHUB_TOKEN, "x-oauth-basic"}}

type Repo struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	FullName string `json:"fullName"`
	Forks    int    `json:"forks"`
	Private  bool   `json:"private"`
}

func getStats(url string) *grequests.Response {
	resp, err := grequests.Get(url, ro)
	// you can modify the request by passing an optional RequestOptions struct
	if err != nil {
		fmt.Println("Unable to make request")
		panic(err)
	}
	return resp
}

func main() {
	var repos []Repo
	var repoUrl = "https://api.github.com/users/rcolejnr/repos"
	resp := getStats(repoUrl)
	resp.JSON(&repos)
	fmt.Println(repos)
}
