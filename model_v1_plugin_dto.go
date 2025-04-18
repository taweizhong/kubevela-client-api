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

type V1PluginDto struct {
	BaseURL       string          `json:"baseURL"`
	Category      string          `json:"category"`
	DefaultNavURL string          `json:"defaultNavURL"`
	Id            string          `json:"id"`
	Includes      []TypesIncludes `json:"includes"`
	Info          *TypesInfo      `json:"info"`
	Module        string          `json:"module"`
	Name          string          `json:"name"`
	SubType       string          `json:"subType"`
	Type_         string          `json:"type"`
}
