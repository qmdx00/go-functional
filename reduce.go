package functional

// DO NOT EDIT.
// Code generated by cmd/generate.go
// source: reduce.tmpl

// ReduceUint ...
func ReduceUint(slice []uint, fn func(x uint, y uint) uint) (ret uint) {
	if len(slice) < 2 {
		panic("reduce: slice too short")
	}
	ret = slice[0]
	for i := 1; i < len(slice); i++ {
		ret = fn(ret, slice[i])
	}
	return
}

// ReduceUint8 ...
func ReduceUint8(slice []uint8, fn func(x uint8, y uint8) uint8) (ret uint8) {
	if len(slice) < 2 {
		panic("reduce: slice too short")
	}
	ret = slice[0]
	for i := 1; i < len(slice); i++ {
		ret = fn(ret, slice[i])
	}
	return
}

// ReduceUint16 ...
func ReduceUint16(slice []uint16, fn func(x uint16, y uint16) uint16) (ret uint16) {
	if len(slice) < 2 {
		panic("reduce: slice too short")
	}
	ret = slice[0]
	for i := 1; i < len(slice); i++ {
		ret = fn(ret, slice[i])
	}
	return
}

// ReduceUint32 ...
func ReduceUint32(slice []uint32, fn func(x uint32, y uint32) uint32) (ret uint32) {
	if len(slice) < 2 {
		panic("reduce: slice too short")
	}
	ret = slice[0]
	for i := 1; i < len(slice); i++ {
		ret = fn(ret, slice[i])
	}
	return
}

// ReduceUint64 ...
func ReduceUint64(slice []uint64, fn func(x uint64, y uint64) uint64) (ret uint64) {
	if len(slice) < 2 {
		panic("reduce: slice too short")
	}
	ret = slice[0]
	for i := 1; i < len(slice); i++ {
		ret = fn(ret, slice[i])
	}
	return
}

// ReduceInt ...
func ReduceInt(slice []int, fn func(x int, y int) int) (ret int) {
	if len(slice) < 2 {
		panic("reduce: slice too short")
	}
	ret = slice[0]
	for i := 1; i < len(slice); i++ {
		ret = fn(ret, slice[i])
	}
	return
}

// ReduceInt8 ...
func ReduceInt8(slice []int8, fn func(x int8, y int8) int8) (ret int8) {
	if len(slice) < 2 {
		panic("reduce: slice too short")
	}
	ret = slice[0]
	for i := 1; i < len(slice); i++ {
		ret = fn(ret, slice[i])
	}
	return
}

// ReduceInt16 ...
func ReduceInt16(slice []int16, fn func(x int16, y int16) int16) (ret int16) {
	if len(slice) < 2 {
		panic("reduce: slice too short")
	}
	ret = slice[0]
	for i := 1; i < len(slice); i++ {
		ret = fn(ret, slice[i])
	}
	return
}

// ReduceInt32 ...
func ReduceInt32(slice []int32, fn func(x int32, y int32) int32) (ret int32) {
	if len(slice) < 2 {
		panic("reduce: slice too short")
	}
	ret = slice[0]
	for i := 1; i < len(slice); i++ {
		ret = fn(ret, slice[i])
	}
	return
}

// ReduceInt64 ...
func ReduceInt64(slice []int64, fn func(x int64, y int64) int64) (ret int64) {
	if len(slice) < 2 {
		panic("reduce: slice too short")
	}
	ret = slice[0]
	for i := 1; i < len(slice); i++ {
		ret = fn(ret, slice[i])
	}
	return
}

// ReduceFloat32 ...
func ReduceFloat32(slice []float32, fn func(x float32, y float32) float32) (ret float32) {
	if len(slice) < 2 {
		panic("reduce: slice too short")
	}
	ret = slice[0]
	for i := 1; i < len(slice); i++ {
		ret = fn(ret, slice[i])
	}
	return
}

// ReduceFloat64 ...
func ReduceFloat64(slice []float64, fn func(x float64, y float64) float64) (ret float64) {
	if len(slice) < 2 {
		panic("reduce: slice too short")
	}
	ret = slice[0]
	for i := 1; i < len(slice); i++ {
		ret = fn(ret, slice[i])
	}
	return
}

// ReduceBool ...
func ReduceBool(slice []bool, fn func(x bool, y bool) bool) (ret bool) {
	if len(slice) < 2 {
		panic("reduce: slice too short")
	}
	ret = slice[0]
	for i := 1; i < len(slice); i++ {
		ret = fn(ret, slice[i])
	}
	return
}

// ReduceString ...
func ReduceString(slice []string, fn func(x string, y string) string) (ret string) {
	if len(slice) < 2 {
		panic("reduce: slice too short")
	}
	ret = slice[0]
	for i := 1; i < len(slice); i++ {
		ret = fn(ret, slice[i])
	}
	return
}