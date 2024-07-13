/*
Copyright 2024.

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

package v1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// Base64ControllerSpec defines the desired state of Base64Controller
type Base64ControllerSpec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "make" to regenerate code after modifying this file

	// Foo is an example field of Base64Controller. Edit base64controller_types.go to remove/update
	Name string `json:"name,omitempty"`
	Text string `json:"text,omitempty"`
}

// Base64ControllerStatus defines the observed state of Base64Controller
type Base64ControllerStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "make" to regenerate code after modifying this file
	EncodedText string `json:"encodedText,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status

// Base64Controller is the Schema for the base64controllers API
type Base64Controller struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   Base64ControllerSpec   `json:"spec,omitempty"`
	Status Base64ControllerStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// Base64ControllerList contains a list of Base64Controller
type Base64ControllerList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Base64Controller `json:"items"`
}

func init() {
	SchemeBuilder.Register(&Base64Controller{}, &Base64ControllerList{})
}
