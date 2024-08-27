package infrastructure_cloud_provider

import "app/entity"

//go:generate mockgen -destination=../../mocks/mock_infrastructure_cloud_provider.go -package=mocks app/infrastructure/cloud_provider ICloudProvider
type ICloudProvider interface {
	Connect(cloudAccount entity.EntityCloudAccount) (cloudProvider ICloudProvider, err error)

	// EC2
	GetInstances() (instances []*entity.EntityInstance, err error)
	GetInstanceByID(string) (instance *entity.EntityInstance, err error)
	StartInstance(instanceID string) (err error)
	StopInstance(instanceID string) (err error)

	//RDS
	GetDBInstances() (dbInstances []*entity.EntityDbinstance, err error)
	GetDBInstanceByID(string) (dbInstance *entity.EntityDbinstance, err error)
	StartDBInstance(dbInstanceID string) (err error)
	StopDBInstance(dbInstanceID string) (err error)

	//AutoScallingGroups
	GetAutoScalingGroups() (autoScalingGroups []*entity.EntityAutoScalingGroup, err error)
	GetAutoScalingGroupByID(string) (autoScalingGroup *entity.EntityAutoScalingGroup, err error)
	StartAutoScalingGroup(autoScalingGroup *entity.EntityAutoScalingGroup) (err error)
	StopAutoScalingGroup(autoScalingGroup *entity.EntityAutoScalingGroup) (err error)
}
