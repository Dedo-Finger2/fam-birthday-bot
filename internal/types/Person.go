package types

type Person struct {
	Name       string  `json:"name" yaml:"name"`
	Age        float64 `json:"age" yaml:"age"`
	Complement string  `json:"complement" yaml:"complement"`
}
