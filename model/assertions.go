package model

import (
	"fmt"
	"github.com/sirupsen/logrus"
)

// assertionKind describes all possible assertions
type assertionKind string

const (
	// StatusAssertion asserts response status
	StatusAssertion assertionKind = "status"
	// BodyAssertion asserts response body
	BodyAssertion assertionKind = "body"
)

// Assertions is a logical aggregation of assertion
type Assertions struct {
	Name       string       `json:"name"`
	Assertions []*Assertion `json:"assertions"`
}

// AddOf adds a new assertion
func (as *Assertions) AddOf(kind assertionKind, expected, actual interface{}, msg string) {
	logrus.WithFields(logrus.Fields{"kind": kind, "expected": expected, "actual": actual}).Debug("assertion added")
	a := &Assertion{
		Kind:     kind,
		Expected: expected,
		Actual:   actual,
		Msg:      msg,
	}

	as.Assertions = append(as.Assertions, a)
}

// NewAssertions creates an empty Assertions
func NewAssertions(name string) *Assertions {
	return &Assertions{
		Name: name,
	}
}

// Assertion is a result of predicate execution
type Assertion struct {
	Kind     assertionKind
	Expected interface{}
	Actual   interface{}
	Msg      string
}

// String formats an Assertion as a string
func (a *Assertion) String() string {
	return fmt.Sprintf("Expected [%v] but got [%v]", a.Expected, a.Actual)
}

func (a *Assertion) True() bool {
	return a.Expected == a.Actual
}
