package util

import (
	"fmt"
	"sort"
)

func Ksort(sortMap map[string]string) string {
	var dataParams string
	var keys []string
	for k := range sortMap {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	//拼接
	for _, k := range keys {
		fmt.Println("key:", k, "Value:", sortMap[k])
		dataParams = dataParams + k + "=" + sortMap[k] + "&"
	}
	fmt.Println(dataParams)
	ff := dataParams[0 : len(dataParams)-1]
	fmt.Println("去掉最后一个&：", ff)

	//对字符串进行sha1哈希
	//h := sha1.New()
	//h.Write([]byte(dataParams))
	//bs := h.Sum(nil)
	//sign := hex.EncodeToString(bs)
	//
	////拼接请求参数
	//dataPost := dataParams + "&sign" + "=" + sign
	return ff
}
