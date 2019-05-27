// Package reltime implements a "time ago" algorithm.
package reltime // import "code.soquee.net/reltime"

import (
	"math"
	"strconv"
	"time"
)

// TimeAgo transforms the difference between a time and the current time into a
// human readable string.
//
// It is a convenience wrapper for Ago(time.Until(t)).
// For more information see the example on Ago.
func TimeAgo(t time.Time) string {
	return Ago(time.Until(t))
}

// Ago transforms durations into human readable strings.
func Ago(d time.Duration) string {
	// Take the absolute value and record the sign.
	sign := d >> 63
	d = (d ^ sign) - sign
	var ago string
	if sign < 0 {
		ago = " ago"
	}

	switch {
	case d < 30*time.Second:
		return "just now"
	case d < time.Minute:
		return "less than a minute" + ago
	case d < time.Minute+(30*time.Second):
		return "about a minute" + ago
	case d < 30*time.Minute:
		return strconv.FormatFloat(math.Round(d.Minutes()), 'f', -1, 64) + " minutes" + ago
	case d < time.Hour:
		return "less than an hour" + ago
	case d < time.Hour+30*time.Minute:
		return "about an hour" + ago
	case d < 24*time.Hour:
		return strconv.FormatFloat(math.Round(d.Hours()), 'f', -1, 64) + " hours" + ago
	case d < 32*time.Hour:
		return "about a day" + ago
	case d < 28*24*time.Hour:
		return strconv.FormatFloat(math.Round(d.Hours()/24), 'f', -1, 64) + " days" + ago
	case d < 45*24*time.Hour:
		return "about a month" + ago
	case d < 12*30*24*time.Hour:
		return strconv.FormatFloat(math.Round(d.Hours()/24/30), 'f', -1, 64) + " months" + ago
	case d < 18*30*24*time.Hour:
		return "about a year" + ago
	}
	return strconv.FormatFloat(math.Round(d.Hours()/24/30/12), 'f', -1, 64) + " years" + ago
}
