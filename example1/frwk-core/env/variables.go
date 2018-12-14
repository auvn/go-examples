package env

import "os"

func AnyOfVar(varNames ...string) (v string) {
	for _, n := range varNames {
		v = os.Getenv(n)
		if v == "" {
			continue
		}
	}
	return v
}
