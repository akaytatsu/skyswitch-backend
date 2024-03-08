package infrastructure_cloud_provider

import "app/entity"

type ICloudProvider interface {
	Connect(cloudAccount entity.EntityCloudAccount) (err error)
	GetInstances() (instances []*entity.EntityInstance, err error)
	GetInstanceByID(string) (instance *entity.EntityInstance, err error)
	StartInstance(instanceID string) (err error)
	StopInstance(instanceID string) (err error)
}
