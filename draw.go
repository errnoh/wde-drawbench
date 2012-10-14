package bench

import (
	"github.com/BurntSushi/xgbutil/xgraphics"
	"github.com/skelterjohn/go.wde/win"
	"image"
)

// This. This seems to be the solution. 0.72ns/op. RGBA -> RGBA is 0.02ns/op. Draw is 180ns/op.
// from https://github.com/BurntSushi/xgbutil/blob/master/xgraphics/convert.go
func convertRGBAtoXgb(dest *xgraphics.Image, src *image.RGBA) {
	var x, y, i, si int

	for x = dest.Rect.Min.X; x < dest.Rect.Max.X; x++ {
		for y = dest.Rect.Min.Y; y < dest.Rect.Max.Y; y++ {
			si = src.PixOffset(x, y)
			i = dest.PixOffset(x, y)
			dest.Pix[i+0] = src.Pix[si+2]
			dest.Pix[i+1] = src.Pix[si+1]
			dest.Pix[i+2] = src.Pix[si+0]
			dest.Pix[i+3] = src.Pix[si+3]
		}
	}
}

func convertRGBAtoWin(dest *win.DIB, src *image.RGBA) {
	var x, y, i, si int

	for x = dest.Rect.Min.X; x < dest.Rect.Max.X; x++ {
		for y = dest.Rect.Min.Y; y < dest.Rect.Max.Y; y++ {
			si = src.PixOffset(x, y)
			i = dest.PixOffset(x, y)
			dest.Pix[i+0] = src.Pix[si+2]
			dest.Pix[i+1] = src.Pix[si+1]
			dest.Pix[i+2] = src.Pix[si+0]
		}
	}
}

/*
// based on https://github.com/skelterjohn/go.wde/blob/master/win/dib_windows.go
func convertRGBAtoWinWithAlpha(dest *win.DIB, src *image.RGBA) {
	var x, y, i, si int

	for x = dest.Rect.Min.X; x < dest.Rect.Max.X; x++ {
		for y = dest.Rect.Min.Y; y < dest.Rect.Max.Y; y++ {
			si = src.PixOffset(x, y)
			i = dest.PixOffset(x, y)
			if src.Pix[si+3] == 0 {
				dest.Pix[i+0] = 0
				dest.Pix[i+1] = 0
				dest.Pix[i+2] = 0
				continue
			} else if src.Pix[si+3] == 255 {
				dest.Pix[i+0] = src.Pix[si+2] //B
				dest.Pix[i+1] = src.Pix[si+1] //G
				dest.Pix[i+2] = src.Pix[si+0] //R
				continue
			}
			// B*a G*a R*a
			//dest.Pix[i+0] = (src.Pix[si+2] / 0xff) * src.Pix[si+3]
			//dest.Pix[i+1] = (src.Pix[si+1] / 0xff) * src.Pix[si+3]
			//dest.Pix[i+2] = (src.Pix[si+0] / 0xff) * src.Pix[si+3]
			dest.Pix[i+0] = uint8((float64(src.Pix[si+2]) / 0xff) * float64(src.Pix[si+3]))
			dest.Pix[i+1] = uint8((float64(src.Pix[si+1]) / 0xff) * float64(src.Pix[si+3]))
			dest.Pix[i+2] = uint8((float64(src.Pix[si+0]) / 0xff) * float64(src.Pix[si+3]))
		}
	}
}
*/
