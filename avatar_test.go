package goavatar_test

import (
	"testing"

	"github.com/mekramy/goavatar"
)

func TestTextAvatar(t *testing.T) {
	factory := goavatar.New().
		DefaultPalettes().
		CircleShape().
		Build()

	johnDoe := factory.NewText("John Doe")

	johnDoe.
		RandomizeShape().
		RandomizePalette()

	err := johnDoe.Save("demo/john_doe.svg")
	if err != nil {
		t.Fatal(err)
	}
}

func TestStickerAvatar(t *testing.T) {
	factory := goavatar.New().
		DefaultPalettes().
		CircleShape().
		Build()

	star := factory.NewSticker("")

	star.
		RandomizeShape().
		RandomizePalette()

	err := star.Save("demo/star.svg")
	if err != nil {
		t.Fatal(err)
	}
}

func TestMaleAvatar(t *testing.T) {
	factory := goavatar.New().
		DefaultAccessories().
		DefaultBeards().
		SuitDress().
		PrescriptionGlasses().
		DarkHairColor().
		DefaultHairs().
		DefaultPalettes().
		CircleShape().
		WhiteSkin().
		Build()

	male := factory.NewMale()

	male.
		RandomizeShape().
		RandomizePalette().
		RandomizeSkinColor().
		RandomizeHairColor().
		RandomizeHair().
		RandomizeBeard().
		RandomizeDress().
		RandomizeEye().
		RandomizeEyebrow().
		RandomizeMouth().
		RandomizeGlasses().
		RandomizeAccessory()

	err := male.Save("demo/male.svg")
	if err != nil {
		t.Fatal(err)
	}
}

func TestFemaleAvatar(t *testing.T) {
	factory := goavatar.New().
		DefaultAccessories().
		DefaultBeards().
		DefaultDresses().
		RoundPrescriptionGlasses().
		// DefaultHairColors().
		DefaultHairs().
		DefaultPalettes().
		CircleShape().
		WhiteSkin().
		Build()

	female := factory.NewFemale()

	female.
		RandomizeShape().
		RandomizePalette().
		RandomizeSkinColor().
		RandomizeHairColor().
		RandomizeHair().
		RandomizeBeard().
		RandomizeDress().
		RandomizeEye().
		RandomizeEyebrow().
		RandomizeMouth().
		RandomizeGlasses().
		RandomizeAccessory()

	err := female.Save("demo/female.svg")
	if err != nil {
		t.Fatal(err)
	}
}
