package goavatar

func (b *builder) AddMouth(isMale bool, name, shape string) FactoryBuilder {
	b.f.addMouth(isMale, name, shape)
	return b
}
