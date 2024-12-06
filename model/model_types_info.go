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

type TypesInfo struct {
	Author *TypesInfoLink `json:"author"`
	Build *TypesBuildInfo `json:"build"`
	Description string `json:"description"`
	Links []TypesInfoLink `json:"links"`
	Logos *TypesLogos `json:"logos"`
	Screenshots []TypesScreenshots `json:"screenshots"`
	Updated string `json:"updated"`
	Version string `json:"version"`
}
