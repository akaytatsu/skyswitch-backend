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
		ValidHoliday: true,
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

// func TestProccessCalendar(t *testing.T) {
// 	ctrl := gomock.NewController(t)
// 	defer ctrl.Finish()

// 	mockUsecaseInstance := mocks.NewMockIUseCaseInstance(ctrl)
// 	mockInfraCloudProvider := mocks.NewMockICloudProvider(ctrl)
// 	mockUsecaseHoliday := mocks.NewMockIUsecaseHoliday(ctrl)
// 	mockUsecaseLog := mocks.NewMockIUsecaseLog(ctrl)

// 	calendar := &entity.EntityCalendar{
// 		ID:         1,
// 		Active:     true,
// 		TypeAction: "on",
// 	}

// 	instance := &entity.EntityInstance{
// 		InstanceID:   "testID",
// 		InstanceName: "testName",
// 		Active:       true,
// 		CloudAccount: entity.EntityCloudAccount{
// 			ID: 1,
// 		},
// 	}

// 	mockUsecaseInstance.EXPECT().GetAllOFCalendar(calendar.ID).Return([]*entity.EntityInstance{instance}, nil)
// 	mockInfraCloudProvider.EXPECT().Connect(instance.CloudAccount).Return(nil)
// 	mockUsecaseHoliday.EXPECT().IsHoliday(gomock.Any()).Return(false, nil)
// 	mockInfraCloudProvider.EXPECT().StartInstance(instance.InstanceID).Return(nil)
// 	mockUsecaseLog.EXPECT().Create(gomock.Any()).Return(nil)

// 	u := &usecase.UsecaseCalendar{
// 		usecaseInstance:    mockUsecaseInstance,
// 		infraCloudProvider: mockInfraCloudProvider,
// 		usecaseHoliday:     mockUsecaseHoliday,
// 		usecaseLog:         mockUsecaseLog,
// 	}

// 	err := u.ProccessCalendar(calendar)
// 	if err != nil {
// 		t.Errorf("Expected no error, but got %v", err)
// 	}
// }
