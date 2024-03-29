package dzpk_service

import (
	"dzpk/model"
	"fmt"
	"sort"
	"time"
)

type pkpar struct {
	//牌的等级
	parkGrade int
	//需要的排序数据
	handPark []model.Hand
}

func GetPkpar() *pkpar {
	return &pkpar{}
}

type packOfCards []model.Hand

//查看sort源码发现重写下面三个方法
func (p packOfCards) Len() int           { return len(p) }
func (p packOfCards) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }
func (p packOfCards) Less(i, j int) bool { return p[i].OriginalFace < p[j].OriginalFace }

//开始打牌喽
func (*pkpar) BeganToPlay(filePath string) {

	Path := filePath + "/match.json"
	file, err := GetReadFileStruct().ReadFile(Path)
	if err != nil {
		fmt.Println("ReadFaile is err or filePath is err" + err.Error())
	}
	start := time.Now()
	for _, value := range file.Matches {
		strA := analysisString(value.PlayerA).JudgeCardType()
		strB := analysisString(value.PlayerB).JudgeCardType()

		if strA.parkGrade > strB.parkGrade {
			value.Result = 1
			//A赢了
		} else if strA.parkGrade == strB.parkGrade {
			//同一种牌 //如果是同一种牌就递归比较
			for i := len(strA.handPark) - 1; ; i-- { //获取改数组的最后的一个开始
				if strA.handPark[i].OriginalFace > strB.handPark[i].OriginalFace {
					value.Result = 1
					break
					//A赢
				} else if strA.handPark[i].OriginalFace == strB.handPark[i].OriginalFace {
					if i == 0 {
						//平局
						value.Result = 0
						break
					}
					continue //继续

				} else {
					//B赢
					value.Result = 2
					break
				}

			}
		} else {
			//B贏了
			value.Result = 2
		}
	}
	fmt.Println(time.Since(start))
	filePath = filePath + "/result.json"
	//把file写入result文件中
	if fileErr := GetReadFileStruct().WhirteFile(filePath, file); fileErr != nil {
		fmt.Println(fileErr.Error())
	}
}

//解析手牌
func analysisString(player string) (events packOfCards) {
	var playerFace int
	var brandSubscript string
	for i := 0; i < len(player); i = i + 2 {
		playerFace = Grade[string(player[i])]
		brandSubscript = string(player[i+1])
		pack := model.Hand{
			OriginalFace:  playerFace,
			OriginalColor: brandSubscript,
		}
		events = append(events, pack)
	}
	sort.Sort(events) //对牌进行排序
	return events
}

//获取手牌的牌型和大小
func (p packOfCards) JudgeCardType() pkpar {
	pkpar := pkpar{}
	pkpar.getInfomation(p)
	return pkpar
}

