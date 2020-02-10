// +build !ignore_autogenerated

// This file was autogenerated by openapi-gen. Do not edit it manually!

package v1beta1

import (
	spec "github.com/go-openapi/spec"
	common "k8s.io/kube-openapi/pkg/common"
)

func GetOpenAPIDefinitions(ref common.ReferenceCallback) map[string]common.OpenAPIDefinition {
	return map[string]common.OpenAPIDefinition{
		"./pkg/apis/openliberty/v1beta1.OpenLibertyApplication":               schema_pkg_apis_openliberty_v1beta1_OpenLibertyApplication(ref),
		"./pkg/apis/openliberty/v1beta1.OpenLibertyApplicationAutoScaling":    schema_pkg_apis_openliberty_v1beta1_OpenLibertyApplicationAutoScaling(ref),
		"./pkg/apis/openliberty/v1beta1.OpenLibertyApplicationService":        schema_pkg_apis_openliberty_v1beta1_OpenLibertyApplicationService(ref),
		"./pkg/apis/openliberty/v1beta1.OpenLibertyApplicationServiceability": schema_pkg_apis_openliberty_v1beta1_OpenLibertyApplicationServiceability(ref),
		"./pkg/apis/openliberty/v1beta1.OpenLibertyApplicationSpec":           schema_pkg_apis_openliberty_v1beta1_OpenLibertyApplicationSpec(ref),
		"./pkg/apis/openliberty/v1beta1.OpenLibertyApplicationStatus":         schema_pkg_apis_openliberty_v1beta1_OpenLibertyApplicationStatus(ref),
		"./pkg/apis/openliberty/v1beta1.OpenLibertyApplicationStorage":        schema_pkg_apis_openliberty_v1beta1_OpenLibertyApplicationStorage(ref),
		"./pkg/apis/openliberty/v1beta1.OpenLibertyDump":                      schema_pkg_apis_openliberty_v1beta1_OpenLibertyDump(ref),
		"./pkg/apis/openliberty/v1beta1.OpenLibertyDumpSpec":                  schema_pkg_apis_openliberty_v1beta1_OpenLibertyDumpSpec(ref),
		"./pkg/apis/openliberty/v1beta1.OpenLibertyDumpStatus":                schema_pkg_apis_openliberty_v1beta1_OpenLibertyDumpStatus(ref),
		"./pkg/apis/openliberty/v1beta1.OpenLibertyTrace":                     schema_pkg_apis_openliberty_v1beta1_OpenLibertyTrace(ref),
		"./pkg/apis/openliberty/v1beta1.OpenLibertyTraceSpec":                 schema_pkg_apis_openliberty_v1beta1_OpenLibertyTraceSpec(ref),
		"./pkg/apis/openliberty/v1beta1.OpenLibertyTraceStatus":               schema_pkg_apis_openliberty_v1beta1_OpenLibertyTraceStatus(ref),
		"./pkg/apis/openliberty/v1beta1.OperatedResource":                     schema_pkg_apis_openliberty_v1beta1_OperatedResource(ref),
		"./pkg/apis/openliberty/v1beta1.OperationStatusCondition":             schema_pkg_apis_openliberty_v1beta1_OperationStatusCondition(ref),
		"./pkg/apis/openliberty/v1beta1.ServiceBindingConsumes":               schema_pkg_apis_openliberty_v1beta1_ServiceBindingConsumes(ref),
		"./pkg/apis/openliberty/v1beta1.ServiceBindingProvides":               schema_pkg_apis_openliberty_v1beta1_ServiceBindingProvides(ref),
		"./pkg/apis/openliberty/v1beta1.StatusCondition":                      schema_pkg_apis_openliberty_v1beta1_StatusCondition(ref),
	}
}

func schema_pkg_apis_openliberty_v1beta1_OpenLibertyApplication(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "OpenLibertyApplication is the Schema for the OpenLibertyApplications API",
				Type:        []string{"object"},
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
							Ref: ref("./pkg/apis/openliberty/v1beta1.OpenLibertyApplicationSpec"),
						},
					},
					"status": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("./pkg/apis/openliberty/v1beta1.OpenLibertyApplicationStatus"),
						},
					},
				},
			},
		},
		Dependencies: []string{
			"./pkg/apis/openliberty/v1beta1.OpenLibertyApplicationSpec", "./pkg/apis/openliberty/v1beta1.OpenLibertyApplicationStatus", "k8s.io/apimachinery/pkg/apis/meta/v1.ObjectMeta"},
	}
}

