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
	"os"
)

func main() {
	// 对外提供的 UDS (Unix Domain Socket) 服务端监听
	if err := os.RemoveAll("/tmp/unix.sock"); err != nil && !os.IsNotExist(err) {
		log.Fatal(err)
	}
	listener, err := net.Listen("unix", "/tmp/unix.sock")
	if err != nil {
		log.Fatal(err)
	}

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Fatal(err)
		}

		go func() {
			defer conn.Close()
			for {
				var buf = make([]byte, 4096)
				n, err := conn.Read(buf)
				log.Println("uds receive:", string(buf), n, err)
				if err != nil {
					//conn.Write([]byte("uds server: " + string(buf[:n]) + "\n"))
					if err == io.EOF {
						conn.Write([]byte("socket 连接已关闭!\n"))
						break
					} else {
						conn.Write([]byte("写入数据失败! error:" + err.Error() + "\n"))
						break
					}
				}
				conn.Write([]byte("uds server: " + string(buf[:n]) + "\n"))
			}
		}()
	}
}
