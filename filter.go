package functional

// DO NOT EDIT.
// Code generated by cmd/generate.go
// source: filter.tmpl

// FilterUint ...
func FilterUint(slice []uint, fn func(uint) bool) (ret []uint) {
	for _, v := range slice {
		if fn(v) {
			ret = append(ret, v)
		}
	}
	return
}

// FilterUint8 ...
func FilterUint8(slice []uint8, fn func(uint8) bool) (ret []uint8) {
	for _, v := range slice {
		if fn(v) {
			ret = append(ret, v)
		}
	}
	return
}

// FilterUint16 ...
func FilterUint16(slice []uint16, fn func(uint16) bool) (ret []uint16) {
	for _, v := range slice {
		if fn(v) {
			ret = append(ret, v)
		}
	}
	return
}

// FilterUint32 ...
func FilterUint32(slice []uint32, fn func(uint32) bool) (ret []uint32) {
	for _, v := range slice {
		if fn(v) {
			ret = append(ret, v)
		}
	}
	return
}

// FilterUint64 ...
func FilterUint64(slice []uint64, fn func(uint64) bool) (ret []uint64) {
	for _, v := range slice {
		if fn(v) {
			ret = append(ret, v)
		}
	}
	return
}

// FilterInt ...
func FilterInt(slice []int, fn func(int) bool) (ret []int) {
	for _, v := range slice {
		if fn(v) {
			ret = append(ret, v)
		}
	}
	return
}

// FilterInt8 ...
func FilterInt8(slice []int8, fn func(int8) bool) (ret []int8) {
	for _, v := range slice {
		if fn(v) {
			ret = append(ret, v)
		}
	}
	return
}

// FilterInt16 ...
func FilterInt16(slice []int16, fn func(int16) bool) (ret []int16) {
	for _, v := range slice {
		if fn(v) {
			ret = append(ret, v)
		}
	}
	return
}

// FilterInt32 ...
func FilterInt32(slice []int32, fn func(int32) bool) (ret []int32) {
	for _, v := range slice {
		if fn(v) {
			ret = append(ret, v)
		}
	}
	return
}

// FilterInt64 ...
func FilterInt64(slice []int64, fn func(int64) bool) (ret []int64) {
	for _, v := range slice {
		if fn(v) {
			ret = append(ret, v)
		}
	}
	return
}

// FilterFloat32 ...
func FilterFloat32(slice []float32, fn func(float32) bool) (ret []float32) {
	for _, v := range slice {
		if fn(v) {
			ret = append(ret, v)
		}
	}
	return
}

// FilterFloat64 ...
func FilterFloat64(slice []float64, fn func(float64) bool) (ret []float64) {
	for _, v := range slice {
		if fn(v) {
			ret = append(ret, v)
		}
	}
	return
}

// FilterBool ...
func FilterBool(slice []bool, fn func(bool) bool) (ret []bool) {
	for _, v := range slice {
		if fn(v) {
			ret = append(ret, v)
		}
	}
	return
}

// FilterString ...
func FilterString(slice []string, fn func(string) bool) (ret []string) {
	for _, v := range slice {
		if fn(v) {
			ret = append(ret, v)
		}
	}
	return
}