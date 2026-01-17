package model

type Supplier struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Type        string `json:"type"`
	ContactInfo string `json:"contact_info"`
}
