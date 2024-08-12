package models

import "time"

type Data struct {
	TotalSeconds      float64 `json:"total_seconds"`
	Text              string  `json:"text"`
	Decimal           string  `json:"decimal"`
	Digital           string  `json:"digital"`
	DailyAverage      float64 `json:"daily_average"`
	IsUpToDate        bool    `json:"is_up_to_date"`
	PercentCalculated int     `json:"percent_calculated"`
	Range             Range   `json:"range"`
	Timeout           int     `json:"timeout"`
}

type Range struct {
	Start     time.Time `json:"start"`
	StartDate string    `json:"start_date"`
	StartText string    `json:"start_text"`
	End       time.Time `json:"end"`
	EndDate   string    `json:"end_date"`
	EndText   string    `json:"end_text"`
	Timezone  string    `json:"timezone"`
}

type Response struct {
	Data Data `json:"data"`
}
