package util

import (
	"flag"
	"os"
	"path/filepath"
)

// ParseK8sConfig 解析kubernetes config 目录
func ParseK8sConfig() *string {
	var k8sConfig *string
	if home := homeDir(); home != "" {
		k8sConfig = flag.String("kubeconfig", filepath.Join(home, ".kube", "config"), "(optional) absolute path to the kubeconfig file")
	} else {
		k8sConfig = flag.String("kubeconfig", "", "absolute path to the kubeconfig file")
	}
	flag.Parse()
	return k8sConfig
}

func homeDir() string {
	if h := os.Getenv("HOME"); h != "" {
		return h
	}
	return os.Getenv("USERPROFILE") // windows
}
