package seeker

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type seeker struct {
	client    http.Client
	userAgent string
}

type UserInfo struct {
	ID             string `json:"id"`
	FollowersCount int    `json:"followersNumber"`
	IsPrivate      bool   `json:"isPrivate"`
	FollowedByUser bool   `json:"followed_by_viewer"`
}

type Story struct {
	MediaType int    `json:"media_type"`
	ImageURL  string `json:"image_url,omitempty"`
	VideoURL  string `json:"video_url,omitempty"`
}

func NewSeeker(client http.Client, userAgent string) Seeker {
	return &seeker{
		client:    client,
		userAgent: userAgent,
	}
}

func (s *seeker) Get(cookies, username string) ([]Story, error) {
	userInfo, err := s.getID(cookies, username)
	if err != nil {
		return nil, err
	}

	if userInfo.IsPrivate && !userInfo.FollowedByUser {
		return nil, fmt.Errorf("The account is private, you cannot get followers info")
	}

	storiesURL := fmt.Sprintf("https://i.instagram.com/api/v1/feed/user/%s/reel_media/", userInfo.ID)
	req, err := http.NewRequest("GET", storiesURL, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("cookie", cookies)
	req.Header.Set("user-agent", s.userAgent)

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var response struct {
		Items []Story `json:"items"`
	}
	if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
		return nil, err
	}

	if len(response.Items) == 0 {
		return nil, fmt.Errorf("There are no stories")
	}

	return response.Items, nil
}

func (s *seeker) getID(cookies, username string) (*UserInfo, error) {
	getInfoURL := fmt.Sprintf("https://i.instagram.com/api/v1/users/web_profile_info/?username=%s", username)
	req, err := http.NewRequest("GET", getInfoURL, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("cookie", cookies)
	req.Header.Set("user-agent", s.userAgent)

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var userInfo UserInfo
	if err := json.NewDecoder(resp.Body).Decode(&userInfo); err != nil {
		return nil, err
	}

	return &userInfo, nil
}
