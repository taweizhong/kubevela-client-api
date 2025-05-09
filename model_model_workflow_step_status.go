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

import (
	"time"
)

type ModelWorkflowStepStatus struct {
	Alias            string            `json:"alias"`
	FirstExecuteTime time.Time         `json:"firstExecuteTime,omitempty"`
	Id               string            `json:"id"`
	LastExecuteTime  time.Time         `json:"lastExecuteTime,omitempty"`
	Message          string            `json:"message,omitempty"`
	Name             string            `json:"name"`
	Phase            string            `json:"phase,omitempty"`
	Reason           string            `json:"reason,omitempty"`
	SubSteps         []ModelStepStatus `json:"subSteps,omitempty"`
	Type_            string            `json:"type,omitempty"`
}
