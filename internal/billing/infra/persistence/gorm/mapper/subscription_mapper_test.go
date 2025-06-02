package mapper_test

import (
	"testing"
	"time"

	"github.com/stretchr/testify/suite"

	"github.com/cristiano-pacheco/goflix/internal/billing/domain/model"
	"github.com/cristiano-pacheco/goflix/internal/billing/infra/persistence/gorm/entity"
	"github.com/cristiano-pacheco/goflix/internal/billing/infra/persistence/gorm/mapper"
)

type SubscriptionMapperTestSuite struct {
	suite.Suite
	sut mapper.SubscriptionMapper
}

func (s *SubscriptionMapperTestSuite) SetupTest() {
	s.sut = mapper.NewSubscriptionMapper()
}

func TestSubscriptionMapperSuite(t *testing.T) {
	suite.Run(t, new(SubscriptionMapperTestSuite))
}

func (s *SubscriptionMapperTestSuite) TestToModel_ValidSubscriptionEntityWithAllFields_ReturnsModel() {
	// Arrange
	now := time.Now().UTC()
	endDate := now.Add(30 * 24 * time.Hour)
	subscriptionEntity := entity.SubscriptionEntity{
		ID:        123,
		UserID:    456,
		PlanID:    789,
		Status:    "Active",
		StartDate: now,
		EndDate:   endDate,
		AutoRenew: true,
		CreatedAt: now,
		UpdatedAt: now,
	}

	// Act
	subscriptionModel, err := s.sut.ToModel(subscriptionEntity)

	// Assert
	s.Require().NoError(err)
	s.Equal(uint64(123), subscriptionModel.ID())
	s.Equal(uint64(456), subscriptionModel.UserID())
	s.Equal(uint64(789), subscriptionModel.PlanID())

	statusEnum := subscriptionModel.Status()
	s.Equal("Active", (&statusEnum).String())

	s.Equal(now.Unix(), subscriptionModel.StartDate().Unix())
	s.NotNil(subscriptionModel.EndDate())
	s.Equal(endDate.Unix(), subscriptionModel.EndDate().Unix())
	s.True(subscriptionModel.AutoRenew())
	s.Equal(now.Unix(), subscriptionModel.CreatedAt().Unix())
	s.Equal(now.Unix(), subscriptionModel.UpdatedAt().Unix())
}

func (s *SubscriptionMapperTestSuite) TestToModel_ValidSubscriptionEntityWithoutEndDate_ReturnsModel() {
	// Arrange
	now := time.Now().UTC()
	subscriptionEntity := entity.SubscriptionEntity{
		ID:        123,
		UserID:    456,
		PlanID:    789,
		Status:    "Active",
		StartDate: now,
		EndDate:   time.Time{},
		AutoRenew: false,
		CreatedAt: now,
		UpdatedAt: now,
	}

	// Act
	subscriptionModel, err := s.sut.ToModel(subscriptionEntity)

	// Assert
	s.Require().NoError(err)
	s.Equal(uint64(123), subscriptionModel.ID())
	s.Equal(uint64(456), subscriptionModel.UserID())
	s.Equal(uint64(789), subscriptionModel.PlanID())

	statusEnum := subscriptionModel.Status()
	s.Equal("Active", (&statusEnum).String())

	s.Equal(now.Unix(), subscriptionModel.StartDate().Unix())
	s.Nil(subscriptionModel.EndDate())
	s.False(subscriptionModel.AutoRenew())
}

func (s *SubscriptionMapperTestSuite) TestToModel_SubscriptionEntityWithDifferentValidStatuses_ReturnsModel() {
	// Arrange
	now := time.Now().UTC()
	validStatuses := []string{"Active", "Inactive", "Cancelled", "Expired", "PastDue"}

	for _, status := range validStatuses {
		subscriptionEntity := entity.SubscriptionEntity{
			ID:        100,
			UserID:    200,
			PlanID:    300,
			Status:    status,
			StartDate: now,
			EndDate:   time.Time{},
			AutoRenew: true,
			CreatedAt: now,
			UpdatedAt: now,
		}

		// Act
		subscriptionModel, err := s.sut.ToModel(subscriptionEntity)

		// Assert
		s.Require().NoError(err)
		statusEnum := subscriptionModel.Status()
		s.Equal(status, (&statusEnum).String())
	}
}

