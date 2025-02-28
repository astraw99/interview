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
	"time"
)

type ClusterSnapshot struct {
	Status ClusterSnapshotStatus `json:"status,omitempty" protobuf:"bytes,3,opt,name=status"`
}

type ClusterSnapshotStatus struct {
	Volumes []VolumeSnapshotItem `json:"volumes" protobuf:"bytes,1,rep,name=volumes"`
}

type VolumeSnapshotItem struct {
	// VolumeSnapshotName
	Name    string `json:"name" protobuf:"bytes,1,opt,name=name"`
	PVCName string `json:"pvcName" protobuf:"bytes,2,opt,name=pvcName"`

	// VolumeSnapshotClassName
	ClassName string `json:"className" protobuf:"bytes,3,opt,name=className"`

	// VolumeSnapshotContentName
	ContentName string    `json:"contentName" protobuf:"bytes,4,opt,name=contentName"`
	ReadyToUse  bool      `json:"readyToUse" protobuf:"varint,5,opt,name=readyToUse"`
	FinishTime  time.Time `json:"finishTime" protobuf:"bytes,6,opt,name=finishTime"`
}

func main() {
	updateCSS := &ClusterSnapshot{
		Status: ClusterSnapshotStatus{
			Volumes: []VolumeSnapshotItem{
				{
					Name:        "test-vss-1",
					PVCName:     "test-pvc-1",
					ClassName:   "test-vss-class-1",
					ContentName: "test-vss-content-1",
					ReadyToUse:  false,
					FinishTime:  time.Now(),
				},
				{
					Name:        "test-vss-2",
					PVCName:     "test-pvc-2",
					ClassName:   "test-vss-class-2",
					ContentName: "test-vss-content-2",
					ReadyToUse:  false,
					FinishTime:  time.Now(),
				},
				{
					Name:        "test-vss-3",
					PVCName:     "test-pvc-3",
					ClassName:   "test-vss-class-3",
					ContentName: "test-vss-content-3",
					ReadyToUse:  false,
					FinishTime:  time.Now(),
				},
				{
					Name:        "test-vss-4",
					PVCName:     "test-pvc-4",
					ClassName:   "test-vss-class-4",
					ContentName: "test-vss-content-4",
					ReadyToUse:  false,
					FinishTime:  time.Now(),
				},
				{
					Name:        "test-vss-5",
					PVCName:     "test-pvc-5",
					ClassName:   "test-vss-class-5",
					ContentName: "test-vss-content-5",
					ReadyToUse:  false,
					FinishTime:  time.Now(),
				},
				{
					Name:        "test-vss-6",
					PVCName:     "test-pvc-6",
					ClassName:   "test-vss-class-6",
					ContentName: "test-vss-content-6",
					ReadyToUse:  false,
					FinishTime:  time.Now(),
				},
				{
					Name:        "test-vss-7",
					PVCName:     "test-pvc-7",
					ClassName:   "test-vss-class-7",
					ContentName: "test-vss-content-7",
					ReadyToUse:  false,
					FinishTime:  time.Now(),
				},
				{
					Name:        "test-vss-8",
					PVCName:     "test-pvc-8",
					ClassName:   "test-vss-class-8",
					ContentName: "test-vss-content-8",
					ReadyToUse:  false,
					FinishTime:  time.Now(),
				},
				{
					Name:        "test-vss-9",
					PVCName:     "test-pvc-9",
					ClassName:   "test-vss-class-9",
					ContentName: "test-vss-content-9",
					ReadyToUse:  false,
					FinishTime:  time.Now(),
				},
				{
					Name:        "test-vss-10",
					PVCName:     "test-pvc-10",
					ClassName:   "test-vss-class-10",
					ContentName: "test-vss-content-10",
					ReadyToUse:  false,
					FinishTime:  time.Now(),
				},
			},
		},
	}

	vssName := "test-vss-5"
	for i, vssItem := range updateCSS.Status.Volumes {
		if vssItem.Name == vssName {
			if !vssItem.ReadyToUse {
				updateCSS.Status.Volumes[i].ReadyToUse = true
				updateCSS.Status.Volumes[i].FinishTime = time.Now()
			}
			break
		}
	}

	for _, vssItem := range updateCSS.Status.Volumes {
		//vssItem := vssItem
		fmt.Println(vssItem.ReadyToUse)
	}
}
