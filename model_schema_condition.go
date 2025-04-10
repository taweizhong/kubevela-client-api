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

type SchemaCondition struct {
	Action  string                `json:"action,omitempty"`
	JsonKey string                `json:"jsonKey"`
	Op      string                `json:"op,omitempty"`
	Value   *SchemaConditionValue `json:"value"`
}
