package main

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	helmclient "github.com/mittwald/go-helm-client"
)

// https://pkg.go.dev/github.com/mittwald/go-helm-client#NewClientFromKubeConf

func main() {
	opt := &helmclient.Options{
		Namespace: "default", // Change this to the namespace you wish the client to operate in.
		// RepositoryCache:  "/tmp/.helmcache",
		// RepositoryConfig: "/tmp/.helmrepo",
		Debug:    true,
		Linting:  true,
		DebugLog: func(format string, v ...interface{}) {},
	}

	hc, err := helmclient.New(opt)
	if err != nil {
		fmt.Println(err)
	}

	router := gin.Default()
	router.GET("/", func(c *gin.Context) {
		// Define the chart to be installed
		chartSpec := helmclient.ChartSpec{
			ReleaseName: "test-server",
			ChartName:   "test-server-0.1.0.tgz",
			Namespace:   "default",
			UpgradeCRDs: true,
		}

		// Install a chart release.
		// Note that helmclient.Options.Namespace should ideally match the namespace in chartSpec.Namespace.
		if _, err := hc.InstallOrUpgradeChart(context.Background(), &chartSpec, nil); err != nil {
			c.String(http.StatusBadRequest, "Install helm chart failed")
		}
		c.String(http.StatusOK, "Welcome Gin Server")
	})

	srv := &http.Server{
		Addr:    ":8080",
		Handler: router,
	}

	// service connections
	if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Fatalf("listen: %s\n", err)
	}

	log.Println("Server exiting")
}