func schema_pkg_apis_openliberty_v1beta1_OpenLibertyApplicationAutoScaling(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "OpenLibertyApplicationAutoScaling ...",
				Type:        []string{"object"},
				Properties: map[string]spec.Schema{
					"targetCPUUtilizationPercentage": {
						SchemaProps: spec.SchemaProps{
							Type:   []string{"integer"},
							Format: "int32",
						},
					},
					"minReplicas": {
						SchemaProps: spec.SchemaProps{
							Type:   []string{"integer"},
							Format: "int32",
						},
					},
					"maxReplicas": {
						SchemaProps: spec.SchemaProps{
							Type:   []string{"integer"},
							Format: "int32",
						},
					},
				},
			},
		},
	}
}

func schema_pkg_apis_openliberty_v1beta1_OpenLibertyApplicationService(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "OpenLibertyApplicationService ...",
				Type:        []string{"object"},
				Properties: map[string]spec.Schema{
					"type": {
						SchemaProps: spec.SchemaProps{
							Type:   []string{"string"},
							Format: "",
						},
					},
					"port": {
						SchemaProps: spec.SchemaProps{
							Type:   []string{"integer"},
							Format: "int32",
						},
					},
					"annotations": {
						SchemaProps: spec.SchemaProps{
							Type: []string{"object"},
							AdditionalProperties: &spec.SchemaOrBool{
								Allows: true,
								Schema: &spec.Schema{
									SchemaProps: spec.SchemaProps{
										Type:   []string{"string"},
										Format: "",
									},
								},
							},
						},
					},
					"consumes": {
						VendorExtensible: spec.VendorExtensible{
							Extensions: spec.Extensions{
								"x-kubernetes-list-type": "atomic",
							},
						},
						SchemaProps: spec.SchemaProps{
							Type: []string{"array"},
							Items: &spec.SchemaOrArray{
								Schema: &spec.Schema{
									SchemaProps: spec.SchemaProps{
										Ref: ref("./pkg/apis/openliberty/v1beta1.ServiceBindingConsumes"),
									},
								},
							},
						},
					},
					"provides": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("./pkg/apis/openliberty/v1beta1.ServiceBindingProvides"),
						},
					},
				},
			},
		},
		Dependencies: []string{
			"./pkg/apis/openliberty/v1beta1.ServiceBindingConsumes", "./pkg/apis/openliberty/v1beta1.ServiceBindingProvides"},
	}
}

func schema_pkg_apis_openliberty_v1beta1_OpenLibertyApplicationServiceability(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "OpenLibertyApplicationServiceability ...",
				Type:        []string{"object"},
				Properties: map[string]spec.Schema{
					"size": {
						SchemaProps: spec.SchemaProps{
							Type:   []string{"string"},
							Format: "",
						},
					},
					"volumeClaimName": {
						SchemaProps: spec.SchemaProps{
							Type:   []string{"string"},
							Format: "",
						},
					},
				},
			},
		},
	}
}

