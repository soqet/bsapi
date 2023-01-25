package bsapi

import (
	"bytes"
	"net/http"
	"net/url"
)

type Player struct {
	Club            PlayerClub `json:"club"`
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

func (api BsApi) AGetPlayerStats(tag string) (Player, error) {
	client := http.Client{}
	req, err := http.NewRequest("GET", apiUrl+"/players/"+url.QueryEscape(tag), nil)
	if err != nil {
		return Player{}, err
	}
	req.Header.Set("Authorization", "Bearer "+api.token)
	if err != nil {
		return Player{}, err
	}
	resp, err := client.Do(req)
	if err != nil {
		return Player{}, err
	}
	defer resp.Body.Close()
	buf := bytes.Buffer{}
	buf.ReadFrom(resp.Body)
	data := buf.Bytes()
	player := Player{}
	json.Unmarshal(data, &player)
	return player, nil
}

func (api BsApi) GetPlayerStats(tag string) (Player, error) {
	url := "/players/"+url.QueryEscape(tag)
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
