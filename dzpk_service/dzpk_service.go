package dzpk_service

import (
	"dzpk/model"
	"fmt"
	"sort"
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

	for _, value := range file.Matches {
		strA := analysisString(value.PlayerA).JudgeCardType()
		strB := analysisString(value.PlayerB).JudgeCardType()

		if strA.parkGrade > strB.parkGrade {
			value.Result = 1
			//A赢了
		} else if strA.parkGrade == strB.parkGrade {
			//同一种牌 //如果是同一种牌就递归比较
			for i := len(strA.handPark) - 1; i < len(strA.handPark); i++ { //获取改数组的最后的一个开始
				if strA.handPark[i].OriginalFace > strB.handPark[i].OriginalFace {
					value.Result = 1
					//A赢
				} else if strA.handPark[i].OriginalFace == strB.handPark[i].OriginalFace {
					if i == 0 {
						//平局
						value.Result = 0
					}
					continue //继续

				} else {
					//B赢
					value.Result = 2
				}

			}
		} else {
			//B贏了
			value.Result = 2
		}
	}
	filePath = filePath + "/result.json"
	//把file写入result文件中
	GetReadFileStruct().WhirteFile(filePath, file)
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
	sort.Sort(events)

	return events
}

//获取手牌的牌型和大小
func (p packOfCards) JudgeCardType() pkpar {
	pkpar := pkpar{}
	pkpar.threeZoneS(p)
	return pkpar
}

//真正获取牌型和牌数据
func (p *pkpar) threeZoneS(site packOfCards) (result bool) {
	isSurpllus := 0                                                             //状态标记，是否是顺子
	if (site[len(site)-1].OriginalFace - site[0].OriginalFace) == len(site)-1 { //顺子
		isSurpllus++
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
	record := 0                    //记录下相同牌出现的次数
	endRecordI := 0                //在出现三张连续牌的时候，记录下牌的下标
	pairEtc := 0                   //记录下一对一对出现的次数
	threeZone := 0                 //记录下三张牌出现的次数
	recoOne := 0
	for i := 1; i < len(site); i++ {
		if docker == site[i].OriginalFace {
			record++
			if record == 3 {
				endRecordI = i
			}
			if record == 2 {
				endRecordI = i
				threeZone++
			}
			if record == 1 {
				endRecordI = i
				pairEtc++
			}

			if pairEtc == 1 {
				recoOne = i
			}
		} else {
			docker = site[i].OriginalFace //如果两张牌不相等，则用后面的牌放入容器中，用后面的牌继续做比价
			record = 0
		}
	}
	spilc := make([]model.Hand, 2) //创建一个容量为2的切片
	if threeZone == 1 {
		startWhere := endRecordI - 2
		spilc = append(site[:startWhere], site[endRecordI+1:]...) //采用数组切片的方式获取剩下的值
	}

	if record == 3 { //判断四条出现的情况  下面的判断顺序是这个程序的关键
		slice := make([]model.Hand, 2)
		slice = append(site[:endRecordI-3], site[endRecordI+1:]...)
		s := append(slice, site[endRecordI])
		p.handPark = s
		p.parkGrade = QUARTIC
		return true
	} else if threeZone == 1 && spilc[0] == spilc[1] { //满堂彩 三带二
		p.handPark = append(append(p.handPark, spilc[0]), site[endRecordI])
		p.parkGrade = THREE_ZONES
		return true
	} else if p.sameeflower(site) { //同花
		p.handPark = site //同花的情况则需要把所有的值加上去
		p.parkGrade = SEQUENCE
		return true
	} else if isSurpllus == 1 { //顺子
		p.handPark[0] = site[0]
		p.parkGrade = STRAIGHT
		return true
	} else if threeZone == 1 { //三带二
		p.handPark = append(append(p.handPark, spilc...), site[endRecordI])
		p.parkGrade = THREE
		return true
	} else if pairEtc == 2 { //两对
		spilc = append(site[:endRecordI], site[endRecordI+1:]...)
		p.handPark = append(spilc[:recoOne], spilc[recoOne+1:]...)
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
		//p.handPark[0] = site[0] //这里记录下那个值
		return true
	}
	return false
}
