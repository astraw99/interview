package main

import (
	"fmt"
	"regexp"
)

func main() {
	str := "delete *"
	fmt.Println(IsSQLInject(str))
}

// IsSQLInject 正则过滤sql注入
func IsSQLInject(str string) bool {
	// 过滤 ‘
	// ORACLE 注解 --  /**/
	// 关键字过滤 update, delete
	// 正则的字符串, 不能用 " " 因为" "里面的内容会转义
	regStr := `(?:')|(?:--)|(/\\*(?:.|[\\n\\r])*?\\*/)|(\b(select|update|and|or|delete|insert|trancate|char|chr|into|substr|ascii|declare|exec|count|master|into|drop|execute)\b)`
	re, err := regexp.Compile(regStr)
	if err != nil {
		panic(err.Error())
		return false
	}
	return re.MatchString(str)
}
