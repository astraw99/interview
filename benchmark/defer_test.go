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

package benchmark

import (
	"sync"
	"testing"
)

type structWithLock struct {
	m sync.Mutex
	n int64
}

func BenchmarkWithoutDeferUnLock(b *testing.B) {
	s := structWithLock{}
	for i := 0; i < b.N; i++ {
		s.addWithoutDefer()
	}
}

func BenchmarkWithDeferUnLock(b *testing.B) {
	s := structWithLock{}
	for i := 0; i < b.N; i++ {
		s.addWithDefer()
	}
}

func (s *structWithLock) addWithoutDefer() {
	s.m.Lock()
	s.n++
	s.m.Unlock()
}

func (s *structWithLock) addWithDefer() {
	s.m.Lock()
	defer s.m.Unlock()
	s.n++
}
