package bsapi

import "net/url"

type BattleResult string
type BrawlerName string

const (
	BattleVictory BattleResult = "victory"
	BattleDefeat  BattleResult = "defeat"
	BattleDraw    BattleResult = "draw"
)

type BattleList struct {
	Items []BattleItem `json:"items"`
}

type BattleItem struct {
	BattleTime string `json:"battleTime"`
	Battle     Battle `json:"battle"`
	Event      Event  `json:"event"`
}

type Battle struct {
	Mode         string         `json:"mode"`
	Type         string         `json:"type"`
	Result       BattleResult   `json:"result"`
	Duration     int            `json:"duration"`
	TrophyChange int            `json:"trophyChange"`
	Teams        [][]TeamMember `json:"teams"`
}

type Event struct {
	Id   int    `json:"id"`
	Mode string `json:"mode"`
	Map  string `json:"map"`
}

type TeamMember struct {
	Tag     string        `json:"tag"`
	Name    string        `json:"name"`
	Brawler BattleBrawler `json:"brawler"`
}

type BattleBrawler struct {
	Id       int         `json:"id"`
	Name     BrawlerName `json:"name"`
	Power    int         `json:"power"`
	Trophies int         `json:"trophies"`
}

func (api BsApi) GetBattleList(playerTag string) (BattleList, error) {
	url := "/players/" + url.QueryEscape(playerTag) + "/battlelog"
	data, err := api.makeRequest(url)
	if err != nil {
		return BattleList{}, err
	}
	battleLog := BattleList{}
	err = json.Unmarshal(data, &battleLog)
	if err != nil {
		return BattleList{}, err
	}
	return battleLog, nil
}
