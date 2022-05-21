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
	"errors"
	"fmt"
)

type myStruct struct {
	name string
	age  int
}

type nestStruct struct {
	nested *myStruct
	sex    string
}

func main() {
	s, err := test("Alice", 20)
	fmt.Println(s, err)

	s2, err := testPointer("Bob", 30)
	fmt.Println(s2, err)

	s3, err := testNestPointer("David", 40)
	fmt.Println(s3, s3.nested, err)
}

func test(name string, age int) (s myStruct, err error) {
	//s = myStruct{} // no need to init struct
	s.name = name
	s.age = age

	err = errors.New("no err")

	return s, err
}

func testPointer(name string, age int) (s *myStruct, err error) {
	s = &myStruct{} // MUST init pointer first, or will panic!!
	s.name = name
	s.age = age

	err = errors.New("no err")

	return s, err
}

func testNestPointer(name string, age int) (s *nestStruct, err error) {
	// MUST init pointer first, or will panic!!
	s = &nestStruct{
		nested: &myStruct{}, // MUST init nested pointer first, or will panic!!
	}
	s.nested.name = name
	s.nested.age = age

	err = errors.New("no err")

	return s, err
}
