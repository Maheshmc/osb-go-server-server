/*
 * Open Service Broker API
 *
 * The Open Service Broker API defines an HTTP(S) interface between Platforms and Service Brokers.
 *
 * API version: master - might contain changes that are not yet released
 * Contact: open-service-broker-api@googlegroups.com
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

type ServiceBindingVolumeMount struct {

	ContainerDir string `json:"container_dir"`

	Device *ServiceBindingVolumeMountDevice `json:"device"`

	DeviceType string `json:"device_type"`

	Driver string `json:"driver"`

	Mode string `json:"mode"`
}
