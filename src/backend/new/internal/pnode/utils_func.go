package pnode

func basicNodeEqual(v1, v2 Node) bool {
	if v1 == nil && v2 == nil {
		return true
	} else if v1 == nil || v2 == nil {
		return false
	}

	return v1.Tag() == v2.Tag()
}
