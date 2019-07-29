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

//记录下需要的信息
type Record struct {
	//原始手牌先保存下来
	Original string
	//记录花色
	OriginalColor [5]string
	//记录下牌面
	OriginalFace [5]uint
}

//組成的牌型
type Hand struct {
	OriginalColor string
	OriginalFace  int
}

