package fbintegration

import (
	"regexp"
)

type (
	// Error pending comment
	Error struct {
		Error string
	}
)

// NewError comment pending
func NewError(err error) *Error {
	return &Error{err.Error()}
}

// IsRateLimited comment pending
func (e *Error) IsRateLimited() bool {
	rateLimit := regexp.MustCompile(`\(#17\)$`)
	return rateLimit.MatchString(e.Error)
}
