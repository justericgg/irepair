package room

type Message struct {
	Id      string `json:"id"`
	Author  string `json:"author"`
	Avatar  string `json:"avatar"`
	Message string `json:"message"`
	Images  string `json:"images"`
	Time    int    `json:"time"`
}
