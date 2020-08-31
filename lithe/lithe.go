package lithe

import (
	"bytes"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func PasswordCrypto(str string) string {
	//使用sha256哈希函数
	h := sha256.New()
	h.Write([]byte(str))
	sum := h.Sum(nil)
	s := hex.EncodeToString(sum)
	//fmt.Println(string(s))
	return s
}

func HealthyLogin(card string, pass string) string {
	data := make(map[string]string)
	data["cardNo"] = card
	data["password"] = PasswordCrypto(pass)
	bytesData, _ := json.Marshal(data)

	res, err := http.Post("http://hmgr.sec.lit.edu.cn/wms/healthyLogin",
		"application/json;charset=utf-8", bytes.NewBuffer([]byte(bytesData)))
	if err != nil {
		fmt.Println("Fatal error ", err.Error())
	}

	defer res.Body.Close()

	content, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println("Fatal error ", err.Error())
	}

	//fmt.Println(string(content))
	//str := (*string)(unsafe.Pointer(&content)) //转化为string,优化内存

	//fmt.Println(*str)
	return string(content)
}
