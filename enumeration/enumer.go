package enumeration

//给牌进行分类，从大到小赋值
const(
	//皇家从花顺
	ROYALFLUSH = 10
	//同花顺
	SEQUENCE   = 9
	//四条
	QUARTIC    = 8
	//三带二
    THREE_ZONES = 7
	//同花
	SAMEFLOWER  = 6
	//顺子
	STRAIGHT    = 5
	//三条
	THREE       = 4
	//两对
	TWOPAIRS    = 3
	//一对
	TWAIN       = 2
	//单张大牌
	SOLA        = 1
)

//创建一个map key为输入的值  value为对应的大小

																																																																																var Cardsize = make(map[string]uint){

}