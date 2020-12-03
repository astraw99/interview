package main

import "C"
import "fmt"

func main()  {
	/*A. str:='abc'+'123'
	B. str:="abc"+"123"
	C. str:='123'+"abc"
	D. str:=`123`+`abc`
	E. str:=`123`+'abc'
	F. str:=`123`+"abc"*/

	// BDF are right.

	//====================================


	str := "abc"
	//+ "def" // 拼接符 + 只能放上一行末尾

	str2 := `123` + `xyz` // 反引号：不进行任何转义的字符串
	type People struct {
		Name string `json:"name"`
		Age int `json:"age"`
	}

	str3 := `test` + "100"

	fmt.Println(str, str2, str3)
}
