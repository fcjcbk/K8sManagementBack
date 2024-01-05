package Node

import (
	"net/http"
	"test_kubernetes/client"
	"test_kubernetes/model/server/v1"
)
import "github.com/gin-gonic/gin"

type NodeController struct {
	cl *client.Client
}

func NewNodeController(cl *client.Client) *NodeController {
	return &NodeController{cl}
}

func (n *NodeController) GetNode(c *gin.Context) {
	// should be change in actual run
	// nodes := MockNode()

	nodes, err := n.cl.GetNodes()
	if err != nil {
		// Todo: err message should change
		c.String(http.StatusOK, err.Error())
		return
	}

	c.JSON(http.StatusOK, nodes)
}

func MockNode() []v1.Node {
	nodes := []v1.Node{
		{
			Name:           "1234",
			CrateTimeStamp: "1011-10-1",
			RemainCpu:      "10%",
			RemainMemory:   "10g",
		},
		{
			Name:           "123rere4",
			CrateTimeStamp: "10fd11-10-1",
			RemainCpu:      "1fd0%",
			RemainMemory:   "10fdg",
		},
	}
	return nodes
}
