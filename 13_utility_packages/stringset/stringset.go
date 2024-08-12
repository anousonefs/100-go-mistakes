package stringset

func New(...string) map[string]struct{} {
	res := make(map[string]struct{})
	return res
}

func Sort(map[string]struct{}) []string {
	res := make([]string, 0)
	return res
}

// use custom type

type Set map[string]struct{}

func New2(...string) Set {
	res := make(Set)
	return res
}

func (s Set) Sort2() []string {
	res := make([]string, 0)
	return res
}
