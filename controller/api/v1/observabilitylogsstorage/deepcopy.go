package observabilitylogsstorage

import "k8s.io/apimachinery/pkg/runtime"

func (in *Crd) DeepCopyInto(out *Crd) {
	out.TypeMeta = in.TypeMeta
	out.ObjectMeta = in.ObjectMeta

	out.Spec = CrdSpec{
		Memory: MemorySpec{
			TTL: in.Spec.Memory.TTL,
		},
		Iceberg: IcebergSpec{
			EndpointURL:     in.Spec.Iceberg.EndpointURL,
			DestinationPath: in.Spec.Iceberg.DestinationPath,
			S3Credentials: S3Credentials{
				AccessKeyIdSecretKey: SecretKeyRef{
					Name: in.Spec.Iceberg.S3Credentials.AccessKeyIdSecretKey.Name,
					Key:  in.Spec.Iceberg.S3Credentials.AccessKeyIdSecretKey.Key,
				},
				SecretAccessKeySecretKey: SecretKeyRef{
					Name: in.Spec.Iceberg.S3Credentials.SecretAccessKeySecretKey.Name,
					Key:  in.Spec.Iceberg.S3Credentials.SecretAccessKeySecretKey.Key,
				},
			},
		},
	}
}

func (in *Crd) DeepCopyObject() runtime.Object {
	out := Crd{}
	in.DeepCopyInto(&out)
	return &out
}

func (in *CrdList) DeepCopyObject() runtime.Object {
	out := CrdList{}
	out.TypeMeta = in.TypeMeta
	out.ListMeta = in.ListMeta

	if in.Items != nil {
		out.Items = make([]Crd, len(in.Items))
		for i := range in.Items {
			in.Items[i].DeepCopyInto(&out.Items[i])
		}
	}

	return &out
}
