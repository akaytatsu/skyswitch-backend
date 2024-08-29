package usecase_autoscalling_groups_test

import (
	"app/entity"
	"app/mocks"
	usecase_autoscalling_groups "app/usecase/autoscalling_groups"
	"errors"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

type MockConfigure struct {
	mockIRepoAutoScalingGroup *mocks.MockIRepositoryAutoScalingGroup
}

func configureMocks(ctrl *gomock.Controller) (*MockConfigure, *usecase_autoscalling_groups.UsecaseAutoScalingGroup) {
	mocks := &MockConfigure{
		mockIRepoAutoScalingGroup: mocks.NewMockIRepositoryAutoScalingGroup(ctrl),
	}

	u := usecase_autoscalling_groups.NewService(
		mocks.mockIRepoAutoScalingGroup,
	)

	return mocks, u

}

func TestCreateOrUpdateCreateWithNonZero(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mocks, u := configureMocks(ctrl)

	autoScallingGroup := &entity.EntityAutoScalingGroup{
		AutoScalingGroupID: "test",
		MinSize:            5,
		MaxSize:            10,
		DesiredCapacity:    5,
	}

	mocks.mockIRepoAutoScalingGroup.EXPECT().GetByID(autoScallingGroup.AutoScalingGroupID).Return(nil, errors.New("error"))

	mocks.mockIRepoAutoScalingGroup.EXPECT().Create(autoScallingGroup).Return(nil)

	err := u.CreateOrUpdate(autoScallingGroup, false)

	assert.Nil(t, err)
	assert.Equal(t, 5, autoScallingGroup.MinSize)
	assert.Equal(t, 10, autoScallingGroup.MaxSize)
	assert.Equal(t, 5, autoScallingGroup.DesiredCapacity)
}

func TestCreateOrUpdateCreateWithZero(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mocks, u := configureMocks(ctrl)

	autoScallingGroup := &entity.EntityAutoScalingGroup{
		AutoScalingGroupID: "test",
		MinSize:            0,
		MaxSize:            0,
		DesiredCapacity:    0,
	}

	mocks.mockIRepoAutoScalingGroup.EXPECT().GetByID(autoScallingGroup.AutoScalingGroupID).Return(nil, errors.New("error"))

	mocks.mockIRepoAutoScalingGroup.EXPECT().Create(autoScallingGroup).Return(nil)

	err := u.CreateOrUpdate(autoScallingGroup, false)

	assert.Nil(t, err)
	assert.Equal(t, 0, autoScallingGroup.MinSize)
	assert.Equal(t, 0, autoScallingGroup.MaxSize)
	assert.Equal(t, 0, autoScallingGroup.DesiredCapacity)
}

func TestCreateOrUpdateUpdateWithNonZero(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mocks, u := configureMocks(ctrl)

	autoScallingGroup := &entity.EntityAutoScalingGroup{
		AutoScalingGroupID: "test",
		MinSize:            0,
		MaxSize:            0,
		DesiredCapacity:    0,
	}

	mocks.mockIRepoAutoScalingGroup.EXPECT().GetByID(autoScallingGroup.AutoScalingGroupID).Return(&entity.EntityAutoScalingGroup{
		AutoScalingGroupID: "test",
		MinSize:            3,
		MaxSize:            7,
		DesiredCapacity:    9,
	}, nil)

	mocks.mockIRepoAutoScalingGroup.EXPECT().Update(autoScallingGroup, gomock.Any()).Return(nil)

	err := u.CreateOrUpdate(autoScallingGroup, false)

	assert.Nil(t, err)
	assert.Equal(t, 3, autoScallingGroup.MinSize)
	assert.Equal(t, 7, autoScallingGroup.MaxSize)
	assert.Equal(t, 9, autoScallingGroup.DesiredCapacity)
}

func TestCreateOrUpdateUpdateWithZero(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mocks, u := configureMocks(ctrl)

	autoScallingGroup := &entity.EntityAutoScalingGroup{
		AutoScalingGroupID: "test",
		MinSize:            0,
		MaxSize:            0,
		DesiredCapacity:    0,
	}

	mocks.mockIRepoAutoScalingGroup.EXPECT().GetByID(autoScallingGroup.AutoScalingGroupID).Return(&entity.EntityAutoScalingGroup{
		AutoScalingGroupID: "test",
		MinSize:            0,
		MaxSize:            0,
		DesiredCapacity:    0,
	}, nil)

	mocks.mockIRepoAutoScalingGroup.EXPECT().Update(autoScallingGroup, gomock.Any()).Return(nil)

	err := u.CreateOrUpdate(autoScallingGroup, false)

	assert.Nil(t, err)
	assert.Equal(t, 0, autoScallingGroup.MinSize)
	assert.Equal(t, 0, autoScallingGroup.MaxSize)
	assert.Equal(t, 0, autoScallingGroup.DesiredCapacity)
}
