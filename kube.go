package main

import (
	"flag"
	"fmt"
	"k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
	"os"
	"path/filepath"
	"regexp"
	"time"
)


var clientset *kubernetes.Clientset
func init()  {
	var kubeconfig *string
	if home := homeDir(); home != "" {
		kubeconfig = flag.String("kubeconfig", filepath.Join(home, ".kube", "config"), "(optional) absolute path to the kubeconfig file")
	} else {
		kubeconfig = flag.String("kubeconfig", "", "absolute path to the kubeconfig file")
	}
	flag.Parse()

	config, err := clientcmd.BuildConfigFromFlags("", *kubeconfig)
	if err != nil {
		panic(err.Error())
	}

	clientset, err = kubernetes.NewForConfig(config)
	if err != nil {
		panic(err.Error())
	}
}
func main()  {
	findPodByNamespace("default", "apollo-649c55f8c5-dsf5f")
	nss, err := clientset.CoreV1().Namespaces().List(metav1.ListOptions{})
	if err != nil {
		panic(err)
	}
	for true  {
		for _,v := range nss.Items {
			findPodsByNamespace(v.Name)
		}
		time.Sleep(3 * time.Minute)
	}

}
func findPodByNamespace(namespace, pod string) {
	_,err := clientset.CoreV1().Pods(namespace).Get(pod, metav1.GetOptions{})
	if errors.IsNotFound(err) {
		fmt.Printf("pods %s in namespace %s not found\n", pod, namespace)
	}else if statusError, isStatus := err.(*errors.StatusError); isStatus {
		fmt.Printf("error getting pods %s in namespace %s: v%\n",
			pod,namespace,statusError.ErrStatus.Message)
	}else if err != nil{
		panic(err.Error())
	} else {
		fmt.Printf("Found pod %s in namespace %s\n", pod, namespace)
	}
}
func findPodsByNamespace(namespace string)  {
	pods, err := clientset.CoreV1().Pods(namespace).List(metav1.ListOptions{})
	if err != nil {
		panic(err.Error())
	}
	//for pod := range pods.Items {
	//	println(pod)
	//}
	for i := 0; i < len(pods.Items); i++ {
		fmt.Println(pods.Items[i].Name)
		//fmt.Println(pods.Items[i].Status.ContainerStatuses)
		for a := 0; a < len(pods.Items[i].Status.ContainerStatuses); a++ {
			status := pods.Items[i].Status.ContainerStatuses[a].Ready
			//fmt.Println(status)
			//fmt.Println(pods.Items[i].Status.StartTime)
			now := time.Now()
			if status != true {
				if now.Sub(pods.Items[i].Status.StartTime.Time).Minutes() > 3 {
					fmt.Println(now.Sub(pods.Items[i].Status.StartTime.Time).Minutes())
					fmt.Println("dingding")
					reg := regexp.MustCompile(`Error|Fail|Cannot`)
					if reg.MatchString(pods.Items[i].Status.ContainerStatuses[0].State.Waiting.Reason) == true {
						send(pods.Items[i].Namespace, pods.Items[i].Name, pods.Items[i].Status.ContainerStatuses[0].State.Waiting.Message)
					}
				}
			}

		}
	}
}

func homeDir() string {
	if h := os.Getenv("HOME"); h != "" {
		return h
	}
	return os.Getenv("USERPROFILE")
}