package dzpk_read_file

import (
	"dzpk/model"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
)

//在go中结构体就是类
type readFile struct{}

//提供统一的获取类的方式
func GetReadFileStruct() *readFile {
	return &readFile{}
}

//根据提供的的地址加载json数据
func (*readFile) ReadFile(filePath string) (Matches *model.Matches, err error) {
	if len(filePath) < 0 {
		return nil, nil
	}
	file, err := ioutil.ReadFile(filePath)
	if err != nil {
		fmt.Println("log:filePath is error" + err.Error())
		return nil, errors.New(err.Error())
	}

	//获取了字符串，把字符串转为对象
	if err := json.Unmarshal(file, &Matches); err != nil {
		fmt.Println("log: readfile unmasrshel err" + err.Error())
		return nil, errors.New(err.Error())
	}

	return Matches, nil
}
