package client

import "github.com/b2wdigital/goignite/v2/core/config"

const (
	root           = "gi.k8s"
	kubeConfigPath = root + ".kubeConfigPath"
)

func init() {
	config.Add(kubeConfigPath, "~/.kube/config", "defines kubeconfig request")
}
