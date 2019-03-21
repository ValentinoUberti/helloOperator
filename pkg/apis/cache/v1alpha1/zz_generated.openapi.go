// +build !ignore_autogenerated

// Code generated by openapi-gen. DO NOT EDIT.

// This file was autogenerated by openapi-gen. Do not edit it manually!

package v1alpha1

import (
	spec "github.com/go-openapi/spec"
	common "k8s.io/kube-openapi/pkg/common"
)

func GetOpenAPIDefinitions(ref common.ReferenceCallback) map[string]common.OpenAPIDefinition {
	return map[string]common.OpenAPIDefinition{
		"github.com/ValentinoUberti/hello-operator/pkg/apis/cache/v1alpha1.JedyKind":       schema_pkg_apis_cache_v1alpha1_JedyKind(ref),
		"github.com/ValentinoUberti/hello-operator/pkg/apis/cache/v1alpha1.JedyKindSpec":   schema_pkg_apis_cache_v1alpha1_JedyKindSpec(ref),
		"github.com/ValentinoUberti/hello-operator/pkg/apis/cache/v1alpha1.JedyKindStatus": schema_pkg_apis_cache_v1alpha1_JedyKindStatus(ref),
	}
}

func schema_pkg_apis_cache_v1alpha1_JedyKind(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "JedyKind is the Schema for the jedykinds API",
				Properties: map[string]spec.Schema{
					"kind": {
						SchemaProps: spec.SchemaProps{
							Description: "Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#types-kinds",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"apiVersion": {
						SchemaProps: spec.SchemaProps{
							Description: "APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#resources",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"metadata": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("k8s.io/apimachinery/pkg/apis/meta/v1.ObjectMeta"),
						},
					},
					"spec": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("github.com/ValentinoUberti/hello-operator/pkg/apis/cache/v1alpha1.JedyKindSpec"),
						},
					},
					"status": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("github.com/ValentinoUberti/hello-operator/pkg/apis/cache/v1alpha1.JedyKindStatus"),
						},
					},
				},
			},
		},
		Dependencies: []string{
			"github.com/ValentinoUberti/hello-operator/pkg/apis/cache/v1alpha1.JedyKindSpec", "github.com/ValentinoUberti/hello-operator/pkg/apis/cache/v1alpha1.JedyKindStatus", "k8s.io/apimachinery/pkg/apis/meta/v1.ObjectMeta"},
	}
}

func schema_pkg_apis_cache_v1alpha1_JedyKindSpec(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "JedyKindSpec defines the desired state of JedyKind",
				Properties:  map[string]spec.Schema{},
			},
		},
		Dependencies: []string{},
	}
}

func schema_pkg_apis_cache_v1alpha1_JedyKindStatus(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "JedyKindStatus defines the observed state of JedyKind",
				Properties:  map[string]spec.Schema{},
			},
		},
		Dependencies: []string{},
	}
}