func schema_pkg_apis_openliberty_v1beta1_OpenLibertyApplicationSpec(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "OpenLibertyApplicationSpec defines the desired state of OpenLibertyApplication",
				Type:        []string{"object"},
				Properties: map[string]spec.Schema{
					"version": {
						SchemaProps: spec.SchemaProps{
							Type:   []string{"string"},
							Format: "",
						},
					},
					"applicationImage": {
						SchemaProps: spec.SchemaProps{
							Type:   []string{"string"},
							Format: "",
						},
					},
					"replicas": {
						SchemaProps: spec.SchemaProps{
							Type:   []string{"integer"},
							Format: "int32",
						},
					},
					"autoscaling": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("./pkg/apis/openliberty/v1beta1.OpenLibertyApplicationAutoScaling"),
						},
					},
					"pullPolicy": {
						SchemaProps: spec.SchemaProps{
							Type:   []string{"string"},
							Format: "",
						},
					},
					"pullSecret": {
						SchemaProps: spec.SchemaProps{
							Type:   []string{"string"},
							Format: "",
						},
					},
					"volumes": {
						VendorExtensible: spec.VendorExtensible{
							Extensions: spec.Extensions{
								"x-kubernetes-list-map-keys": "name",
								"x-kubernetes-list-type":     "map",
							},
						},
						SchemaProps: spec.SchemaProps{
							Type: []string{"array"},
							Items: &spec.SchemaOrArray{
								Schema: &spec.Schema{
									SchemaProps: spec.SchemaProps{
										Ref: ref("k8s.io/api/core/v1.Volume"),
									},
								},
							},
						},
					},
					"volumeMounts": {
						VendorExtensible: spec.VendorExtensible{
							Extensions: spec.Extensions{
								"x-kubernetes-list-map-keys": "name",
								"x-kubernetes-list-type":     "map",
							},
						},
						SchemaProps: spec.SchemaProps{
							Type: []string{"array"},
							Items: &spec.SchemaOrArray{
								Schema: &spec.Schema{
									SchemaProps: spec.SchemaProps{
										Ref: ref("k8s.io/api/core/v1.VolumeMount"),
									},
								},
							},
						},
					},
					"resourceConstraints": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("k8s.io/api/core/v1.ResourceRequirements"),
						},
					},
					"readinessProbe": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("k8s.io/api/core/v1.Probe"),
						},
					},
					"livenessProbe": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("k8s.io/api/core/v1.Probe"),
						},
					},
					"service": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("./pkg/apis/openliberty/v1beta1.OpenLibertyApplicationService"),
						},
					},
					"expose": {
						SchemaProps: spec.SchemaProps{
							Type:   []string{"boolean"},
							Format: "",
						},
					},
					"envFrom": {
						VendorExtensible: spec.VendorExtensible{
							Extensions: spec.Extensions{
								"x-kubernetes-list-type": "atomic",
							},
						},
						SchemaProps: spec.SchemaProps{
							Type: []string{"array"},
							Items: &spec.SchemaOrArray{
								Schema: &spec.Schema{
									SchemaProps: spec.SchemaProps{
										Ref: ref("k8s.io/api/core/v1.EnvFromSource"),
									},
								},
							},
						},
					},
					"env": {
						VendorExtensible: spec.VendorExtensible{
							Extensions: spec.Extensions{
								"x-kubernetes-list-map-keys": "name",
								"x-kubernetes-list-type":     "map",
							},
						},
						SchemaProps: spec.SchemaProps{
							Type: []string{"array"},
							Items: &spec.SchemaOrArray{
								Schema: &spec.Schema{
									SchemaProps: spec.SchemaProps{
										Ref: ref("k8s.io/api/core/v1.EnvVar"),
									},
								},
							},
						},
					},
					"serviceAccountName": {
						SchemaProps: spec.SchemaProps{
							Type:   []string{"string"},
							Format: "",
						},
					},
					"architecture": {
						VendorExtensible: spec.VendorExtensible{
							Extensions: spec.Extensions{
								"x-kubernetes-list-type": "set",
							},
						},
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
					"storage": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("./pkg/apis/openliberty/v1beta1.OpenLibertyApplicationStorage"),
						},
					},
					"createKnativeService": {
						SchemaProps: spec.SchemaProps{
							Type:   []string{"boolean"},
							Format: "",
						},
					},
					"monitoring": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("./pkg/apis/openliberty/v1beta1.OpenLibertyApplicationMonitoring"),
						},
					},
					"createAppDefinition": {
						SchemaProps: spec.SchemaProps{
							Type:   []string{"boolean"},
							Format: "",
						},
					},
					"initContainers": {
						VendorExtensible: spec.VendorExtensible{
							Extensions: spec.Extensions{
								"x-kubernetes-list-map-keys": "name",
								"x-kubernetes-list-type":     "map",
							},
						},
						SchemaProps: spec.SchemaProps{
							Type: []string{"array"},
							Items: &spec.SchemaOrArray{
								Schema: &spec.Schema{
									SchemaProps: spec.SchemaProps{
										Ref: ref("k8s.io/api/core/v1.Container"),
									},
								},
							},
						},
					},
					"serviceability": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("./pkg/apis/openliberty/v1beta1.OpenLibertyApplicationServiceability"),
						},
					},
				},
				Required: []string{"applicationImage"},
			},
		},
		Dependencies: []string{
			"./pkg/apis/openliberty/v1beta1.OpenLibertyApplicationAutoScaling", "./pkg/apis/openliberty/v1beta1.OpenLibertyApplicationMonitoring", "./pkg/apis/openliberty/v1beta1.OpenLibertyApplicationService", "./pkg/apis/openliberty/v1beta1.OpenLibertyApplicationServiceability", "./pkg/apis/openliberty/v1beta1.OpenLibertyApplicationStorage", "k8s.io/api/core/v1.Container", "k8s.io/api/core/v1.EnvFromSource", "k8s.io/api/core/v1.EnvVar", "k8s.io/api/core/v1.Probe", "k8s.io/api/core/v1.ResourceRequirements", "k8s.io/api/core/v1.Volume", "k8s.io/api/core/v1.VolumeMount"},
	}
}

