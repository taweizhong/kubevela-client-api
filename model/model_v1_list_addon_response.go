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

type V1ListAddonResponse struct {
	Addons []V1AddonInfo `json:"addons"`
	Message string `json:"message,omitempty"`
}
