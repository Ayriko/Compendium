package globals

var Secret = []byte("secret")

const Userkey = "user"

type User struct {
	Id       string `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
}

type Campaign struct {
	Id          int    `json:"id"`
	Title       string `json:"name"`
	Description string `json:"description"`
	Image       string `json:"image"`
}

type World struct {
	Id          int    `json:"id"`
	Title       string `json:"name"`
	Content     string `json:"content"`
	Image       string `json:"image"`
	Campaign_id int    `json:"campaign_id"`
}

type Characters struct {
	Id          int    `json:"id"`
	Name        string `json:"name"`
	Content     string `json:"content"`
	Image       string `json:"image"`
	Campaign_id int    `json:"campaign_id"`
}

type Lore struct {
	Id          int    `json:"id"`
	Title       string `json:"name"`
	Content     string `json:"content"`
	Image       string `json:"image"`
	Campaign_id int    `json:"campaign_id"`
}
