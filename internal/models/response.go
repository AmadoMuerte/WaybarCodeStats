package models

type Data struct {
	GrandTotal GrandTotal `json:"grand_total"`
}

type GrandTotal struct {
	TotalSeconds      float64  `json:"total_seconds"`
}

type Response struct {
	Data Data `json:"data"`
}
