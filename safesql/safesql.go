package safesql

import "github.com/empijei/def-prog-exercises/safesql/internal/raw"

type compileTimeConstant string

type TrustedSQL struct {
	s string
}

func init() {
	raw.TrustedSQLCtor =
		func(unsafe string) TrustedSQL {
			return TrustedSQL{unsafe}
		}
}

func New(text compileTimeConstant) TrustedSQL {
	return TrustedSQL{string(text)}
}
