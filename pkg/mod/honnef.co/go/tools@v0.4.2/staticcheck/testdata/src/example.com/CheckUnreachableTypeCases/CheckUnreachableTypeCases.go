package pkg

import "io"

type T struct{}

func (T) Read(b []byte) (int, error) { return 0, nil }
func (T) something() string          { return "non-exported method" }

type V error
type U error

func fn1() {
	var (
		v   interface{}
		err error
	)

	switch v.(type) {
	case io.Reader:
		println("io.Reader")
	case io.ReadCloser: //@ diag(`unreachable case clause: io.Reader will always match before io.ReadCloser`)
		println("io.ReadCloser")
	}

	switch v.(type) {
	case io.Reader:
		println("io.Reader")
	case T: //@ diag(`unreachable case clause: io.Reader will always match before example.com/CheckUnreachableTypeCases.T`)
		println("T")
	}

	switch v.(type) {
	case io.Reader:
		println("io.Reader")
	case io.ReadCloser: //@ diag(`unreachable case clause: io.Reader will always match before io.ReadCloser`)
		println("io.ReadCloser")
	case T: //@ diag(`unreachable case clause: io.Reader will always match before example.com/CheckUnreachableTypeCases.T`)
		println("T")
	}

	switch v.(type) {
	case io.Reader:
		println("io.Reader")
	case io.ReadCloser, T: //@ diag(`unreachable case clause: io.Reader will always match before io.ReadCloser`)
		println("io.ReadCloser or T")
	}

	switch v.(type) {
	case io.ReadCloser, io.Reader:
		println("io.ReadCloser or io.Reader")
	case T: //@ diag(`unreachable case clause: io.Reader will always match before example.com/CheckUnreachableTypeCases.T`)
		println("T")
	}

	switch v.(type) {
	default:
		println("something else")
	case io.Reader:
		println("io.Reader")
	case T: //@ diag(`unreachable case clause: io.Reader will always match before example.com/CheckUnreachableTypeCases.T`)
		println("T")
	}

	switch v.(type) {
	case interface{}:
		println("interface{}")
	case nil, T: //@ diag(`unreachable case clause: interface{} will always match before example.com/CheckUnreachableTypeCases.T`)
		println("nil or T")
	}

	switch err.(type) {
	case V:
		println("V")
	case U: //@ diag(`unreachable case clause: example.com/CheckUnreachableTypeCases.V will always match before example.com/CheckUnreachableTypeCases.U`)
		println("U")
	}

	switch err.(type) {
	case U:
		println("U")
	case V: //@ diag(`unreachable case clause: example.com/CheckUnreachableTypeCases.U will always match before example.com/CheckUnreachableTypeCases.V`)
		println("V")
	}
}

func fn3() {
	var (
		v   interface{}
		err error
	)

	switch v.(type) {
	case T:
		println("T")
	case io.Reader:
		println("io.Reader")
	}

	switch v.(type) {
	case io.ReadCloser:
		println("io.ReadCloser")
	case T:
		println("T")
	}

	switch v.(type) {
	case io.ReadCloser:
		println("io.ReadCloser")
	case io.Reader:
		println("io.Reader")
	}

	switch v.(type) {
	case T:
		println("T")
	}

	switch err.(type) {
	case V, U:
		println("V or U")
	case io.Reader:
		println("io.Reader")
	}

	switch v.(type) {
	default:
		println("something")
	}

	switch v.(type) {
	case interface{}:
		println("interface{}")
	case nil:
		println("nil")
	}
}
