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

	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/resource"
)

func main() {
	requests := corev1.ResourceList{
		corev1.ResourceCPU:    resource.MustParse("1"),
		corev1.ResourceMemory: resource.MustParse("1Gi"),
	}

	requests2 := corev1.ResourceList{
		corev1.ResourceCPU:    resource.MustParse("100.1m"),
		corev1.ResourceMemory: resource.MustParse("0.1Gi"),
	}

	fmt.Println(requests.Cpu().MilliValue())  // Cpu.Value is decimal = 1000
	fmt.Println(requests2.Cpu().MilliValue()) // Cpu.Value is decimal = 100
	fmt.Println(requests.Cpu().Value())       // Cpu.Value is decimal = 1
	fmt.Println(requests2.Cpu().Value())      // Cpu.Value is decimal = 1

	fmt.Println(requests.Memory().Value())                      // Memory.Value is byte = 1073741824
	fmt.Println(requests2.Memory().Value())                     // Memory.Value is byte = 107374183
	fmt.Println(requests.Memory().Value() / 1024 / 1024 / 1024) // = 1G
}
