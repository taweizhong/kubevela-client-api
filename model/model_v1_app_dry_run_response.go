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

type V1AppDryRunResponse struct {
	Message string `json:"message,omitempty"`
	Success bool `json:"success"`
	Yaml string `json:"yaml"`
}
