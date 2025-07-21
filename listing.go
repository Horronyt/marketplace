package marketplace

type Listing struct {
	Id          int    `json:"id" db:"id"`
	User_id     int    `json:"user_id" db:"user_id"`
	Title       string `json:"title" binding:"required" db:"title"`
	Description string `json:"description" db:"description"`
	Img_path    string `json:"img_path" db:"img_path"`
	Price       int    `json:"price" db:"price"`
}

type ListingOutputFormat struct {
	Title       string `json:"title" binding:"required" db:"title"`
	Description string `json:"description" db:"description"`
	Img_path    string `json:"img_path" db:"img_path"`
	Price       int    `json:"price" db:"price"`
	Username    string `json:"username" db:"username"`
	Belonging   bool   `json:"belonging" db:"belonging"`
}

type ListingOutputFormatAnon struct {
	Title       string `json:"title" binding:"required" db:"title"`
	Description string `json:"description" db:"description"`
	Img_path    string `json:"img_path" db:"img_path"`
	Price       int    `json:"price" db:"price"`
	Username    string `json:"username" db:"username"`
}
