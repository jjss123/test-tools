//package main
//
//import (
//	"flag"
//	"fmt"
//	"log"
//	"path/filepath"
//	"context"
//
//	"k8s.io/apimachinery/pkg/api/meta"
//	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
//	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
//	"k8s.io/apimachinery/pkg/runtime/schema"
//	"k8s.io/apimachinery/pkg/runtime/serializer/yaml"
//	"k8s.io/client-go/discovery"
//	"k8s.io/client-go/discovery/cached/memory"
//	"k8s.io/client-go/dynamic"
//	"k8s.io/client-go/rest"
//	"k8s.io/client-go/restmapper"
//	"k8s.io/client-go/tools/clientcmd"
//	"k8s.io/client-go/util/homedir"
//)
//
//// 自定义数据
//const metaCRD = `
//apiVersion: apiextensions.k8s.io/v1beta1
//kind: CustomResourceDefinition
//metadata:
//  annotations:
//    controller-gen.kubebuilder.io/version: v0.2.5
//  creationTimestamp: null
//  name: domains.jvessel.jdcloud.com
//spec:
//  additionalPrinterColumns:
//  - JSONPath: .spec.dualStack
//    name: DualStack
//    type: string
//  - JSONPath: .status.domainStatus
//    name: Status
//    type: string
//  group: jvessel.jdcloud.com
//  names:
//    kind: Domain
//    listKind: DomainList
//    plural: domains
//    singular: domain
//  scope: Namespaced
//  subresources:
//    status: {}
//  validation:
//    openAPIV3Schema:
//      description: Domain is the Schema for the domains API
//      properties:
//        apiVersion:
//          description: 'APIVersion defines the versioned schema of this representation
//            of an object. Servers should convert recognized schemas to the latest
//            internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources'
//          type: string
//        kind:
//          description: 'Kind is a string value representing the REST resource this
//            object represents. Servers may infer this from the endpoint the client
//            submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds'
//          type: string
//        metadata:
//          type: object
//        spec:
//          description: PublicDomainSpec defines the desired state of PublicDomain
//          properties:
//            dualStack:
//              type: string
//            ips:
//              additionalProperties:
//                properties:
//                  ip:
//                    type: string
//                  ttl:
//                    type: integer
//                  type:
//                    type: string
//                  viewValue:
//                    type: integer
//                  weight:
//                    type: integer
//                required:
//                - ttl
//                - viewValue
//                - weight
//                type: object
//              type: object
//            serviceName:
//              type: string
//            serviceNamespace:
//              type: string
//          type: object
//        status:
//          description: PublicDomainStatus defines the observed state of PublicDomain
//          properties:
//            domainRecord:
//              additionalProperties:
//                properties:
//                  ip:
//                    type: string
//                  resourceRecordId:
//                    type: string
//                  status:
//                    type: string
//                  ttl:
//                    type: integer
//                  type:
//                    type: string
//                  weight:
//                    type: integer
//                required:
//                - resourceRecordId
//                type: object
//              type: object
//            domainStatus:
//              type: string
//          type: object
//      type: object
//  version: v1alpha1
//  versions:
//  - name: v1alpha1
//    served: true
//    storage: true
//`
//
//func GetK8sConfig() (config *rest.Config, err error) {
//	// 获取k8s rest config
//	var kubeconfig *string
//	if home := homedir.HomeDir(); home != "" {
//		kubeconfig = flag.String("kubeconfig", filepath.Join(home, ".kube", "config"), "(optional) absolute path to the kubeconfig file")
//	} else {
//		kubeconfig = flag.String("kubeconfig", "", "absolute path to the kubeconfig file")
//	}
//	flag.Parse()
//
//	config, err = clientcmd.BuildConfigFromFlags("", *kubeconfig)
//	if err != nil {
//		panic(err)
//	}
//	return
//}
//
//func GetGVRdyClient(gvk *schema.GroupVersionKind, namespace string) (dr dynamic.ResourceInterface, err error) {
//
//	config, err := GetK8sConfig()
//	if err != nil {
//		return
//	}
//
//	// 创建discovery客户端
//	discoveryClient, err := discovery.NewDiscoveryClientForConfig(config)
//	if err != nil {
//		return
//	}
//
//	// 获取GVK GVR 映射
//	mapperGVRGVK := restmapper.NewDeferredDiscoveryRESTMapper(memory.NewMemCacheClient(discoveryClient))
//
//	// 根据资源GVK 获取资源的GVR GVK映射
//	resourceMapper, err := mapperGVRGVK.RESTMapping(gvk.GroupKind(), gvk.Version)
//	if err != nil {
//		return
//	}
//
//	// 创建动态客户端
//	dynamicClient, err := dynamic.NewForConfig(config)
//	if err != nil {
//		return
//	}
//
//	if resourceMapper.Scope.Name() == meta.RESTScopeNameNamespace {
//		// 获取gvr对应的动态客户端
//		dr = dynamicClient.Resource(resourceMapper.Resource).Namespace(namespace)
//	} else {
//		// 获取gvr对应的动态客户端
//		dr = dynamicClient.Resource(resourceMapper.Resource)
//	}
//
//	return
//}
//
//func main() {
//
//	var (
//		err       error
//		objGET    *unstructured.Unstructured
//		//objCreate *unstructured.Unstructured
//		//objUpdate *unstructured.Unstructured
//		gvk       *schema.GroupVersionKind
//		dr        dynamic.ResourceInterface
//	)
//	obj := &unstructured.Unstructured{}
//	_, gvk, err = yaml.NewDecodingSerializer(unstructured.UnstructuredJSONScheme).Decode([]byte(metaCRD), nil, obj)
//	if err != nil {
//		panic(fmt.Errorf("failed to get GVK: %v", err))
//	}
//
//	dr, err = GetGVRdyClient(gvk, obj.GetNamespace())
//	if err != nil {
//		panic(fmt.Errorf("failed to get dr: %v", err))
//	}
//
//	////创建
//	//objCreate, err = dr.Create(obj, metav1.CreateOptions{})
//	//if err != nil {
//	//	//panic(fmt.Errorf("Create resource ERROR: %v", err))
//	//	log.Print(err)
//	//}
//	//log.Print("Create: : ", objCreate)
//
//	// 查询
//	objGET, err = dr.Get(context.Background(), obj.GetName(), metav1.GetOptions{}, "json`group`")
//	if err != nil {
//		panic(fmt.Errorf("select resource ERROR: %v", err))
//	}
//	log.Print("GET: ", objGET)
//
//	////更新
//	//retryErr := retry.RetryOnConflict(retry.DefaultRetry, func() (err error) {
//	//	// 查询resource是否存在
//	//	result, getErr := dr.Get(obj.GetName(), metav1.GetOptions{})
//	//	if getErr != nil {
//	//		panic(fmt.Errorf("failed to get latest version of : %v", getErr))
//	//	}
//	//
//	//	// 提取obj 的 spec 期望值
//	//	spec, found, err := unstructured.NestedMap(obj.Object, "spec")
//	//	if err != nil || !found || spec == nil {
//	//		panic(fmt.Errorf(" not found or error in spec: %v", err))
//	//	}
//	//	// 更新 存在资源的spec
//	//	if err := unstructured.SetNestedMap(result.Object, spec, "spec"); err != nil {
//	//		panic(err)
//	//	}
//	//	// 更新资源
//	//	objUpdate, err = dr.Update(result, metav1.UpdateOptions{})
//	//	log.Print("update : ", objUpdate)
//	//	return err
//	//})
//	//if retryErr != nil {
//	//	panic(fmt.Errorf("update failed: %v", retryErr))
//	//} else {
//	//	log.Print("更新成功")
//	//}
//
//	// 删除
//	//err = dr.Delete(obj.GetName(),&metav1.DeleteOptions{})
//	//if err != nil {
//	//	panic(fmt.Errorf("delete resource ERROR : %v", err))
//	//} else {
//	//	log.Print("删除成功")
//	//}
//
//}
