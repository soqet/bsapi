package bsapi

import (
	"context"
	"net/url"
	"strings"
	"time"
)

type BattleResult string

const (
	BattleVictory BattleResult = "victory"
	BattleDefeat  BattleResult = "defeat"
	BattleDraw    BattleResult = "draw"
)

type BattleList struct {
	Items []BattleItem `json:"items"`
}

type BattleItem struct {
	// 	battle time in supercell's format: yyyymmddThhmmss.000Z
	// 	use BattleItem.ParseTime to get time
	BattleTime string `json:"battleTime"`
	Battle     Battle `json:"battle"`
	Event      Event  `json:"event"`
}

type Battle struct {
	Mode   string       `json:"mode"`
	Type   string       `json:"type"`
	Result BattleResult `json:"result"`
	//  in seconds
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
	Id       int    `json:"id"`
	Name     string `json:"name"`
	// power lvl of brawler
	Power    int    `json:"power"`
	Trophies int    `json:"trophies"`
}

func (api BsApi) GetBattleList(ctx context.Context, playerTag string) (BattleList, error) {
	url := "/players/" + url.QueryEscape(playerTag) + "/battlelog"
	data, err := api.makeRequest(ctx, url)
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

func (bi BattleItem) ParseTime() (time.Time, error) {
	stringTime := bi.BattleTime
	var b strings.Builder
	//				01234567890123456789
	// format:		yyyymmddThhmmss.000Z
	//				[  ][][   ][][     ]
	b.WriteString(stringTime[:4])
	b.WriteString("-")
	b.WriteString(stringTime[4:6])
	b.WriteString("-")
	b.WriteString(stringTime[6:11])
	b.WriteString(":")
	b.WriteString(stringTime[11:13])
	b.WriteString(":")
	b.WriteString(stringTime[13:])
	t, err := time.Parse(time.RFC3339, b.String())
	if err != nil {
		return t, err
	}
	return t, nil
}