func schema_pkg_apis_openliberty_v1beta1_OpenLibertyApplicationStatus(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "OpenLibertyApplicationStatus defines the observed state of OpenLibertyApplication",
				Type:        []string{"object"},
				Properties: map[string]spec.Schema{
					"conditions": {
						VendorExtensible: spec.VendorExtensible{
							Extensions: spec.Extensions{
								"x-kubernetes-list-map-keys": "type",
								"x-kubernetes-list-type":     "map",
							},
						},
						SchemaProps: spec.SchemaProps{
							Type: []string{"array"},
							Items: &spec.SchemaOrArray{
								Schema: &spec.Schema{
									SchemaProps: spec.SchemaProps{
										Ref: ref("./pkg/apis/openliberty/v1beta1.StatusCondition"),
									},
								},
							},
						},
					},
					"consumedServices": {
						SchemaProps: spec.SchemaProps{
							Type: []string{"object"},
							AdditionalProperties: &spec.SchemaOrBool{
								Allows: true,
								Schema: &spec.Schema{
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
							},
						},
					},
				},
			},
		},
		Dependencies: []string{
			"./pkg/apis/openliberty/v1beta1.StatusCondition"},
	}
}

func schema_pkg_apis_openliberty_v1beta1_OpenLibertyApplicationStorage(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "OpenLibertyApplicationStorage ...",
				Type:        []string{"object"},
				Properties: map[string]spec.Schema{
					"size": {
						SchemaProps: spec.SchemaProps{
							Type:   []string{"string"},
							Format: "",
						},
					},
					"mountPath": {
						SchemaProps: spec.SchemaProps{
							Type:   []string{"string"},
							Format: "",
						},
					},
					"volumeClaimTemplate": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("k8s.io/api/core/v1.PersistentVolumeClaim"),
						},
					},
				},
			},
		},
		Dependencies: []string{
			"k8s.io/api/core/v1.PersistentVolumeClaim"},
	}
}

func schema_pkg_apis_openliberty_v1beta1_OpenLibertyDump(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "OpenLibertyDump is the Schema for the openlibertydumps API",
				Type:        []string{"object"},
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
							Ref: ref("./pkg/apis/openliberty/v1beta1.OpenLibertyDumpSpec"),
						},
					},
					"status": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("./pkg/apis/openliberty/v1beta1.OpenLibertyDumpStatus"),
						},
					},
				},
			},
		},
		Dependencies: []string{
			"./pkg/apis/openliberty/v1beta1.OpenLibertyDumpSpec", "./pkg/apis/openliberty/v1beta1.OpenLibertyDumpStatus", "k8s.io/apimachinery/pkg/apis/meta/v1.ObjectMeta"},
	}
}

func schema_pkg_apis_openliberty_v1beta1_OpenLibertyDumpSpec(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "OpenLibertyDumpSpec defines the desired state of OpenLibertyDump",
				Type:        []string{"object"},
				Properties: map[string]spec.Schema{
					"podName": {
						SchemaProps: spec.SchemaProps{
							Type:   []string{"string"},
							Format: "",
						},
					},
					"include": {
						VendorExtensible: spec.VendorExtensible{
							Extensions: spec.Extensions{
								"x-kubernetes-list-type": "set",
							},
						},
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
				},
				Required: []string{"podName"},
			},
		},
	}
}

