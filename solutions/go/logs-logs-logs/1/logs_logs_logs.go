package logs

import (
	"fmt"
	"strings"
	"unicode/utf8"
)

// Application identifies the application emitting the given log.
func Application(log string) string {
	recommendationRune := '❗'
	searchRune := '🔍'
	weatherRune := '☀'
	for _, runeValue := range log {
		switch {
		case runeValue == recommendationRune:
			return "recommendation"
		case runeValue == searchRune:
			return "search"
		case runeValue == weatherRune:
			return "weather"
		}
	}
	return "default"
}

// Replace replaces all occurrences of old with new, returning the modified log
// to the caller.
func Replace(log string, oldRune, newRune rune) string {
	var newLog strings.Builder
	newLog.Grow(len(log))
	for _, runeValue := range log {
		if runeValue == oldRune {
			fmt.Fprintf(&newLog, "%c", newRune)
		} else {
			fmt.Fprintf(&newLog, "%c", runeValue)
		}

	}
	return newLog.String()
}

// WithinLimit determines whether or not the number of characters in log is
// within the limit.
func WithinLimit(log string, limit int) bool {
	return utf8.RuneCountInString(log) <= limit
}
