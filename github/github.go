package github

import (
	"context"
	"github.com/google/go-github/v35/github"
	"log"
)

type Template struct {
	Name               *string
	SortDesc           *string
	Author             *string
	SupportedPlatforms []string
	TemplateIcon       string
}

var githubClient *github.Client

func GetRepository(authToken string, owner string, repoString string) (*Template, error) {
	log.Printf("Getting %s repository details", repoString)
	repository, _, err := githubClient.Repositories.Get(context.Background(), owner, repoString)
	if err != nil {
		log.Printf("Getting %s repository details Error ", repoString)
		log.Print(err)
		return nil, err
	}

	template := new(Template)
	template.Name = repository.Name
	template.SortDesc = repository.Description
	template.Author = repository.Owner.Name
	template.SupportedPlatforms = repository.Topics
	template.TemplateIcon = ""
	log.Printf("Getting %s repository details successful ", repoString)
	return template, nil
}

func init() {
	githubClient = github.NewClient(nil)
}
