// Code generated by "fmgo"; DO NOT EDIT.
package e2e

func NewSimpleStruct(id string, name string) Interface {
	return &SimpleStruct{id: id, name: name}
}

type SimpleStructOption func(*SimpleStruct)

func NewSimpleStructOptions(opts ...SimpleStructOption) Interface {
	i := &SimpleStruct{}
	for _, o := range opts {
		o(i)
	}
	return i
}
func WithId(id string) SimpleStructOption {
	return func(i *SimpleStruct) {
		i.id = id
	}
}
func WithName(name string) SimpleStructOption {
	return func(i *SimpleStruct) {
		i.name = name
	}
}
