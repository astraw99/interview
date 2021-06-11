package main

import (
	"fmt"

	"github.com/shopspring/decimal"
)

func main() {
	numStr := "8.589934592e+09"
	decimalNum, err := decimal.NewFromString(numStr)
	if err != nil {
		fmt.Printf("decimal.NewFromString error, numStr:%s, err:%v", numStr, err)
		return
	}

	fmt.Println(decimalNum.IntPart())
	fmt.Println(decimalNum.String())
}
