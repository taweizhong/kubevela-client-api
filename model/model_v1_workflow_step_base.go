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

type V1WorkflowStepBase struct {
	Alias string `json:"alias,omitempty"`
	DependsOn []string `json:"dependsOn,omitempty"`
	Description string `json:"description,omitempty"`
	If_ string `json:"if,omitempty"`
	Inputs []V1alpha1InputItem `json:"inputs,omitempty"`
	Meta *V1alpha1WorkflowStepMeta `json:"meta,omitempty"`
	Name string `json:"name"`
	Outputs []V1alpha1OutputItem `json:"outputs,omitempty"`
	Properties *interface{} `json:"properties,omitempty"`
	Timeout string `json:"timeout,omitempty"`
	Type_ string `json:"type"`
}
