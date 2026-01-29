package request

type Post struct {
	ProductID string `json:"product_id" validate:"required,uuid4"`
	Type      string `json:"type" validate:"required"`
	Value     int    `json:"value" validate:"required"`
}
