package goavatar

func (b *builder) AddHaircolor(name string, color SVGHairColor) FactoryBuilder {
	b.f.addHairColor(name, color)
	return b
}

func (b *builder) DefaultHairColors() FactoryBuilder {
	b.BrownHairColor()
	b.LightHairColor()
	b.DarkHairColor()
	return b
}

func (b *builder) BrownHairColor() FactoryBuilder {
	b.f.addHairColor(
		string(HairBrown),
		NewHairColor(
			NewHex("#3a292f"),
			NewHex("#191215"),
			NewHex("#664954"),
		),
	)
	return b
}

func (b *builder) LightHairColor() FactoryBuilder {
	b.f.addHairColor(
		string(HairBrown),
		NewHairColor(
			NewHex("#744819"),
			NewHex("#563517"),
			NewHex("#825329"),
		),
	)
	return b
}

func (b *builder) DarkHairColor() FactoryBuilder {
	b.f.addHairColor(
		string(HairBrown),
		NewHairColor(
			NewHex("#432818"),
			NewHex("#140b06"),
			NewHex("#884139"),
		),
	)
	return b
}
