package main

import (
	"fmt"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"regexp"
	"time"
)

var clientset *kubernetes.Clientset

func init() {
	config, err := rest.InClusterConfig()
	if err != nil {
		panic(err.Error())
	}
	clientset, err = kubernetes.NewForConfig(config)
	if err != nil {
		panic(err.Error())
	}
}
func main() {
	//findPodByNamespace("default", "apollo-649c55f8c5-dsf5f")
	nss, err := clientset.CoreV1().Namespaces().List(metav1.ListOptions{})
	if err != nil {
		panic(err)
	}
	for true {
		for _, v := range nss.Items {
			findPodsByNamespace(v.Name)
		}
		time.Sleep(3 * time.Minute)
	}

}
func findPodsByNamespace(namespace string) {
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
					//fmt.Println("dingding")
					reg := regexp.MustCompile(`Error|Fail|Cannot`)
					if reg.MatchString(pods.Items[i].Status.ContainerStatuses[0].State.Waiting.Reason) == true {
						send(pods.Items[i].Namespace, pods.Items[i].Name, pods.Items[i].Status.ContainerStatuses[0].State.Waiting.Message)
					}
				}
			}

		}
	}
}
