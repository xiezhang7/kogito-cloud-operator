// Copyright 2020 Red Hat, Inc. and/or its affiliates
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// KogitoTrustyCRDName is the name of the Kogito Trusty CRD in the cluster.
const KogitoTrustyCRDName = "kogitotrusties.app.kiegroup.org"

// KogitoTrustySpec defines the desired state of KogitoTrusty.
// +k8s:openapi-gen=true
type KogitoTrustySpec struct {
	InfinispanMeta    `json:",inline"`
	KafkaMeta         `json:",inline"`
	KogitoServiceSpec `json:",inline"`
	// Add custom validation using kubebuilder tags: https://book-v1.book.kubebuilder.io/beyond_basics/generating_crd.html
}

// GetRuntime ...
func (d *KogitoTrustySpec) GetRuntime() RuntimeType {
	return QuarkusRuntimeType
}

// KogitoTrustyStatus defines the observed state of KogitoTrusty.
// +k8s:openapi-gen=true
type KogitoTrustyStatus struct {
	KogitoServiceStatus `json:",inline"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// KogitoTrusty defines the Trusty Service infrastructure deployment.
// +k8s:openapi-gen=true
// +kubebuilder:resource:path=kogitotrusties,scope=Namespaced
// +kubebuilder:printcolumn:name="Replicas",type="integer",JSONPath=".spec.replicas",description="Number of replicas set for this service"
// +kubebuilder:printcolumn:name="Image",type="string",JSONPath=".status.image",description="Base image for this service"
// +kubebuilder:printcolumn:name="Endpoint",type="string",JSONPath=".status.externalURI",description="External URI to access this service"
// +operator-sdk:gen-csv:customresourcedefinitions.displayName="Kogito Trusty"
// +operator-sdk:gen-csv:customresourcedefinitions.resources="Deployment,apps/v1,\"A Kubernetes Deployment\""
// +operator-sdk:gen-csv:customresourcedefinitions.resources="Route,route.openshift.io/v1,\"A Openshift Route\""
// +operator-sdk:gen-csv:customresourcedefinitions.resources="Service,v1,\"A Kubernetes Service\""
// +operator-sdk:gen-csv:customresourcedefinitions.resources="KafkaTopic,kafka.strimzi.io/v1beta1,\"A Kafka topic\""
type KogitoTrusty struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   KogitoTrustySpec   `json:"spec,omitempty"`
	Status KogitoTrustyStatus `json:"status,omitempty"`
}

// GetSpec ...
func (k *KogitoTrusty) GetSpec() KogitoServiceSpecInterface {
	return &k.Spec
}

// GetStatus ...
func (k *KogitoTrusty) GetStatus() KogitoServiceStatusInterface {
	return &k.Status
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// KogitoTrustyList contains a list of KogitoTrusty.
type KogitoTrustyList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// +listType=atomic
	Items []KogitoTrusty `json:"items"`
}

// GetItemsCount ...
func (l *KogitoTrustyList) GetItemsCount() int {
	return len(l.Items)
}

// GetItemAt ...
func (l *KogitoTrustyList) GetItemAt(index int) KogitoService {
	if len(l.Items) > index {
		return KogitoService(&l.Items[index])
	}
	return nil
}

func init() {
	SchemeBuilder.Register(&KogitoTrusty{}, &KogitoTrustyList{})
}
