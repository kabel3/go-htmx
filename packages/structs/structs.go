package structs

type Film struct {
	Id       int
	Title    string
	Director string
	Genre    string
	GenreId  int
}

type Genre struct {
	Id          int
	Description string
}

type Option struct {
	Id       string
	Value    string
	Text     string
	Selected bool
}
