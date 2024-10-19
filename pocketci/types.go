package pocketci

import (
	"encoding/json"

	"dagger.io/dagger"
	"github.com/google/go-github/v61/github"
)

type Webhook struct {
	Vendor    string
	EventType string
	Payload   json.RawMessage
}

// CreatePipelineRequest is the payload received on pipeline creation.
type PipelineDoneRequest struct {
	ID int `json:"id"`
}

// PipelineClaimRequest is the payload received when a runner wants to claim
// a pipeline.
type PipelineClaimRequest struct {
	RunnerName string `json:"runner_name"`
}

const (
	GithubPullRequest = "pull_request"
	GithubPush        = "push"
	GithubRelease     = "release"
)

// GithubEvent is a wrapper of a github webhook. It centralizes all information
// used to handle a github event.
type GithubEvent struct {
	RawPayload []byte

	EventType string   `json:"event_type"`
	Changes   []string `json:"changes"`

	Repository     *dagger.Directory `json:"-"`
	RepositoryName string            `json:"repository_name"`

	PullRequestEvent *github.PullRequestEvent
	PushEvent        *github.PushEvent

	GitInfo
}

// Pipeline is a user-defined pipeline generated by pocketci's vendor modules.
type Pipeline struct {
	Repository   string   `json:"repository"`
	Runner       string   `json:"runner"`
	Changes      []string `json:"changes"`
	Module       string   `json:"module"`
	Name         string   `json:"name"`
	Actions      []string `json:"pr_actions"`
	OnPR         bool     `json:"on_pr"`
	BaseBranches []string `json:"on_pr_against"`
	OnPush       bool     `json:"on_push"`
	Branches     []string `json:"branches"`
	Exec         []string `json:"exec"`
	PipelineDeps []string `json:"after"`
}

// GitInfo collects all relevant git information that is sent attached to a given
// set of pipelines.
type GitInfo struct {
	Branch string `json:"branch"`
	SHA    string `json:"sha"`

	// Configured when the event comes from a pull request.
	BaseBranch string `json:"base_branch"`
	BaseSHA    string `json:"base_sha"`
	PRNumber   int    `json:"pr_number"`
}
