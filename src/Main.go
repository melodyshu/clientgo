package main

import (
	"context"
	"fmt"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
)

func main() {
	config, err := clientcmd.BuildConfigFromFlags("", "admin.conf")
	if err != nil {
		panic(err)
	}
	client, _ := kubernetes.NewForConfig(config)
	ctx := context.Background()
	pods, err := client.CoreV1().Pods("default").List(ctx, metav1.ListOptions{})
	if err != nil {
		fmt.Println(err)
		return
	}

	for _, v := range pods.Items {
		fmt.Printf(" 命名空间是：%v\n pod名字：%v\n IP：%v\n\n", v.Namespace, v.Name, v.Status.PodIP)
	}
}
