/*
 * Copyright (C) 2022 https://github.com/astraw99. All rights reserved.
 *
 * Licensed under the BSD 3-Clause License (the "License"); you may not use this
 * file except in compliance with the License. You may obtain a copy of the
 * License at
 *
 * https://opensource.org/licenses/BSD-3-Clause
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS, WITHOUT
 * WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied. See the
 * License for the specific language governing permissions and limitations under
 * the License.
 */

package main

import (
	"crypto/sha256"
	"fmt"
)

func main() {
	message := []byte("hello world")
	hashCode := GetSHA256HashCode(message)
	fmt.Println(hashCode)
}

func GetSHA256HashCode(message []byte) string {
	// 方法一：
	// 创建一个基于SHA256算法的hash.Hash接口的对象
	/*hash := sha256.New()
	// 输入数据
	hash.Write(message)
	// 计算哈希值
	bytes := hash.Sum(nil)
	// 将字符串编码为16进制格式,返回字符串
	hashCode := hex.EncodeToString(bytes)
	// 返回哈希值
	return hashCode*/

	// 方法二：
	// 计算哈希值，返回一个长度为32的数组
	//bytes2:=sha256.Sum256(message)
	// 将数组转换成切片，转换成16进制，返回字符串
	//hashCode2:=hex.EncodeToString(bytes2[:])
	//return hashCode2

	// 方法三：
	// 计算哈希值，返回一个长度为32的数组
	bytes3 := sha256.Sum256(message)
	println(bytes3[:])
	println(string(bytes3[:]))
	return fmt.Sprintf("%x\n", bytes3) // b94d27b9934d3e08a52e52d7da7dabfac484efe37a5380ee9088f7ace2efcde9
}
