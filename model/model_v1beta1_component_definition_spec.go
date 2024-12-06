/*
 * Kubevela api doc
 *
 * Kubevela api doc
 *
 * API version: v1beta1
 * Contact: feedback@mail.kubevela.io
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

type V1beta1ComponentDefinitionSpec struct {
	ChildResourceKinds []CommonChildResourceKind `json:"childResourceKinds,omitempty"`
	Extension string `json:"extension,omitempty"`
	PodSpecPath string `json:"podSpecPath,omitempty"`
	RevisionLabel string `json:"revisionLabel,omitempty"`
	Schematic *CommonSchematic `json:"schematic,omitempty"`
	Status *CommonStatus `json:"status,omitempty"`
	Workload *CommonWorkloadTypeDescriptor `json:"workload"`
}
