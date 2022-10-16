package framework

type Slice struct {
}

// ContainsInt : Returns true if the slice contains the integer
//
//	in find, otherwise returns false.
func (s *Slice) ContainsInt(slice []int, find int) bool {
	for _, a := range slice {
		if a == find {
			return true
		}
	}
	return false
}

// IndexOfInt : Returns the index if the slice contains the integer
//
//	in find, otherwise returns -1.
func (s *Slice) IndexOfInt(slice []int, find int) int {
	for i, a := range slice {
		if a == find {
			return i
		}
	}
	return -1
}

// RemoveIntAt : Removes the integer at the specified index
//
//	from the slice
func (s *Slice) RemoveIntAt(slice []int, index int) []int {
	return append(slice[:index], slice[index+1:]...)
}

// RemoveIntAtFast : Removes the integer at the specified index
//
//	from the slice. IMPORTANT : This function
//	reorders the items in the slice!
func (s *Slice) RemoveIntAtFast(slice []int, index int) []int {
	slice[index] = slice[len(slice)-1]
	return slice[:len(slice)-1]
}

// ContainsString : Returns true if the slice contains the string
//
//	in find, otherwise returns false.
func (s *Slice) ContainsString(slice []string, find string) bool {
	for _, a := range slice {
		if a == find {
			return true
		}
	}
	return false
}

// IndexOfString : Returns index if the slice contains the string
//
//	in find, otherwise returns -1.
func (s *Slice) IndexOfString(slice []string, find string) int {
	for i, a := range slice {
		if a == find {
			return i
		}
	}
	return -1
}

// RemoveStringAt : Removes the string at the specified index
//
//	from the slice
func (s *Slice) RemoveStringAt(slice []string, index int) []string {
	return append(slice[:index], slice[index+1:]...)
}

// RemoveStringAtFast : Removes the string at the specified index
//
//	from the slice. IMPORTANT : This function
//	reorders the items in the slice!
func (s *Slice) RemoveStringAtFast(slice []string, index int) []string {
	slice[index] = slice[len(slice)-1]
	return slice[:len(slice)-1]
}
