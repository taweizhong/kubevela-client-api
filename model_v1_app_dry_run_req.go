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

type V1AppDryRunReq struct {
	DryRunType string `json:"dryRunType"`
	Env        string `json:"env"`
	Version    string `json:"version"`
	Workflow   string `json:"workflow"`
}