func (s *SubscriptionMapperTestSuite) TestToModel_SubscriptionEntityWithZeroUserID_ReturnsError() {
	// Arrange
	now := time.Now().UTC()
	subscriptionEntity := entity.SubscriptionEntity{
		ID:        123,
		UserID:    0,
		PlanID:    789,
		Status:    "Active",
		StartDate: now,
		EndDate:   time.Time{},
		AutoRenew: true,
		CreatedAt: now,
		UpdatedAt: now,
	}

	// Act
	_, err := s.sut.ToModel(subscriptionEntity)

	// Assert
	s.Require().Error(err)
	s.Contains(err.Error(), "user ID is required")
}

func (s *SubscriptionMapperTestSuite) TestToModel_SubscriptionEntityWithZeroPlanID_ReturnsError() {
	// Arrange
	now := time.Now().UTC()
	subscriptionEntity := entity.SubscriptionEntity{
		ID:        123,
		UserID:    456,
		PlanID:    0,
		Status:    "Active",
		StartDate: now,
		EndDate:   time.Time{},
		AutoRenew: true,
		CreatedAt: now,
		UpdatedAt: now,
	}

	// Act
	_, err := s.sut.ToModel(subscriptionEntity)

	// Assert
	s.Require().Error(err)
	s.Contains(err.Error(), "plan ID is required")
}

func (s *SubscriptionMapperTestSuite) TestToModel_SubscriptionEntityWithInvalidStatus_ReturnsError() {
	// Arrange
	now := time.Now().UTC()
	subscriptionEntity := entity.SubscriptionEntity{
		ID:        123,
		UserID:    456,
		PlanID:    789,
		Status:    "InvalidStatus",
		StartDate: now,
		EndDate:   time.Time{},
		AutoRenew: true,
		CreatedAt: now,
		UpdatedAt: now,
	}

	// Act
	_, err := s.sut.ToModel(subscriptionEntity)

	// Assert
	s.Require().Error(err)
	s.Contains(err.Error(), "invalid subscription status")
}

func (s *SubscriptionMapperTestSuite) TestToModel_SubscriptionEntityWithZeroStartDate_ReturnsError() {
	// Arrange
	subscriptionEntity := entity.SubscriptionEntity{
		ID:        123,
		UserID:    456,
		PlanID:    789,
		Status:    "Active",
		StartDate: time.Time{},
		EndDate:   time.Time{},
		AutoRenew: true,
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
	}

	// Act
	_, err := s.sut.ToModel(subscriptionEntity)

	// Assert
	s.Require().Error(err)
	s.Contains(err.Error(), "start date is required")
}

func (s *SubscriptionMapperTestSuite) TestToEntity_ValidSubscriptionModelWithAllFields_ReturnsEntity() {
	// Arrange
	now := time.Now().UTC()
	endDate := now.Add(30 * 24 * time.Hour)
	subscriptionModel, err := model.RestoreSubscriptionModel(
		123,
		456,
		789,
		"Active",
		now,
		&endDate,
		true,
		now,
		now,
	)
	s.Require().NoError(err)

	// Act
	subscriptionEntity := s.sut.ToEntity(subscriptionModel)

	// Assert
	s.Equal(uint64(123), subscriptionEntity.ID)
	s.Equal(uint64(456), subscriptionEntity.UserID)
	s.Equal(uint64(789), subscriptionEntity.PlanID)
	s.Equal("Active", subscriptionEntity.Status)
	s.Equal(now.Unix(), subscriptionEntity.StartDate.Unix())
	s.Equal(endDate.Unix(), subscriptionEntity.EndDate.Unix())
	s.True(subscriptionEntity.AutoRenew)
	s.Equal(now.Unix(), subscriptionEntity.CreatedAt.Unix())
	s.Equal(now.Unix(), subscriptionEntity.UpdatedAt.Unix())
}

func (s *SubscriptionMapperTestSuite) TestToEntity_ValidSubscriptionModelWithoutEndDate_ReturnsEntity() {
	// Arrange
	now := time.Now().UTC()
	subscriptionModel, err := model.RestoreSubscriptionModel(
		123,
		456,
		789,
		"Cancelled",
		now,
		nil,
		false,
		now,
		now,
	)
	s.Require().NoError(err)

	// Act
	subscriptionEntity := s.sut.ToEntity(subscriptionModel)

	// Assert
	s.Equal(uint64(123), subscriptionEntity.ID)
	s.Equal(uint64(456), subscriptionEntity.UserID)
	s.Equal(uint64(789), subscriptionEntity.PlanID)
	s.Equal("Cancelled", subscriptionEntity.Status)
	s.Equal(now.Unix(), subscriptionEntity.StartDate.Unix())
	s.True(subscriptionEntity.EndDate.IsZero())
	s.False(subscriptionEntity.AutoRenew)
}

