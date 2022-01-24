package finders

import (
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/client-go/kubernetes"
)

type Finder interface {
	Find(namespace, name string) runtime.Object
}

func CreateFinder(clientSet kubernetes.Interface, kind string) Finder {
	if kind == "deployments" {
		return CreateDeploymentFinder(clientSet)
	} else {
		return CreatePodFinder(clientSet)
	}
}
