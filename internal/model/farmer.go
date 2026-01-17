package model

type Farmer struct {
	ID       string  `json:"id"`
	Name     string  `json:"name"`
	Location string  `json:"location"`
	CropType string  `json:"crop_type"`
	LandSize float64 `json:"land_size"`
}
