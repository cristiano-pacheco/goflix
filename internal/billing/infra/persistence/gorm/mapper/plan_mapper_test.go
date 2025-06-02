package mapper_test

import (
	"testing"
	"time"

	"github.com/stretchr/testify/suite"

	"github.com/cristiano-pacheco/goflix/internal/billing/domain/model"
	"github.com/cristiano-pacheco/goflix/internal/billing/infra/persistence/gorm/entity"
	"github.com/cristiano-pacheco/goflix/internal/billing/infra/persistence/gorm/mapper"
)

type PlanMapperTestSuite struct {
	suite.Suite
	sut mapper.PlanMapper
}

func (s *PlanMapperTestSuite) SetupTest() {
	s.sut = mapper.NewPlanMapper()
}

func TestPlanMapperSuite(t *testing.T) {
	suite.Run(t, new(PlanMapperTestSuite))
}

func (s *PlanMapperTestSuite) TestToModel_ValidPlanEntityWithAllFields_ReturnsModel() {
	// Arrange
	now := time.Now().UTC()
	planEntity := entity.PlanEntity{
		ID:          123,
		Name:        "Premium Plan",
		Description: "Premium subscription plan",
		AmountCents: 2999,
		Currency:    "USD",
		Interval:    "Month",
		TrialPeriod: 7,
		CreatedAt:   now,
		UpdatedAt:   now,
	}

	// Act
	planModel, err := s.sut.ToModel(planEntity)

	// Assert
	s.Require().NoError(err)
	s.Equal(uint64(123), planModel.ID())

	nameModel := planModel.Name()
	s.Equal("Premium Plan", (&nameModel).String())

	s.NotNil(planModel.Description())
	s.Equal("Premium subscription plan", planModel.Description().String())

	amountModel := planModel.Amount()
	s.Equal(uint(2999), (&amountModel).Cents())

	currencyModel := planModel.Currency()
	s.Equal("USD", currencyModel.Code())

	intervalModel := planModel.Interval()
	s.Equal("Month", intervalModel.String())

	s.NotNil(planModel.TrialPeriod())
	s.Equal(uint(7), planModel.TrialPeriod().Days())

	s.Equal(now.Unix(), planModel.CreatedAt().Unix())
	s.Equal(now.Unix(), planModel.UpdatedAt().Unix())
}

func (s *PlanMapperTestSuite) TestToModel_ValidPlanEntityWithoutDescription_ReturnsModel() {
	// Arrange
	now := time.Now().UTC()
	planEntity := entity.PlanEntity{
		ID:          456,
		Name:        "Basic Plan",
		Description: "",
		AmountCents: 999,
		Currency:    "EUR",
		Interval:    "Year",
		TrialPeriod: 14,
		CreatedAt:   now,
		UpdatedAt:   now,
	}

	// Act
	planModel, err := s.sut.ToModel(planEntity)

	// Assert
	s.Require().NoError(err)
	s.Equal(uint64(456), planModel.ID())

	nameModel := planModel.Name()
	s.Equal("Basic Plan", (&nameModel).String())

	s.Nil(planModel.Description())

	amountModel := planModel.Amount()
	s.Equal(uint(999), (&amountModel).Cents())

	currencyModel := planModel.Currency()
	s.Equal("EUR", currencyModel.Code())

	intervalModel := planModel.Interval()
	s.Equal("Year", intervalModel.String())

	s.NotNil(planModel.TrialPeriod())
	s.Equal(uint(14), planModel.TrialPeriod().Days())
}

func (s *PlanMapperTestSuite) TestToModel_ValidPlanEntityWithoutTrialPeriod_ReturnsModel() {
	// Arrange
	now := time.Now().UTC()
	planEntity := entity.PlanEntity{
		ID:          789,
		Name:        "Enterprise Plan",
		Description: "Enterprise subscription plan",
		AmountCents: 9999,
		Currency:    "USD",
		Interval:    "Month",
		TrialPeriod: 0,
		CreatedAt:   now,
		UpdatedAt:   now,
	}

	// Act
	planModel, err := s.sut.ToModel(planEntity)

	// Assert
	s.Require().NoError(err)
	s.Equal(uint64(789), planModel.ID())

	nameModel := planModel.Name()
	s.Equal("Enterprise Plan", (&nameModel).String())

	s.NotNil(planModel.Description())
	s.Equal("Enterprise subscription plan", planModel.Description().String())

	s.Nil(planModel.TrialPeriod())
}

func (s *PlanMapperTestSuite) TestToModel_PlanEntityWithDifferentValidIntervals_ReturnsModel() {
	// Arrange
	now := time.Now().UTC()
	validIntervals := []string{"Day", "Week", "Month", "Year"}

	for _, interval := range validIntervals {
		planEntity := entity.PlanEntity{
			ID:          100,
			Name:        "Test Plan",
			Description: "Test description",
			AmountCents: 1999,
			Currency:    "USD",
			Interval:    interval,
			TrialPeriod: 7,
			CreatedAt:   now,
			UpdatedAt:   now,
		}

		// Act
		planModel, err := s.sut.ToModel(planEntity)

		// Assert
		s.Require().NoError(err)
		intervalModel := planModel.Interval()
		s.Equal(interval, intervalModel.String())
	}
}

