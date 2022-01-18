package a

import "b"

func f() error {
	return b.B("err")
}
