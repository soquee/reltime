package reltime_test

import (
	"strconv"
	"testing"
	"time"

	"code.soquee.net/reltime"
)

var testCases = [...]struct {
	in  time.Duration
	out string
}{
	0: {
		in:  15 * time.Second,
		out: "just now",
	},
	1:  {in: (-9 * 24 * time.Hour) - (12 * time.Hour), out: "10 days ago"},
	2:  {in: 15 * time.Second, out: "just now"},
	3:  {in: 45 * time.Second, out: "less than a minute"},
	4:  {in: -time.Minute, out: "about a minute ago"},
	5:  {in: -3*time.Minute + (28 * time.Second), out: "3 minutes ago"},
	6:  {in: 29 * time.Minute, out: "29 minutes"},
	7:  {in: 39 * time.Minute, out: "less than an hour"},
	8:  {in: -time.Hour, out: "about an hour ago"},
	9:  {in: -2 * time.Hour, out: "2 hours ago"},
	10: {in: 25 * time.Hour, out: "about a day"},
	11: {in: 3 * 24 * time.Hour, out: "3 days"},
	12: {in: 25 * 24 * time.Hour, out: "25 days"},
	13: {in: 30 * 24 * time.Hour, out: "about a month"},
	14: {in: -7 * 30 * 24 * time.Hour, out: "7 months ago"},
	15: {in: 10 * 30 * 24 * time.Hour, out: "10 months"},
	16: {in: 12 * 30 * 24 * time.Hour, out: "about a year"},
	17: {in: -19 * 30 * 24 * time.Hour, out: "2 years ago"},
	18: {in: -20 * 12 * 30 * 24 * time.Hour, out: "20 years ago"},
	19: {in: 100 * 12 * 30 * 24 * time.Hour, out: "100 years"},
	20: {out: "just now"},
	21: {in: 10 * 31 * 24 * time.Hour, out: "10 months"},
}

func TestAgo(t *testing.T) {
	for i, tc := range testCases {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			out := reltime.Ago(tc.in)
			if out != tc.out {
				t.Errorf("want=%q, got=%q", tc.out, out)
			}
		})
	}
}
