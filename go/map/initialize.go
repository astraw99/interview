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

type teacher struct {
	Labels map[string]string
	Name   string
	Age    int
}

type people struct {
	Labels map[string]string
	Name   string
}

func main() {
	t := teacher{
		Labels: map[string]string{"abc": "def"},
		Name:   "name",
		Age:    10,
	}

	/*// panic case 1
	t := teacher{
		Labels: nil, // panic: assignment to entry in nil map
		Name:   "",
		Age:    0,
	}*/

	p := people{
		Labels: t.Labels,
		Name:   "",
	}

	/*// panic case 2
	p := people{
		Labels: nil, // panic: assignment to entry in nil map
		Name:   "",
	}*/

	p.Labels["x"] = "y"

	fmt.Println(p)
}
