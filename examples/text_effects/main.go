package main

import(
  "gopkg.in/gographics/imagick.v2/imagick"
  "os/exec"
  "log"
  "fmt"
) // note: in go you want to have log and fmt (and generally all your favourite util packages) at the bottom of the import #golang

var dargs  = []float64{ 120. }
var d_args = []float64{ -0.02, 0.0, 0.0, 1.02, 0.0, 0.0, -0.5, 1.9 }

func main() {
	drawText()
}

// var image_w = 320
// var image_h = 100

var image_w = 320
var image_h = 100

// var IMAGE_W = 320
// var IMAGE_H = 100


// Text effect 5 and 6 - Plain text and then Barrel distortion
func drawText() {
	imagick.Initialize()
	defer imagick.Terminate()

  // var text
  text := "asd"

  mw := imagick.NewMagickWand()
	dw := imagick.NewDrawingWand()
	pw := imagick.NewPixelWand()

  // pw.SetColor("none")
	pw.SetColor("white")
	mw.NewImage(320, 100, pw)

	dw.SetFont("Libra Mono")
	dw.SetFontSize(20)

	dw.Annotation(25, 65, text)

	mw.DrawImage(dw)
	mw.WriteImage("frame1.png")

  // then combine all the frames in bash
  // file_size - print file size

  path := "/home/makevoid/goapps/src/github.com/gographics/imagick/examples/text_effects"

  cmd := path+"/rb/image_size.rb"

  exe()
}

func exe() {
  out, err := exec.Command(cmd).Output()
  if err != nil {
      log.Fatal(err)
  }
  // fmt.Printf("executing '#{cmd}':\n\n#{out}\n\n")
  fmt.Printf("executing '%s':\n\n%s\n\n", cmd, out)
}
