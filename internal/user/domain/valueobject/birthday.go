package valueobject

import (
	"database/sql/driver"
	"errors"
	"github.com/take0fit/ddd-clean-temp/pkg/util"
	"time"
)

type Birthday struct {
	Time  time.Time
	Valid bool
}

func NewBirthday(t time.Time) Birthday {
	return Birthday{Time: t}
}

func (b Birthday) Age() *int {
	if !b.Valid {
		return nil
	}

	now := time.Now()
	years := now.Year() - b.Time.Year()
	if util.IsBefore(b.Time, now) {
		years--
	}
	return &years
}

func (b Birthday) Value() (driver.Value, error) {
	if !b.Valid {
		return nil, nil // NULL を返す
	}
	return b.Time, nil
}

func (b *Birthday) Scan(value interface{}) error {
	if value == nil {
		b.Valid = false
		return nil
	}

	b.Valid = true
	switch v := value.(type) {
	case time.Time:
		b.Time = v
	default:
		return errors.New("birthday must be a time.Time or nil")
	}
	return nil
}
