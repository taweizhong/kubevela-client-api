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

type V1alpha1WorkflowStep struct {
	DependsOn []string `json:"dependsOn,omitempty"`
	If_ string `json:"if,omitempty"`
	Inputs []V1alpha1InputItem `json:"inputs,omitempty"`
	Meta *V1alpha1WorkflowStepMeta `json:"meta,omitempty"`
	Mode string `json:"mode,omitempty"`
	Name string `json:"name,omitempty"`
	Outputs []V1alpha1OutputItem `json:"outputs,omitempty"`
	Properties string `json:"properties,omitempty"`
	SubSteps []V1alpha1WorkflowStepBase `json:"subSteps,omitempty"`
	Timeout string `json:"timeout,omitempty"`
	Type_ string `json:"type"`
}
