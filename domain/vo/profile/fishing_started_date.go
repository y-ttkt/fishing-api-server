package profile

import (
	"errors"
	"fmt"
	date "github.com/yusuke-takatsu/fishing-api-server/config/date"
	"time"
)

type FishingStartedDate time.Time

func NewFishingStartedDate(value string) (FishingStartedDate, error) {
	if value == "" {
		return FishingStartedDate(time.Time{}), errors.New("value is empty")
	}

	t, err := time.Parse(date.DateFormat, value)
	if err != nil {
		return FishingStartedDate(time.Time{}), fmt.Errorf("invalid date format: %w", err)
	}

	if t.After(time.Now()) {
		return FishingStartedDate(time.Time{}), errors.New("釣りを始めた日は未来に設定できません。")
	}

	return FishingStartedDate(t), nil
}

func (d FishingStartedDate) Value() time.Time {
	return time.Time(d)
}
