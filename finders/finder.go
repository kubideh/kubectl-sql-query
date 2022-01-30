package finders

import (
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/client-go/kubernetes"
)

// Finder provides a way to find K8s API objects by namespace and name.
type Finder interface {
	Find(namespace, name string) runtime.Object
}

// Create returns the correct Finder based on K8s object kind.
func Create(clientSet kubernetes.Interface, kind string) Finder {
	if kind == "deployments" {
		return CreateDeploymentFinder(clientSet)
	} else {
		return CreatePodFinder(clientSet)
	}
}
