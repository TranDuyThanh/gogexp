package gogexp

import (
	"regexp"
)

var (
	IsNumbersOnly = regexp.MustCompile(`\A[0-9]+\z`)
	// Ex: 12.234 or 12
	IsFloatWithDot = regexp.MustCompile(`\A[0-9]+(\.[0-9]+)?\z`)
	// Ex: 12,234 or 12
	IsFloatWithComma = regexp.MustCompile(`\A[0-9]+(,[0-9]+)?\z`)
	// Ex: 12,234,543.99
	IsBigFloatNumber = regexp.MustCompile(`\A\d+((,|\.)\d{3})*((,|\.)\d+)?\z`)
	//Ex: 31/07/2017   1/5/2017   19/4/99
	IsDateGeneral = regexp.MustCompile(`\A[0123]?\d/[01]?\d/\d{2,4}\z`)
	//Ex: 07/2017   5/2017   4/99
	IsDateWithoutDay = regexp.MustCompile(`\A[0123]?\d(/\d{2,4})\z`)
	//Ex: 31/07   1/5   19/4
	IsDateWithoutYear = regexp.MustCompile(`\A[0123]?\d/[01]?\d\z`)
	HasSpaces         = regexp.MustCompile(`\s+`)
)
