package main

import (
	"fmt"
	"strconv"
)

func main()  {
	code := 335000043
	status_code := ""
	status_code = "ES" + strconv.Itoa(code)

	// 下单改地址错误号处理：业务错误不报警，如 ES335000043
	if code > 4000000 {
		strCode := fmt.Sprint(code)
		if strCode[:2] == "33" {
			status_code = "WW" + strconv.Itoa(code)
		}
	}

	fmt.Println(status_code)
}
