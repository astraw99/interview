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
	"strings"
)

func main() {
	statusArr := strings.Split("delete,offline", ",")

	var tempArr []string
	for _, v := range statusArr {
		//tempArr = append(tempArr, fmt.Sprintf("%q", v)) // double quote
		tempArr = append(tempArr, "'"+v+"'") // single quote
	}
	inStatus := strings.Join(tempArr, ",")

	fmt.Printf("%v", inStatus)
}