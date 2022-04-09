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

package benchmark

import (
	"testing"
)

var (
	s = []string{"a"}
	//s = []string{"a", "b"}
	//s       = []string{"a", "b", "c"}
	//s       = []string{"a", "b", "c", "d"}
	sCount = 0

	m = map[string]string{"a": "1"}
	//m = map[string]string{"a": "1", "b": "1"}
	//m       = map[string]string{"a": "1", "b": "1", "c": "1"}
	//m       = map[string]string{"a": "1", "b": "1", "c": "1", "d": "1"}
	mCount = 0

	findStr = "a"
)

func BenchmarkWithSlice(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for i, _ := range s {
			if s[i] == findStr {
				sCount++
			}
		}
	}
}

func BenchmarkWithMap(b *testing.B) {
	for i := 0; i < b.N; i++ {
		if _, ok := m[findStr]; ok {
			mCount++
		}
	}
}

func BenchmarkWithMapRange(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, v := range m {
			if v == findStr {
				mCount++
			}
		}
	}
}
