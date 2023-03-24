package main

import (
	"encoding/json"
	"fmt"
	"github.com/levigross/grequests"
	"github.com/urfave/cli"
	"io/ioutil"
	"os"
)

var GITHUB_TOKEN = os.Getenv("GITHUB_TOKEN")
var ro = &grequests.RequestOptions{}

// struct to hold response of repos fetched
type Repo struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	FullName string `json:"fullName"`
	Forks    int    `json:"forks"`
	Private  bool   `json:"private"`
}

// Structs for modelling json body in create gist
type File struct {
	Content string `json:"content"`
}

type Gist struct {
	Description string          `json:"description"`
	Public      bool            `json:"public"`
	Files       map[string]File `json:"files"`
}

// Fetches the repos for the given Github users
// getStats fetches the github api for repos
func getStats(url string) *grequests.Response {
	resp, err := grequests.Get(url, ro)
	// you can modify the request by passing an optional RequestOptions struct
	if err != nil {
		fmt.Println("Unable to make request")
		panic(err)
	}
	return resp
}

// Read the files provided and creates Gist on GitHub
func createGist(url string, args []string) *grequests.Response {
	// get first two arguments
	description := args[0]
	// remaining arguments are file names with path
	var fileContents = make(map[string]File)
	for i := 1; i < len(args); i++ {
		dat, err := ioutil.ReadFile(args[i])
		if err != nil {
			fmt.Println("Please chech the  Absolute path (or) same directory are allowed")
			return nil
		}
		var file File
		file.Content = string(dat)
		fileContents[args[i]] = file
	}

	var gist = Gist{Description: description, Public: true, Files: fileContents}

	// Add data to json field
	var postBody, _ = json.Marshal(gist)
	var roCopy = ro
	roCopy.JSON = string(postBody)
	// make a Post request to github
	resp, err := grequests.Post(url, roCopy)
	if err != nil {
		fmt.Println("Creat request failed for github api")
	}
	return resp
}

func main() {
	app := cli.NewApp()
	// define command for our client
	app.Commands = []cli.Command{
		{
			Name:    "fetch",
			Aliases: []string{"f"},
			Usage:   "Fetch the repo details with user, [Usage]: goTool fetch user",
			Action:  actionFetch,
		},
		{
			Name:    "create",
			Aliases: []string{"c"},
			Usage:   "Creates a gist from the given text. [Usage]: goTol name `description` sample.txt",
			Action:  actionCreate,
		},
	}
	app.Version = "1.0"
	app.Run(os.Args)
}

func actionFetch(c *cli.Context) error {
	if c.NArg() > 0 {
		// GITHUB API logic
		var repos []Repo
		var repoUrl = "https://api.github.com/users/rcolejnr/repos"
		resp := getStats(repoUrl)
		resp.JSON(&repos)
		fmt.Println(repos)
	} else {
		fmt.Println("Please give a username. See -h to see help")
	}
	return nil
}

func actionCreate(c *cli.Context) error {
	if c.NArg() > 1 {
		// Github API logic
		args := c.Args()
		var postUrl = "https://api.github.com/gists"
		resp := createGist(postUrl, args)
		fmt.Println(resp.String())
	} else {
		fmt.Println("Please give sufficient arguments. See -h to see help")
	}
	return nil
}
