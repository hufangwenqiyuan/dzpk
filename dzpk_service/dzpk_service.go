package dzpk_service

import (
	"dzpk/model"
	"fmt"
	"giutils/logger"
	"sort"
)

type pkpar struct {
}

func GetPkpar() *pkpar {
	return &pkpar{}
}

type packOfCards []model.Hand
type maxPack model.MaxHandPark

//查看sort源码发现重写下面三个方法
func (p packOfCards) Len() int           { return len(p) }
func (p packOfCards) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }
func (p packOfCards) Less(i, j int) bool { return p[i].OriginalFace < p[j].OriginalFace }

//开始打牌喽
func (*pkpar) BeganToPlay(filePath string) {
	file, err := GetReadFileStruct().ReadFile(filePath)
	if err != nil {
		fmt.Println("ReadFaile is err or filePath is err" + err.Error())
		logger.Info("ReadFaile is err or filePath is err" + err.Error())
	}

	for index, value := range file.Matches {
		if err := compare(value.PlayerA, value.PlayerB); err != nil {
			fmt.Println("err in %d ", index, value.PlayerA, value.PlayerB)
		}
	}
}

//比较两幅牌
func compare(playerA, PlayerB string) (err error) {
	//由于"alice":"AsKhQcJsTc","bob":"As5s6s8sTs"牌是这种格式,需要先解析出来 并且排好序 解析为对象的数组
	strA := analysisString(playerA)
	strB := analysisString(PlayerB)

	//比较比较大的牌型
	go JudgeCardType(&strA)

	go JudgeCardType(&strB)
	return nil
}

//解析手牌
func analysisString(player string) (events packOfCards) {
	//var playerFace uint
	//var brandSubscript string
	//orginal := model.Record{Original:player}
	////有两种可能，5张牌和7张牌
	//if len(player) == 10 {
	//	//先判断牌属于哪个等级
	//	for i := 0;i < len(player);i++{
	//		if i%2 == 0 {
	//			playerFace = Grade[string(player[i])]
	//			//记录下牌面
	//			orginal.OriginalFace[i] = playerFace
	//			continue
	//		}
	//		brandSubscript = string(player[i])
	//		orginal.OriginalColor[i] = brandSubscript
	//	}
	//}
	////下面是七张牌
	//return &orginal
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

func JudgeCardType(site *packOfCards) {

	//判斷是否是皇家从花顺
	if maxPack.isroyalflush(site) {

	} else if maxPack.sequence(site) { ////同花顺

	} else if maxPack.quartic(site) { ////四条

	} else if maxPack.threeZoneS(site) { ////三带二

	} else if maxPack.sameeflower(site) { ////同花

	} else if maxPack.straight(site) { //顺子

	} else if maxPack.three(site) { //三条

	} else if maxPack.twopairs(site) { //两对

	} else if maxPack.twopairs(site) { //一对

	} else if maxPack.sola(site) { //单张大牌

	}
}

func (judge *maxPack) isroyalflush(site *packOfCards) (result bool) {

	return true
}
