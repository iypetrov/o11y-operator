package controller

import (
	"context"

	"github.com/go-logr/logr"
	"github.com/iypetrov/o11y-operator/controller/api/v1/observabilitylogsstorage"
	kerr "k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/runtime"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

type ObservabilityLogsStorageReconciler struct {
	Scheme *runtime.Scheme
	Log    logr.Logger

	client.Client
}

func (r *ObservabilityLogsStorageReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&observabilitylogsstorage.Crd{}).
		Complete(r)
}

func (r *ObservabilityLogsStorageReconciler) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {
	r.Log.Info("reconciling ObservabilityLogsStorage", "name", req.Name, "namespace", req.Namespace)

	wp := observabilitylogsstorage.Crd{}

	err := r.Client.Get(ctx, req.NamespacedName, &wp)
	if err != nil && kerr.IsNotFound(err) {
		return ctrl.Result{}, nil
	} else if err != nil {
		return ctrl.Result{}, err
	}

	// reconciliation logic
	r.Log.Info("ObservabilityLogsStorage reconciled")

	return ctrl.Result{}, nil
}
