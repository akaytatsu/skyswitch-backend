package infrastructure_cloud_provider_aws

import (
	"app/entity"
	infrastructure_cloud_provider "app/infrastructure/cloud_provider"
	"log"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ec2"
	"github.com/aws/aws-sdk-go/service/rds"

)

type AWSCloudProvider struct {
	awsSession   *session.Session
	cloudAccount entity.EntityCloudAccount
}

func NewAWSCloudProvider() infrastructure_cloud_provider.ICloudProvider {
	return &AWSCloudProvider{}
}

func (a *AWSCloudProvider) Connect(cloudAccount entity.EntityCloudAccount) (cloudProvider infrastructure_cloud_provider.ICloudProvider, err error) {

	sess := session.Must(session.NewSession(&aws.Config{
		Region:      aws.String(cloudAccount.Region),
		Credentials: credentials.NewStaticCredentials(cloudAccount.AccessKeyID, cloudAccount.SecretAccessKey, ""),
	}))

	cloudProviderReturn := &AWSCloudProvider{
		awsSession:   sess,
		cloudAccount: cloudAccount,
	}

	return cloudProviderReturn, nil
}

//EC2
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
		log.Println("Error on stop instance: ", err.Error())
		return err
	}

	return nil
}

// GetDBInstances retorna todas as instâncias de banco de dados no RDS
func (a *AWSCloudProvider) GetDBInstances() (dbInstances []*entity.EntityDBInstance, err error) {
	dbInstances = make([]*entity.EntityDBInstance, 0)

	// Cria um novo cliente RDS
	svc := rds.New(a.awsSession)

	// Descreve as instâncias de banco de dados
	result, err := svc.DescribeDBInstances(nil)
	if err != nil {
		log.Println("Error describing DB instances: ", err.Error())
		return dbInstances, err
	}

	for _, dbInstance := range result.DBInstances {
		dbInstances = append(dbInstances, &entity.EntityDBInstance{
			CloudAccountID:  a.cloudAccount.ID,
			DBInstanceID:    *dbInstance.DBInstanceIdentifier,
			DBInstanceClass: *dbInstance.DBInstanceClass,
			DBInstanceState: *dbInstance.DBInstanceStatus,
			Endpoint:        *dbInstance.Endpoint.Address,
			Port:            *dbInstance.Endpoint.Port,
			Engine:          *dbInstance.Engine,
			Active:          true,
		})
	}

	return dbInstances, nil
}

// GetDBInstanceByID retorna uma instância específica de banco de dados pelo ID
func (a *AWSCloudProvider) GetDBInstanceByID(dbInstanceID string) (dbInstance *entity.EntityDBInstance, err error) {
	dbInstances, err := a.GetDBInstances()
	if err != nil {
		return nil, err
	}

	for _, dbInstance := range dbInstances {
		if dbInstance.DBInstanceID == dbInstanceID {
			return dbInstance, nil
		}
	}

	return nil, nil
}

// StartDBInstance inicia uma instância de banco de dados no RDS
func (a *AWSCloudProvider) StartDBInstance(dbInstanceID string) (err error) {
	svc := rds.New(a.awsSession)

	input := &rds.StartDBInstanceInput{
		DBInstanceIdentifier: aws.String(dbInstanceID),
	}

	_, err = svc.StartDBInstance(input)
	if err != nil {
		log.Println("Error starting DB instance: ", err.Error())
		return err
	}

	return nil
}

// StopDBInstance para uma instância de banco de dados no RDS
func (a *AWSCloudProvider) StopDBInstance(dbInstanceID string) (err error) {
	svc := rds.New(a.awsSession)

	input := &rds.StopDBInstanceInput{
		DBInstanceIdentifier: aws.String(dbInstanceID),
	}

	_, err = svc.StopDBInstance(input)
	if err != nil {
		log.Println("Error stopping DB instance: ", err.Error())
		return err
	}

	return nil
}