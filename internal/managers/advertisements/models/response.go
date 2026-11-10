package models

type AdvertisementsList struct {
	Items      []Advertisement `json:"items"`
	Total      int             `json:"total"`
	HasNext    bool            `json:"has_next"`
	NextCursor string          `json:"next_cursor"`
}

type Advertisement struct {
	Id          string  `json:"id"`
	Title       string  `json:"title"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
	Quantity    int     `json:"quantity"`
	UserId      string  `json:"user_id"`
	CreatedAt   string  `json:"created_at"`
	UpdatedAt   string  `json:"updated_at"`
}
