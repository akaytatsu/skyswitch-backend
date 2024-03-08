package infrastructure_cloud_provider_aws

import (
	"app/entity"
	infrastructure_cloud_provider "app/infrastructure/cloud_provider"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ec2"
)

type AWSCloudProvider struct {
	awsSession   *session.Session
	cloudAccount entity.EntityCloudAccount
}

func NewAWSCloudProvider() infrastructure_cloud_provider.ICloudProvider {
	return &AWSCloudProvider{}
}

func (a *AWSCloudProvider) Connect(cloudAccount entity.EntityCloudAccount) (err error) {

	sess := session.Must(session.NewSession(&aws.Config{
		Region:      aws.String(cloudAccount.Region),
		Credentials: credentials.NewStaticCredentials(cloudAccount.AccessKeyID, cloudAccount.SecretAccessKey, ""),
	}))

	a.awsSession = sess
	a.cloudAccount = cloudAccount

	return nil
}

func (a *AWSCloudProvider) GetInstances() (instances []*entity.EntityInstance, err error) {
	instances = make([]*entity.EntityInstance, 0)

	// Create new EC2 client
	svc := ec2.New(a.awsSession)

	// Call to get detailed information on each instance
	result, err := svc.DescribeInstances(nil)
	if err != nil {
		return instances, err
	}

	for _, reservations := range result.Reservations {
		for _, instance := range reservations.Instances {

			var name string

			for _, tag := range instance.Tags {
				if *tag.Key == "Name" {
					name = *tag.Value
				}
			}

			instances = append(instances, &entity.EntityInstance{
				CloudAccountID: a.cloudAccount.ID,
				InstanceID:     *instance.InstanceId,
				InstanceType:   *instance.InstanceType,
				InstanceName:   name,
				InstanceRegion: *instance.Placement.AvailabilityZone,
				InstanceState:  *instance.State.Name,
				Active:         true,
			})
		}
	}

	return instances, nil
}

func (a *AWSCloudProvider) GetInstanceByID(instanceID string) (instance *entity.EntityInstance, err error) {
	instances, _ := a.GetInstances()

	for _, instance := range instances {
		if instance.InstanceID == instanceID {
			return instance, nil
		}
	}

	return instance, nil
}

func (a *AWSCloudProvider) StartInstance(instanceID string) (err error) {
	svc := ec2.New(a.awsSession)

	input := &ec2.StartInstancesInput{
		InstanceIds: []*string{
			aws.String(instanceID),
		},
	}

	_, err = svc.StartInstances(input)
	if err != nil {
		return err
	}

	return nil
}

func (a *AWSCloudProvider) StopInstance(instanceID string) (err error) {
	svc := ec2.New(a.awsSession)

	input := &ec2.StopInstancesInput{
		InstanceIds: []*string{
			aws.String(instanceID),
		},
	}

	_, err = svc.StopInstances(input)
	if err != nil {
		return err
	}

	return nil
}
