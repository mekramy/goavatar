package goavatar

func (b *builder) SetBodyShape(shape string) FactoryBuilder {
	b.f.setBodyShape(shape)
	return b
}
