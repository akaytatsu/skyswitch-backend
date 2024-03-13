package usecase_calendar_test

import (
	"app/entity"
	"app/mocks"
	usecase_calendar "app/usecase/calendar"
	"errors"
	"testing"
	"time"

	"github.com/go-co-op/gocron"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

type MockConfigure struct {
	mockIRepoCalendar       *mocks.MockIRepositoryCalendar
	mockUseCaseInstance     *mocks.MockIUseCaseInstance
	mockInfraCloudProvider  *mocks.MockICloudProvider
	mockUseCaseCloudAccount *mocks.MockIUsecaseCloudAccount
	mockUsecaseHoliday      *mocks.MockIUsecaseHoliday
	mockUsecaseLog          *mocks.MockIUsecaseLog
}

func configureMocks(ctrl *gomock.Controller) (*MockConfigure, *usecase_calendar.UsecaseCalendar) {
	mocks := &MockConfigure{
		mockIRepoCalendar:       mocks.NewMockIRepositoryCalendar(ctrl),
		mockUseCaseInstance:     mocks.NewMockIUseCaseInstance(ctrl),
		mockInfraCloudProvider:  mocks.NewMockICloudProvider(ctrl),
		mockUseCaseCloudAccount: mocks.NewMockIUsecaseCloudAccount(ctrl),
		mockUsecaseHoliday:      mocks.NewMockIUsecaseHoliday(ctrl),
		mockUsecaseLog:          mocks.NewMockIUsecaseLog(ctrl),
	}

	schedule := gocron.NewScheduler(time.UTC)

	u := usecase_calendar.NewService(
		mocks.mockIRepoCalendar,
		schedule,
		mocks.mockUseCaseInstance,
		mocks.mockInfraCloudProvider,
		mocks.mockUseCaseCloudAccount,
		mocks.mockUsecaseHoliday,
		mocks.mockUsecaseLog,
	)

	return mocks, u

}

func TestProcessInstance_ErrorConnectionCloud(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	instance := &entity.EntityInstance{
		InstanceID:   "testID",
		InstanceName: "testName",
		Active:       true,
		CloudAccount: entity.EntityCloudAccount{
			ID: 1,
		},
	}

	calendar := &entity.EntityCalendar{
		ID:         1,
		Active:     true,
		TypeAction: "on",
	}

	mocksConfig, u := configureMocks(ctrl)

	mocksConfig.mockInfraCloudProvider.EXPECT().Connect(instance.CloudAccount).Return(
		errors.New("error connecting to cloud provider"),
	)

	err := u.ProcessInstance(instance, calendar)

	assert.Error(t, err)
}

func TestProcessInstance_instanceActiveFalse(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	instance := &entity.EntityInstance{
		InstanceID:   "testID",
		InstanceName: "testName",
		Active:       false,
		CloudAccount: entity.EntityCloudAccount{
			ID: 1,
		},
	}

	calendar := &entity.EntityCalendar{
		ID:         1,
		Active:     true,
		TypeAction: "on",
	}

	mocksConfig, u := configureMocks(ctrl)

	mocksConfig.mockInfraCloudProvider.EXPECT().Connect(instance.CloudAccount).Return(nil)

	err := u.ProcessInstance(instance, calendar)

	assert.Error(t, err)
}

func TestProcessInstance_CalendarActiveFalse(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	instance := &entity.EntityInstance{
		InstanceID:   "testID",
		InstanceName: "testName",
		Active:       true,
		CloudAccount: entity.EntityCloudAccount{
			ID: 1,
		},
	}

	calendar := &entity.EntityCalendar{
		ID:         1,
		Active:     false,
		TypeAction: "on",
	}

	mocksConfig, u := configureMocks(ctrl)

	mocksConfig.mockInfraCloudProvider.EXPECT().Connect(instance.CloudAccount).Return(nil)

	err := u.ProcessInstance(instance, calendar)

	assert.Error(t, err)
}

func TestProcessInstance_IsHolidayValid(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	instance := &entity.EntityInstance{
		InstanceID:   "testID",
		InstanceName: "testName",
		Active:       true,
		CloudAccount: entity.EntityCloudAccount{
			ID: 1,
		},
	}

	calendar := &entity.EntityCalendar{
		ID:           1,
		Active:       true,
		TypeAction:   "on",
		ValidHoliday: false,
	}

	mocksConfig, u := configureMocks(ctrl)

	u.Now = func() time.Time {
		return time.Date(2024, 1, 1, 0, 0, 0, 0, time.UTC)
	}

	mocksConfig.mockInfraCloudProvider.EXPECT().Connect(instance.CloudAccount).Return(nil)
	mocksConfig.mockUsecaseHoliday.EXPECT().IsHoliday(gomock.Any()).Return(true, nil)

	err := u.ProcessInstance(instance, calendar)

	assert.Error(t, err)
	assert.Equal(t, "today is holiday", err.Error())
}

func TestProcessInstance_TypeActionONError(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	instance := &entity.EntityInstance{
		InstanceID:   "testID",
		InstanceName: "testName",
		Active:       true,
		CloudAccount: entity.EntityCloudAccount{
			ID: 1,
		},
	}

	calendar := &entity.EntityCalendar{
		ID:         1,
		Active:     true,
		TypeAction: "on",
	}

	mocksConfig, u := configureMocks(ctrl)

	u.Now = func() time.Time {
		return time.Date(2024, 1, 1, 0, 0, 0, 0, time.UTC)
	}

	mocksConfig.mockInfraCloudProvider.EXPECT().Connect(instance.CloudAccount).Return(nil)
	mocksConfig.mockUsecaseHoliday.EXPECT().IsHoliday(gomock.Any()).Return(false, nil)
	mocksConfig.mockInfraCloudProvider.EXPECT().StartInstance(instance.InstanceID).Return(errors.New("error starting instance"))
	mocksConfig.mockUsecaseLog.EXPECT().Create(gomock.Any()).Return(nil)

	err := u.ProcessInstance(instance, calendar)

	assert.Error(t, err)
	assert.Equal(t, "error starting instance", err.Error())
}

func TestProcessInstance_TypeActionOFFNError(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	instance := &entity.EntityInstance{
		InstanceID:   "testID",
		InstanceName: "testName",
		Active:       true,
		CloudAccount: entity.EntityCloudAccount{
			ID: 1,
		},
	}

	calendar := &entity.EntityCalendar{
		ID:         1,
		Active:     true,
		TypeAction: "off",
	}

	mocksConfig, u := configureMocks(ctrl)

	mocksConfig.mockInfraCloudProvider.EXPECT().Connect(instance.CloudAccount).Return(nil)
	mocksConfig.mockUsecaseHoliday.EXPECT().IsHoliday(gomock.Any()).Return(false, nil)
	mocksConfig.mockInfraCloudProvider.EXPECT().StopInstance(instance.InstanceID).Return(errors.New("error stopping instance"))
	mocksConfig.mockUsecaseLog.EXPECT().Create(gomock.Any()).Return(nil)

	err := u.ProcessInstance(instance, calendar)

	assert.Error(t, err)
	assert.Equal(t, "error stopping instance", err.Error())
}
