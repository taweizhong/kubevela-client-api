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

type V1UpdateConfigRequest struct {
	Alias string `json:"alias"`
	Description string `json:"description"`
	Properties string `json:"properties,omitempty"`
}
