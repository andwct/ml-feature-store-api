package models

type Feature struct {
	HistDate string  `json:"entity_id" binding:"required"`
	Name     string  `json:"name" binding:"required"`
	Value    float64 `json:"value"`
}
