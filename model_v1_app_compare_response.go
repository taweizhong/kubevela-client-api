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

type V1AppCompareResponse struct {
	BaseAppYAML   string `json:"baseAppYAML"`
	DiffReport    string `json:"diffReport"`
	IsDiff        bool   `json:"isDiff"`
	TargetAppYAML string `json:"targetAppYAML"`
}
