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

type V1ApplicationBase struct {
	Alias string `json:"alias"`
	Annotations map[string]string `json:"annotations,omitempty"`
	CreateTime time.Time `json:"createTime"`
	Description string `json:"description"`
	Icon string `json:"icon"`
	Labels map[string]string `json:"labels,omitempty"`
	Name string `json:"name"`
	Project *V1ProjectBase `json:"project"`
	ReadOnly bool `json:"readOnly,omitempty"`
	UpdateTime time.Time `json:"updateTime"`
}
