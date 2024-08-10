package types

type Birthday struct {
	Date   string   `json:"date"`
	People []Person `json:"people"`
}