func (s *PlanMapperTestSuite) TestToModel_PlanEntityWithInvalidName_ReturnsError() {
	// Arrange
	now := time.Now().UTC()
	planEntity := entity.PlanEntity{
		ID:          123,
		Name:        "A",
		Description: "Valid description",
		AmountCents: 999,
		Currency:    "USD",
		Interval:    "Month",
		TrialPeriod: 7,
		CreatedAt:   now,
		UpdatedAt:   now,
	}

	// Act
	_, err := s.sut.ToModel(planEntity)

	// Assert
	s.Require().Error(err)
	s.Contains(err.Error(), "plan name must be at least 2 characters long")
}

func (s *PlanMapperTestSuite) TestToModel_PlanEntityWithInvalidCurrency_ReturnsError() {
	// Arrange
	now := time.Now().UTC()
	planEntity := entity.PlanEntity{
		ID:          123,
		Name:        "Valid Plan",
		Description: "Valid description",
		AmountCents: 999,
		Currency:    "INVALID",
		Interval:    "Month",
		TrialPeriod: 7,
		CreatedAt:   now,
		UpdatedAt:   now,
	}

	// Act
	_, err := s.sut.ToModel(planEntity)

	// Assert
	s.Require().Error(err)
	s.Contains(err.Error(), "currency code must be exactly 3 characters")
}

func (s *PlanMapperTestSuite) TestToModel_PlanEntityWithInvalidInterval_ReturnsError() {
	// Arrange
	now := time.Now().UTC()
	planEntity := entity.PlanEntity{
		ID:          123,
		Name:        "Valid Plan",
		Description: "Valid description",
		AmountCents: 999,
		Currency:    "USD",
		Interval:    "Invalid",
		TrialPeriod: 7,
		CreatedAt:   now,
		UpdatedAt:   now,
	}

	// Act
	_, err := s.sut.ToModel(planEntity)

	// Assert
	s.Require().Error(err)
	s.Contains(err.Error(), "invalid plan interval")
}

func (s *PlanMapperTestSuite) TestToEntity_ValidPlanModelWithAllFields_ReturnsEntity() {
	// Arrange
	now := time.Now().UTC()
	trialPeriod := uint(7)
	planModel, err := model.RestorePlanModel(
		123,
		"Premium Plan",
		"Premium subscription plan",
		"USD",
		"Month",
		2999,
		&trialPeriod,
		now,
		now,
	)
	s.Require().NoError(err)

	// Act
	planEntity := s.sut.ToEntity(planModel)

	// Assert
	s.Equal(uint64(123), planEntity.ID)
	s.Equal("Premium Plan", planEntity.Name)
	s.Equal("Premium subscription plan", planEntity.Description)
	s.Equal(uint(2999), planEntity.AmountCents)
	s.Equal("USD", planEntity.Currency)
	s.Equal("Month", planEntity.Interval)
	s.Equal(uint(7), planEntity.TrialPeriod)
	s.Equal(now.Unix(), planEntity.CreatedAt.Unix())
	s.Equal(now.Unix(), planEntity.UpdatedAt.Unix())
}

func (s *PlanMapperTestSuite) TestToEntity_ValidPlanModelWithoutDescription_ReturnsEntity() {
	// Arrange
	now := time.Now().UTC()
	trialPeriod := uint(14)
	planModel, err := model.RestorePlanModel(
		456,
		"Basic Plan",
		"",
		"EUR",
		"Year",
		999,
		&trialPeriod,
		now,
		now,
	)
	s.Require().NoError(err)

	// Act
	planEntity := s.sut.ToEntity(planModel)

	// Assert
	s.Equal(uint64(456), planEntity.ID)
	s.Equal("Basic Plan", planEntity.Name)
	s.Empty(planEntity.Description)
	s.Equal(uint(999), planEntity.AmountCents)
	s.Equal("EUR", planEntity.Currency)
	s.Equal("Year", planEntity.Interval)
	s.Equal(uint(14), planEntity.TrialPeriod)
}

func (s *PlanMapperTestSuite) TestToEntity_ValidPlanModelWithoutTrialPeriod_ReturnsEntity() {
	// Arrange
	now := time.Now().UTC()
	planModel, err := model.RestorePlanModel(
		789,
		"Enterprise Plan",
		"Enterprise subscription plan",
		"USD",
		"Month",
		9999,
		nil,
		now,
		now,
	)
	s.Require().NoError(err)

	// Act
	planEntity := s.sut.ToEntity(planModel)

	// Assert
	s.Equal(uint64(789), planEntity.ID)
	s.Equal("Enterprise Plan", planEntity.Name)
	s.Equal("Enterprise subscription plan", planEntity.Description)
	s.Equal(uint(9999), planEntity.AmountCents)
	s.Equal("USD", planEntity.Currency)
	s.Equal("Month", planEntity.Interval)
	s.Equal(uint(0), planEntity.TrialPeriod)
}

