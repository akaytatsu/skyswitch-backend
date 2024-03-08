package usecase_cloud_account

import (
	"app/entity"
	usecase_instance "app/usecase/instance"
	"log"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ec2"
)

type UseCaseAWSCloudAccount struct {
	repo             IRepositoryCloudAccount
	useCaseInstances usecase_instance.IUseCaseInstance
}

func NewAWSService(repository IRepositoryCloudAccount, usecaseInstances usecase_instance.IUseCaseInstance) *UseCaseAWSCloudAccount {
	return &UseCaseAWSCloudAccount{repo: repository, useCaseInstances: usecaseInstances}
}

func (u *UseCaseAWSCloudAccount) GetAll(searchParams entity.SearchEntityCloudAccountParams) (response []entity.EntityCloudAccount, totalRegisters int64, err error) {
	return u.repo.GetAll(searchParams)
}

func (u *UseCaseAWSCloudAccount) GetByID(id int64) (cloudAccount *entity.EntityCloudAccount, err error) {
	return u.repo.GetByID(id)
}

func (u *UseCaseAWSCloudAccount) CreateCloudAccount(cloudAccount *entity.EntityCloudAccount) error {
	err := u.repo.CreateCloudAccount(cloudAccount)

	if err != nil {
		return err
	}

	go u.UpdateAllInstancesOnCloudAccountProvider(cloudAccount)

	return nil
}

func (u *UseCaseAWSCloudAccount) UpdateCloudAccount(cloudAccount *entity.EntityCloudAccount) error {
	err := u.repo.UpdateCloudAccount(cloudAccount)

	if err != nil {
		return err
	}

	go u.UpdateAllInstancesOnCloudAccountProvider(cloudAccount)

	return nil
}

func (u *UseCaseAWSCloudAccount) DeleteCloudAccount(cloudAccount *entity.EntityCloudAccount) error {
	return u.repo.DeleteCloudAccount(cloudAccount)
}

func (u *UseCaseAWSCloudAccount) ActiveDeactiveCloudAccount(id int64, status bool) (cloudAccount *entity.EntityCloudAccount, err error) {
	return u.repo.ActiveDeactiveCloudAccount(id, status)
}

func (u *UseCaseAWSCloudAccount) UpdateAllInstancesOnAllCloudAccountProvider() (instances []*entity.EntityInstance, err error) {

	params := entity.SearchEntityCloudAccountParams{
		Page:     0,
		PageSize: 10000,
	}

	cloudAccounts, _, err := u.repo.GetAll(params)
	if err != nil {
		return nil, err
	}

	for _, cloudAccount := range cloudAccounts {
		instances, err = u.UpdateAllInstancesOnCloudAccountProvider(&cloudAccount)
		if err != nil {
			log.Println("Error updating all instances on cloud account provider: ", err)
		}
	}

	return instances, nil
}

func (u *UseCaseAWSCloudAccount) UpdateAllInstancesOnCloudAccountProvider(cloudAccount *entity.EntityCloudAccount) (instances []*entity.EntityInstance, err error) {

	aws_access_key := cloudAccount.AccessKeyID
	aws_secret_key := cloudAccount.SecretAccessKey

	client := u.getAwsClient(aws_access_key, aws_secret_key)

	instances, err = u.getAwsEC2AllInstances(cloudAccount, client)

	if err != nil {
		return instances, err
	}

	for _, instance := range instances {

		err = u.useCaseInstances.CreateOrUpdateInstance(instance, false)

		if err != nil {
			log.Println("Error creating or updating instance: ", err)
		}
	}

	return make([]*entity.EntityInstance, 0), nil
}

func (u *UseCaseAWSCloudAccount) getAwsClient(aws_access_key string, aws_secret_key string) *session.Session {

	sess := session.Must(session.NewSession(&aws.Config{
		Region:      aws.String("us-east-1"),                                              // Substitua pela sua região desejada
		Credentials: credentials.NewStaticCredentials(aws_access_key, aws_secret_key, ""), // Substitua com suas chaves de acesso e segredo
	}))

	return sess

}

func (u *UseCaseAWSCloudAccount) getAwsEC2AllInstances(cloudAccount *entity.EntityCloudAccount, sess *session.Session) (instances []*entity.EntityInstance, err error) {

	instances = make([]*entity.EntityInstance, 0)

	// Create new EC2 client
	svc := ec2.New(sess)

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
				CloudAccountID: cloudAccount.ID,
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

func (u *UseCaseAWSCloudAccount) UpdateAllInstancesOnCloudAccountProviderFromID(id int) (instances []*entity.EntityInstance, err error) {
	cloudAccount, err := u.GetByID(int64(id))
	if err != nil {
		return nil, err
	}

	return u.UpdateAllInstancesOnCloudAccountProvider(cloudAccount)
}
