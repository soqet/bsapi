package bsapi

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
	Mode         string       `json:"mode"`
	Type         string       `json:"type"`
	Result       BattleResult `json:"result"`
	Duration     int          `json:"duration"`
	TrophyChange int          `json:"trophyChange"`
	Teams        []Team       `json:"teams"`
}

type Event struct {
	Id   int    `json:"id"`
	Mode string `json:"mode"`
	Map  string `json:"map"`
}

type Team struct {
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
