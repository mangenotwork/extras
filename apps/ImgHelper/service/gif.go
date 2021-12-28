package service

import (
	"bytes"
	"image"
	"image/color"
	"image/color/palette"
	"image/draw"
	"image/gif"
	"mime/multipart"

	"github.com/mangenotwork/extras/common/logger"
)

func CompositeGif(file []*multipart.FileHeader) (outByte []byte, err error) {
	var (
		disposals []byte
		images []*image.Paletted
	 	delays []int
	)

	for i,handler := range file{
		fileName:=handler.Filename
		fileSize:=handler.Size
		logger.Info(i, fileName, fileSize)
		f, fErr := file[i].Open()
		if fErr != nil {
			err = fErr
			return
		}
		g, _, err := image.Decode(f)
		if err != nil {
			logger.Error(err)
		}
		_ = f.Close()
		cp :=  getPalette(g)
		//cp:=append(palette.WebSafe,color.Transparent)
		disposals = append(disposals, gif.DisposalBackground)//透明图片需要设置
		p := image.NewPaletted(image.Rect(0, 0, 640,996),cp)
		draw.Draw(p, p.Bounds(), g, image.ZP, draw.Src)
		images = append(images, p)
		delays = append(delays,4)
	}

	g := &gif.GIF{
		Image:     images,
		Delay:     delays,
		LoopCount: -1,
		Disposal: disposals,
	}

	out := new(bytes.Buffer)
	_=gif.EncodeAll(out, g)
	outByte = out.Bytes()
	return
}

func getPalette(m image.Image) color.Palette {
	p := color.Palette{color.RGBA{0x00,0x00,0x00,0x00}}
	p9 := color.Palette(palette.Plan9)
	b := m.Bounds()
	black := false
	for y := b.Min.Y; y < b.Max.Y; y++ {
		for x := b.Min.X; x < b.Max.X; x++ {
			c := m.At(x, y)
			cc := p9.Convert(c)
			if cc == p9[0] {
				black = true
			}
			if isInPalette(p, cc) == -1 {
				p = append(p, cc)
			}
		}
	}
	if len(p) < 256 && black == true {
		p[0] = color.RGBA{0x00,0x00,0x00,0x00} // transparent
		p = append(p, p9[0])
	}
	return p
}

func isInPalette(p color.Palette, c color.Color) int {
	ret := -1
	for i, v := range p {
		if v == c {
			return i
		}
	}
	return ret
}