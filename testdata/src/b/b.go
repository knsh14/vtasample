package b

type e struct {
	s string
}

func (v *e) Error() string {
	return v.s
}

func B(s string) error {
	return &e{s: s}
}
