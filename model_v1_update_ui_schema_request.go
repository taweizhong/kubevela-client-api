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

type V1UpdateUiSchemaRequest struct {
	Type_    string              `json:"type"`
	UiSchema []SchemaUiParameter `json:"uiSchema"`
}
