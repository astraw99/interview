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
	"io"
	"log"
	"net"
	"time"
)

func main() {
	conn, err := net.DialTimeout("tcp", ":8080", 5*time.Second)
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	go func() {
		for {
			var buf = make([]byte, 4096)
			// 从服务器接收数据
			_, err := conn.Read(buf)
			//log.Println("client receive:", string(buf))
			if err != nil {
				if err == io.EOF {
					conn.Write([]byte("socket 连接已关闭!\n"))
					break
				} else {
					conn.Write([]byte("写入数据失败! error:" + err.Error() + "\n"))
					break
				}
			}
			log.Println("client receive:", string(buf))
		}
	}()

	for i := 0; i < 5; i++ {
		// 发送数据给服务器端
		conn.Write([]byte("hello socket\n"))
		time.Sleep(1 * time.Second)
	}
}