func schema_pkg_apis_openliberty_v1beta1_OpenLibertyDumpStatus(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "OpenLibertyDumpStatus defines the observed state of OpenLibertyDump",
				Type:        []string{"object"},
				Properties: map[string]spec.Schema{
					"conditions": {
						VendorExtensible: spec.VendorExtensible{
							Extensions: spec.Extensions{
								"x-kubernetes-list-type": "atomic",
							},
						},
						SchemaProps: spec.SchemaProps{
							Type: []string{"array"},
							Items: &spec.SchemaOrArray{
								Schema: &spec.Schema{
									SchemaProps: spec.SchemaProps{
										Ref: ref("./pkg/apis/openliberty/v1beta1.OperationStatusCondition"),
									},
								},
							},
						},
					},
					"dumpFile": {
						SchemaProps: spec.SchemaProps{
							Type:   []string{"string"},
							Format: "",
						},
					},
				},
			},
		},
		Dependencies: []string{
			"./pkg/apis/openliberty/v1beta1.OperationStatusCondition"},
	}
}

func schema_pkg_apis_openliberty_v1beta1_OpenLibertyTrace(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "OpenLibertyTrace is the schema for the openlibertytraces API",
				Type:        []string{"object"},
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
							Ref: ref("./pkg/apis/openliberty/v1beta1.OpenLibertyTraceSpec"),
						},
					},
					"status": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("./pkg/apis/openliberty/v1beta1.OpenLibertyTraceStatus"),
						},
					},
				},
			},
		},
		Dependencies: []string{
			"./pkg/apis/openliberty/v1beta1.OpenLibertyTraceSpec", "./pkg/apis/openliberty/v1beta1.OpenLibertyTraceStatus", "k8s.io/apimachinery/pkg/apis/meta/v1.ObjectMeta"},
	}
}

func schema_pkg_apis_openliberty_v1beta1_OpenLibertyTraceSpec(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "OpenLibertyTraceSpec defines the desired state of OpenLibertyTrace",
				Type:        []string{"object"},
				Properties: map[string]spec.Schema{
					"podName": {
						SchemaProps: spec.SchemaProps{
							Type:   []string{"string"},
							Format: "",
						},
					},
					"traceSpecification": {
						SchemaProps: spec.SchemaProps{
							Type:   []string{"string"},
							Format: "",
						},
					},
					"maxFileSize": {
						SchemaProps: spec.SchemaProps{
							Type:   []string{"integer"},
							Format: "int32",
						},
					},
					"maxFiles": {
						SchemaProps: spec.SchemaProps{
							Type:   []string{"integer"},
							Format: "int32",
						},
					},
					"disable": {
						SchemaProps: spec.SchemaProps{
							Type:   []string{"boolean"},
							Format: "",
						},
					},
				},
				Required: []string{"podName", "traceSpecification"},
			},
		},
	}
}

func schema_pkg_apis_openliberty_v1beta1_OpenLibertyTraceStatus(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "OpenLibertyTraceStatus defines the observed state of OpenLibertyTrace operation",
				Type:        []string{"object"},
				Properties: map[string]spec.Schema{
					"conditions": {
						VendorExtensible: spec.VendorExtensible{
							Extensions: spec.Extensions{
								"x-kubernetes-list-type": "atomic",
							},
						},
						SchemaProps: spec.SchemaProps{
							Type: []string{"array"},
							Items: &spec.SchemaOrArray{
								Schema: &spec.Schema{
									SchemaProps: spec.SchemaProps{
										Ref: ref("./pkg/apis/openliberty/v1beta1.OperationStatusCondition"),
									},
								},
							},
						},
					},
					"operatedResource": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("./pkg/apis/openliberty/v1beta1.OperatedResource"),
						},
					},
				},
			},
		},
		Dependencies: []string{
			"./pkg/apis/openliberty/v1beta1.OperatedResource", "./pkg/apis/openliberty/v1beta1.OperationStatusCondition"},
	}
}

