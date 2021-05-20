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
	"bytes"
	"encoding/json"
	"fmt"
)

func main() {
	unmarshal() // 会转为科学计数法: 8.589934592e+09
	decode()    // ok
}

func unmarshal() {
	body := `{"message": "success", "data": 8589934592}`
	v := make(map[string]interface{})

	// To unmarshal JSON into an interface value, Unmarshal stores one of these in the interface value:
	// float64, for JSON numbers
	err := json.Unmarshal([]byte(body), &v)
	fmt.Println(v, err)
}

func decode() {
	body := `{"message": "success", "data": 8589934592}`
	v := make(map[string]interface{})

	d := json.NewDecoder(bytes.NewReader([]byte(body)))
	d.UseNumber()
	err := d.Decode(&v)
	fmt.Println(v, err)
}