//真正获取牌型和牌数据  尽可能减少循环
func (p *pkpar) getInfomation(site packOfCards) (result bool) {
	var record, endRecordI, pairEtc, threeZone, recoOne, isSurpllus int //状态标记，是否是顺子
	//record       //记录下相同牌出现的次数
	//endRecordI            //在出现三张连续牌的时候，记录下牌的下标
	//pairEtc                    //记录下一对一对出现的次数
	//threeZone               //记录下三张牌出现的次数
	//isSurpllus int 	//状态标记，是否是顺子
	//recoOne
	if (site[len(site)-1].OriginalFace - site[0].OriginalFace) == len(site)-1 { //顺子
		isSurpllus = 1
		if site[len(site)-1].OriginalFace == 13 { //如果是顺子，并且最后一张牌是13（A）则表示是皇家
			if p.sameeflower(site) { //判断其是否是同花顺 返回true 表示是皇家同花顺
				p.parkGrade = ROYALFLUSH //如果是皇家同花顺则不用比较
				return true
			}

		} else {
			if p.sameeflower(site) { //这里判断出 是同花顺不是皇家同花顺
				p.parkGrade = SEQUENCE //如果是顺则只需要比较相同坐标下的值
				return true
			}
		}
	}

	docker := site[0].OriginalFace //先初始化第一张牌的牌面
	for i := 1; i < len(site); i++ {
		if docker == site[i].OriginalFace {
			endRecordI = i
			record++
			if record == 3 {
				threeZone--
			} else if record == 2 {
				threeZone++
			} else if record == 1 {
				pairEtc++
			}
			if pairEtc == 1 {
				recoOne = i
				isSurpllus = 0
			}
		} else {
			docker = site[i].OriginalFace //如果两张牌不相等，则用后面的牌放入容器中，用后面的牌继续做比价
			record = 0
		}
	}
	if record == 3 { //判断四条出现的情况  下面的判断顺序是这个程序的关键
		p.handPark = append(site[:endRecordI-3], site[endRecordI:]...)
		p.parkGrade = QUARTIC
		return true
	} else if threeZone == 1 { //满堂彩 三带二
		spilc := make([]model.Hand, 2) //创建一个容量为2的切片
		if threeZone == 1 {
			startWhere := endRecordI - 2
			spilc = append(site[:startWhere], site[endRecordI+1:]...) //采用数组切片的方式获取剩下的值
		}
		if spilc[0].OriginalFace == spilc[1].OriginalFace {
			p.handPark = append(append(p.handPark, spilc[0]), site[endRecordI])
			p.parkGrade = THREE_ZONES
			return true
		} else {
			p.handPark = append(append(p.handPark, spilc...), site[endRecordI])
			p.parkGrade = THREE
			return true
		}

	} else if p.sameeflower(site) { //同花
		p.handPark = site //同花的情况则需要把所有的值加上去
		p.parkGrade = SAMEFLOWER
		return true
	} else if isSurpllus == 1 { //顺子
		p.handPark = append(p.handPark, site[0])
		p.parkGrade = STRAIGHT
		return true
	} else if pairEtc == 2 { //两对
		spilct := append(site[:endRecordI], site[endRecordI+1:]...)
		p.handPark = append(spilct[:recoOne], spilct[recoOne+1:]...)
		p.parkGrade = TWOPAIRS
		return true
	} else if pairEtc == 1 { //一对
		p.handPark = append(site[:endRecordI], site[endRecordI+1:]...)
		p.parkGrade = TWAIN
		return true
	} else { //如果上面都不符合则表示是散牌
		p.parkGrade = SOLA
		p.handPark = site
		return true
	}
}

//判断是否是同花
func (p *pkpar) sameeflower(site packOfCards) (result bool) {
	docker := site[0].OriginalColor  //获取第一张牌的花色
	recode := 0                      //记录出现的次数
	for i := 1; i < len(site); i++ { //遍历剩下的花色
		if docker == site[i].OriginalColor {
			recode++ //因为同花是所有的牌是一种花色，如果出现相等则加一
		}
	}
	if recode == len(site)-1 { //如果全部都是同一种花色则返回true
		return true
	}
	return false
}

func imputVoltage(pkpA, pkpB string) {
	//fmt.Println("请输入你要参加者A的牌号：")
	//var pkpA string
	//var pkpB string
	//fmt.Scan("%s",&pkpA)
	//fmt.Println("请输入你要参加者B的牌号：")
	//fmt.Scan("%s",&pkpB)
	strA := analysisString(pkpA).JudgeCardType()
	strB := analysisString(pkpB).JudgeCardType()

	if strA.parkGrade > strB.parkGrade {
		fmt.Println("A同学赢了：恭喜A同学")
		//A赢了
	} else if strA.parkGrade == strB.parkGrade {
		//同一种牌 //如果是同一种牌就递归比较
		for i := len(strA.handPark) - 1; ; i-- { //获取改数组的最后的一个开始
			if strA.handPark[i].OriginalFace > strB.handPark[i].OriginalFace {
				fmt.Println("A同学赢了：恭喜A同学")
				break
				//A赢
			} else if strA.handPark[i].OriginalFace == strB.handPark[i].OriginalFace {
				if i == 0 {
					//平局
					fmt.Println("平局：厉害打了平局")
					break
				}
				continue //继续

			} else {
				//B赢
				fmt.Println("B同学赢了：恭喜B同学")
				break
			}

		}
	} else {
		//B贏了
		fmt.Println("B同学赢了：恭喜B同学")
	}
}
