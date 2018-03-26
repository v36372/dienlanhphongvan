package view

type Url struct {
	Value string `json:"url"`
}

type Urls struct {
	Data []Url `json:"data"`
}
