package usecase_cloud_account

import (
	"app/entity"
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ec2"
)

type UseCaseAWSCloudAccount struct {
	repo IRepositoryCloudAccount
}

func NewAWSService(repository IRepositoryCloudAccount) *UseCaseAWSCloudAccount {
	return &UseCaseAWSCloudAccount{repo: repository}
}

func (u *UseCaseAWSCloudAccount) GetAll() (cloudAccounts []*entity.EntityCloudAccount, err error) {
	return u.repo.GetAll()
}

func (u *UseCaseAWSCloudAccount) GetByID(id int64) (cloudAccount *entity.EntityCloudAccount, err error) {
	return u.repo.GetByID(id)
}

func (u *UseCaseAWSCloudAccount) CreateCloudAccount(cloudAccount *entity.EntityCloudAccount) error {
	return u.repo.CreateCloudAccount(cloudAccount)
}

func (u *UseCaseAWSCloudAccount) UpdateCloudAccount(cloudAccount *entity.EntityCloudAccount) error {
	return u.repo.UpdateCloudAccount(cloudAccount)
}

func (u *UseCaseAWSCloudAccount) DeleteCloudAccount(cloudAccount *entity.EntityCloudAccount) error {
	return u.repo.DeleteCloudAccount(cloudAccount)
}

func (u *UseCaseAWSCloudAccount) ActiveDeactiveCloudAccount(id int64, status bool) (cloudAccount *entity.EntityCloudAccount, err error) {
	return u.repo.ActiveDeactiveCloudAccount(id, status)
}

func (u *UseCaseAWSCloudAccount) UpdateAllInstancesOnCloudAccountProvider(cloudAccount *entity.EntityCloudAccount) (instances []*entity.EntityInstance, err error) {

	aws_access_key := cloudAccount.AccessKeyID
	aws_secret_key := cloudAccount.SecretAccessKey

	client := u.getAwsClient(aws_access_key, aws_secret_key)

	u.getAwsEC2AllInstances(client)

	return make([]*entity.EntityInstance, 0), nil
}

func (u *UseCaseAWSCloudAccount) getAwsClient(aws_access_key string, aws_secret_key string) *session.Session {

	sess := session.Must(session.NewSession(&aws.Config{
		Region:      aws.String("us-east-1"),                                              // Substitua pela sua região desejada
		Credentials: credentials.NewStaticCredentials(aws_access_key, aws_secret_key, ""), // Substitua com suas chaves de acesso e segredo
	}))

	return sess

}

func (u *UseCaseAWSCloudAccount) getAwsEC2AllInstances(sess *session.Session) {

	// Create new EC2 client
	svc := ec2.New(sess)

	// Call to get detailed information on each instance
	result, err := svc.DescribeInstances(nil)
	if err != nil {
		fmt.Println("Error", err)
	}

	for _, reservations := range result.Reservations {
		for _, instance := range reservations.Instances {
			fmt.Println("Instance ID: " + *instance.InstanceId)
			fmt.Println("Instance Type: " + *instance.InstanceType)
			fmt.Println("Public IP Address: " + *instance.PublicIpAddress)
			fmt.Println("Private IP Address: " + *instance.PrivateIpAddress)
			fmt.Println("Instance State: " + *instance.State.Name)
			fmt.Println("DNS Name: " + *instance.PublicDnsName)
			fmt.Println("Key Name: " + *instance.KeyName)
			fmt.Println("AMI ID: " + *instance.ImageId)
			fmt.Println("Launch Time: " + (*instance.LaunchTime).Format("2006-01-02 15:04:05 Monday"))
			fmt.Println("Tags:")
			for _, tag := range instance.Tags {
				fmt.Println("  " + *tag.Key + ": " + *tag.Value)
			}
			fmt.Println("")
		}
	}

}