func schema_pkg_apis_openliberty_v1beta1_OperatedResource(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "OperatedResource ...",
				Type:        []string{"object"},
				Properties: map[string]spec.Schema{
					"resourceType": {
						SchemaProps: spec.SchemaProps{
							Type:   []string{"string"},
							Format: "",
						},
					},
					"resourceName": {
						SchemaProps: spec.SchemaProps{
							Type:   []string{"string"},
							Format: "",
						},
					},
				},
			},
		},
	}
}

func schema_pkg_apis_openliberty_v1beta1_OperationStatusCondition(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "OperationStatusCondition ...",
				Type:        []string{"object"},
				Properties: map[string]spec.Schema{
					"lastTransitionTime": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("k8s.io/apimachinery/pkg/apis/meta/v1.Time"),
						},
					},
					"lastUpdateTime": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("k8s.io/apimachinery/pkg/apis/meta/v1.Time"),
						},
					},
					"reason": {
						SchemaProps: spec.SchemaProps{
							Type:   []string{"string"},
							Format: "",
						},
					},
					"message": {
						SchemaProps: spec.SchemaProps{
							Type:   []string{"string"},
							Format: "",
						},
					},
					"status": {
						SchemaProps: spec.SchemaProps{
							Type:   []string{"string"},
							Format: "",
						},
					},
					"type": {
						SchemaProps: spec.SchemaProps{
							Type:   []string{"string"},
							Format: "",
						},
					},
				},
			},
		},
		Dependencies: []string{
			"k8s.io/apimachinery/pkg/apis/meta/v1.Time"},
	}
}

func schema_pkg_apis_openliberty_v1beta1_ServiceBindingConsumes(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "ServiceBindingConsumes represents a service to be consumed",
				Type:        []string{"object"},
				Properties: map[string]spec.Schema{
					"name": {
						SchemaProps: spec.SchemaProps{
							Type:   []string{"string"},
							Format: "",
						},
					},
					"namespace": {
						SchemaProps: spec.SchemaProps{
							Type:   []string{"string"},
							Format: "",
						},
					},
					"category": {
						SchemaProps: spec.SchemaProps{
							Type:   []string{"string"},
							Format: "",
						},
					},
					"mountPath": {
						SchemaProps: spec.SchemaProps{
							Type:   []string{"string"},
							Format: "",
						},
					},
				},
				Required: []string{"name", "category"},
			},
		},
	}
}

func schema_pkg_apis_openliberty_v1beta1_ServiceBindingProvides(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "ServiceBindingProvides represents information about",
				Type:        []string{"object"},
				Properties: map[string]spec.Schema{
					"category": {
						SchemaProps: spec.SchemaProps{
							Type:   []string{"string"},
							Format: "",
						},
					},
					"context": {
						SchemaProps: spec.SchemaProps{
							Type:   []string{"string"},
							Format: "",
						},
					},
					"protocol": {
						SchemaProps: spec.SchemaProps{
							Type:   []string{"string"},
							Format: "",
						},
					},
					"auth": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("./pkg/apis/openliberty/v1beta1.ServiceBindingAuth"),
						},
					},
				},
				Required: []string{"category"},
			},
		},
		Dependencies: []string{
			"./pkg/apis/openliberty/v1beta1.ServiceBindingAuth"},
	}
}

func schema_pkg_apis_openliberty_v1beta1_StatusCondition(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "StatusCondition ...",
				Type:        []string{"object"},
				Properties: map[string]spec.Schema{
					"lastTransitionTime": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("k8s.io/apimachinery/pkg/apis/meta/v1.Time"),
						},
					},
					"lastUpdateTime": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("k8s.io/apimachinery/pkg/apis/meta/v1.Time"),
						},
					},
					"reason": {
						SchemaProps: spec.SchemaProps{
							Type:   []string{"string"},
							Format: "",
						},
					},
					"message": {
						SchemaProps: spec.SchemaProps{
							Type:   []string{"string"},
							Format: "",
						},
					},
					"status": {
						SchemaProps: spec.SchemaProps{
							Type:   []string{"string"},
							Format: "",
						},
					},
					"type": {
						SchemaProps: spec.SchemaProps{
							Type:   []string{"string"},
							Format: "",
						},
					},
				},
			},
		},
		Dependencies: []string{
			"k8s.io/apimachinery/pkg/apis/meta/v1.Time"},
	}
}