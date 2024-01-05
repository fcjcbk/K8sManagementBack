package main

import (
	"fmt"
	"reflect"
	"test_kubernetes/server"
)

func main() {

	server.Run()

	// Create the pod object
	// pod := &corev1.Pod{
	// 	ObjectMeta: metav1.ObjectMeta{
	// 		Name:      "my-pod",
	// 		Namespace: "default",
	// 	},
	// 	Spec: corev1.PodSpec{
	// 		Containers: []corev1.Container{
	// 			{
	// 				Name:  "my-container",
	// 				Image: "nginx:latest",
	// 			},
	// 		},
	// 		NodeName: "",
	// 	},
	// }

	// create pod
	//client.getPods()
	//client.CreatePod(pod)
	//client.UseFileCreatePod("./pod.yaml")

	// delete pod
	//client.DeletePod(pod)

}

func printStructKeyAndValue(s interface{}) {
	v := reflect.ValueOf(s)

	for i := 0; i < v.NumField(); i++ {
		fmt.Printf("%s : %v\n", v.Type().Field(i).Name, v.Field(i).Interface())
	}
}
