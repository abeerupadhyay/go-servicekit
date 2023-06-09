package str

var exists = struct{}{}

type Set map[string]struct{}

func (s Set) Add(element string) {
	s[element] = exists
}

func (s Set) AddN(elements ...string) {
	for _, el := range elements {
		s[el] = exists
	}
}

func (s Set) Remove(element string) {
	if s.Exists(element) {
		delete(s, element)
	}
}

func (s Set) Exists(element string) bool {
	_, ok := s[element]
	return ok
}

func (s Set) List() []string {
	r := make([]string, 0, len(s))
	for element := range s {
		r = append(r, element)
	}
	return r
}

func (s Set) Clear() {
	for element := range s {
		delete(s, element)
	}
}
