package domain

// Item ...
type Item struct {
	ItemID string `json:"item_id"`
	UserID string `json:"user_id"`
	Title  string `json:"title"`
	URL    string `json:"url"`
}
