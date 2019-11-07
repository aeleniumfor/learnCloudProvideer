package main

import (
	"fmt"
	"math/rand"
	"net/http"
	"os"
	"time"

	_ "github.com/learnCloudProvideer/provider"

	"k8s.io/apiserver/pkg/server/healthz"
	"k8s.io/component-base/logs"
	"k8s.io/klog"
	"k8s.io/kubernetes/cmd/cloud-controller-manager/app"

	_ "k8s.io/kubernetes/pkg/client/metrics/prometheus"
	_ "k8s.io/kubernetes/pkg/version/prometheus"
)

func init() {
	fmt.Fprintf(os.Stderr, "start: %v\n", "init")
	healthz.InstallHandler(http.DefaultServeMux)
}

func main() {
	fmt.Fprintf(os.Stderr, "start: %v\n", "main")
	rand.Seed(time.Now().UnixNano())

	fmt.Fprintf(os.Stderr, "start: %v\n", "command")
	command := app.NewCloudControllerManagerCommand()


	fmt.Fprintf(os.Stderr, "start: %v\n", "init")
	logs.InitLogs()
	defer logs.FlushLogs()

	fmt.Fprintf(os.Stderr, "start: %v\n", "klog.info")
	klog.V(1).Infof("cloud-controller-manager version: %s", "1")

	if err := command.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}
}