func (s *PlanMapperTestSuite) TestToEntity_PlanModelWithDifferentValidIntervals_ReturnsEntity() {
	// Arrange
	now := time.Now().UTC()
	validIntervals := []string{"Day", "Week", "Month", "Year"}

	for _, interval := range validIntervals {
		planModel, err := model.RestorePlanModel(
			100,
			"Test Plan",
			"Test description",
			"USD",
			interval,
			1999,
			nil,
			now,
			now,
		)
		s.Require().NoError(err)

		// Act
		planEntity := s.sut.ToEntity(planModel)

		// Assert
		s.Equal(interval, planEntity.Interval)
	}
}

func (s *PlanMapperTestSuite) TestToEntity_PlanModelWithZeroID_ReturnsEntity() {
	// Arrange
	planModel, err := model.CreatePlanModel(
		"New Plan",
		"New plan description",
		"USD",
		"Month",
		1999,
		nil,
	)
	s.Require().NoError(err)

	// Act
	planEntity := s.sut.ToEntity(planModel)

	// Assert
	s.Equal(uint64(0), planEntity.ID)
	s.Equal("New Plan", planEntity.Name)
	s.Equal("New plan description", planEntity.Description)
	s.Equal(uint(1999), planEntity.AmountCents)
	s.Equal("USD", planEntity.Currency)
	s.Equal("Month", planEntity.Interval)
	s.Equal(uint(0), planEntity.TrialPeriod)
}

func (s *PlanMapperTestSuite) TestRoundTrip_EntityToModelToEntity_PreservesData() {
	// Arrange
	now := time.Now().UTC()
	originalEntity := entity.PlanEntity{
		ID:          123,
		Name:        "Round Trip Plan",
		Description: "Round trip test description",
		AmountCents: 2999,
		Currency:    "USD",
		Interval:    "Month",
		TrialPeriod: 7,
		CreatedAt:   now,
		UpdatedAt:   now,
	}

	// Act
	planModel, err := s.sut.ToModel(originalEntity)
	s.Require().NoError(err)

	resultEntity := s.sut.ToEntity(planModel)

	// Assert
	s.Equal(originalEntity.ID, resultEntity.ID)
	s.Equal(originalEntity.Name, resultEntity.Name)
	s.Equal(originalEntity.Description, resultEntity.Description)
	s.Equal(originalEntity.AmountCents, resultEntity.AmountCents)
	s.Equal(originalEntity.Currency, resultEntity.Currency)
	s.Equal(originalEntity.Interval, resultEntity.Interval)
	s.Equal(originalEntity.TrialPeriod, resultEntity.TrialPeriod)
	s.Equal(originalEntity.CreatedAt.Unix(), resultEntity.CreatedAt.Unix())
	s.Equal(originalEntity.UpdatedAt.Unix(), resultEntity.UpdatedAt.Unix())
}

func (s *PlanMapperTestSuite) TestRoundTrip_ModelToEntityToModel_PreservesData() {
	// Arrange
	now := time.Now().UTC()
	trialPeriod := uint(14)
	originalModel, err := model.RestorePlanModel(
		456,
		"Round Trip Model",
		"Model round trip test",
		"EUR",
		"Year",
		9999,
		&trialPeriod,
		now,
		now,
	)
	s.Require().NoError(err)

	// Act
	planEntity := s.sut.ToEntity(originalModel)
	resultModel, err := s.sut.ToModel(planEntity)

	// Assert
	s.Require().NoError(err)
	s.Equal(originalModel.ID(), resultModel.ID())

	originalName := originalModel.Name()
	resultName := resultModel.Name()
	s.Equal((&originalName).String(), (&resultName).String())

	s.NotNil(originalModel.Description())
	s.NotNil(resultModel.Description())
	s.Equal(originalModel.Description().String(), resultModel.Description().String())

	originalAmount := originalModel.Amount()
	resultAmount := resultModel.Amount()
	s.Equal((&originalAmount).Cents(), (&resultAmount).Cents())

	originalCurrency := originalModel.Currency()
	resultCurrency := resultModel.Currency()
	s.Equal(originalCurrency.Code(), resultCurrency.Code())

	originalInterval := originalModel.Interval()
	resultInterval := resultModel.Interval()
	s.Equal(originalInterval.String(), resultInterval.String())

	s.NotNil(originalModel.TrialPeriod())
	s.NotNil(resultModel.TrialPeriod())
	s.Equal(originalModel.TrialPeriod().Days(), resultModel.TrialPeriod().Days())

	s.Equal(originalModel.CreatedAt().Unix(), resultModel.CreatedAt().Unix())
	s.Equal(originalModel.UpdatedAt().Unix(), resultModel.UpdatedAt().Unix())
}
