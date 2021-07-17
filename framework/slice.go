package framework

type Slice struct {

}

// ContainsInt : Returns true if the slice contains the integer
//               in find. Otherwise returns false.
func (s *Slice) ContainsInt(slice []int, find int) bool {
	for _, a := range slice {
		if a == find {
			return true
		}
	}
	return false
}

// RemoveIntAt : Removes the integer at the specified index
//               from the slice
func (s *Slice) RemoveIntAt(slice []int, index int) []int {
	return append(slice[:index], slice[index+1:]...)
}


// RemoveIntAtFast : Removes the integer at the specified index
//                   from the slice. IMPORTANT : This function
//                   reorders the items in the slice!
func (s *Slice) RemoveIntAtFast(slice []int, index int) []int {
	slice[index] = slice[len(slice)-1]
	return slice[:len(slice)-1]
}

// ContainsString : Returns true if the slice contains the string
//                  in find. Otherwise returns false.
func (s *Slice) ContainsString(slice []string, find string) bool {
	for _, a := range slice {
		if a == find {
			return true
		}
	}
	return false
}

// RemoveStringAt : Removes the string at the specified index
//                  from the slice
func (s *Slice) RemoveStringAt(slice []string, index int) []string {
	return append(slice[:index], slice[index+1:]...)
}

// RemoveStringAtFast : Removes the string at the specified index
//                      from the slice. IMPORTANT : This function
//                      reorders the items in the slice!
func (s *Slice) RemoveStringAtFast(slice []string, index int) []string {
	slice[index] = slice[len(slice)-1]
	return slice[:len(slice)-1]
}
