package main

import (
	"fmt"
	"io"
	"strings"
)

func main() {
	var reader1 strings.Reader
	fmt.Printf("reader size %d, reader len %d\n", reader1.Size(), reader1.Len())

	offset2 := int64(17)
	// Size() 字符串总长度；Len() 未读部分剩余长度
	expectedIndex := reader1.Size() - int64(reader1.Len()) + offset2
	fmt.Printf("Seek with offset %d and whence %d ...\n", offset2, io.SeekCurrent)

	// Seek() 设置下一次读取起始索引值
	readingIndex, _ := reader1.Seek(offset2, io.SeekCurrent)
	fmt.Printf("The reading index in reader: %d (returned by Seek)\n", readingIndex)
	fmt.Printf("The reading index in reader: %d (computed by me)\n", expectedIndex)
}
