package bsapi

import (
	"net/url"
)

type Player struct {
	Club            PlayerClub `json:"club"`
	TrioVictories   int        `json:"3vs3Victories"`
	SoloVictories   int        `json:"soloVictories"`
	DuoVictories    int        `json:"duoVictories"`
	Icon            PlayerIcon `json:"icon"`
	ExpLevel        int        `json:"expLevel"`
	ExpPoints       int        `json:"expPoints"`
	Tag             string     `json:"tag"`
	Name            string     `json:"name"`
	Trophies        int        `json:"trophies"`
	HighestTrophies int        `json:"highestTrophies"`
	NameColor       string     `json:"nameColor"`
}

type PlayerClub struct {
	Tag  string `json:"tag"`
	Name string `json:"name"`
}

type PlayerIcon struct {
	Id int `json:"id"`
}

func (api BsApi) GetPlayerStats(tag string) (Player, error) {
	url := "/players/" + url.QueryEscape(tag)
	data, err := api.makeRequest(url)
	if err != nil {
		return Player{}, err
	}
	player := Player{}
	err = json.Unmarshal(data, &player)
	if err != nil {
		return Player{}, err
	}
	return player, nil
}
