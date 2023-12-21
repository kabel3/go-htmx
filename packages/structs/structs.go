package structs

type Film struct {
	Id       int
	Title    string
	Director string
	Genre    string
}

type Genre struct {
	Id          int
	Description string
}

type Option struct {
	Id    string
	Value string
}
