package dzpk_service

import (
	"testing"
)

//测试程序，文件件会写入match_template文件夹下面
func Test_dzpk_test(t *testing.T) {
	//获取时间戳
	GetPkpar().BeganToPlay("./../match_template")
}
