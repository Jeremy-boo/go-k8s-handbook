package indexer

import (
	"fmt"
	"k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/tools/cache"
	"strings"
)

// UsersIndexFunc ...
func UsersIndexFunc(obj interface{}) ([]string, error) {
	pod := obj.(*v1.Pod)
	usersString := pod.Annotations["users"]
	return strings.Split(usersString, ","), nil
}

//  Indexer
func Indexer() error {
	index := cache.NewIndexer(cache.MetaNamespaceKeyFunc, cache.Indexers{"byUsers": UsersIndexFunc})
	pod1 := &v1.Pod{ObjectMeta: metav1.ObjectMeta{Name: "one",
		Annotations: map[string]string{"users": "ernie,bert"}}}
	pod2 := &v1.Pod{ObjectMeta: metav1.ObjectMeta{Name: "two",
		Annotations: map[string]string{"users": "bert,oscar"}}}
	pod3 := &v1.Pod{ObjectMeta: metav1.ObjectMeta{Name: "three",
		Annotations: map[string]string{"users": "ernie,elmo"}}}
	index.Add(pod1)
	index.Add(pod2)
	index.Add(pod3)

	erniePods, err := index.ByIndex("byUsers", "ernie")
	if err != nil {
		return err
	}
	for _, val := range erniePods {
		fmt.Println(val.(*v1.Pod).Name)
	}
	return nil
}
