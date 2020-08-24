package v1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// ManagedClusterAddon is the Custom Resource object which holds the current state
// of an operator. This object is used by operators to convey their state to
// the rest of the cluster.
type ManagedClusterAddon struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata"`

	// spec holds configuration that could apply to any operator.
	// +kubebuilder:validation:Required
	// +required
	Spec ManagedClusterAddonSpec `json:"spec"`

	// status holds the information about the state of an operator.  It is consistent with status information across
	// the Kubernetes ecosystem.
	// +optional
	Status ManagedClusterAddonStatus `json:"status"`
}

// ManagedClusterAddonSpec is empty for now, but you could imagine holding information like "pause".
type ManagedClusterAddonSpec struct {
	// UpdateApproved ...
	UpdateApproved bool `json:"updateApproved"`
}

// ManagedClusterAddonStatus provides information about the status of the operator.
// +k8s:deepcopy-gen=true
type ManagedClusterAddonStatus struct {
	// conditions describes the state of the operator's managed and monitored components.
	// +patchMergeKey=type
	// +patchStrategy=merge
	// +optional
	Conditions []AddonStatusCondition `json:"conditions,omitempty"  patchStrategy:"merge" patchMergeKey:"type"`

	// relatedObjects is a reference to the configuration resource that drives the operator
	// +optional
	RelatedObject ObjectReference `json:"relatedObject,omitempty"`

	// LatestRelease indicates the latest available release for the addon
	LatestRelease Release `json:"availableUpdate,omitempty"`

	// CurrentRelease ...
	CurrentRelease Release `json:"currentVersion,omitempty"`

	// extension contains any additional status information specific to the
	// operator which owns this status object.
	// +nullable
	// +optional
	// +kubebuilder:pruning:PreserveUnknownFields
	Extension runtime.RawExtension `json:"extension"`
}

// ObjectReference contains enough information to let you inspect or modify the referred object.
type ObjectReference struct {
	// group of the referent.
	// +kubebuilder:validation:Required
	// +required
	Group string `json:"group"`
	// resource of the referent.
	// +kubebuilder:validation:Required
	// +required
	Resource string `json:"resource"`
	// name of the referent.
	// +kubebuilder:validation:Required
	// +required
	Name string `json:"name"`
}

type ConditionStatus string

// These are valid condition statuses. "ConditionTrue" means a resource is in the condition.
// "ConditionFalse" means a resource is not in the condition. "ConditionUnknown" means kubernetes
// can't decide if a resource is in the condition or not. In the future, we could add other
// intermediate conditions, e.g. ConditionDegraded.
const (
	ConditionTrue    ConditionStatus = "True"
	ConditionFalse   ConditionStatus = "False"
	ConditionUnknown ConditionStatus = "Unknown"
)

// AddonStatusCondition represents the state of the operator's
// managed and monitored components.
// +k8s:deepcopy-gen=true
type AddonStatusCondition struct {
	// type specifies the aspect reported by this condition.
	// +kubebuilder:validation:Required
	// +required
	Type AddonStatusConditionType `json:"type"`

	// status of the condition, one of True, False, Unknown.
	// +kubebuilder:validation:Required
	// +required
	Status ConditionStatus `json:"status"`

	// lastTransitionTime is the time of the last update to the current status property.
	// +kubebuilder:validation:Required
	// +required
	LastTransitionTime metav1.Time `json:"lastTransitionTime"`

	// reason is the CamelCase reason for the condition's current status.
	// +optional
	Reason string `json:"reason,omitempty"`

	// message provides additional information about the current condition.
	// This is only to be consumed by humans.
	// +optional
	Message string `json:"message,omitempty"`
}

// Release represents an OpenShift release image and associated metadata.
// +k8s:deepcopy-gen=true
type Release struct {
	// version is a semantic versioning identifying the update version. When this
	// field is part of spec, version is optional if image is specified.
	// +required
	Version string `json:"version"`

	// imageManifest is list of images will be used for the specific version
	// this will be use for backward compatability (i.e hub upgraded but agent have not)
	// +required
	ImageManifest []Image `json:"imageManifest"`
}

type Image struct {
	Name  string `json: "name"`
	Image string `json: "image"`
}

// AddonStatusConditionType is an aspect of operator state.
type AddonStatusConditionType string

const (
	// Available indicates that the operand (eg: openshift-apiserver for the
	// openshift-apiserver-operator), is functional and available in the cluster.
	Available AddonStatusConditionType = "Available"

	// Progressing indicates that the operator is actively rolling out new code,
	// propagating config changes, or otherwise moving from one steady state to
	// another.  Operators should not report progressing when they are reconciling
	// a previously known state.
	Progressing AddonStatusConditionType = "Progressing"

	// Degraded indicates that the operator's current state does not match its
	// desired state over a period of time resulting in a lower quality of service.
	// The period of time may vary by component, but a Degraded state represents
	// persistent observation of a condition.  As a result, a component should not
	// oscillate in and out of Degraded state.  A service may be Available even
	// if its degraded.  For example, your service may desire 3 running pods, but 1
	// pod is crash-looping.  The service is Available but Degraded because it
	// may have a lower quality of service.  A component may be Progressing but
	// not Degraded because the transition from one state to another does not
	// persist over a long enough period to report Degraded.  A service should not
	// report Degraded during the course of a normal upgrade.  A service may report
	// Degraded in response to a persistent infrastructure failure that requires
	// administrator intervention.  For example, if a control plane host is unhealthy
	// and must be replaced.  An operator should report Degraded if unexpected
	// errors occur over a period, but the expectation is that all unexpected errors
	// are handled as operators mature.
	Degraded AddonStatusConditionType = "Degraded"
)

// ManagedClusterAddonList is a list of OperatorStatus resources.
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
type ManagedClusterAddonList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata"`

	Items []ManagedClusterAddon `json:"items"`
}
