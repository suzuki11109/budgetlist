package budgetlist

import "testing"

func TestQueryBudget(t *testing.T) {
	tests := []struct {
		name   string
		budget Budget
		start  string
		end    string
		out    float64
	}{
		{
			name:   "Zero",
			budget: Budget{},
			start:  "2018-06-01",
			end:    "2018-06-30",
			out:    0,
		},
		{
			name:   "Full month",
			budget: Budget{"2018-06": 3000},
			start:  "2018-06-01",
			end:    "2018-06-30",
			out:    3000,
		},
		{
			name: "Two months",
			budget: Budget{
				"2018-06": 3000,
				"2018-07": 3500,
			},
			start: "2018-06-01",
			end:   "2018-07-31",
			out:   6500,
		},
		{
			name: "Two NonConsecutive Months",
			budget: Budget{
				"2018-06": 3000,
				"2018-08": 3500,
			},
			start: "2018-06-01",
			end:   "2018-08-31",
			out:   6500,
		},
		{
			name: "Partial In One Month",
			budget: Budget{
				"2018-06": 3000,
			},
			start: "2018-06-01",
			end:   "2018-06-15",
			out:   1500,
		},
		{
			name: "Partial In Two Month",
			budget: Budget{
				"2018-06": 3000,
				"2018-07": 3100,
			},
			start: "2018-06-16",
			end:   "2018-07-15",
			out:   3000,
		},
		{
			name: "Partial In Feb",
			budget: Budget{
				"2018-02": 2800,
				"2018-03": 3100,
			},
			start: "2018-02-15",
			end:   "2018-03-31",
			out:   4500,
		},
		{
			name: "Very Small Budget",
			budget: Budget{
				"2018-03": 1,
			},
			start: "2018-03-01",
			end:   "2018-03-01",
			out:   0.03,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := tt.budget.QueryBudget(tt.start, tt.end)
			if result != tt.out {
				t.Error("want", tt.out)
				t.Error("got", result)
			}
		})
	}
}
