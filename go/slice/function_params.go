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
	s := []string{"a", "b"}

	res, modified := addFinalizer(s, "c")
	fmt.Println(res, s, modified)

	res, modified = addFinalizerNew(s, "c")
	fmt.Println(res, s, modified)

	res2, modified2 := removeFinalizer(s, "b")
	fmt.Println(res2 == nil, s, modified2)

	res2, modified2 = removeFinalizerNew(s, "b")
	fmt.Println(res2, s, modified2)
}

func addFinalizerNew(finalizers []string, finalizerToAdd string) ([]string, bool) {
	for _, finalizer := range finalizers {
		if finalizer == finalizerToAdd {
			// finalizer already exists
			return finalizers, false
		}
	}

	return append(finalizers, finalizerToAdd), true
}

func removeFinalizerNew(finalizers []string, finalizerToRemove string) ([]string, bool) {
	for i, finalizer := range finalizers {
		if finalizer == finalizerToRemove {
			return append(finalizers[:i], finalizers[i+1:]...), true
		}
	}

	return finalizers, false
}

func addFinalizer(finalizers []string, finalizerToAdd string) ([]string, bool) {
	modifiedFinalizers := make([]string, 0)
	for _, finalizer := range finalizers {
		if finalizer == finalizerToAdd {
			// finalizer already exists
			return finalizers, false
		}
	}
	modifiedFinalizers = append(modifiedFinalizers, finalizers...)
	modifiedFinalizers = append(modifiedFinalizers, finalizerToAdd)
	return modifiedFinalizers, true
}

func removeFinalizer(finalizers []string, finalizerToRemove string) ([]string, bool) {
	modified := false
	modifiedFinalizers := make([]string, 0)
	modifiedFinalizers2 := make([]string, 0)
	for _, finalizer := range finalizers {
		if finalizer != finalizerToRemove {
			modifiedFinalizers = append(modifiedFinalizers, finalizer)
		}
	}

	//modifiedFinalizers2 = nil
	//fmt.Println("is nil:", modifiedFinalizers2 == nil) // false

	if len(modifiedFinalizers) == 0 {
		modifiedFinalizers = nil
	}

	if len(modifiedFinalizers) != len(finalizers) {
		modified = true
	}
	//modifiedFinalizers2 = nil
	return modifiedFinalizers2, modified
}
