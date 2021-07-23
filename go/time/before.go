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

func main() {
	expireDuration := 7 * 24 * time.Hour
	deleteTime, err := time.Parse("2006-01-02 15:04:05", "2021-07-16 16:03:34")
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(deleteTime.Add(expireDuration)) // UTC format

	fmt.Println(time.Now().After(deleteTime.Add(expireDuration)))
}
