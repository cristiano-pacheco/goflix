package mapper

import (
	"time"

	"github.com/cristiano-pacheco/goflix/internal/billing/domain/enum"
)

type EndDateMapper interface {
	Map(startDate time.Time, planInterval enum.PlanIntervalEnum) *time.Time
}

type endDateMapper struct {
}

func NewEndDateMapper() EndDateMapper {
	return &endDateMapper{}
}

func (m *endDateMapper) Map(startDate time.Time, planInterval enum.PlanIntervalEnum) *time.Time {
	const (
		daysInWeek = 7
		daysInDay  = 1
	)
	var endDate *time.Time

	switch planInterval.String() {
	case enum.EnumPlanIntervalMonth:
		// add 1 month
		end := startDate.AddDate(0, 1, 0)
		endDate = &end
	case enum.EnumPlanIntervalYear:
		// add 1 year
		end := startDate.AddDate(1, 0, 0)
		endDate = &end
	case enum.EnumPlanIntervalWeek:
		// add 1 week
		end := startDate.AddDate(0, 0, daysInWeek)
		endDate = &end
	case enum.EnumPlanIntervalDay:
		// add 1 day
		end := startDate.AddDate(0, 0, daysInDay)
		endDate = &end
	default:
		// For lifetime or other intervals, leave endDate as nil
		endDate = nil
	}

	return endDate
}
