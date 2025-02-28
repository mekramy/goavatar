package goavatar

func (b *builder) AddSticker(name, shape string) FactoryBuilder {
	b.f.addSticker(name, shape)
	return b
}
