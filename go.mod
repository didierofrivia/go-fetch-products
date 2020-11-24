module github.com/didierofrivia/go-fetch-products

go 1.13

require (
	k8s.io/apimachinery v0.17.2
	k8s.io/client-go v0.17.2

)

replace github.com/didierofrivia/go-fetch-products/pkg/apis/k8sinitiative.3scale.net/v1alpha1 => ./pkg/apis/k8sinitiative.3scale.net/v1alpha1
