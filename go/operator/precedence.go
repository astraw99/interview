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
)

func main() {
	if !test1() || !test2() && test3() {
		fmt.Println("111111111")
		return
	}

	fmt.Println("22222222")
}

func test1() bool {
	fmt.Println("test1")
	return true
}

func test2() bool {
	fmt.Println("test2")
	return true
}

func test3() bool {
	fmt.Println("test3")
	return true
}
