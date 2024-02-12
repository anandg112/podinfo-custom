/*
Copyright 2024 AnandGautam.

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

type Image struct {
	Repository string `json:"repository"`
	Tag        string `json:"tag"`
}

type Resources struct {
	MemoryLimit string `json:"memoryLimit,omitempty"`
	CpuRequest  string `json:"cpuRequest,omitempty"`
}

type Redis struct {
	Enabled bool `json:"enabled,omitempty"`
}

type UI struct {
	Color   string `json:"color,omitempty"`
	Message string `json:"message,omitempty"`
}

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// MyAppResourceSpec defines the desired state of MyAppResource
type MyAppResourceSpec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "make" to regenerate code after modifying this file

	// Size defines the number of MyAppResource instances
	// The following markers will use OpenAPI v3 schema to validate the value
	// More info: https://book.kubebuilder.io/reference/markers/crd-validation.html

	// Port defines the port that will be used to init the container with the image
	ContainerPort int32 `json:"containerPort,omitempty"`

	ReplicaCount int32 `json:"replicaCount,omitempty"`

	// Resources defines the resources that will be used to init the container
	Resources Resources `json:"resources,omitempty"`

	// Image defines the container image that will be used to init the container
	Image Image `json:"image"`

	// UI defines the UI that will be used to init the container
	UI UI `json:"ui,omitempty"`

	//Redis defines the Redis that will be used to cache values in the container
	Redis Redis `json:"redis,omitempty"`
}

// MyAppResourceStatus defines the observed state of MyAppResource
type MyAppResourceStatus struct {
	// Represents the observations of a MyAppResource's current state.
	// MyAppResource.status.conditions.type are: "Available", "Progressing", and "Degraded"
	// MyAppResource.status.conditions.status are one of True, False, Unknown.
	// MyAppResource.status.conditions.reason the value should be a CamelCase string and producers of specific
	// condition types may define expected values and meanings for this field, and whether the values
	// are considered a guaranteed API.
	// MyAppResource.status.conditions.Message is a human readable message indicating details about the transition.
	// For further information see: https://github.com/kubernetes/community/blob/master/contributors/devel/sig-architecture/api-conventions.md#typical-status-properties

	Conditions []metav1.Condition `json:"conditions,omitempty" patchStrategy:"merge" patchMergeKey:"type" protobuf:"bytes,1,rep,name=conditions"`
}

//+kubebuilder:object:root=true
//+kubebuilder:subresource:status

// MyAppResource is the Schema for the myappresources API
type MyAppResource struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   MyAppResourceSpec   `json:"spec,omitempty"`
	Status MyAppResourceStatus `json:"status,omitempty"`
}

//+kubebuilder:object:root=true

// MyAppResourceList contains a list of MyAppResource
type MyAppResourceList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []MyAppResource `json:"items"`
}

func init() {
	SchemeBuilder.Register(&MyAppResource{}, &MyAppResourceList{})
}
