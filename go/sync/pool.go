/*
 * Copyright (C) 2024 THL A29 Limited, a Tencent company. All rights reserved.
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
	"runtime/debug"
	"sync"
	"sync/atomic"
)

var (
	count    int32
	initFunc = func() interface{} {
		return atomic.AddInt32(&count, 1)
	}

	pool = sync.Pool{New: initFunc}
)

func main() {
	debug.SetGCPercent(debug.SetGCPercent(-1))

	v1 := pool.Get()
	fmt.Printf("value 1: %v\n", v1)
	pool.Put(10)
	pool.Put(11)
	pool.Put(12)
	pool.Put(13)
	v2 := pool.Get()
	fmt.Printf("value 2: %v\n", v2)

	// enable GC will change the result?
	//runtime.GC()

	v3 := pool.Get()
	fmt.Printf("value 3: %v\n", v3)
	pool.New = nil
	v4 := pool.Get()
	fmt.Printf("value 4: %v\n", v4)

	v5 := pool.Get()
	fmt.Printf("value 5: %v\n", v5)
}

/*// 书中的输出结果
value 1: 1
value 2: 10
value 3: 2
value 4: <nil>

// 我实际的输出结果
value 1: 1
value 2: 10
value 3: 11
value 4: 12*/
