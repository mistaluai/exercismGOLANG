package expenses

import "errors"

// Record represents an expense record.
type Record struct {
	Day      int
	Amount   float64
	Category string
}

// DaysPeriod represents a period of days for expenses.
type DaysPeriod struct {
	From int
	To   int
}

// Filter returns the records for which the predicate function returns true.
func Filter(in []Record, predicate func(Record) bool) []Record {
	var records []Record = make([]Record, 0)
	for _, r := range in {
		if predicate(r) {
			records = append(records, r)
		}
	}
	return records
}

// ByDaysPeriod returns predicate function that returns true when
// the day of the record is inside the period of day and false otherwise.
func ByDaysPeriod(p DaysPeriod) func(Record) bool {
	return func(record Record) bool {
		d := record.Day
		return d >= p.From && d <= p.To
	}
}

// ByCategory returns predicate function that returns true when
// the category of the record is the same as the provided category
// and false otherwise.
func ByCategory(c string) func(Record) bool {
	return func(record Record) bool {
		return record.Category == c
	}
}

// TotalByPeriod returns total amount of expenses for records
// inside the period p.
func TotalByPeriod(in []Record, p DaysPeriod) float64 {
	records := Filter(in, ByDaysPeriod(p))
	var expenses float64 = 0
	for _, r := range records {
		expenses += r.Amount
	}
	return expenses
}

// CategoryExpenses returns total amount of expenses for records
// in category c that are also inside the period p.
// An error must be returned only if there are no records in the list that belong
// to the given category, regardless of period of time.
func CategoryExpenses(in []Record, p DaysPeriod, c string) (float64, error) {
	var expenses float64 = 0
	var hasExpensesCategory bool
	var err error
	for _, r := range in {
		if !hasExpensesCategory {
			hasExpensesCategory = r.Category == c
		}
		if r.Category == c && ByDaysPeriod(p)(r) == true {
			expenses += r.Amount
		}
	}
	if !hasExpensesCategory {
		err = errors.New("unknown category " + c)
	}
	return expenses, err
}
