package bench

// CPU1 is a Windows workstation, CPU2 is a Linux server.

import (
	"github.com/BurntSushi/xgbutil/xgraphics"
	"github.com/skelterjohn/go.wde/win"
	"image"
	"image/draw"
	"math"
	"math/rand"
	"testing"
)

func getrgba(r image.Rectangle) *image.RGBA {
	rgba := image.NewRGBA(r)
	// Random data
	for i := 0; i < len(rgba.Pix); i++ {
		rgba.Pix[i] = uint8(rand.Intn(256))
	}
	// Single color
	//draw.Draw(rgba, rgba.Bounds(), image.Black, image.ZP, draw.Src)
	return rgba
}

func compareImages(first, second image.Image) (ok bool) {
	for x := 0; x < first.Bounds().Max.X; x++ {
		for y := 0; y < first.Bounds().Max.Y; y++ {
			r1, g1, b1, _ := first.At(x, y).RGBA()
			r2, g2, b2, _ := second.At(x, y).RGBA()
			if r1 != r2 || g1 != g2 || b1 != b2 {
				return
			}
		}
	}
	ok = true
	return
}

// DEBUG: CPU1 - 0.03ns/op
//        CPU2 - 0.03ns/op
func BenchmarkDrawToRGBA(b *testing.B) {
	b.StopTimer()
	diameter := int(math.Sqrt(float64(b.N)))
	if diameter > 5000 {
		diameter = 5000
	}
	r := image.Rect(0, 0, diameter, diameter)
	source := getrgba(r)
	target := image.NewRGBA(r)
	b.StartTimer()

	draw.Draw(target, target.Bounds(), source, image.ZP, draw.Src)

	b.StopTimer()
	if !compareImages(source, target) {
		b.Fail()
	}
}

// DEBUG: CPU1 - 219ns/op
//        CPU2 - 194ns/op
func BenchmarkDrawToXgb(b *testing.B) {
	b.StopTimer()
	diameter := int(math.Sqrt(float64(b.N)))
	if diameter > 5000 {
		diameter = 5000
	}
	r := image.Rect(0, 0, diameter, diameter)
	source := getrgba(r)
	target := &xgraphics.Image{
		X:      nil,
		Pixmap: 0,
		Pix:    make([]uint8, 4*r.Dx()*r.Dy()),
		Stride: 4 * r.Dx(),
		Rect:   r,
		Subimg: false,
	}
	b.StartTimer()

	draw.Draw(target, target.Bounds(), source, image.ZP, draw.Src)

	b.StopTimer()
	if !compareImages(source, target) {
		b.Fail()
	}
	target.Destroy()
}

// DEBUG: CPU1 - 45ns/op
// DEBUG: CPU2 - 0.75ns/op
func BenchmarkConvertXgb(b *testing.B) {
	b.StopTimer()
	diameter := int(math.Sqrt(float64(b.N)))
	if diameter > 5000 {
		diameter = 5000
	}
	r := image.Rect(0, 0, diameter, diameter)
	source := getrgba(r)
	target := &xgraphics.Image{
		X:      nil,
		Pixmap: 0,
		Pix:    make([]uint8, 4*r.Dx()*r.Dy()),
		Stride: 4 * r.Dx(),
		Rect:   r,
		Subimg: false,
	}
	b.StartTimer()

	convertRGBAtoXgb(target, source)

	b.StopTimer()
	if !compareImages(source, target) {
		b.Fail()
	}
	target.Destroy()
}

// DEBUG: CPU1 - 276ns/op
func BenchmarkDrawToWin(b *testing.B) {
	b.StopTimer()
	diameter := int(math.Sqrt(float64(b.N)))
	if diameter > 5000 {
		diameter = 5000
	}
	r := image.Rect(0, 0, diameter, diameter)
	source := getrgba(r)
	target := win.NewDIB(r)
	b.StartTimer()

	draw.Draw(target, target.Bounds(), source, image.ZP, draw.Src)

	b.StopTimer()
	if !compareImages(source, target) {
		b.Fail()
	}
}

// DEBUG: CPU1 - 35ns/op without alpha
//               165ns/op if alpha is also taken into account (with floats)
func BenchmarkConvertWin(b *testing.B) {
	b.StopTimer()
	diameter := int(math.Sqrt(float64(b.N)))
	if diameter > 5000 {
		diameter = 5000
	}
	r := image.Rect(0, 0, diameter, diameter)
	source := getrgba(r)
	target := win.NewDIB(r)
	b.StartTimer()

	convertRGBAtoWin(target, source)

	b.StopTimer()
	if !compareImages(source, target) {
		b.Fail()
	}
}
