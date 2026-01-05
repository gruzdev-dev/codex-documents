package domain

type Identifier struct {
	System string
	Value  string
}

type Codeable struct {
	Code    string
	System  string
	Display string
}
