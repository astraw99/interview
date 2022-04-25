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

type System struct {
	//Sidecars map[string]Container `json:"sidecars,omitempty" protobuf:"bytes,1,rep,name=sidecars"` // non-pointer will copy a new Container
	Sidecars map[string]*Container `json:"sidecars,omitempty" protobuf:"bytes,1,rep,name=sidecars"` // pointer will use the same  Container
}

type Container struct {
	Image   string            `json:"image,omitempty" protobuf:"bytes,1,opt,name=image"`
	Command []string          `json:"command,omitempty" protobuf:"bytes,2,rep,name=command"`
	Args    []string          `json:"args,omitempty" protobuf:"bytes,3,rep,name=args"`
	Envs    map[string]string `json:"envs,omitempty" protobuf:"bytes,4,rep,name=envs"`
}

func main() {
	target := System{
		Sidecars: map[string]*Container{
			"a": {
				//Image:   "image - target",
				Command: []string{"command", "target"},
			},
			"b": {
				Command: []string{"command", "target"},
			},
		},
	}

	preset := System{
		Sidecars: map[string]*Container{
			"a": {
				Image: "image - preset",
			},
			"b": {
				Command: []string{"command", "preset"},
			},
		},
	}

	for name := range preset.Sidecars {
		if len(target.Sidecars) == 0 {
			target.Sidecars = make(map[string]*Container)
		}
		targetContainer, ok1 := target.Sidecars[name]
		presetContainer, ok2 := preset.Sidecars[name]
		if ok1 && ok2 {
			populateContainerPreset(targetContainer, presetContainer)
			fmt.Println(targetContainer, "-----", target.Sidecars[name])
			fmt.Printf("%v=====\n", target.Sidecars["a"])
		} else if ok2 {
			target.Sidecars[name] = presetContainer
		}
	}

	fmt.Printf("%v", target)
}

func populateContainerPreset(target, preset *Container) {
	if target.Image == "" {
		target.Image = preset.Image
	}
	if len(target.Command) == 0 {
		target.Command = preset.Command
	}
	if len(target.Args) == 0 {
		target.Args = preset.Args
	}

	// merge environments map
	populateStringMap(target.Envs, preset.Envs)
}

func populateStringMap(to, from map[string]string) {
	for key, value := range from {
		if _, ok := to[key]; !ok {
			to[key] = value
		}
	}
}
