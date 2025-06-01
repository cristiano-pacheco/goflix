package mapper_test

import (
	"testing"
	"time"

	"github.com/stretchr/testify/suite"

	"github.com/cristiano-pacheco/goflix/internal/billing/domain/enum"
	"github.com/cristiano-pacheco/goflix/internal/billing/domain/mapper"
)

type EndDateMapperTestSuite struct {
	suite.Suite
	sut mapper.EndDateMapper
}

func (s *EndDateMapperTestSuite) SetupTest() {
	s.sut = mapper.NewEndDateMapper()
}

func TestEndDateMapperSuite(t *testing.T) {
	suite.Run(t, new(EndDateMapperTestSuite))
}

func (s *EndDateMapperTestSuite) TestMap_MonthInterval_AddsOneMonthToStartDate() {
	// Arrange
	startDate := time.Date(2024, 1, 15, 10, 30, 0, 0, time.UTC)
	planInterval, err := enum.NewPlanIntervalEnum(enum.EnumPlanIntervalMonth)
	s.Require().NoError(err)

	// Act
	result := s.sut.Map(startDate, planInterval)

	// Assert
	s.NotNil(result)
	expectedEndDate := time.Date(2024, 2, 15, 10, 30, 0, 0, time.UTC)
	s.Equal(expectedEndDate, *result)
}

func (s *EndDateMapperTestSuite) TestMap_YearInterval_AddsOneYearToStartDate() {
	// Arrange
	startDate := time.Date(2024, 3, 10, 14, 45, 30, 0, time.UTC)
	planInterval, err := enum.NewPlanIntervalEnum(enum.EnumPlanIntervalYear)
	s.Require().NoError(err)

	// Act
	result := s.sut.Map(startDate, planInterval)

	// Assert
	s.NotNil(result)
	expectedEndDate := time.Date(2025, 3, 10, 14, 45, 30, 0, time.UTC)
	s.Equal(expectedEndDate, *result)
}

func (s *EndDateMapperTestSuite) TestMap_WeekInterval_AddsSevenDaysToStartDate() {
	// Arrange
	startDate := time.Date(2024, 6, 1, 9, 0, 0, 0, time.UTC)
	planInterval, err := enum.NewPlanIntervalEnum(enum.EnumPlanIntervalWeek)
	s.Require().NoError(err)

	// Act
	result := s.sut.Map(startDate, planInterval)

	// Assert
	s.NotNil(result)
	expectedEndDate := time.Date(2024, 6, 8, 9, 0, 0, 0, time.UTC)
	s.Equal(expectedEndDate, *result)
}

func (s *EndDateMapperTestSuite) TestMap_DayInterval_AddsOneDayToStartDate() {
	// Arrange
	startDate := time.Date(2024, 12, 31, 23, 59, 59, 0, time.UTC)
	planInterval, err := enum.NewPlanIntervalEnum(enum.EnumPlanIntervalDay)
	s.Require().NoError(err)

	// Act
	result := s.sut.Map(startDate, planInterval)

	// Assert
	s.NotNil(result)
	expectedEndDate := time.Date(2025, 1, 1, 23, 59, 59, 0, time.UTC)
	s.Equal(expectedEndDate, *result)
}

func (s *EndDateMapperTestSuite) TestMap_MonthInterval_HandlesLeapYearCorrectly() {
	// Arrange
	startDate := time.Date(2024, 1, 29, 12, 0, 0, 0, time.UTC)
	planInterval, err := enum.NewPlanIntervalEnum(enum.EnumPlanIntervalMonth)
	s.Require().NoError(err)

	// Act
	result := s.sut.Map(startDate, planInterval)

	// Assert
	s.NotNil(result)
	expectedEndDate := time.Date(2024, 2, 29, 12, 0, 0, 0, time.UTC)
	s.Equal(expectedEndDate, *result)
}

func (s *EndDateMapperTestSuite) TestMap_MonthInterval_HandlesMonthWithFewerDays() {
	// Arrange
	startDate := time.Date(2024, 1, 31, 8, 15, 0, 0, time.UTC)
	planInterval, err := enum.NewPlanIntervalEnum(enum.EnumPlanIntervalMonth)
	s.Require().NoError(err)

	// Act
	result := s.sut.Map(startDate, planInterval)

	// Assert
	s.NotNil(result)
	expectedEndDate := time.Date(2024, 3, 2, 8, 15, 0, 0, time.UTC)
	s.Equal(expectedEndDate, *result)
}

func (s *EndDateMapperTestSuite) TestMap_YearInterval_HandlesLeapYearCorrectly() {
	// Arrange
	startDate := time.Date(2024, 2, 29, 16, 30, 0, 0, time.UTC)
	planInterval, err := enum.NewPlanIntervalEnum(enum.EnumPlanIntervalYear)
	s.Require().NoError(err)

	// Act
	result := s.sut.Map(startDate, planInterval)

	// Assert
	s.NotNil(result)
	expectedEndDate := time.Date(2025, 3, 1, 16, 30, 0, 0, time.UTC)
	s.Equal(expectedEndDate, *result)
}

func (s *EndDateMapperTestSuite) TestMap_WeekInterval_HandlesMonthBoundaryCorrectly() {
	// Arrange
	startDate := time.Date(2024, 4, 28, 11, 45, 0, 0, time.UTC)
	planInterval, err := enum.NewPlanIntervalEnum(enum.EnumPlanIntervalWeek)
	s.Require().NoError(err)

	// Act
	result := s.sut.Map(startDate, planInterval)

	// Assert
	s.NotNil(result)
	expectedEndDate := time.Date(2024, 5, 5, 11, 45, 0, 0, time.UTC)
	s.Equal(expectedEndDate, *result)
}

func (s *EndDateMapperTestSuite) TestMap_InvalidInterval_ReturnsNilEndDate() {
	// Arrange
	startDate := time.Date(2024, 5, 15, 10, 0, 0, 0, time.UTC)
	invalidInterval := enum.PlanIntervalEnum{}

	// Act
	result := s.sut.Map(startDate, invalidInterval)

	// Assert
	s.Nil(result)
}

func (s *EndDateMapperTestSuite) TestMap_ZeroStartDateWithMonthInterval_ReturnsCorrectEndDate() {
	// Arrange
	startDate := time.Time{}
	planInterval, err := enum.NewPlanIntervalEnum(enum.EnumPlanIntervalMonth)
	s.Require().NoError(err)

	// Act
	result := s.sut.Map(startDate, planInterval)

	// Assert
	s.NotNil(result)
	expectedEndDate := time.Date(1, 2, 1, 0, 0, 0, 0, time.UTC)
	s.Equal(expectedEndDate, *result)
}

func (s *EndDateMapperTestSuite) TestMap_AllValidIntervals_ProduceNonNilResults() {
	// Arrange
	startDate := time.Date(2024, 6, 15, 12, 0, 0, 0, time.UTC)
	validIntervals := []string{
		enum.EnumPlanIntervalDay,
		enum.EnumPlanIntervalWeek,
		enum.EnumPlanIntervalMonth,
		enum.EnumPlanIntervalYear,
	}

	for _, intervalValue := range validIntervals {
		// Arrange
		planInterval, err := enum.NewPlanIntervalEnum(intervalValue)
		s.Require().NoError(err)

		// Act
		result := s.sut.Map(startDate, planInterval)

		// Assert
		s.NotNil(result, "interval %s should produce non-nil result", intervalValue)
		s.True(
			result.After(startDate),
			"end date should be after start date for interval %s",
			intervalValue,
		)
	}
}
