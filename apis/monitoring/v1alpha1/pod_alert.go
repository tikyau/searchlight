package v1alpha1

import (
	"fmt"
	"time"

	apiv1 "k8s.io/api/core/v1"
)

const (
	ResourceKindPodAlert = "PodAlert"
	ResourceNamePodAlert = "pod-alert"
	ResourceTypePodAlert = "podalerts"
)

var _ Alert = &PodAlert{}

func (a PodAlert) GetName() string {
	return a.Name
}

func (a PodAlert) GetNamespace() string {
	return a.Namespace
}

func (a PodAlert) Command() string {
	return string(a.Spec.Check)
}

func (a PodAlert) GetCheckInterval() time.Duration {
	return a.Spec.CheckInterval.Duration
}

func (a PodAlert) GetAlertInterval() time.Duration {
	return a.Spec.AlertInterval.Duration
}

func (a PodAlert) IsValid() (bool, error) {
	cmd, ok := PodCommands[a.Spec.Check]
	if !ok {
		return false, fmt.Errorf("%s is not a valid pod check command.", a.Spec.Check)
	}
	for k := range a.Spec.Vars {
		if _, ok := cmd.Vars[k]; !ok {
			return false, fmt.Errorf("Var %s is unsupported for check command %s.", k, a.Spec.Check)
		}
	}
	for _, rcv := range a.Spec.Receivers {
		found := false
		for _, state := range cmd.States {
			if state == rcv.State {
				found = true
				break
			}
		}
		if !found {
			return false, fmt.Errorf("State %s is unsupported for check command %s.", rcv.State, a.Spec.Check)
		}
	}
	return true, nil
}

func (a PodAlert) GetNotifierSecretName() string {
	return a.Spec.NotifierSecretName
}

func (a PodAlert) GetReceivers() []Receiver {
	return a.Spec.Receivers
}

func (a PodAlert) ObjectReference() *apiv1.ObjectReference {
	return &apiv1.ObjectReference{
		APIVersion:      SchemeGroupVersion.String(),
		Kind:            ResourceKindPodAlert,
		Namespace:       a.Namespace,
		Name:            a.Name,
		UID:             a.UID,
		ResourceVersion: a.ResourceVersion,
	}
}
