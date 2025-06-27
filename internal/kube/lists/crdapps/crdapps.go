package crdapps

import (
	"context"
	"fmt"

	"github.com/mahyarmirrashed/hajimari/internal/config"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/client-go/dynamic"
)

// List struct is used to list ingresses
type List struct {
	appConfig config.Config
	err       error // Used for forwarding errors
	items     []unstructured.Unstructured
	dynClient dynamic.Interface
}

var appResource = schema.GroupVersionResource{
	Group:    "hajimari.io",
	Version:  "v1alpha1",
	Resource: "applications",
}

// NewList creates an List object that you can use to query ingresses
func NewList(dynClient dynamic.Interface, appConfig config.Config, items ...unstructured.Unstructured) *List {
	return &List{
		dynClient: dynClient,
		appConfig: appConfig,
		items:     items,
	}
}

// Populate function returns a list of ingresses
func (il *List) Populate(namespaces ...string) *List {
	if il.dynClient == nil {
		il.err = fmt.Errorf("dynamic Kubernetes client is nil")
		return il
	}

	for _, namespace := range namespaces {
		apps, err := il.dynClient.
			Resource(appResource).
			Namespace(namespace).
			List(context.Background(), metav1.ListOptions{})
		if err != nil {
			il.err = err
			return il
		}
		il.items = append(il.items, apps.Items...)
	}

	return il
}

// Get function returns the ingresses currently present in List
func (il *List) Get() ([]unstructured.Unstructured, error) {
	return il.items, il.err
}
