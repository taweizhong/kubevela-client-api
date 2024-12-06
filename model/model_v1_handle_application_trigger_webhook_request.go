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

type V1HandleApplicationTriggerWebhookRequest struct {
	Action string `json:"action,omitempty"`
	CodeInfo *ModelCodeInfo `json:"codeInfo,omitempty"`
	Step string `json:"step,omitempty"`
	Upgrade map[string]ModelJsonStruct `json:"upgrade,omitempty"`
}
