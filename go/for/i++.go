/*
 * Copyright (C) 2024 https://github.com/astraw99. All rights reserved.
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

import "fmt"

func main() {
	count := 0
	for i := 0; i < 3; i++ {
		println("before =", i)
		i--
		println("after =", i)

		count++
		if count == 3 {
			break
		}
	}

	RemoveFinalizer([]string{"b", "b", "c"}, "b")
}

func RemoveFinalizer(all []string, e string) {
	for i := 0; i < len(all); i++ {
		println("start =", i)
		if all[i] == e {
			println("before =", i)
			all = append(all[:i], all[i+1:]...)
			fmt.Println(all)
			i--
			println("after =", i)
		}
	}
	fmt.Println(all)
}
