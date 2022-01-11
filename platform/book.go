package platform

type Book struct {
	Id                string `json:"id"`
	Name              string `json:"name"`
	Author            string `json:"author"`
	Isbn              string `json:"isbn"`
	Publisher         string `json:"publisher"`
	Genre             string `json:"genre"`
	YearOfPublication int    `json:"year_of_publication"`
	Pages             int    `json:"pages"`
}
