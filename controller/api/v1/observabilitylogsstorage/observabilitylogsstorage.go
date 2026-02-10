package observabilitylogsstorage

import metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

type Crd struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata"`
	Spec              CrdSpec `json:"spec"`
}

type CrdList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata"`
	Items           []Crd `json:"items"`
}

type CrdSpec struct {
	Memory  MemorySpec  `json:"memory"`
	Iceberg IcebergSpec `json:"iceberg"`
}

type MemorySpec struct {
	TTL metav1.Duration `json:"ttl,omitempty"`
}

type IcebergSpec struct {
	EndpointURL     string        `json:"endpointURL,omitempty"`
	DestinationPath string        `json:"destinationPath,omitempty"`
	S3Credentials   S3Credentials `json:"s3Credentials"`
}

type S3Credentials struct {
	AccessKeyIdSecretKey     SecretKeyRef `json:"accessKeyIdSecretKey"`
	SecretAccessKeySecretKey SecretKeyRef `json:"secretAccessKeySecretKey"`
}

type SecretKeyRef struct {
	Name string `json:"name,omitempty"`
	Key  string `json:"key,omitempty"`
}
