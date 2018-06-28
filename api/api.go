package api

type Article struct {
	ID    string   `json:"id"`
	Body  string   `json:"body"`
	Date  string   `json:"date"`
	Tags  []string `json:"tags"`
	Title string   `json:"title"`
}

type Tag struct {
	Articles    []string `json:"articles"`
	Count       int      `json:"count"`
	RelatedTags []string `json:"related_tags"`
	Tag         string   `json:"tag"`
}
