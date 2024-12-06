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

type V1Descriptor struct {
	Annotations map[string]string `json:"annotations,omitempty"`
	ArtifactType string `json:"artifactType,omitempty"`
	Data string `json:"data,omitempty"`
	Digest string `json:"digest"`
	MediaType string `json:"mediaType"`
	Platform *V1Platform `json:"platform,omitempty"`
	Size int64 `json:"size"`
	Urls []string `json:"urls,omitempty"`
}
