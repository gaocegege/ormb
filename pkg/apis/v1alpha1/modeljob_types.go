/*
Copyright 2020 Caicloud Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// ModelJobSpec defines the desired state of ModelJob
type ModelJobSpec struct {
	// Important: Run "make" to regenerate code after modifying this file

	Model      string  `json:"model,omitempty"`
	DesiredTag *string `json:"desiredTag,omitempty"`

	ModelJobSource `json:",inline"`
}

type ModelJobSource struct {
	Conversion *ConversionSource `json:"conversion,omitempty"`
}

type ConversionSource struct {
	MMdnn *MMdnnSpec `json:"mmdnn,omitempty"`
}

type MMdnnSpec struct {
	ConversionBaseSpec `json:",inline"`
}

type ConversionBaseSpec struct {
	From Format `json:"from,omitempty"`
	To   Format `json:"to,omitempty"`
}

type Format string

const (
	SavedModel Format = "SavedModel"
	ONNX              = "ONNX"
)

// ModelJobStatus defines the observed state of ModelJob
type ModelJobStatus struct {
	// Important: Run "make" to regenerate code after modifying this file
}

// +kubebuilder:object:root=true

// ModelJob is the Schema for the modeljobs API
type ModelJob struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   ModelJobSpec   `json:"spec,omitempty"`
	Status ModelJobStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ModelJobList contains a list of ModelJob
type ModelJobList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ModelJob `json:"items"`
}

func init() {
	SchemeBuilder.Register(&ModelJob{}, &ModelJobList{})
}
