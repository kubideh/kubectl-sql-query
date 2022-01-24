package finders

import (
	"context"

	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/client-go/kubernetes"
)

type PodFinder struct {
	clientSet kubernetes.Interface
}

func CreatePodFinder(clientSet kubernetes.Interface) *PodFinder {
	return &PodFinder{
		clientSet: clientSet,
	}
}

func (f *PodFinder) Find(namespace, name string) runtime.Object {
	var podList *corev1.PodList

	if name != "" {
		podList = findPod(f.clientSet, namespace, name)
	} else {
		podList = findAllPods(f.clientSet, namespace)
	}

	return podList
}

func findPod(clientSet kubernetes.Interface, namespace, name string) *corev1.PodList {
	var podList *corev1.PodList

	pod, err := clientSet.CoreV1().Pods(namespace).Get(context.TODO(), name, metav1.GetOptions{})

	if err != nil && !errors.IsNotFound(err) {
		panic(err.Error())
	}

	if errors.IsNotFound(err) {
		podList = &corev1.PodList{}
	} else {
		podList = &corev1.PodList{
			Items: []corev1.Pod{
				*pod,
			},
		}
	}

	return podList
}

func findAllPods(clientSet kubernetes.Interface, namespace string) *corev1.PodList {
	podList, err := clientSet.CoreV1().Pods(namespace).List(context.TODO(), metav1.ListOptions{})

	if err != nil {
		panic(err.Error())
	}

	return podList
}
