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

type V1alpha1WorkflowRunStatus struct {
	Conditions []ConditionCondition `json:"conditions,omitempty"`
	ContextBackend *V1ObjectReference `json:"contextBackend,omitempty"`
	EndTime string `json:"endTime,omitempty"`
	Finished bool `json:"finished"`
	Message string `json:"message,omitempty"`
	Mode *V1alpha1WorkflowExecuteMode `json:"mode"`
	StartTime string `json:"startTime,omitempty"`
	Status string `json:"status"`
	Steps []V1alpha1WorkflowStepStatus `json:"steps,omitempty"`
	Suspend bool `json:"suspend"`
	SuspendState string `json:"suspendState,omitempty"`
	Terminated bool `json:"terminated"`
}
