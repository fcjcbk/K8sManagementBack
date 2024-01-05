package Pod

import (
	"fmt"
	"io"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/util/yaml"
	"net/http"
	"strconv"
	"test_kubernetes/client"
	v1 "test_kubernetes/model/server/v1"
)
import "github.com/gin-gonic/gin"

type PodController struct {
	cl   *client.Client
	pods []corev1.Pod
}

func NewPodController(cl *client.Client) *PodController {
	return &PodController{cl, nil}
}

func (p *PodController) GetPods(c *gin.Context) {

	var err error
	p.pods, err = p.cl.GetPods("default")

	if err != nil {
		// Todo: error message
		c.String(http.StatusBadGateway, "Get nodes fail")
		return
	}
	var resPods []v1.Pod

	for _, pod := range p.pods {

		fmt.Println("podname: " + pod.Name)
		var pd v1.Pod
		pd.Name = pod.Name

		pd.Namespace = pod.ObjectMeta.Namespace
		pd.Nodename = pod.Spec.NodeName

		pd.Status.Phase = string(pod.Status.Phase)
		pd.Status.HostIP = pod.Status.HostIP
		pd.Status.PodIP = pod.Status.PodIP
		if pod.Status.StartTime != nil {
			pd.Status.StartTime = pod.Status.StartTime.String()
		}

		for _, container := range pod.Spec.Containers {
			pd.Containers = append(pd.Containers, v1.Container{
				Name:  container.Name,
				Image: container.Image,
			})
		}
		fmt.Println(pd)
		resPods = append(resPods, pd)

	}

	c.JSON(http.StatusOK, resPods)
}

func (p *PodController) DeletePod(c *gin.Context) {
	// get array index from post
	id, _ := strconv.ParseInt(c.PostForm("deleteIndex"), 10, 64)
	fmt.Printf("delete index: %d\n", id)
	if id < 0 || id >= int64(len(p.pods)) {
		c.String(http.StatusOK, "Invalid delete")
		return
	}

	if err := p.cl.DeletePod(&p.pods[id]); err != nil {
		c.String(http.StatusOK, "Delete fail")
	}

	p.pods = append(p.pods[:id], p.pods[id+1:]...)
	//p.cl.DeletePod()
	c.String(http.StatusOK, "Delete success")

}

func (p *PodController) CreatePod(c *gin.Context) {

	file, _ := c.FormFile("file")
	src, err := file.Open()
	defer src.Close()

	nodeName := c.PostForm("nodeName")

	if err != nil {
		c.String(http.StatusOK, "can not solve the file")
		return
	}

	podYAML, err := io.ReadAll(src)
	pod := &corev1.Pod{}
	err = yaml.Unmarshal(podYAML, pod)
	if err != nil {
		c.String(http.StatusOK, err.Error())
		return
	}

	pod.Spec.NodeName = nodeName
	err = p.cl.CreatePod(pod)
	if err != nil {
		c.String(http.StatusOK, err.Error())
		return
	}

	c.String(http.StatusOK, "create pod success")
}

//func MockPod() []v1.Pod {
//	pods := []v1.Pod{
//		v1.Pod{
//			Name:      "my-nginx",
//			Namespace: "default",
//			Status: v1.Status{
//				Phase:     "running",
//				HostIP:    "1342",
//				PodIP:     "127.0.0.1",
//				StartTime: "2023",
//			},
//			Nodename: "node1",
//			Containers: []v1.Container{
//				v1.Container{
//					Name:  "nmae",
//					Image: "image",
//				},
//			},
//		},
//	}
//	return pods
//}
