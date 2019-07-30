package dzpk_service

//给牌进行分类，从大到小赋值
const (
	//单张大牌
	SOLA        = iota
	TWAIN       //一对
	TWOPAIRS    //两对
	THREE       //三条
	STRAIGHT    //顺子
	SAMEFLOWER  //同花
	THREE_ZONES //三带二
	QUARTIC     //四条
	SEQUENCE    //同花顺
	ROYALFLUSH  //皇家从花顺
)

//不能使用枚举？？？？？？？？  行吧使用map
//const(
//   s  = iota
//	"3"
//	"4"
//	"5"
//	"6"
//	"7"
//	"8"
//	"9"
//	"T"
//	"J"
//	"Q"
//	"K"
//	"A"
//)
//创建一个map key为输入的值  value为对应的大小
var Grade = map[string]int{
	"2": 1,
	"3": 2,
	"4": 3,
	"5": 4,
	"6": 5,
	"7": 6,
	"8": 7,
	"9": 8,
	"T": 9,
	"J": 10,
	"Q": 11,
	"K": 12,
	"A": 13,
}
