package main

type StringIntMap struct {
	StrInMap map[string]int
}

func main() {

}

func newStringIntMap() *StringIntMap {
	return &StringIntMap{
		StrInMap: make(map[string]int),
	}
}

func (s *StringIntMap) Add(key string, value int) {
	s.StrInMap[key] = value
}

func (s *StringIntMap) Remove(key string) {
	delete(s.StrInMap, key)
}

func (s *StringIntMap) Copy() map[string]int {
	newMap := make(map[string]int)
	for key, value := range s.StrInMap {
		newMap[key] = value
	}
	return newMap
}

func (s *StringIntMap) Exists(key string) bool {
	_, exists := s.StrInMap[key]
	return exists
}

func (s *StringIntMap) Get(key string) (int, bool) {
	value, exists := s.StrInMap[key]
	return value, exists
}
