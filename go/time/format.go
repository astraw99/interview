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
	"time"
)

func main()  {
	day := time.Now().Weekday().String()
	fmt.Printf("%T\n%v\n", day, day)

	fmt.Println(time.Now().Format("15:04"))

	fmt.Println(time.Now().Format("15:04") >= "03:07")

	t, err := time.Parse("2006-01-02 15:04:05", time.Now().Format("2006-01-02")+" 11:20"+":00")

	fmt.Println(t, err)
	fmt.Println(t.Add(time.Duration(200 * 1e9)))

	fmt.Println(time.Now().Add(10 * time.Hour))

	fmt.Println(time.Now().Unix())
}
