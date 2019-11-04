package main

import (
	"fmt"
	"math/rand"
	"net/http"
	"os"
	"time"

	"k8s.io/apiserver/pkg/server/healthz"
	"k8s.io/component-base/logs"
	"k8s.io/klog"
	"k8s.io/kubernetes/cmd/cloud-controller-manager/app"
)

func init() {
	fmt.Fprintf(os.Stderr, "start: %v\n", "init")
	healthz.InstallHandler(http.DefaultServeMux)
}

func main() {
	fmt.Fprintf(os.Stderr, "start: %v\n", "main")
	rand.Seed(time.Now().UnixNano())

	command := app.NewCloudControllerManagerCommand()

	logs.InitLogs()
	defer logs.FlushLogs()

	klog.V(1).Infof("cloud-controller-manager version: %s", "1")

	if err := command.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}
}
