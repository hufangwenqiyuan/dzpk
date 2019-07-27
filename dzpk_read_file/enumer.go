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
var Grade = map[string]uint{
	"A"   : 13,
	"2"   : 1,
	"3"   : 2,
	"4"   : 3,
	"5"   : 4,
	"6"   : 5,
	"7"   : 6,
	"8"   : 7,
	"9"   : 8,
	"10"  : 9,
	"J"   :  10,
	"Q"   : 11,
	"K"   : 12,
}

//记录下需要的信息
type Record struct {
	//原始手牌先保存下来
	original string
	//记录花色
	originalColor [5]string
	//记录下牌面
	originalFace  [5]uint
}
