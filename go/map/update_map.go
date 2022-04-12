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

import "fmt"

type Attribute struct {
	Name string
	Age  int
}

type Teacher struct {
	//TeacherMap map[string]*Attribute // pointer value will be updated synchronously
	TeacherMap map[string]Attribute // non-pointer value will not be updated synchronously
}

type People struct {
	Pointer    *Teacher
	NonPointer Teacher
}

func main() {

	p := People{
		Pointer: &Teacher{
			map[string]Attribute{
				"test": {
					Name: "foo",
					Age:  20,
				},
			},
		},
		NonPointer: Teacher{
			map[string]Attribute{
				"test2": {
					Name: "foo2",
					Age:  22,
				},
			},
		},
	}

	fmt.Println(p, "============")

	update := p.Pointer.TeacherMap["test"]
	update2 := p.NonPointer.TeacherMap["test2"]

	(&update).Name = "bar"
	(&update).Age = 30

	(&update2).Name = "bar2"
	(&update2).Age = 32

	fmt.Println(p.Pointer.TeacherMap["test"], update, p.NonPointer.TeacherMap["test2"], update2)
}
