package bsapi

import (
	"bytes"
	"net/http"
	"net/url"
)

type Club struct {
	Tag              string       `json:"tag"`
	Name             string       `json:"name"`
	Description      string       `json:"description"`
	Trophies         int          `json:"trophies"`
	RequiredTrophies int          `json:"requiredTrophies"`
	Members          []ClubMember `json:"members"`
	Type             string       `json:"type"`
	BadgeId          int          `json:"badgeId"`
}

type ClubMember struct {
	IconId    int    `json:"icon"`
	Tag       string `json:"tag"`
	Name      string `json:"name"`
	Trophies  int    `json:"trophies"`
	Role      string `json:"role"`
	NameColor string `json:"nameColor"`
}

func (api BsApi) GetClubStats(tag string) (Club, error) {
	client := http.Client{}
	req, err := http.NewRequest("GET", apiUrl+"/clubs/"+url.QueryEscape(tag), nil)
	if err != nil {
		return Club{}, err
	}
	req.Header.Set("Authorization", "Bearer "+api.token)
	if err != nil {
		return Club{}, err
	}
	resp, err := client.Do(req)
	if err != nil {
		return Club{}, err
	}
	defer resp.Body.Close()
	buf := bytes.Buffer{}
	buf.ReadFrom(resp.Body)
	data := buf.Bytes()
	club := Club{}
	json.Unmarshal(data, &club)
	return club, nil
}
