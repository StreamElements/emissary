/*


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

package v2

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

type AuthService_AuthIncludeBody struct {
	MaxBytes     int32 `json:"max_bytes,omitempty"`
	AllowPartial bool  `json:"allow_partial,omitempty"`
}

// Why isn't this just an int??
type AuthService_AuthStatusOnError struct {
	Code int32 `json:"code,omitempty"`
}

// AuthServiceSpec defines the desired state of AuthService
type AuthServiceSpec struct {
	AmbassadorID AmbassadorID `json:"ambassador_id,omitempty"`

	AuthService string `json:"auth_service,omitempty"`
	PathPrefix  string `json:"path_prefix,omitempty"`
	TLS         string `json:"tls,omitempty"`
	// +kubebuilder:validation:Enum={"http","grpc"}
	Proto                       string                         `json:"proto,omitempty"`
	TimeoutMs                   int32                          `json:"timeout_ms,omitempty"`
	AllowedRequestHeaders       []string                       `json:"allowed_request_headers,omitempty"`
	AllowedAuthorizationHeaders []string                       `json:"allowed_authorization_headers,omitempty"`
	AddAuthHeaders              map[string]string              `json:"add_auth_headers,omitempty"`
	AllowRequestBody            bool                           `json:"allow_request_body,omitempty"`
	AddLinkerdHeaders           bool                           `json:"add_linkerd_headers,omitempty"`
	FailureModeAllow            bool                           `json:"failure_mode_allow,omitempty"`
	IncludeBody                 *AuthService_AuthIncludeBody   `json:"include_body,omitempty"`
	StatusOnError               *AuthService_AuthStatusOnError `json:"status_on_error,omitempty"`
}

// AuthServiceStatus defines the observed state of AuthService
type AuthServiceStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "make" to regenerate code after modifying this file
}

// +kubebuilder:object:root=true

// AuthService is the Schema for the authservices API
type AuthService struct {
	metav1.TypeMeta   `json:""`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   AuthServiceSpec   `json:"spec,omitempty"`
	Status AuthServiceStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// AuthServiceList contains a list of AuthService
type AuthServiceList struct {
	metav1.TypeMeta `json:""`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []AuthService `json:"items"`
}

func init() {
	SchemeBuilder.Register(&AuthService{}, &AuthServiceList{})
}
