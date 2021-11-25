/*
 * Copyright (C) 2021 https://github.com/astraw99. All rights reserved.
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
	var m = map[string]string{}
	//m["key"] = "val"

	for k, v := range m { // range will skip nil map
		fmt.Println(k, v)
	}

	fmt.Println("nil map")
}
