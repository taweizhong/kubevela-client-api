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

type V1CreateRoleRequest struct {
	Alias       string   `json:"alias"`
	Name        string   `json:"name"`
	Permissions []string `json:"permissions"`
}
