package bsapi

import (
	"net/url"
)

type ClubRole string

const (
	ClubRolePresident     string = "president"
	ClubRoleVicePresident string = "vicePresident"
	ClubRoleSenior        string = "senior"
	ClubRoleMember        string = "member"
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
	Icon      PlayerIcon `json:"icon"`
	Tag       string     `json:"tag"`
	Name      string     `json:"name"`
	Trophies  int        `json:"trophies"`
	Role      ClubRole   `json:"role"`
	NameColor string     `json:"nameColor"`
}

func (api BsApi) GetClubStats(tag string) (Club, error) {
	url := "/clubs/" + url.QueryEscape(tag)
	data, err := api.makeRequest(url)
	if err != nil {
		return Club{}, err
	}
	club := Club{}
	err = json.Unmarshal(data, &club)
	if err != nil {
		return Club{}, err
	}
	return club, nil
}
