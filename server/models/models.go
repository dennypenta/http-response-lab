package models

import (
	"encoding/json"
	"time"
)

type TimeNanosec int64

func (t *TimeNanosec) UnmarshalJSON(data []byte) error {
	var tt time.Time
	if err := json.Unmarshal(data, &tt); err != nil {
		return err
	}

	*t = TimeNanosec(tt.UnixNano())
	return nil
}

type AutoGenerated struct {
	CreatedAt TimeNanosec `json:"created_at"`
	Public    bool        `json:"public"`
	ID        string      `json:"id"`
	Type      string      `json:"type"`
	Actor     struct {
		ID         int    `json:"id"`
		Login      string `json:"login"`
		GravatarID string `json:"gravatar_id"`
		URL        string `json:"url"`
		AvatarURL  string `json:"avatar_url"`
	} `json:"actor"`
	Repo struct {
		ID   int    `json:"id"`
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"repo"`
	Payload struct {
		Ref          string `json:"ref"`
		RefType      string `json:"ref_type"`
		MasterBranch string `json:"master_branch"`
		Description  string `json:"description"`
		PusherType   string `json:"pusher_type"`
	} `json:"payload"`
}
