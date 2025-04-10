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

type V1SystemInfoResponse struct {
	DexUserDefaultPlatformRoles []string          `json:"dexUserDefaultPlatformRoles,omitempty"`
	DexUserDefaultProjects      []ModelProjectRef `json:"dexUserDefaultProjects,omitempty"`
	EnableCollection            bool              `json:"enableCollection"`
	InstallTime                 time.Time         `json:"installTime,omitempty"`
	LoginType                   string            `json:"loginType"`
	PlatformID                  string            `json:"platformID"`
	StatisticInfo               *V1StatisticInfo  `json:"statisticInfo,omitempty"`
	SystemVersion               *V1SystemVersion  `json:"systemVersion"`
}
