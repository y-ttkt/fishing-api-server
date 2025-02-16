package profile

import (
	"errors"
	"fmt"
	"github.com/yusuke-takatsu/fishing-api-server/config/date"
	"time"
)

type DateOfBirth time.Time

func NewDateOfBirth(value string) (DateOfBirth, error) {
	if value == "" {
		return DateOfBirth(time.Time{}), errors.New("value is empty")
	}

	t, err := time.Parse(date.DateFormat, value)
	if err != nil {
		return DateOfBirth(time.Time{}), fmt.Errorf("invalid date format: %w", err)
	}

	if t.After(time.Now()) {
		return DateOfBirth(time.Time{}), errors.New("誕生日は未来日に設定できません。")
	}

	return DateOfBirth(t), nil
}

func (d DateOfBirth) Value() time.Time {
	return time.Time(d)
}
