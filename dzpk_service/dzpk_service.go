package dzpk_service

import (
	"dzpk/dzpk_read_file"
	"dzpk/model"
	"fmt"
	"giutils/logger"
)

type pkpar struct {
}

func GetPkpar() *pkpar {
	return &pkpar{}
}

//开始打牌喽
func (*pkpar) BeganToPlay(filePath string) {
	file, err := dzpk_read_file.GetReadFileStruct().ReadFile(filePath)
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
	//由于"alice":"AsKhQcJsTc","bob":"As5s6s8sTs"牌是这种格式,需要先解析出来
	analysisString(playerA)
	analysisString(PlayerB)
}

//解析手牌
func analysisString(player string) *model.Analysis {
	//有两种可能，5张牌和7张牌
	if len(player) == 10 {
		//先判断牌属于哪个等级

	}
	//下面是七张牌
}
