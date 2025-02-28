# GoAvatar

GoAvatar is a Go library for generating customizable avatars. It supports creating avatars for text, stickers, and persons with various customization options.

| Letter                                                                               | Sticker                                                                           | Male                                                                           | Female                                                                             |
| ------------------------------------------------------------------------------------ | --------------------------------------------------------------------------------- | ------------------------------------------------------------------------------ | ---------------------------------------------------------------------------------- |
| [Letter](https://github.com/mekramy/goavatar/blob/master/demo/john_doe.svg?raw=true) | [Sticker](https://github.com/mekramy/goavatar/blob/master/demo/star.svg?raw=true) | [Male](https://github.com/mekramy/goavatar/blob/master/demo/male.svg?raw=true) | [Female](https://github.com/mekramy/goavatar/blob/master/demo/female.svg?raw=true) |

## Installation

To install GoAvatar, use `go get`:

```sh
go get github.com/mekramy/goavatar
```

## Usage

### Creating a Text Avatar

```go
package main

import (
    "github.com/mekramy/goavatar"
    "log"
)

func main() {
    factory := goavatar.
        New().
        DefaultPalettes().
        CircleShape().
        Build()

    johnDoe := factory.NewText("John Doe")

    johnDoe.
        RandomizeShape().
        RandomizePalette()

    err := johnDoe.Save("john_doe.svg")
    if err != nil {
        log.Fatal(err)
    }
}
```

### Creating a Sticker Avatar

```go
package main

import (
    "github.com/mekramy/goavatar"
    "log"
)

func main() {
    factory := goavatar.
        New().
        DefaultPalettes().
        CircleShape().
        Build()

    star := factory.NewSticker("")

    star.
        RandomizeShape().
        RandomizePalette()

    err := star.Save("star.svg")
    if err != nil {
        log.Fatal(err)
    }
}
```

### Creating a Person Avatar

```go
package main

import (
    "github.com/mekramy/goavatar"
    "log"
)

func main() {
    factory := goavatar.
        New().
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
    err := male.Save("male.svg")
    if err != nil {
        log.Fatal(err)
    }

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
    err := female.Save("female.svg")
    if err != nil {
        log.Fatal(err)
    }
}
```
