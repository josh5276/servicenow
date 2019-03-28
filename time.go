package servicenow

import (
	"bytes"
	"time"
)

// SNTime is a servicenow package specific time type to describe
// time formats received from SNOW
type SNTime struct {
	time.Time
}

const snFormat = "2006-01-02 15:04:05"

// UnmarshalJSON method will take a SNTime pointer receiver and format the
// Time into a Time object
func (s *SNTime) UnmarshalJSON(bs []byte) error {
	st := string(bytes.Trim(bs, "\""))

	if len(st) < len(snFormat) {
		return nil
	}

	var err error
	s.Time, err = time.Parse(snFormat, st)
	return err
}

// MarshalJSON method will take a SNTime pointer receiver and format the
// Time into a []byte
func (s *SNTime) MarshalJSON() ([]byte, error) {
	return []byte(s.Time.Format(snFormat)), nil
}
