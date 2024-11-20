package service

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/giuszeppe/github-activity-go-cli/api"
)

type Activity struct {
	ID    string `json:"id"`
	Type  string `json:"type"`
	Actor struct {
		ID           int    `json:"id"`
		Login        string `json:"login"`
		DisplayLogin string `json:"display_login"`
		GravatarID   string `json:"gravatar_id"`
		URL          string `json:"url"`
		AvatarURL    string `json:"avatar_url"`
	} `json:"actor"`
	Repo struct {
		ID   int    `json:"id"`
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"repo"`
	Payload struct {
		Ref          any    `json:"ref"`
		RefType      string `json:"ref_type"`
		MasterBranch string `json:"master_branch"`
		Description  any    `json:"description"`
		PusherType   string `json:"pusher_type"`
	} `json:"payload"`
	Public    bool      `json:"public"`
	CreatedAt time.Time `json:"created_at"`
}

func GetActivityForUsername(username string) ([]string, error) {
	api := api.GetAPI()
	out, err := api.Fetch("/users/" + username + "/events")
	if err != nil {
		return []string{}, err
	}

	activities := parseActivityJson(out)
	result := formatActivitiesString(activities)

	return result, nil
}

func parseActivityJson(text []byte) []Activity {
	m := []Activity{}
	json.Unmarshal(text, &m)
	return m
}

func formatActivitiesString(activities []Activity) []string {
	res := []string{}
	for _, activity := range activities {
		prefix := "- "
		switch activity.Type {
		case "CreateEvent":
			res = append(res, fmt.Sprintf("%sCreated repo %v", prefix, activity.Repo.Name))
		case "DeleteEvent":
			res = append(res, fmt.Sprintf("%sDeleted event in repo %s", prefix, activity.Repo.Name))
		case "ForkEvent":
			res = append(res, fmt.Sprintf("%sForked repo %s", prefix, activity.Repo.Name))
		case "GollumEvent":
			res = append(res, fmt.Sprintf("%sUpdated wiki in repo %s", prefix, activity.Repo.Name))
		case "IssueCommentEvent":
			res = append(res, fmt.Sprintf("%sCommented on issue in repo %s", prefix, activity.Repo.Name))
		case "IssuesEvent":
			res = append(res, fmt.Sprintf("%sIssue event in repo %s", prefix, activity.Repo.Name))
		case "MemberEvent":
			res = append(res, fmt.Sprintf("%sMember event in repo %s", prefix, activity.Repo.Name))
		case "PublicEvent":
			res = append(res, fmt.Sprintf("%sMade repo %s public", prefix, activity.Repo.Name))
		case "PullRequestEvent":
			res = append(res, fmt.Sprintf("%sPull request event in repo %s", prefix, activity.Repo.Name))
		case "PullRequestReviewEvent":
			res = append(res, fmt.Sprintf("%sReviewed pull request in repo %s", prefix, activity.Repo.Name))
		case "PullRequestReviewCommentEvent":
			res = append(res, fmt.Sprintf("%sCommented on pull request review in repo %s", prefix, activity.Repo.Name))
		case "PullRequestReviewThreadEvent":
			res = append(res, fmt.Sprintf("%sPull request review thread event in repo %s", prefix, activity.Repo.Name))
		case "PushEvent":
			res = append(res, fmt.Sprintf("%sPushed to repo %s", prefix, activity.Repo.Name))
		case "ReleaseEvent":
			res = append(res, fmt.Sprintf("%sReleased in repo %s", prefix, activity.Repo.Name))
		case "SponsorshipEvent":
			res = append(res, fmt.Sprintf("%sSponsorship event in repo %s", prefix, activity.Repo.Name))
		case "WatchEvent":
			res = append(res, fmt.Sprintf("%sWatched repo %s", prefix, activity.Repo.Name))
		}

	}
	return res
}
