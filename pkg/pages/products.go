package products

import (
	"html/template"
	"log"
	"os"

	"github.com/didierofrivia/go-fetch-products/pkg/apis/k8sinitiative.3scale.net/v1alpha1"
)

// Index prints the list of products
func Index(productList v1alpha1.ProductList) {
	const tpl = `
	<!DOCTYPE html>
	<html>
		<head>
			<meta charset="UTF-8">
			<title>Products</title>
		</head>
		<body>
			{{range .Items}}<div>{{ .Spec }}</div>{{else}}<div><strong>no rows</strong></div>{{end}}
		</body>
	</html>`

	check := func(err error) {
		if err != nil {
			log.Fatal(err)
		}
	}
	t, err := template.New("webpage").Parse(tpl)
	check(err)

	err = t.Execute(os.Stdout, productList)
	check(err)

	noItems := v1alpha1.ProductList{}

	err = t.Execute(os.Stdout, noItems)
	check(err)
}
