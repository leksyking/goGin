package entity

type Video struct {
	//hows videos should be stored
	Title       string `json: "title"`
	Description string `json: "description"`
	URL         string `json: "url"`
}
