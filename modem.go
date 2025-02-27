package marshal

type Marshaler[V any] interface {
	Marshal(v V, buf []byte) (err error)
}
type UnMarshaler[V any] interface {
	Unmarshal(buf []byte, v *V) (err error)
}
type Sizeofer[V any] interface {
	Sizeof(v V) int
}
type Modem[V any] interface {
	Marshaler[V]
	UnMarshaler[V]
	Sizeofer[V]
}

//type ModemGeneratorFn[V any] func() Modem[V]
