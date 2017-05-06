package base

//方法执行时，向response中写入map包装的数据，包括result=true，和返回的数据信息。
import (
	"encoding/json"
	"math/rand"
	"net/http"
)

func RightJsonToRsp(rw http.ResponseWriter, resultmap map[string]interface{}) {
	resultmap["result"] = true
	b, _ := json.Marshal(resultmap)
	rw.Write(b)
}

//发生错误时，向response中写入map包装的数据，包括result=false，和错误信息msg。
func WrongJsonToRsp(rw http.ResponseWriter, msg string) {
	resultmap := make(map[string]interface{}, 2)
	resultmap["result"] = false
	resultmap["msg"] = msg
	b, _ := json.Marshal(resultmap)
	rw.Write(b)
}

//发生错误时，向response中写入map包装的数据，包括result=false，和错误信息msg。
func WrongCodeJsonToRsp(rw http.ResponseWriter, msg string, errcode int) {
	resultmap := make(map[string]interface{}, 3)
	resultmap["result"] = false
	resultmap["msg"] = msg
	resultmap["errcode"] = errcode
	b, _ := json.Marshal(resultmap)
	rw.Write(b)
}

//产生随机字符串
func GenRandomStr(base int, length int) (ranstr string) {
	base1 := "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz1234567890"
	base2 := "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	base3 := "abcdefghijklmnopqrstuvwxyz"
	base4 := "ABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890"
	base5 := "1234567890"
	var basestr string
	switch base {
	case 1:
		basestr = base1
	case 2:
		basestr = base2
	case 3:
		basestr = base3
	case 4:
		basestr = base4
	case 5:
		basestr = base5
	}
	byts := []byte(basestr)
	bytslen := len(byts)
	ranbytes := make([]byte, length)
	for i := 0; i < length; i++ {
		ran := rand.Intn(bytslen)
		ranbytes[i] = byts[ran]
	}
	return string(ranbytes)
}

func Getformvalue(formkey string, r *http.Request) (value string) {
	defer func() {
		if err := recover(); err != nil {
			value = ""
		}
	}()
	value = r.Form[formkey][0]
	return
}
