package structs

type Film struct {
	Id       uint
	Title    string
	Director string
	Genre    string
	GenreId  uint
	Starred  bool
}

type Genre struct {
	Id          uint
	Description string
}

type Option struct {
	Id       string
	Value    string
	Text     string
	Selected bool
}
