package marketplace

type Listing struct {
	Id          int    `json:"id"`
	User_id     int    `json:"user_id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Img_path    string `json:"img_path"`
	Price       int    `json:"price"`
}
