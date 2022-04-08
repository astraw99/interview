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

func main() {
	var ann map[string]string
	//var ann = map[string]string{"key": "val"}
	fmt.Println(HaveMapKey(ann), ann, len(ann))
}

func HaveMapKey(annotations map[string]string) bool {
	if _, ok := annotations["key"]; ok {
		return true
	}

	return false
}
