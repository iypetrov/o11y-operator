package main

import (
	"os"

	"github.com/iypetrov/o11y-operator/controller"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/log/zap"
)

func main() {
	logger := zap.New()
	ctrl.SetLogger(logger)
	log := ctrl.Log.WithName("main")
	log.Info("set up manager")

	mgr, err := ctrl.NewManager(ctrl.GetConfigOrDie(), ctrl.Options{
		Scheme: controller.Scheme(),
	})
	if err != nil {
		log.Error(err, "unable to start manager")
		return
	}

	ols := controller.ObservabilityLogsStorageReconciler{
		Client: mgr.GetClient(),
		Scheme: mgr.GetScheme(),
		Log:    log.WithName("observability-logs-storage"),
	}
	err = ols.SetupWithManager(mgr)
	if err != nil {
		log.Error(err, "unable to set up ObservabilityLogsStorageReconciler")
		return
	}

	ctx := ctrl.SetupSignalHandler()
	if err = mgr.Start(ctx); err != nil {
		log.Error(err, "problem running manager")
		os.Exit(1)
	}
}
