package model

//10000条数据的结构体
type Match struct {
	PlayerA string `json:"alice" name:"游戏者A"`
	PlayerB string `json:"bob" name:"游戏者B"`
	Result  int    `json:"result" name:"比较结果"`
}

type Matches struct {
	Matches []*Match `json:"matches"`
}

//組成的牌型
type Hand struct {
	OriginalColor string
	OriginalFace  int
}