func (s *SubscriptionMapperTestSuite) TestToEntity_SubscriptionModelWithDifferentValidStatuses_ReturnsEntity() {
	// Arrange
	now := time.Now().UTC()
	validStatuses := []string{"Active", "Inactive", "Cancelled", "Expired", "PastDue"}

	for _, status := range validStatuses {
		subscriptionModel, err := model.RestoreSubscriptionModel(
			100,
			200,
			300,
			status,
			now,
			nil,
			true,
			now,
			now,
		)
		s.Require().NoError(err)

		// Act
		subscriptionEntity := s.sut.ToEntity(subscriptionModel)

		// Assert
		s.Equal(status, subscriptionEntity.Status)
	}
}

func (s *SubscriptionMapperTestSuite) TestToEntity_SubscriptionModelWithZeroID_ReturnsEntity() {
	// Arrange
	now := time.Now().UTC()
	subscriptionModel, err := model.CreateSubscriptionModel(
		456,
		789,
		now,
		nil,
	)
	s.Require().NoError(err)

	// Act
	subscriptionEntity := s.sut.ToEntity(subscriptionModel)

	// Assert
	s.Equal(uint64(0), subscriptionEntity.ID)
	s.Equal(uint64(456), subscriptionEntity.UserID)
	s.Equal(uint64(789), subscriptionEntity.PlanID)
	s.Equal("Active", subscriptionEntity.Status)
	s.Equal(now.Unix(), subscriptionEntity.StartDate.Unix())
	s.True(subscriptionEntity.EndDate.IsZero())
	s.True(subscriptionEntity.AutoRenew)
}

func (s *SubscriptionMapperTestSuite) TestRoundTrip_EntityToModelToEntity_PreservesData() {
	// Arrange
	now := time.Now().UTC()
	endDate := now.Add(30 * 24 * time.Hour)
	originalEntity := entity.SubscriptionEntity{
		ID:        123,
		UserID:    456,
		PlanID:    789,
		Status:    "Active",
		StartDate: now,
		EndDate:   endDate,
		AutoRenew: true,
		CreatedAt: now,
		UpdatedAt: now,
	}

	// Act
	subscriptionModel, err := s.sut.ToModel(originalEntity)
	s.Require().NoError(err)

	resultEntity := s.sut.ToEntity(subscriptionModel)

	// Assert
	s.Equal(originalEntity.ID, resultEntity.ID)
	s.Equal(originalEntity.UserID, resultEntity.UserID)
	s.Equal(originalEntity.PlanID, resultEntity.PlanID)
	s.Equal(originalEntity.Status, resultEntity.Status)
	s.Equal(originalEntity.StartDate.Unix(), resultEntity.StartDate.Unix())
	s.Equal(originalEntity.EndDate.Unix(), resultEntity.EndDate.Unix())
	s.Equal(originalEntity.AutoRenew, resultEntity.AutoRenew)
	s.Equal(originalEntity.CreatedAt.Unix(), resultEntity.CreatedAt.Unix())
	s.Equal(originalEntity.UpdatedAt.Unix(), resultEntity.UpdatedAt.Unix())
}

func (s *SubscriptionMapperTestSuite) TestRoundTrip_ModelToEntityToModel_PreservesData() {
	// Arrange
	now := time.Now().UTC()
	endDate := now.Add(60 * 24 * time.Hour)
	originalModel, err := model.RestoreSubscriptionModel(
		456,
		789,
		123,
		"Expired",
		now,
		&endDate,
		false,
		now,
		now,
	)
	s.Require().NoError(err)

	// Act
	subscriptionEntity := s.sut.ToEntity(originalModel)
	resultModel, err := s.sut.ToModel(subscriptionEntity)

	// Assert
	s.Require().NoError(err)
	s.Equal(originalModel.ID(), resultModel.ID())
	s.Equal(originalModel.UserID(), resultModel.UserID())
	s.Equal(originalModel.PlanID(), resultModel.PlanID())

	originalStatus := originalModel.Status()
	resultStatus := resultModel.Status()
	s.Equal((&originalStatus).String(), (&resultStatus).String())

	s.Equal(originalModel.StartDate().Unix(), resultModel.StartDate().Unix())

	s.NotNil(originalModel.EndDate())
	s.NotNil(resultModel.EndDate())
	s.Equal(originalModel.EndDate().Unix(), resultModel.EndDate().Unix())

	s.Equal(originalModel.AutoRenew(), resultModel.AutoRenew())
	s.Equal(originalModel.CreatedAt().Unix(), resultModel.CreatedAt().Unix())
	s.Equal(originalModel.UpdatedAt().Unix(), resultModel.UpdatedAt().Unix())
}
