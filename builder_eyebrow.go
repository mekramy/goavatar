package goavatar

func (b *builder) AddEyebrow(isMale bool, name, shape string) FactoryBuilder {
	b.f.addEyebrow(isMale, name, shape)
	return b
}
