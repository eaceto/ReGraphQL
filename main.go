/*
 * ReGraphQL - Proxy
 * This is the proxy service of project ReGraphQL
 *
 * Contact: ezequiel.aceto+regraphql@gmail.com
 */

package main

import (
	"flag"
	"github.com/eaceto/ReGraphQL/app"
	"github.com/eaceto/ReGraphQL/helpers"
	serviceAPI "github.com/eaceto/ReGraphQL/services"
	"github.com/gorilla/mux"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
	"k8s.io/klog/v2"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

/*
 * ReGraphQL
 * A simple (yet effective) REST / HTTP to GraphQL router
 *
 * API version: 0.0.1
 * Contact: ezequiel.aceto+regraphql@gmail.com
 */

func main() {
	klog.InitFlags(nil) // initializing the flags
	flag.Parse()
	defer klog.Flush()

	flag.String(app.ServerHostKey, app.ServerHostDefaultValue, "Server host")
	flag.Int(app.ServerPortKey, app.ServerPortDefaultValue, "Server port")

	pflag.CommandLine.AddGoFlagSet(flag.CommandLine)
	pflag.Parse()
	err := viper.BindPFlags(pflag.CommandLine)
	if err != nil {
		klog.Fatal(err)
	}

	router := mux.NewRouter()

	application, err := app.NewApplication(router)
	if err != nil {
		klog.Fatal(err)
	}

	serviceAPI.AddServiceEndpoints(application, router)

	if application.Configuration.DebugEnabled {
		helpers.LogEndpoints(router)
	}

	srv := &http.Server{
		Handler:      router,
		Addr:         application.Configuration.ServerAddr,
		ReadTimeout:  application.Configuration.ServerReadTimeout,
		WriteTimeout: application.Configuration.ServerWriteTimeout,
	}

	done := make(chan os.Signal, 1)
	signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)
	go func() {
		klog.InfoS("Server status", "status", "starting")
		if application.Configuration.DebugEnabled {
			klog.InfoS("Server props", "addr", application.Configuration.ServerAddr, "readTimeout", application.Configuration.ServerReadTimeout, "writeTimeout", application.Configuration.ServerWriteTimeout)
		}
		if err := srv.ListenAndServe(); err != http.ErrServerClosed {
			klog.InfoS("Server status", "status", "error")
			klog.Fatalf("ListenAndServe(): %v", err)
		}
	}()

	<-done
	klog.InfoS("Server status", "status", "stopped")
}
