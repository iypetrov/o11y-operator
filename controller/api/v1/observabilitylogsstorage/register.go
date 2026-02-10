package observabilitylogsstorage

import (
	apiv1 "github.com/iypetrov/o11y-operator/controller/api/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
)

var (
	SchemeBuilder = runtime.NewSchemeBuilder(addKnownTypes)
	AddToScheme   = SchemeBuilder.AddToScheme
)

func addKnownTypes(scheme *runtime.Scheme) error {
	scheme.AddKnownTypes(
		SchemaGroupVersion,
		&Crd{},
		&CrdList{},
	)

	metav1.AddToGroupVersion(scheme, SchemaGroupVersion)
	return nil
}

var SchemaGroupVersion = schema.GroupVersion{
	Group:   apiv1.GroupName,
	Version: apiv1.GroupVersion,
}
