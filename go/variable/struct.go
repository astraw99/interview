/*
 * Copyright (C) 2025 https://github.com/astraw99. All rights reserved.
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
	"encoding/json"
	"fmt"
)

var (
	stu  *student
	stu2 *student
)

type student struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func main() {
	data := &student{
		Name: "abc",
		Age:  20,
	}

	stu = data
	fmt.Println("stu:", stu)

	// json to struct
	stuVal := `{"name":"abc2","age":30}`
	stu2 = &student{}
	if err := json.Unmarshal([]byte(stuVal), stu2); err != nil {
		panic(err)
	}
	fmt.Println("stu2:", stu2)
}
