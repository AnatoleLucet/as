package as

type V struct {
	value any
}

func Value(v any) V {
	return V{v}
}

func (g V) Value() any {
	return g.value
}

func (g V) String() (string, error) {
	return String(g.value)
}

func (g V) Rune() (rune, error) {
	return Rune(g.value)
}

func (g V) Bool() (bool, error) {
	return Bool(g.value)
}

func (g V) Int() (int, error) {
	return Int(g.value)
}
func (g V) Int8() (int8, error) {
	return Int8(g.value)
}
func (g V) Int16() (int16, error) {
	return Int16(g.value)
}
func (g V) Int32() (int32, error) {
	return Int32(g.value)
}
func (g V) Int64() (int64, error) {
	return Int64(g.value)
}

func (g V) Uint() (uint, error) {
	return Uint(g.value)
}
func (g V) Uint8() (uint8, error) {
	return Uint8(g.value)
}
func (g V) Uint16() (uint16, error) {
	return Uint16(g.value)
}
func (g V) Uint32() (uint32, error) {
	return Uint32(g.value)
}
func (g V) Uint64() (uint64, error) {
	return Uint64(g.value)
}

func (g V) Float32() (float32, error) {
	return Float32(g.value)
}
func (g V) Float64() (float64, error) {
	return Float64(g.value)
}
