package goutils

type foo struct {
	bar string
}

func (f foo) Clone() foo {
	return foo{f.bar}
}
