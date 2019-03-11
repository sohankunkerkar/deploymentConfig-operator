package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// DCWatcherSpec defines the desired state of DCWatcher
// +k8s:openapi-gen=true
type DCWatcherSpec struct {
	Size  int32  `json:"size"`
	Image string `json:"image"`
}

// DCWatcherStatus defines the observed state of DCWatcher
// +k8s:openapi-gen=true
type DCWatcherStatus struct {
	Status string `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// DCWatcher is the Schema for the dcwatchers API
// +k8s:openapi-gen=true
type DCWatcher struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   DCWatcherSpec   `json:"spec,omitempty"`
	Status DCWatcherStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// DCWatcherList contains a list of DCWatcher
type DCWatcherList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []DCWatcher `json:"items"`
}

func init() {
	SchemeBuilder.Register(&DCWatcher{}, &DCWatcherList{})
}
