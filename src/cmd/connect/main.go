package main

import (
	"context"
	"fmt"

	"flag"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/client-go/dynamic"
	"k8s.io/client-go/util/homedir"
	"path/filepath"

	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
	"log"
)

func GetK8sConfig() (config *rest.Config, err error) {
	// 获取k8s rest config
	var kubeconfig *string
	if home := homedir.HomeDir(); home != "" {
		kubeconfig = flag.String("kubeconfig", filepath.Join(home, ".kube", "config"), "(optional) absolute path to the kubeconfig file")
	} else {
		kubeconfig = flag.String("kubeconfig", "", "absolute path to the kubeconfig file")
	}
	flag.Parse()

	config, err = clientcmd.BuildConfigFromFlags("", *kubeconfig)
	if err != nil {
		panic(err)
	}
	return
}

func main() {
	log.SetFlags(log.Llongfile)
	flag.Parse()
	//获取Config
	config, err := GetK8sConfig()
	if err != nil {
		return
	}
	clthird, err := dynamic.NewForConfig(config)
	if err != nil {
		log.Println(err)
	}

	gvr := schema.GroupVersionResource{"jvessel.jdcloud.com", "v1alpha1", "exposeservices"}

	meta := metav1.TypeMeta{"ExposeService", "jvessel.jdcloud.com/v1alpha1"}
	optionsGet := metav1.GetOptions{meta, "266785193"}

	objthird, err := clthird.Resource(gvr).Namespace("jiashuo-38149f645").Get(context.Background(), "myweb-clusterip1", optionsGet)
	if err != nil {
		log.Println(err)
	}
	fmt.Println(objthird)


	optionsList := new(metav1.ListOptions)
	optionsList.TypeMeta = meta
	objthird1, err := clthird.Resource(gvr).Namespace("jiashuo-38149f645").List(context.Background(), *optionsList)
	if err != nil {
		log.Println(err)
	}
	fmt.Println(objthird1)
}
