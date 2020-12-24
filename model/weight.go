package model

import (
	"encoding/json"
	"fmt"
	"time"
)

type Weight struct {
	BMI    float64
	Weight float64
	LogId  int
	Time   time.Time
	Source string
}

func (w *Weight) UnmarshalJSON(b []byte) error {
	tmp := struct {
		BMI    *float64 `json:"bmi"`
		Weight *float64 `json:"weight"`
		LogId  *int     `json:"logId"`
		Time   string   `json:"time"`
		Date   string   `json:"date"`
		Source *string  `json:"source"`
	}{
		BMI:    &w.BMI,
		Weight: &w.Weight,
		LogId:  &w.LogId,
		Source: &w.Source,
	}
	if err := json.Unmarshal(b, &tmp); err != nil {
		return fmt.Errorf("unmarshalling Weight: %w", err)
	}

	t, err := time.Parse("2006-01-0215:04:05", tmp.Date+tmp.Time)
	if err != nil {
		return fmt.Errorf("parsing date time: %w", err)
	}
	w.Time = t
	return nil
}
