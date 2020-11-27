package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/runtime/serializer"
	"k8s.io/client-go/kubernetes/scheme"
	"k8s.io/client-go/rest"

	"github.com/didierofrivia/go-fetch-products/pkg/apis/k8sinitiative.3scale.net/v1alpha1"
	products "github.com/didierofrivia/go-fetch-products/pkg/pages"
)

const homepageEndPoint = "/"
const productsEndPoint = "/products"

// StartWebServer the webserver
func StartWebServer() {
	http.HandleFunc(homepageEndPoint, handleHomePage)
	http.HandleFunc(productsEndPoint, handleProductsPage)
	port := os.Getenv("PORT")
	if len(port) == 0 {
		panic("Environment variable PORT is not set")
	}

	log.Printf("Starting web server to listen on endpoints [%s] and port %s",
		homepageEndPoint, port)
	if err := http.ListenAndServe(":"+port, nil); err != nil {
		panic(err)
	}
}

func handleHomePage(w http.ResponseWriter, r *http.Request) {
	urlPath := r.URL.Path
	log.Printf("Web request received on url path %s", urlPath)
	msg := "Hello world"
	_, err := w.Write([]byte(msg))
	if err != nil {
		fmt.Printf("Failed to write response, err: %s", err)
	}
}

func handleProductsPage(w http.ResponseWriter, r *http.Request) {
	// creates the in-cluster config
	config, err := rest.InClusterConfig()
	if err != nil {
		panic(err.Error())
	}

	v1alpha1.AddToScheme(scheme.Scheme)

	crdConfig := *config
	crdConfig.ContentConfig.GroupVersion = &schema.GroupVersion{Group: v1alpha1.GroupName, Version: v1alpha1.GroupVersion}
	crdConfig.APIPath = "/apis"
	crdConfig.NegotiatedSerializer = serializer.NewCodecFactory(scheme.Scheme)
	crdConfig.UserAgent = rest.DefaultKubernetesUserAgent()

	threescaleRestClient, err := rest.RESTClientFor(&crdConfig)
	if err != nil {
		panic(err)
	}
	result := v1alpha1.ProductList{}

	getErr := threescaleRestClient.
		Get().
		Namespace("k8sinitiative").
		Resource("products").
		Do().
		Into(&result)

	fmt.Printf("%d results found: %+v\n", len(result.Items), result)
	fmt.Println(getErr)
	products.Index(result)
	/* msg := fmt.Sprintf("%d results found: %+v\n", len(result.Items), result)
	_, writeErr := w.Write([]byte(msg))
	if writeErr != nil {
		fmt.Printf("Failed to write response, err: %s", writeErr)
	}
	*/
}

func main() {
	StartWebServer()
}
