package util

import "testing"

func TestParseK8sConfig(t *testing.T) {
	path := ParseK8sConfig()
	if *path == "" {
		t.Error("get kube config error")
		return
	}
	t.Log(*path)
}
