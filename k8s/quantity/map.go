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

import (
	"fmt"

	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/resource"
)

func main() {

	var (
		capacity    resource.Quantity
		resourceReq corev1.ResourceRequirements
	)

	resourceReq.Requests = make(corev1.ResourceList) // need to make map first, or will panic!!

	resourceReq.Requests[corev1.ResourceCPU] = resource.MustParse("100m")

	fmt.Printf("%v\n", resourceReq)
	fmt.Printf("%v\n", capacity)
	fmt.Printf("%v\n", &capacity)
}
