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

type V1AddonStatusResponse struct {
	AllClusters []V1NameAlias `json:"allClusters,omitempty"`
	AppStatus *CommonAppStatus `json:"appStatus,omitempty"`
	Args *interface{} `json:"args"`
	Clusters map[string]V1AddonStatusResponseClusters `json:"clusters,omitempty"`
	EnablingProgress *V1EnablingProgress `json:"enabling_progress,omitempty"`
	InstalledVersion string `json:"installedVersion,omitempty"`
	Name string `json:"name"`
	Phase string `json:"phase"`
}
