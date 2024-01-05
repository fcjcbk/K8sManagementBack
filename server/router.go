package server

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"test_kubernetes/client"
	"test_kubernetes/server/controller/v1/Node"
	"test_kubernetes/server/controller/v1/Pod"
	"time"
)

func Run() {

	cl := client.MakeClient()

	router := gin.Default()

	server := http.Server{
		Addr:              ":8080",
		Handler:           router,
		ReadHeaderTimeout: 10 * time.Second,
		WriteTimeout:      10 * time.Second,
		MaxHeaderBytes:    1 << 20,
	}

	router.Use(Secure)
	router.Use(Options)
	router.Use(NoCache)

	router.NoRoute(func(c *gin.Context) {
		c.String(http.StatusOK, "NoRoute")
	})

	node := Node.NewNodeController(cl)
	pod := Pod.NewPodController(cl)

	router.GET("/v1/nodes", node.GetNode)
	router.GET("/v1/pods", pod.GetPods)

	router.MaxMultipartMemory = 8 << 20
	router.POST("/v1/CreatePod", pod.CreatePod)

	router.DELETE("/v1/DeletePod", pod.DeletePod)

	go func() {
		fmt.Println("Start server!")
		if err := server.ListenAndServe(); err != nil {
			log.Panicln(err)
		}
	}()

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	go func() {
		if err := myPingServer(ctx); err != nil {
			fmt.Println(err)
		}
	}()

	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	<-quit

	fmt.Println("Receive system interrupt! Start Kill!")

	if err := server.Shutdown(ctx); err != nil {
		fmt.Println(err)
	}
}

func myPingServer(ctx context.Context) error {
	for {
		req, err := http.NewRequestWithContext(ctx, http.MethodGet, "http://127.0.0.1:8080/", nil)
		if err != nil {
			return err
		}

		resq, err := http.DefaultClient.Do(req)
		if err == nil && resq.StatusCode == http.StatusOK {
			resq.Body.Close()
			return nil
		}

		time.Sleep(1 * time.Second)

		select {
		case <-ctx.Done():
			return fmt.Errorf("can not ping the server in given time")
		default:
		}
	}
}

func UPLoadFile(c *gin.Context) {
	// test command:
	// curl -X POST http://localhost:8080/upload -F "file=@.\script.sh" -H "Content-Type: multipart/form-data"

	file, err := c.FormFile("file")
	if err != nil {
		log.Println(err)
		c.String(http.StatusForbidden, "wrong")
		return
	}
	log.Println(file.Filename)

	c.SaveUploadedFile(file, fmt.Sprintf("./%s", file.Filename))

	c.String(http.StatusOK, fmt.Sprintln("upload success"))
}

// may be fix
func Secure(c *gin.Context) {
	c.Header("Access-Control-Allow-Origin", "*")
	c.Header("X-Frame-Options", "DENY")
	c.Header("X-Content-Type-Options", "nosniff")
	c.Header("X-XSS-Protection", "1; mode=block")
	if c.Request.TLS != nil {
		c.Header("Strict-Transport-Security", "max-age=31536000")
	}

	// Also consider adding Content-Security-Policy headers
	// c.Header("Content-Security-Policy", "script-src 'self' https://cdnjs.cloudflare.com")
}

// NoCache is a middleware function that appends headers
// to prevent the client from caching the HTTP response.
func NoCache(c *gin.Context) {
	c.Header("Cache-Control", "no-cache, no-store, max-age=0, must-revalidate, value")
	c.Header("Expires", "Thu, 01 Jan 1970 00:00:00 GMT")
	c.Header("Last-Modified", time.Now().UTC().Format(http.TimeFormat))
	c.Next()
}

// Options is a middleware function that appends headers
// for options requests and aborts then exits the middleware
// chain and ends the request.
func Options(c *gin.Context) {
	if c.Request.Method != "OPTIONS" {
		c.Next()
	} else {
		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Methods", "GET,POST,PUT,PATCH,DELETE,OPTIONS")
		c.Header("Access-Control-Allow-Headers", "authorization, origin, content-type, accept")
		c.Header("Allow", "HEAD,GET,POST,PUT,PATCH,DELETE,OPTIONS")
		c.Header("Content-Type", "application/json")
		c.AbortWithStatus(200)
	}
}
