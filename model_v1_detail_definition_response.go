/*
 * Kubevela api doc
 *
 * Kubevela api doc
 *
 * API version: v1beta1
 * Contact: feedback@mail.kubevela.io
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package velaclient

type V1DetailDefinitionResponse struct {
	Alias        string                             `json:"alias"`
	Category     string                             `json:"category"`
	Component    *V1beta1ComponentDefinitionSpec    `json:"component,omitempty"`
	Description  string                             `json:"description"`
	Icon         string                             `json:"icon"`
	Labels       map[string]string                  `json:"labels"`
	Name         string                             `json:"name"`
	OwnerAddon   string                             `json:"ownerAddon"`
	Policy       *V1beta1PolicyDefinitionSpec       `json:"policy,omitempty"`
	Schema       string                             `json:"schema"`
	Status       string                             `json:"status"`
	Trait        *V1beta1TraitDefinitionSpec        `json:"trait,omitempty"`
	UiSchema     []SchemaUiParameter                `json:"uiSchema"`
	WorkflowStep *V1beta1WorkflowStepDefinitionSpec `json:"workflowStep,omitempty"`
	WorkloadType string                             `json:"workloadType,omitempty"`
}
