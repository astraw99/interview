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

import (
	"fmt"
	"time"
)

func main() {
	expireDuration := 7 * 24 * time.Hour
	loc, err := time.LoadLocation("Local")
	if err != nil {
		panic(err)
	}
	fmt.Println(loc)

	deleteTimeUTC, err := time.Parse("2006-01-02 15:04:05", "2021-07-20 16:00:34")
	if err != nil {
		panic(err)
	}
	fmt.Println(deleteTimeUTC) // UTC format

	deleteTime, err := time.ParseInLocation("2006-01-02 15:04:05", "2021-07-20 16:00:34", time.Local)
	if err != nil {
		panic(err)
	}

	fmt.Println(deleteTime) // CST format
	fmt.Println(time.Now()) // CST format

	fmt.Println(time.Now().After(deleteTime.Add(expireDuration)))
}
