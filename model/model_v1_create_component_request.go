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

type V1CreateComponentRequest struct {
	Alias string `json:"alias,omitempty"`
	ComponentType string `json:"componentType"`
	DependsOn []string `json:"dependsOn,omitempty"`
	Description string `json:"description,omitempty"`
	Icon string `json:"icon,omitempty"`
	Inputs []V1alpha1InputItem `json:"inputs,omitempty"`
	Labels map[string]string `json:"labels,omitempty"`
	Name string `json:"name"`
	Outputs []V1alpha1OutputItem `json:"outputs,omitempty"`
	Properties string `json:"properties,omitempty"`
	Traits []V1CreateApplicationTraitRequest `json:"traits,omitempty"`
}
