// Copyright 2021 Ezequiel (Kimi) Aceto. All rights reserved.

package main

import (
	"flag"
	"github.com/eaceto/ReGraphQL/app"
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
	viper.BindPFlags(pflag.CommandLine)

	app, err := app.NewApp()
	if err != nil {
		klog.Fatal(err)
	}

	logAppConfiguration(app)

	routes, err := app.LoadRoutesFromFiles()
	if err != nil {
		klog.Fatal(err)
	}

	if len(routes) == 0 {
		klog.Fatalf("No routes available in config path: %s", app.RouterConfigsPath)
	}

	subRouter, err := app.GetServiceHTTPRouter(routes)
	if subRouter == nil {
		klog.Fatalf("Could not create HTTP router")
	}
	logEndpoints(subRouter)

	srv := &http.Server{
		Handler:      subRouter,
		Addr:         app.ServerAddr,
		ReadTimeout:  app.ServerReadTimeout,
		WriteTimeout: app.ServerWriteTimeout,
	}

	done := make(chan os.Signal, 1)
	signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)
	go func() {
		klog.Infoln("Server starting...")
		if app.DebugEnabled {
			klog.InfoS("Server props", "addr", app.ServerAddr, "readTimeout", app.ServerReadTimeout, "writeTimeout", app.ServerWriteTimeout)
		}
		if err := srv.ListenAndServe(); err != http.ErrServerClosed {
			klog.Fatalf("ListenAndServe(): %v", err)
		}
	}()

	<-done
	klog.Info("Server stopped.")
}

func logAppConfiguration(appConfiguration *app.App) {
	if !appConfiguration.DebugEnabled {
		return
	}
	klog.Warningln("Debug Enabled")

	klog.Infof("Config files: %s\n", appConfiguration.RouterConfigsPath)
	klog.Infof("Service path: %s\n", appConfiguration.ServicePath)

	if appConfiguration.TraceCallsEnabled {
		klog.Infoln("TraceCalls Enabled")
	}
}

func logEndpoints(r *mux.Router) {
	r.Walk(func(route *mux.Route, router *mux.Router, ancestors []*mux.Route) error {
		path, err := route.GetPathTemplate()
		if err != nil {
			return nil
		}
		methods, err := route.GetMethods()
		if err != nil {
			return nil
		}
		klog.Infof("%v %s\n", methods, path)
		return nil
	})
}
