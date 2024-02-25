package scheme

import (
	group01V1alpha1 "kubernetes-operator/api/project01/group01/v1alpha1"

	metaV1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/runtime/serializer"
	utilruntime "k8s.io/apimachinery/pkg/util/runtime"
)

var Scheme = runtime.NewScheme()
var Codecs = serializer.NewCodecFactory(Scheme)
var ParameterCodec = runtime.NewParameterCodec(Scheme)
var localSchemeBuilder = runtime.SchemeBuilder{
	group01V1alpha1.AddToScheme,
}
var AddToScheme = localSchemeBuilder.AddToScheme

func init() {
	metaV1.AddToGroupVersion(Scheme, schema.GroupVersion{Version: "v1alpha1"})
	utilruntime.Must(AddToScheme(Scheme))
}
