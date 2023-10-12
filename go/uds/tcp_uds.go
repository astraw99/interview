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
	// 对外提供的 TCP 服务端监听
	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatal(err)
	}

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Fatal(err)
		}

		// 处理每一个连接
		go handleConnection(conn)
	}
}

func handleConnection(conn net.Conn) {
	defer conn.Close()

	udsConn, err := net.DialTimeout("unix", "/tmp/unix.sock", 5*time.Second)
	if err != nil {
		log.Fatal(err)
	}
	defer udsConn.Close()

	// 将目标 UDS 输出流转发到来源 TCP 输出流
	go io.Copy(conn, udsConn)

	// 将来源 TCP 输入流转发到目标 UDS 输入流
	io.Copy(udsConn, conn)
}
