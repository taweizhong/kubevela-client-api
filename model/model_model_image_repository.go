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

type ModelImageRepository struct {
	CreateTime time.Time `json:"createTime,omitempty"`
	FullName string `json:"fullName"`
	Name string `json:"name"`
	Namespace string `json:"namespace"`
	Region string `json:"region,omitempty"`
	Type_ string `json:"type"`
}
