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
import (
	"time"
)

type RepoChartVersion struct {
	Metadata *ChartMetadata `json:"Metadata"`
	Checksum string `json:"checksum,omitempty"`
	Created time.Time `json:"created,omitempty"`
	Digest string `json:"digest,omitempty"`
	Engine string `json:"engine,omitempty"`
	Removed bool `json:"removed,omitempty"`
	TillerVersion string `json:"tillerVersion,omitempty"`
	Url string `json:"url,omitempty"`
	Urls []string `json:"urls"`
}
