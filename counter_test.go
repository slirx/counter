package main

import (
	"testing"
	"time"
)

func TestParse(t *testing.T) {
	type Input struct {
		date time.Time
		diff Diff
	}

	type TableTest struct {
		in   Input
		want string
	}

	table := []TableTest{
		{
			in: Input{
				date: time.Now().Add(time.Hour * 24 * 3),
				diff: DiffDays,
			},
			want: "3 days",
		},
		{
			in: Input{
				date: time.Now().Add(time.Hour * 24 * -2),
				diff: DiffDays,
			},
			want: "2 days",
		},
		{
			in: Input{
				date: time.Now().Add(time.Hour * 24 * 23),
				diff: DiffDays,
			},
			want: "23 days",
		},
		{
			in: Input{
				date: time.Now().Add(time.Hour * 24 * 2),
				diff: DiffWeeks,
			},
			want: "0 weeks",
		},
		{
			in: Input{
				date: time.Now().Add(time.Hour * 24 * 8),
				diff: DiffWeeks,
			},
			want: "1 week",
		},
		{
			in: Input{
				date: time.Now().Add(time.Hour * 24 * 15),
				diff: DiffWeeks,
			},
			want: "2 weeks",
		},
	}

	var actual string

	for _, item := range table {
		actual = parse(item.in.date, item.in.diff)
		if actual != item.want {
			t.Fatalf("want: %s, actual: %s", item.want, actual)
		}
	}
}
