package types

type BirthDate struct {
	Date   string   `json:"date" yaml:"date"`
	People []Person `json:"people" yaml:"people"`
}
