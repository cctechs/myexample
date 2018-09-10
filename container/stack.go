package container


type Stack struct {
	s []interface{}
}

func (s * Stack) Push(v int){
	s.s = append(s.s, v)
}

func (s * Stack) Pop() interface{}{
	n := len(s.s)
	if 0 == n{
		return nil
	}
	v := s.s[n - 1]
	s.s = s.s[n - 1:]
	return v
}

func (s* Stack)Empty() bool  {
	if 0 == len(s.s){
		return true
	}
	return false
}