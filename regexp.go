package gogexp

import (
	"regexp"
)

var (
	IsNumbersOnly    = regexp.MustCompile(`\A[0-9]+\z`)
	IsNumbersWithDot = regexp.MustCompile(`\A[0-9]+(\.[0-9]+)?\z`)
	IsDateGeneral    = regexp.MustCompile(`\A[0123]?\d/[0123]?\d(/\d{2,4})\z`)
	HasSpaces        = regexp.MustCompile(`\s+`)
)
