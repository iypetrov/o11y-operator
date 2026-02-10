package controller

import (
	"github.com/iypetrov/o11y-operator/controller/api/v1/observabilitylogsstorage"
	"k8s.io/apimachinery/pkg/runtime"
	utilruntime "k8s.io/apimachinery/pkg/util/runtime"
)

var (
	scheme = runtime.NewScheme()
)

func Scheme() *runtime.Scheme {
	return scheme
}

func init() {
	utilruntime.Must(observabilitylogsstorage.AddToScheme(scheme))
}
