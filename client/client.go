package client

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"log"
	"os"
	"reflect"
	v1 "test_kubernetes/model/server/v1"

	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/util/yaml"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
)

type Client struct {
	ClientSet *kubernetes.Clientset
}

func MakeClient() *Client {
	// Todo: build config may be change
	// Path to the kubeconfig file
	config, err := clientcmd.BuildConfigFromFlags("", "./config")
	if err != nil {
		panic(err.Error())
	}

	// create the clientset
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		panic(err.Error())
	}

	client := Client{ClientSet: clientset}
	return &client
}

func (cl *Client) GetNamespace() []string {
	namespaces, err := cl.ClientSet.CoreV1().Namespaces().List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		panic(err.Error())
	}

	var res []string

	for _, ns := range namespaces.Items {
		fmt.Println(ns.Name)
		res = append(res, ns.Name)
	}
	return res
}

func (cl *Client) GetNodes() ([]v1.Node, error) {
	nodes, err := cl.ClientSet.CoreV1().Nodes().List(context.TODO(), metav1.ListOptions{})

	if err != nil {
		log.Println(err)
		return nil, err
	}

	// name address createTimestamp
	var res []v1.Node
	for _, node := range nodes.Items {
		fmt.Printf("Name: %s\n", node.ObjectMeta.Name)
		fmt.Printf("Node Name: %s\n", node.CreationTimestamp)

		fmt.Printf("label: %v", node.Status.Allocatable.Cpu().String())
		fmt.Printf("label: %v", node.Status.Allocatable.Memory().String())

		// printStructKeyAndValue(node)
		res = append(res, v1.Node{
			Name:           node.ObjectMeta.Name,
			CrateTimeStamp: node.CreationTimestamp.String(),
			RemainCpu:      node.Status.Allocatable.Cpu().String(),
			RemainMemory:   node.Status.Allocatable.Memory().String(),
		})
	}

	return res, nil
}

func (cl *Client) UseFileCreatePod(filename string) bool {
	podYAML, err := os.ReadFile(filename)
	if err != nil {
		log.Println(err)
		return false
	}

	//file, _ := c.FormFile("file")
	//src, err := file.Open()
	//defer src.Close()
	//if err != nil {
	//
	//}
	//
	//content, err := io.ReadAll(src)

	pod := &corev1.Pod{}
	err = yaml.Unmarshal(podYAML, pod)
	if err != nil {
		log.Println(err)
		return false
	}

	return true
	//cl.CreatePod(pod)
	//cl.DeletePod(pod)
	//time.Sleep(10 * time.Second)
	//cl.getPods()
}

func (cl *Client) GetPodLog(pod *corev1.Pod, namespace string) string {
	podLogOpts := &corev1.PodLogOptions{}
	req := cl.ClientSet.CoreV1().Pods(namespace).GetLogs(pod.ObjectMeta.Name, podLogOpts)
	podLogs, err := req.Stream(context.TODO())
	if err != nil {
		panic(err.Error())
	}
	defer podLogs.Close()
	buf := new(bytes.Buffer)
	_, err = io.Copy(buf, podLogs)
	if err != nil {
		panic(err.Error())
	}
	fmt.Println(buf.String())
	return buf.String()
}

func (cl *Client) CreatePod(pod *corev1.Pod) error {
	fmt.Println("Start to create pod!")
	result, err := cl.ClientSet.CoreV1().Pods(pod.Namespace).Create(context.TODO(), pod, metav1.CreateOptions{})
	if err != nil {
		//panic(err.Error())
		return err
	}
	fmt.Printf("Created pod %q in namespace %q\n", result.GetObjectMeta().GetName(), result.GetObjectMeta().GetNamespace())
	return nil
}

func (cl *Client) DeletePod(pod *corev1.Pod) error {
	fmt.Println("Start to delete pod!")
	err := cl.ClientSet.CoreV1().Pods(pod.Namespace).Delete(context.TODO(), pod.Name, metav1.DeleteOptions{})
	if err != nil {
		//panic(err.Error())
		return err
	}

	fmt.Println("Delete pod success")
	return nil
}

func (cl *Client) GetPods(namespace string) ([]corev1.Pod, error) {

	pods, err := cl.ClientSet.CoreV1().Pods(namespace).List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		//panic(err.Error())
		return nil, err
	}

	fmt.Printf("There are %d pods in the cluster\n", len(pods.Items))

	return pods.Items, nil
	//fmt.Printf("Name: %s\n", pod.ObjectMeta.Name)
	//fmt.Printf("Namespace: %s\n", pod.ObjectMeta.Namespace)
	//fmt.Printf("Status: %s\n", pod.Status.Phase)
	//fmt.Println("Containers:")
	//for _, container := range pod.Spec.Containers {
	//	fmt.Printf("- Name: %s\n", container.Name)
	//	fmt.Printf("  Image: %s\n", container.Image)
	//}
	//fmt.Println("------")
	//printStructKeyAndValue(pod.ObjectMeta)
	//printStructKeyAndValue(pod.Status)

	//// Examples for error handling:
	//// - Use helper functions like e.g. errors.IsNotFound()
	//// - And/or cast to StatusError and use its properties like e.g. ErrStatus.Message
	//pod := "my-pod"
	//_, err = cl.ClientSet.CoreV1().Pods(namespace).Get(context.TODO(), pod, metav1.GetOptions{})
	//if errors.IsNotFound(err) {
	//	fmt.Printf("Pod %s in namespace %s not found\n", pod, namespace)
	//} else if statusError, isStatus := err.(*errors.StatusError); isStatus {
	//	fmt.Printf("Error getting pod %s in namespace %s: %v\n",
	//		pod, namespace, statusError.ErrStatus.Message)
	//} else if err != nil {
	//	panic(err.Error())
	//} else {
	//	fmt.Printf("Found pod %s in namespace %s\n", pod, namespace)
	//}

}

func printStructKeyAndValue(s interface{}) {
	v := reflect.ValueOf(s)

	for i := 0; i < v.NumField(); i++ {
		fmt.Printf("%s : %v\n", v.Type().Field(i).Name, v.Field(i).Interface())
	}
}
