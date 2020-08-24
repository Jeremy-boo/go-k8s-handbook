package informer

import (
	"github.com/go-k8s-handbook/util"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/informers"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/cache"
	"k8s.io/client-go/tools/clientcmd"
	"log"
	"time"
)

func InitInformerExample() {
	kubeConfig := util.ParseK8sConfig()
	config, err := clientcmd.BuildConfigFromFlags("", *kubeConfig)
	if err != nil {
		panic(err)
	}
	clientSet, err := kubernetes.NewForConfig(config)
	if err != nil {
		panic(err)
	}
	stopCh := make(chan struct{})
	defer close(stopCh)
	// 实例化sharedInformerFactory
	sharedInformerFactory := informers.NewSharedInformerFactory(clientSet, time.Minute)
	// get pod Informer
	podInformer := sharedInformerFactory.Core().V1().Pods().Informer()
	// 添加监听event 事件
	podInformer.AddEventHandler(cache.ResourceEventHandlerFuncs{
		AddFunc: func(obj interface{}) {
			mObj := obj.(v1.Object)
			log.Printf("New pod Added to Store:%s", mObj.GetName())
		},
		UpdateFunc: func(oldObj, newObj interface{}) {
			oObj := oldObj.(v1.Object)
			nObj := newObj.(v1.Object)
			log.Printf("%s Pod Updated to %s", oObj.GetName(), nObj.GetName())
		},
		DeleteFunc: func(obj interface{}) {
			mObj := obj.(v1.Object)
			log.Printf("Pod deleted From Store:%s", mObj.GetName())
		},
	})
	podInformer.Run(stopCh)
}
