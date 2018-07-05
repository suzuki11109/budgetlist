package budgetlist

import (
	"math"
	"time"
)

const oneDay = 24 * time.Hour

type Budget map[string]float64

func getDayInMonth(t time.Time) int {
	return time.Date(t.Year(), t.Month()+1, 0, 0, 0, 0, 0, time.UTC).Day()
}

func (b *Budget) QueryBudget(start, end string) float64 {
	startDate, _ := time.Parse("2006-01-02", start)
	endDate, _ := time.Parse("2006-01-02", end)

	var total float64

	for cursor := startDate; !cursor.After(endDate); cursor = cursor.Add(oneDay) {
		value := (*b)[cursor.Format("2006-01")]
		total += value / float64(getDayInMonth(cursor))
	}

	return math.Floor(total*100) / 100
}
