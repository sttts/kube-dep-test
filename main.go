package main

import (
	"k8s.io/client-go/kubernetes"
)

func main() {
	_ = kubernetes.Clientset{}
}
