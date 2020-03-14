package category

type PostField struct {
	Name string `json:"name"`
}

type PutField struct {
	Id string `json:"id"`
	Name string `json:"name"`
}

type GetField struct {
	Id string `json:"id"`
	Name string `json:"name"`
	Num string `json:"num"`
}

type DeleteField struct {
	Id string `json:"id"`
}

