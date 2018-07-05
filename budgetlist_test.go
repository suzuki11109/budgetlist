package budgetlist

import "testing"

func TestQueryBudgetEmpty(t *testing.T) {
	budget := &Budget{}
	result := budget.QueryBudget("2018-06-01", "2018-06-30")

	if result != 0 {
		t.Error("want", 0)
		t.Error("got", result)
	}
}

func TestQueryBudgetFullMonth(t *testing.T) {
	budget := &Budget{
		"2018-06": 3000,
	}
	result := budget.QueryBudget("2018-06-01", "2018-06-30")

	if result != 3000 {
		t.Error("want", 3000)
		t.Error("got", result)
	}
}

func TestQueryBudgetTwoMonths(t *testing.T) {
	budget := &Budget{
		"2018-06": 3000,
		"2018-07": 3500,
	}
	result := budget.QueryBudget("2018-06-01", "2018-07-31")

	if result != 6500 {
		t.Error("want", 6500)
		t.Error("got", result)
	}
}

func TestQueryBudgetTwoNonConsecutiveMonths(t *testing.T) {
	budget := &Budget{
		"2018-06": 3000,
		"2018-08": 3500,
	}
	result := budget.QueryBudget("2018-06-01", "2018-08-31")

	if result != 6500 {
		t.Error("want", 6500)
		t.Error("got", result)
	}
}

func TestQueryBudgetPartialInOneMonth(t *testing.T) {
	budget := &Budget{
		"2018-06": 3000,
	}
	result := budget.QueryBudget("2018-06-01", "2018-06-15")

	if result != 1500 {
		t.Error("want", 1500)
		t.Error("got", result)
	}
}

func TestQueryBudgetPartialInTwoMonth(t *testing.T) {
	budget := &Budget{
		"2018-06": 3000,
		"2018-07": 3100,
	}
	result := budget.QueryBudget("2018-06-16", "2018-07-15")

	if result != 3000 {
		t.Error("want", 3000)
		t.Error("got", result)
	}
}

func TestQueryBudgetPartialInFeb(t *testing.T) {
	budget := &Budget{
		"2018-02": 2800,
		"2018-03": 3100,
	}
	result := budget.QueryBudget("2018-02-15", "2018-03-31")

	if result != 4500 {
		t.Error("want", 4500)
		t.Error("got", result)
	}
}

func TestQueryBudgetVerySmallBudget(t *testing.T) {
	budget := &Budget{
		"2018-03": 1,
	}
	result := budget.QueryBudget("2018-03-01", "2018-03-01")

	if result != 0.03 {
		t.Error("want", 0.03)
		t.Error("got", result)
	}
}
