/*
 * Copyright (C) 2021 THL A29 Limited, a Tencent company. All rights reserved.
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
	"fmt"
	"math/rand"
	"os"
	"time"
)

//var logExitFunc func(error)

func main() {
	test()
}

func test() {
	defer fmt.Println("defer called1")
	defer fmt.Println("defer called2")

	fmt.Println("running")

	//os.Exit(1) // exit will not call defer!!

	defer func() {
		if err := recover(); err != nil {
			fmt.Println("recovered:", err)
		}
	}()

	//panic("test panic")

	test2()

	fmt.Println("stopped")
}

func test2() {
	fmt.Println("test2 exit")

	/*logExitFunc = func(error) {}
	fmt.Println(logExitFunc)

	if logExitFunc != nil {
		return
	}*/

	rand.Seed(time.Now().Unix())
	if rand.Intn(100)%2 == 0 {
		return
	}

	os.Exit(2)
}
