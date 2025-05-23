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

import (
	"time"
)

type V1PipelineMeta struct {
	Alias       string       `json:"alias"`
	CreateTime  time.Time    `json:"createTime"`
	Description string       `json:"description"`
	Name        string       `json:"name"`
	Project     *V1NameAlias `json:"project"`
}
