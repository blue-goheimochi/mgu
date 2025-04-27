package commands

import (
	"github.com/AlecAivazis/survey/v2"
	"github.com/blue-goheimochi/mgu/pkg/git"
)

// defaultRepositoryFactory creates a new Git repository
func defaultRepositoryFactory() git.Repository {
	return git.NewLocalRepository()
}

// For testing - allow injecting a repository
var repositoryFactory = defaultRepositoryFactory

// Function type for survey.Ask to allow mocking
type SurveyAskFunc func(qs []*survey.Question, response interface{}) error

// Default implementation of SurveyAskFunc that delegates to actual survey.Ask
func defaultSurveyAsk(qs []*survey.Question, response interface{}) error {
	return survey.Ask(qs, response)
}

// For testing - allow injecting a survey.Ask implementation
var askFunc SurveyAskFunc = defaultSurveyAsk

// Function type for survey.AskOne to allow mocking
type SurveyAskOneFunc func(prompt survey.Prompt, response interface{}, opts ...survey.AskOpt) error

// Default implementation of SurveyAskOneFunc that delegates to actual survey.AskOne
func defaultSurveyAskOne(prompt survey.Prompt, response interface{}, opts ...survey.AskOpt) error {
	return survey.AskOne(prompt, response, opts...)
}

// For testing - allow injecting a survey.AskOne implementation
var askOneFunc SurveyAskOneFunc = defaultSurveyAskOne