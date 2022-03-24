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

	batchv1 "k8s.io/api/batch/v1"

	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func main() {

	podTempSpec := corev1.PodTemplateSpec{
		ObjectMeta: metav1.ObjectMeta{
			Name: "test-name",
		},
		Spec: corev1.PodSpec{},
	}

	job := batchv1.Job{
		TypeMeta:   metav1.TypeMeta{},
		ObjectMeta: metav1.ObjectMeta{},
		Spec: batchv1.JobSpec{
			Template: podTempSpec,
		},
		Status: batchv1.JobStatus{},
	}

	fmt.Println(job.Spec.Template.ObjectMeta.Name) // same as below
	fmt.Println(job.Spec.Template.Name)            // same as above

	fmt.Printf("%+v\n", job.Spec.Template.ObjectMeta)
	fmt.Printf("%+v\n", podTempSpec)
	fmt.Printf("%+v\n", job)
}
