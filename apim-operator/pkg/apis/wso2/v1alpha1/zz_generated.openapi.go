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
		"github.com/apim-crd/apim-operator/pkg/apis/wso2/v1alpha1.API":       schema_pkg_apis_wso2_v1alpha1_API(ref),
		"github.com/apim-crd/apim-operator/pkg/apis/wso2/v1alpha1.APISpec":   schema_pkg_apis_wso2_v1alpha1_APISpec(ref),
		"github.com/apim-crd/apim-operator/pkg/apis/wso2/v1alpha1.APIStatus": schema_pkg_apis_wso2_v1alpha1_APIStatus(ref),
	}
}

func schema_pkg_apis_wso2_v1alpha1_API(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "API is the Schema for the apis API",
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
							Ref: ref("github.com/apim-crd/apim-operator/pkg/apis/wso2/v1alpha1.APISpec"),
						},
					},
					"status": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("github.com/apim-crd/apim-operator/pkg/apis/wso2/v1alpha1.APIStatus"),
						},
					},
				},
			},
		},
		Dependencies: []string{
			"github.com/apim-crd/apim-operator/pkg/apis/wso2/v1alpha1.APISpec", "github.com/apim-crd/apim-operator/pkg/apis/wso2/v1alpha1.APIStatus", "k8s.io/apimachinery/pkg/apis/meta/v1.ObjectMeta"},
	}
}

func schema_pkg_apis_wso2_v1alpha1_APISpec(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "APISpec defines the desired state of API",
				Properties: map[string]spec.Schema{
					"name": {
						SchemaProps: spec.SchemaProps{
							Description: "INSERT ADDITIONAL SPEC FIELDS - desired state of cluster Important: Run \"operator-sdk generate k8s\" to regenerate code after modifying this file Add custom validation using kubebuilder tags: https://book.kubebuilder.io/beyond_basics/generating_crd.html",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"context": {
						SchemaProps: spec.SchemaProps{
							Type:   []string{"string"},
							Format: "",
						},
					},
					"version": {
						SchemaProps: spec.SchemaProps{
							Type:   []string{"string"},
							Format: "",
						},
					},
					"description": {
						SchemaProps: spec.SchemaProps{
							Type:   []string{"string"},
							Format: "",
						},
					},
					"tags": {
						SchemaProps: spec.SchemaProps{
							Type: []string{"array"},
							Items: &spec.SchemaOrArray{
								Schema: &spec.Schema{
									SchemaProps: spec.SchemaProps{
										Type:   []string{"string"},
										Format: "",
									},
								},
							},
						},
					},
					"endpoints": {
						SchemaProps: spec.SchemaProps{
							Type: []string{"array"},
							Items: &spec.SchemaOrArray{
								Schema: &spec.Schema{
									SchemaProps: spec.SchemaProps{
										Ref: ref("github.com/apim-crd/apim-operator/pkg/apis/wso2/v1alpha1.Endpoint"),
									},
								},
							},
						},
					},
					"requestInterceptor": {
						SchemaProps: spec.SchemaProps{
							Type:   []string{"string"},
							Format: "",
						},
					},
					"responseInterceptor": {
						SchemaProps: spec.SchemaProps{
							Type:   []string{"string"},
							Format: "",
						},
					},
					"authorizationHeader": {
						SchemaProps: spec.SchemaProps{
							Type:   []string{"string"},
							Format: "",
						},
					},
					"labels": {
						SchemaProps: spec.SchemaProps{
							Type: []string{"array"},
							Items: &spec.SchemaOrArray{
								Schema: &spec.Schema{
									SchemaProps: spec.SchemaProps{
										Type:   []string{"string"},
										Format: "",
									},
								},
							},
						},
					},
					"urlPatterns": {
						SchemaProps: spec.SchemaProps{
							Type: []string{"array"},
							Items: &spec.SchemaOrArray{
								Schema: &spec.Schema{
									SchemaProps: spec.SchemaProps{
										Ref: ref("github.com/apim-crd/apim-operator/pkg/apis/wso2/v1alpha1.URLPattern"),
									},
								},
							},
						},
					},
					"security": {
						SchemaProps: spec.SchemaProps{
							Type: []string{"array"},
							Items: &spec.SchemaOrArray{
								Schema: &spec.Schema{
									SchemaProps: spec.SchemaProps{
										Type:   []string{"string"},
										Format: "",
									},
								},
							},
						},
					},
					"subscriptionTiers": {
						SchemaProps: spec.SchemaProps{
							Type: []string{"array"},
							Items: &spec.SchemaOrArray{
								Schema: &spec.Schema{
									SchemaProps: spec.SchemaProps{
										Type:   []string{"string"},
										Format: "",
									},
								},
							},
						},
					},
					"advancedThrottlingPolicy": {
						SchemaProps: spec.SchemaProps{
							Type:   []string{"string"},
							Format: "",
						},
					},
					"businessInformation": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("github.com/apim-crd/apim-operator/pkg/apis/wso2/v1alpha1.BusinessInformation"),
						},
					},
					"apiProperties": {
						SchemaProps: spec.SchemaProps{
							Type: []string{"array"},
							Items: &spec.SchemaOrArray{
								Schema: &spec.Schema{
									SchemaProps: spec.SchemaProps{
										Ref: ref("github.com/apim-crd/apim-operator/pkg/apis/wso2/v1alpha1.APIProperty"),
									},
								},
							},
						},
					},
					"mode": {
						SchemaProps: spec.SchemaProps{
							Type:   []string{"string"},
							Format: "",
						},
					},
					"replicaCount": {
						SchemaProps: spec.SchemaProps{
							Type:   []string{"integer"},
							Format: "int32",
						},
					},
				},
				Required: []string{"name", "context", "version", "description", "tags", "endpoints", "requestInterceptor", "responseInterceptor", "authorizationHeader", "labels", "urlPatterns", "security", "subscriptionTiers", "advancedThrottlingPolicy", "businessInformation", "apiProperties", "mode", "replicaCount"},
			},
		},
		Dependencies: []string{
			"github.com/apim-crd/apim-operator/pkg/apis/wso2/v1alpha1.APIProperty", "github.com/apim-crd/apim-operator/pkg/apis/wso2/v1alpha1.BusinessInformation", "github.com/apim-crd/apim-operator/pkg/apis/wso2/v1alpha1.Endpoint", "github.com/apim-crd/apim-operator/pkg/apis/wso2/v1alpha1.URLPattern"},
	}
}

func schema_pkg_apis_wso2_v1alpha1_APIStatus(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "APIStatus defines the observed state of API",
				Properties:  map[string]spec.Schema{},
			},
		},
		Dependencies: []string{},
	}
}
