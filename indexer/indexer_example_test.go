package indexer

import "testing"

func TestIndexer(t *testing.T) {
	err := Indexer()
	if err != nil {
		t.Error(err)
		return
	}
	t.Log("success")
}
