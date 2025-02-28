package goavatar

func (b *builder) AddEye(isMale bool, name, shape string) FactoryBuilder {
	b.f.addEye(isMale, name, shape)
	return b
}
