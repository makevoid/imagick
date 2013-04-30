package imagick

/*
#cgo pkg-config: MagickWand
#include <wand/MagickWand.h>
*/
import "C"
import (
	"os"
	"unsafe"
)

// Returns the current image from the magick wand
func (mw *MagickWand) GetImageFromMagickWand() *Image {
	return newImageFromCAPI(C.GetImageFromMagickWand(mw.wand))
}

// Adaptively blurs the image by blurring less intensely near image edges and more intensely far
// from edges. We blur the image with a Gaussian operator of the given radius and standard deviation
// (sigma). For reasonable results, radius should be larger than sigma. Use a radius of 0 and
// AdaptiveBlurImage() selects a suitable radius for you.
// radius: the radius of the Gaussian, in pixels, not counting the center pixel
// sigma: the standard deviation of the Gaussian, in pixels
func (mw *MagickWand) AdaptiveBlurImage(radius, sigma float64) error {
	C.MagickAdaptiveBlurImage(mw.wand, C.double(radius), C.double(sigma))
	return mw.GetLastError()
}

// Adaptively blurs the image by blurring less intensely near image edges and more intensely far
// from edges. We blur the image with a Gaussian operator of the given radius and standard deviation
// (sigma). For reasonable results, radius should be larger than sigma. Use a radius of 0 and
// AdaptiveBlurImage() selects a suitable radius for you.
// radius: the radius of the Gaussian, in pixels, not counting the center pixel
// sigma: the standard deviation of the Gaussian, in pixels
func (mw *MagickWand) AdaptiveBlurImageChannel(channel ChannelType, radius, sigma float64) error {
	C.MagickAdaptiveBlurImageChannel(mw.wand, C.ChannelType(channel), C.double(radius), C.double(sigma))
	return mw.GetLastError()
}

// Adaptively resize image with data dependent triangulation
func (mw *MagickWand) AdaptiveResizeImage(columns, rows uint) error {
	C.MagickAdaptiveResizeImage(mw.wand, C.size_t(columns), C.size_t(rows))
	return mw.GetLastError()
}

// Adaptively sharpens the image by sharpening more intensely near image edges and less intensely far from edges.
// We sharpen the image with a Gaussian operator of the given radius and standard deviation (sigma). For reasonable
// results, radius should be larger than sigma. Use a radius of 0 and AdaptiveSharpenImage() selects a suitable
// radius for you.
// radius: the radius of the Gaussian, in pixels, not counting the center pixel
// sigma: the standard deviation of the Gaussian, in pixels.
func (mw *MagickWand) AdaptiveSharpenImage(radius, sigma float64) error {
	C.MagickAdaptiveSharpenImage(mw.wand, C.double(radius), C.double(sigma))
	return mw.GetLastError()
}

// Adaptively sharpens the image by sharpening more intensely near image edges and less intensely far from edges.
// We sharpen the image with a Gaussian operator of the given radius and standard deviation (sigma). For reasonable
// results, radius should be larger than sigma. Use a radius of 0 and AdaptiveSharpenImage() selects a suitable
// radius for you.
// radius: the radius of the Gaussian, in pixels, not counting the center pixel
// sigma: the standard deviation of the Gaussian, in pixels.
func (mw *MagickWand) AdaptiveSharpenImageChannel(channel ChannelType, radius, sigma float64) error {
	C.MagickAdaptiveSharpenImageChannel(mw.wand, C.ChannelType(channel), C.double(radius), C.double(sigma))
	return mw.GetLastError()
}

// Selects an individual threshold for each pixel based on the range of intensity values in its local neighborhood.
// This allows for thresholding of an image whose global intensity histogram doesn't contain distinctive peaks.
func (mw *MagickWand) AdaptiveThresholdImage(width, height uint, offset int) error {
	C.MagickAdaptiveThresholdImage(mw.wand, C.size_t(width), C.size_t(height), C.ssize_t(offset))
	return mw.GetLastError()
}

// Adds a clone of the images from the second wand and inserts them into the first wand.
// Use SetLastIterator(), to append new images into an existing wand, current image will be set to last image so
// later adds with also be appened to end of wand.
// Use SetFirstIterator() to prepend new images into wand, any more images added will also be prepended before other
// images in the wand. However the order of a list of new images will not change.
// Otherwise the new images will be inserted just after the current image, and any later image will also be added
// after this current image but before the previously added images. Caution is advised when multiple image adds are
// inserted into the middle of the wand image list.
func (mw *MagickWand) AddImage(wand *MagickWand) error {
	C.MagickAddImage(mw.wand, wand.wand)
	return mw.GetLastError()
}

// Adds random noise to the image
func (mw *MagickWand) AddNoiseImage(noiseType NoiseType) error {
	C.MagickAddNoiseImage(mw.wand, C.NoiseType(noiseType))
	return mw.GetLastError()
}

// Adds random noise to the image's channel
func (mw *MagickWand) AddNoiseImageChannel(channel ChannelType, noiseType NoiseType) error {
	C.MagickAddNoiseImageChannel(mw.wand, C.ChannelType(channel), C.NoiseType(noiseType))
	return mw.GetLastError()
}

// Transforms an image as dictaded by the affine matrix of the drawing wand
func (mw *MagickWand) AffineTransformImage(drawingWand *DrawingWand) error {
	C.MagickAffineTransformImage(mw.wand, drawingWand.draw)
	return mw.GetLastError()
}

// Annotates an image with text
// x: ordinate to left of text
// y: ordinate to text baseline
// angle: rotate text relative to this angle
func (mw *MagickWand) AnnotateImage(drawingWand *DrawingWand, x, y, angle float64, text string) error {
	cstext := C.CString(text)
	defer C.free(unsafe.Pointer(cstext))
	C.MagickAnnotateImage(mw.wand, drawingWand.draw, C.double(x), C.double(y), C.double(angle), cstext)
	return mw.GetLastError()
}

// Animates an image or image sequence
func (mw *MagickWand) AnimateImages(server string) error {
	csserver := C.CString(server)
	defer C.free(unsafe.Pointer(csserver))
	C.MagickAnimateImages(mw.wand, csserver)
	return mw.GetLastError()
}

// Append the images in a wand from the current image onwards, creating a new wand with the single image result.
// This is affected by the gravity and background setting of the first image.
// Typically you would call either ResetIterator() or SetFirstImage() before calling this function to ensure that
// all the images in the wand's image list will be appended together.
// By default, images are stacked left-to-right. Set topToBottom to true to stack them top-to-bottom.
func (mw *MagickWand) AppendImages(topToBottom bool) *MagickWand {
	stack := C.MagickBooleanType(0)
	if topToBottom {
		stack = C.MagickBooleanType(1)
	}
	return &MagickWand{wand: C.MagickAppendImages(mw.wand, stack)}
}

// Extracts the 'mean' from the image and adjust the image to try make set it's gamma appropriatally
func (mw *MagickWand) AutoGammaImage() error {
	C.MagickAutoGammaImage(mw.wand)
	return mw.GetLastError()
}

// Extracts the 'mean' from the image's channel and adjust the image to try make set it's gamma appropriatally
func (mw *MagickWand) AutoGammaImageChannel(channel ChannelType) error {
	C.MagickAutoGammaImageChannel(mw.wand, C.ChannelType(channel))
	return mw.GetLastError()
}

// Adjust the levels of a particular image by scaling the minimum and maximum values to the full quantum range.
func (mw *MagickWand) AutoLevelImage() error {
	C.MagickAutoLevelImage(mw.wand)
	return mw.GetLastError()
}

// Adjust the levels of a particular image channel by scaling the minimum and maximum values to the full quantum range.
func (mw *MagickWand) AutoLevelImageChannel(channel ChannelType) error {
	C.MagickAutoLevelImageChannel(mw.wand, C.ChannelType(channel))
	return mw.GetLastError()
}

// This is like ThresholdImage() but forces all pixels below the threshold into black while leaving all
// pixels above the threshold unchanged.
func (mw *MagickWand) BlackThresholdImage(threshold *PixelWand) error {
	C.MagickBlackThresholdImage(mw.wand, threshold.pixel)
	return mw.GetLastError()
}

// Mutes the colors of the image to simulate a scene at nighttime in the moonlight.
func (mw *MagickWand) BlueShiftImage(factor float64) error {
	C.MagickBlueShiftImage(mw.wand, C.double(factor))
	return mw.GetLastError()
}

// Blurs an image. We convolve the image with a gaussian operator of the given radius and standard deviation (sigma).
// For reasonable results, the radius should be larger than sigma. Use a radius of 0 and BlurImage() selects a suitable
// radius for you.
// radius: the radius of the, in pixels, not counting the center pixel.
// sigma: the standard deviation of the, in pixels
func (mw *MagickWand) BlurImage(radius, sigma float64) error {
	C.MagickBlurImage(mw.wand, C.double(radius), C.double(sigma))
	return mw.GetLastError()
}

// Blurs an image's channel. We convolve the image with a gaussian operator of the given radius and standard deviation (sigma).
// For reasonable results,
// the radius should be larger than sigma. Use a radius of 0 and BlurImage() selects a suitable radius for you.
// radius: the radius of the, in pixels, not counting the center pixel.
// sigma: the standard deviation of the, in pixels
func (mw *MagickWand) BlurImageChannel(channel ChannelType, radius, sigma float64) error {
	C.MagickBlurImageChannel(mw.wand, C.ChannelType(channel), C.double(radius), C.double(sigma))
	return mw.GetLastError()
}

// Surrounds the image with a border of the color defined by the bordercolor pixel wand.
func (mw *MagickWand) BorderImage(borderColor *PixelWand, width, height uint) error {
	C.MagickBorderImage(mw.wand, borderColor.pixel, C.size_t(width), C.size_t(height))
	return mw.GetLastError()
}

// Use this to change the brightness and/or contrast of an image. It converts the brightness and contrast
// brighness: the brightness percent (-100 .. 100)
// contrast: the brightness percent (-100 .. 100)
func (mw *MagickWand) BrightnessContrastImage(brightness, contrast float64) error {
	C.MagickBrightnessContrastImage(mw.wand, C.double(brightness), C.double(contrast))
	return mw.GetLastError()
}

// Use this to change the brightness and/or contrast of an image's channel. It converts the brightness and contrast
// brighness: the brightness percent (-100 .. 100)
// contrast: the brightness percent (-100 .. 100)
func (mw *MagickWand) BrightnessContrastImageChannel(channel ChannelType, brightness, contrast float64) error {
	C.MagickBrightnessContrastImageChannel(mw.wand, C.ChannelType(channel), C.double(brightness), C.double(contrast))
	return mw.GetLastError()
}

// Simulates a charcoal drawing
// radius: the radius of the Gaussian, in pixels, not counting the center pixel
// sigma: the standard deviation of the Gaussian, in pixels
func (mw *MagickWand) CharcoalImage(radius, sigma float64) error {
	C.MagickCharcoalImage(mw.wand, C.double(radius), C.double(sigma))
	return mw.GetLastError()
}

// Removes a region of an image and collapses the image to occupy the removed portion
// width, height: the region width and height
// x, y: the region x and y offsets
func (mw *MagickWand) ChopImage(width, height uint, x, y int) error {
	C.MagickChopImage(mw.wand, C.size_t(width), C.size_t(height), C.ssize_t(x), C.ssize_t(y))
	return mw.GetLastError()
}

// Restricts the color range from 0 to the quantum depth
func (mw *MagickWand) ClampImage() error {
	C.MagickClampImage(mw.wand)
	return mw.GetLastError()
}

// Restricts the color range from 0 to the quantum depth
func (mw *MagickWand) ClampImageChannel(channel ChannelType) error {
	C.MagickClampImageChannel(mw.wand, C.ChannelType(channel))
	return mw.GetLastError()
}

// Clips along the first path from the 8BIM profile, if present
func (mw *MagickWand) ClipImage() error {
	C.MagickClipImage(mw.wand)
	return mw.GetLastError()
}

// Clips along the named paths from the 8BOM profile, if present. Later operations take
// effect inside the path. Id may be a number if preceded with #, to work on a numbered
// path, e.g. "#1" to use the first path.
// pathname: name of clipping path resource. If name is preceded by #, use clipping path numbered by name
// inside: if true, later operations take effect inside clipping path. Otherwise later operations take effect outside clipping path
func (mw *MagickWand) ClipImagePath(pathname string, inside bool) error {
	cspathname := C.CString(pathname)
	defer C.free(unsafe.Pointer(cspathname))
	csinside := 0
	if inside {
		csinside = 1
	}
	C.MagickClipImagePath(mw.wand, cspathname, C.MagickBooleanType(csinside))
	return mw.GetLastError()
}

// Replaces colors in the image from a color lookup table
func (mw *MagickWand) ClutImage(clut *MagickWand) error {
	C.MagickClutImage(mw.wand, clut.wand)
	return mw.GetLastError()
}

// Replaces colors in the image's channel from a color lookup table
func (mw *MagickWand) ClutImageChannel(channel ChannelType, clut *MagickWand) error {
	C.MagickClutImageChannel(mw.wand, C.ChannelType(channel), clut.wand)
	return mw.GetLastError()
}

// Composites a set of images while respecting any page offsets and disposal methods. GIF, MIFF, and MNG
// animation sequences typically start with an image background and each subsequent image varies in size
// and offset. CoalesceImages() returns a new sequence where each image in the sequence is the same size
// as the first and composited with the next image in the sequence.
func (mw *MagickWand) CoalesceImages() *MagickWand {
	return &MagickWand{wand: C.MagickCoalesceImages(mw.wand)}
}

// Accepts a lightweight Color Correction Collection (CCC) file which solely contains one or more color
// corrections and applies the color correction to the image. Here is a sample CCC file content:
// <colorcorrectioncollection xmlns="urn:ASC:CDL:v1.2">
//   <colorcorrection id="cc03345">
//     <sopnode>
//       <slope> 0.9 1.2 0.5 </slope>
//       <offset> 0.4 -0.5 0.6 </offset>
//       <power> 1.0 0.8 1.5 </power>
//     </sopnode>
//     <satnode>
//       <saturation> 0.85 </saturation>
//     </satnode>
//   </colorcorrection>
// </colorcorrectioncollection>
func (mw *MagickWand) ColorDecisionListImage(cccXML string) error {
	cscccXML := C.CString(cccXML)
	defer C.free(unsafe.Pointer(cscccXML))
	C.MagickColorDecisionListImage(mw.wand, cscccXML)
	return mw.GetLastError()
}

// Blends the fill color with each pixel in the image
func (mw *MagickWand) ColorizeImage(colorize, opacity *PixelWand) error {
	C.MagickColorizeImage(mw.wand, colorize.pixel, opacity.pixel)
	return mw.GetLastError()
}

// Apply color transformation to an image. The method permits saturation changes, hue rotation, luminance
// to alpha, and various other effects. Although variable-sized transformation matrices can be used,
// typically one uses a 5x5 matrix for an RGBA image and a 6x6 for CMYKA (or RGBA with offsets). The matrix
// is similar to those used by Adobe Flash except offsets are in column 6 rather than 5 (in support of CMYKA
// images) and offsets are normalized (divide Flash offset by 255).
func (mw *MagickWand) ColorMatrixImage(colorMatrix *KernelInfo) error {
	C.MagickColorMatrixImage(mw.wand, colorMatrix.info)
	return mw.GetLastError()
}

// Combines one or more images into a single image. The grayscale value of the pixels of each image in the
// sequence is assigned in order to the specified hannels of the combined image. The typical ordering would
// be image 1 => Red, 2 => Green, 3 => Blue, etc.
func (mw *MagickWand) CombineImages(channel ChannelType) *MagickWand {
	return &MagickWand{C.MagickCombineImages(mw.wand, C.ChannelType(channel))}
}

// Adds a comment to your image
func (mw *MagickWand) CommentImage(comment string) error {
	cscomment := C.CString(comment)
	defer C.free(unsafe.Pointer(cscomment))
	C.MagickCommentImage(mw.wand, cscomment)
	return mw.GetLastError()
}

// Compares one or more image channels of an image to a reconstructed image and returns the difference image
func (mw *MagickWand) CompareImageChannels(reference *MagickWand, channel ChannelType, metric MetricType) (wand *MagickWand, distortion float64) {
	cdistortion := C.double(0)
	cmw := C.MagickCompareImageChannels(mw.wand, reference.wand, C.ChannelType(channel), C.MetricType(metric), &cdistortion)
	wand = &MagickWand{cmw}
	distortion = float64(cdistortion)
	return
}

// Compares each image with the next in a sequence and returns the maximum bounding region of any pixel differences it discovers.
func (mw *MagickWand) CompareImageLayers(method ImageLayerMethod) *MagickWand {
	return &MagickWand{C.MagickCompareImageLayers(mw.wand, C.ImageLayerMethod(method))}
}

// CompareImages() compares an image to a reconstructed image and returns the specified difference image.
// Returns the new MagickWand and the computed distortion between the images
func (mw *MagickWand) CompareImages(reference *MagickWand, metric MetricType) (wand *MagickWand, distortion float64) {
	cdistortion := C.double(0)
	cmw := C.MagickCompareImages(mw.wand, reference.wand, C.MetricType(metric), &cdistortion)
	wand = &MagickWand{cmw}
	distortion = float64(cdistortion)
	return
}

// Composite one image onto another at the specified offset.
// source: The magick wand holding source image.
// compose: This operator affects how the composite is applied to the image. The default is Over.
// x: the column offset of the composited image.
// y: the row offset of the composited image.
func (mw *MagickWand) CompositeImage(source *MagickWand, compose CompositeOperator, x, y int) error {
	C.MagickCompositeImage(mw.wand, source.wand, C.CompositeOperator(compose), C.ssize_t(x), C.ssize_t(y))
	return mw.GetLastError()
}

// Composite one image onto another at the specified offset.
// source: The magick wand holding source image.
// compose: This operator affects how the composite is applied to the image. The default is Over.
// x: the column offset of the composited image.
// y: the row offset of the composited image.
func (mw *MagickWand) CompositeImageChannel(channel ChannelType, source *MagickWand, compose CompositeOperator, x, y int) error {
	C.MagickCompositeImageChannel(mw.wand, C.ChannelType(channel), source.wand, C.CompositeOperator(compose), C.ssize_t(x), C.ssize_t(y))
	return mw.GetLastError()
}

// Composite the images in the source wand over the images in the destination wand in sequence,
// starting with the current image in both lists.
// Each layer from the two image lists are composted together until the end of one of the image
// lists is reached. The offset of each composition is also adjusted to match the virtual canvas
// offsets of each layer. As such the given offset is relative to the virtual canvas, and not
// the actual image.
// Composition uses given x and y offsets, as the 'origin' location of the source images virtual
// canvas (not the real image) allowing you to compose a list of 'layer images' into the
// destination images. This makes it well sutiable for directly composing 'Clears Frame Animations'
// or 'Coaleased Animations' onto a static or other 'Coaleased Animation' destination image list.
// GIF disposal handling is not looked at.
// Special case: If one of the image sequences is the last image (just a single image remaining),
// that image is repeatally composed with all the images in the other image list. Either the source
// or destination lists may be the single image, for this situation.
// In the case of a single destination image (or last image given), that image will ve cloned to
// match the number of images remaining in the source image list.
// This is equivelent to the "-layer Composite" Shell API operator.
// source: the wand holding the source images
// compose, x, y: composition arguments
func (mw *MagickWand) CompositeLayers(source *MagickWand, compose CompositeOperator, x, y int) error {
	C.MagickCompositeLayers(mw.wand, source.wand, C.CompositeOperator(compose), C.ssize_t(x), C.ssize_t(y))
	return mw.GetLastError()
}

// Enhances the intensity differences between the lighter and darker elements of the image. Set sharpen
// to a value other than 0 to increase the image contrast otherwise the contrast is reduced.
// sharpen: increase or decrease image contrast
func (mw *MagickWand) ContrastImage(sharpen bool) error {
	cssharpen := C.MagickBooleanType(0)
	if sharpen {
		cssharpen = 1
	}
	C.MagickContrastImage(mw.wand, cssharpen)
	return mw.GetLastError()
}

// Enhances the contrast of a color image by adjusting the pixels color to span the entire range of colors
// available. You can also reduce the influence of a particular channel with a gamma value of 0.
func (mw *MagickWand) ContrastStretchImage(blackPoint, whitePoint float64) error {
	C.MagickContrastStretchImage(mw.wand, C.double(blackPoint), C.double(whitePoint))
	return mw.GetLastError()
}

// Enhances the contrast of a color image's channel by adjusting the pixels color to span the entire range of colors
// available. You can also reduce the influence of a particular channel with a gamma value of 0.
func (mw *MagickWand) ContrastStretchImageChannel(channel ChannelType, blackPoint, whitePoint float64) error {
	C.MagickContrastStretchImageChannel(mw.wand, C.ChannelType(channel), C.double(blackPoint), C.double(whitePoint))
	return mw.GetLastError()
}

// Applies a custom convolution kernel to the image.
// order: the number of columns and rows in the filter kernel
// kernel: an array of doubles, representing the convolution kernel
func (mw *MagickWand) ConvolveImage(order uint, kernel []float64) error {
	ckernel := []C.double{}
	for i, k := range kernel {
		ckernel[i] = C.double(k)
	}
	C.MagickConvolveImage(mw.wand, C.size_t(order), &ckernel[0])
	return mw.GetLastError()
}

// Applies a custom convolution kernel to the image's channel.
// order: the number of columns and rows in the filter kernel
// kernel: an array of doubles, representing the convolution kernel
func (mw *MagickWand) ConvolveImageChannel(channel ChannelType, order uint, kernel []float64) error {
	ckernel := []C.double{}
	for i, k := range kernel {
		ckernel[i] = C.double(k)
	}
	C.MagickConvolveImageChannel(mw.wand, C.ChannelType(channel), C.size_t(order), &ckernel[0])
	return mw.GetLastError()
}

// Extracts a region of the image
func (mw *MagickWand) CropImage(width, height uint, x, y int) error {
	C.MagickCropImage(mw.wand, C.size_t(width), C.size_t(height), C.ssize_t(x), C.ssize_t(y))
	return mw.GetLastError()
}

// Displaces an Image's colormap by a given number of positions. If you cycle the colormap a number of
// times you can produce a psychodelic effect.
func (mw *MagickWand) CycleColormapImage(displace int) error {
	C.MagickCycleColormapImage(mw.wand, C.ssize_t(displace))
	return mw.GetLastError()
}

// Adds an image to the wand comprised of the pixel data you supply. The pixel data must be in scanline
// order top-to-bottom.
// storage: Define the data type of the pixels. Float and double types are expected to be normalized [0..1]
// otherwise [0..QuantumRange].
// pmap: This string reflects the expected ordering of the pixel array. It can be any combination or order
// of R = red, G = green, B = blue, A = alpha (0 is transparent), O = opacity (0 is opaque), C = cyan,
// Y = yellow, M = magenta, K = black, I = intensity (for grayscale), P = pad.
// pixels: This array of values contain the pixel components as defined by the type.
func (mw *MagickWand) ConstituteImageFromBytes(cols, rows uint, pmap string, pixels []uint8) error {
	cspmap := C.CString(pmap)
	defer C.free(unsafe.Pointer(cspmap))
	C.MagickConstituteImage(mw.wand, C.size_t(cols), C.size_t(rows), cspmap,
		C.StorageType(CharPixel), unsafe.Pointer(&pixels[0]))
	return mw.GetLastError()
}

// Adds an image to the wand comprised of the pixel data you supply. The pixel data must be in scanline
// order top-to-bottom.
// storage: Define the data type of the pixels. Float and double types are expected to be normalized [0..1]
// otherwise [0..QuantumRange].
// pmap: This string reflects the expected ordering of the pixel array. It can be any combination or order
// of R = red, G = green, B = blue, A = alpha (0 is transparent), O = opacity (0 is opaque), C = cyan,
// Y = yellow, M = magenta, K = black, I = intensity (for grayscale), P = pad.
// pixels: This array of values contain the pixel components as defined by the type.
func (mw *MagickWand) ConstituteImageFromShort(cols, rows uint, pmap string, pixels []uint16) error {
	cspmap := C.CString(pmap)
	defer C.free(unsafe.Pointer(cspmap))
	C.MagickConstituteImage(mw.wand, C.size_t(cols), C.size_t(rows), cspmap,
		C.StorageType(ShortPixel), unsafe.Pointer(&pixels[0]))
	return mw.GetLastError()
}

// Adds an image to the wand comprised of the pixel data you supply. The pixel data must be in scanline
// order top-to-bottom.
// storage: Define the data type of the pixels. Float and double types are expected to be normalized [0..1]
// otherwise [0..QuantumRange].
// pmap: This string reflects the expected ordering of the pixel array. It can be any combination or order
// of R = red, G = green, B = blue, A = alpha (0 is transparent), O = opacity (0 is opaque), C = cyan,
// Y = yellow, M = magenta, K = black, I = intensity (for grayscale), P = pad.
// pixels: This array of values contain the pixel components as defined by the type.
func (mw *MagickWand) ConstituteImageFromInt(cols, rows uint, pmap string, pixels []uint32) error {
	cspmap := C.CString(pmap)
	defer C.free(unsafe.Pointer(cspmap))
	C.MagickConstituteImage(mw.wand, C.size_t(cols), C.size_t(rows), cspmap,
		C.StorageType(IntegerPixel), unsafe.Pointer(&pixels[0]))
	return mw.GetLastError()
}

// Adds an image to the wand comprised of the pixel data you supply. The pixel data must be in scanline
// order top-to-bottom.
// storage: Define the data type of the pixels. Float and double types are expected to be normalized [0..1]
// otherwise [0..QuantumRange].
// pmap: This string reflects the expected ordering of the pixel array. It can be any combination or order
// of R = red, G = green, B = blue, A = alpha (0 is transparent), O = opacity (0 is opaque), C = cyan,
// Y = yellow, M = magenta, K = black, I = intensity (for grayscale), P = pad.
// pixels: This array of values contain the pixel components as defined by the type.
func (mw *MagickWand) ConstituteImageFromFloat32(cols, rows uint, pmap string, pixels []float32) error {
	cspmap := C.CString(pmap)
	defer C.free(unsafe.Pointer(cspmap))
	C.MagickConstituteImage(mw.wand, C.size_t(cols), C.size_t(rows), cspmap,
		C.StorageType(FloatPixel), unsafe.Pointer(&pixels[0]))
	return mw.GetLastError()
}

// Adds an image to the wand comprised of the pixel data you supply. The pixel data must be in scanline
// order top-to-bottom.
// storage: Define the data type of the pixels. Float and double types are expected to be normalized [0..1]
// otherwise [0..QuantumRange].
// pmap: This string reflects the expected ordering of the pixel array. It can be any combination or order
// of R = red, G = green, B = blue, A = alpha (0 is transparent), O = opacity (0 is opaque), C = cyan,
// Y = yellow, M = magenta, K = black, I = intensity (for grayscale), P = pad.
// pixels: This array of values contain the pixel components as defined by the type.
func (mw *MagickWand) ConstituteImageFromFloat64(cols, rows uint, pmap string, pixels []float64) error {
	cspmap := C.CString(pmap)
	defer C.free(unsafe.Pointer(cspmap))
	C.MagickConstituteImage(mw.wand, C.size_t(cols), C.size_t(rows), cspmap,
		C.StorageType(DoublePixel), unsafe.Pointer(&pixels[0]))
	return mw.GetLastError()
}

// Converts cipher pixels to plain pixels
func (mw *MagickWand) DecipherImage(passphrase string) error {
	cspassphrase := C.CString(passphrase)
	defer C.free(unsafe.Pointer(cspassphrase))
	C.MagickDecipherImage(mw.wand, cspassphrase)
	return mw.GetLastError()
}

// Compares each image with the next in a sequence and returns the maximum bouding region of any pixel
// differences it discovers
func (mw *MagickWand) DeconstructImages() *MagickWand {
	return &MagickWand{C.MagickDeconstructImages(mw.wand)}
}

// Removes skew from the image. Skew is an artifact that occurs in scanned images because of the camera
// being misaligned, imperfections in the scanning or surface, or simply because the paper was not placed
// completely flat when scanned.
// threshold: separate background from foreground
func (mw *MagickWand) DeskewImage(threshold float64) error {
	C.MagickDeskewImage(mw.wand, C.double(threshold))
	return mw.GetLastError()
}

// Reduces the speckle noise in an image while perserving the edges of the original image.
func (mw *MagickWand) DespeckleImage() error {
	C.MagickDespeckleImage(mw.wand)
	return mw.GetLastError()
}

// Displays and image
func (mw *MagickWand) DisplayImage(server string) error {
	cstring := C.CString(server)
	defer C.free(unsafe.Pointer(cstring))
	C.MagickDisplayImage(mw.wand, cstring)
	return mw.GetLastError()
}

// Displays and image or image sequence
func (mw *MagickWand) DisplayImages(server string) error {
	cstring := C.CString(server)
	defer C.free(unsafe.Pointer(cstring))
	C.MagickDisplayImages(mw.wand, cstring)
	return mw.GetLastError()
}

// DistortImage() distorts an image using various distortion methods, by mapping color lookups of the
// source image to a new destination image usally of the same size as the source image, unless 'bestfit'
// is set to true. If 'bestfit' is enabled, and distortion allows it, the destination image is adjusted
// to ensure the whole source 'image' will just fit within the final destination image, which will be
// sized and offset accordingly. Also in many cases the virtual offset of the source image will be taken
// into account in the mapping.
// method: the method of image distortion.
// ArcDistortion always ignores the source image offset, and always 'bestfit' the destination
// image with the top left corner offset relative to the polar mapping center.
// Bilinear has no simple inverse mapping so it does not allow 'bestfit' style of image distortion.
// Affine, Perspective, and Bilinear, do least squares fitting of the distortion when more than the
// minimum number of control point pairs are provided.
// Perspective, and Bilinear, falls back to a Affine distortion when less that 4 control point pairs
// are provided. While Affine distortions let you use any number of control point pairs, that is Zero
// pairs is a no-Op (viewport only) distrotion, one pair is a translation and two pairs of control
// points do a scale-rotate-translate, without any shearing.
// args: the arguments for this distortion method.
// bestfit: Attempt to resize destination to fit distorted source.
func (mw *MagickWand) DistortImage(method DistortImageMethod, args []float64, bestfit bool) error {
	cbestfit := C.MagickBooleanType(0)
	if bestfit {
		cbestfit = C.MagickBooleanType(1)
	}
	C.MagickDistortImage(mw.wand, C.DistortImageMethod(method), C.size_t(len(args)), (*C.double)(&args[0]), cbestfit)
	return mw.GetLastError()
}

// Renders the drawing wand on the current image
func (mw *MagickWand) DrawImage(drawingWand *DrawingWand) error {
	C.MagickDrawImage(mw.wand, drawingWand.draw)
	return mw.GetLastError()
}

// Enhance edges within the image with a convolution filter of the
// given radius. Use a radius of 0 and Edge() selects a suitable radius
// for you.
// radius: the radius of the pixel neighborhood
func (mw *MagickWand) EdgeImage(radius float64) error {
	C.MagickEdgeImage(mw.wand, C.double(radius))
	return mw.GetLastError()
}

// Returns a grayscale image with a three-dimensional effect. We convolve the image with a Gaussian operator
// of the given radius and standard deviation (sigma). For reasonable results, radius should be larger than
// sigma. Use a radius of 0 and Emboss() selects a suitable radius for you.
// radius: the radius of the Gaussian, in pixels, not counting the center pixel
// sigma: the standard deviation of the Gaussian, in pixels
func (mw *MagickWand) EmbossImage(radius, sigma float64) error {
	C.MagickEmbossImage(mw.wand, C.double(radius), C.double(sigma))
	return mw.GetLastError()
}

// Converts plain pixels to cipher pixels
func (mw *MagickWand) EncipherImage(passphrase string) error {
	cspassphrase := C.CString(passphrase)
	defer C.free(unsafe.Pointer(cspassphrase))
	C.MagickEncipherImage(mw.wand, cspassphrase)
	return mw.GetLastError()
}

// Applies a digital filter that improves the quality of a noisy image
//func (mw *MagickWand) EnhanceImage() error {
//C.MagickEnhangeImage(mw.wand)
//return mw.GetLastError()
//}
// Equalizes the image histogram.
func (mw *MagickWand) EqualizeImage() error {
	C.MagickEqualizeImage(mw.wand)
	return mw.GetLastError()
}

// Equalizes the image's channel histogram.
func (mw *MagickWand) EqualizeImageChannel(channel ChannelType) error {
	C.MagickEqualizeImageChannel(mw.wand, C.ChannelType(channel))
	return mw.GetLastError()
}

// Applys an arithmetic, relational, or logical expression to an image.
// Use these operators to lighten or darken an image, to increase or
// decrease contrast in an image, or to produce the "negative" of an image.
func (mw *MagickWand) EvaluateImage(op EvaluateOperator, value float64) error {
	C.MagickEvaluateImage(mw.wand, C.MagickEvaluateOperator(op), C.double(value))
	return mw.GetLastError()
}

// Applys an arithmetic, relational, or logical expression to an image.
// Use these operators to lighten or darken an image, to increase or
// decrease contrast in an image, or to produce the "negative" of an image.
func (mw *MagickWand) EvaluateImages(op EvaluateOperator) error {
	C.MagickEvaluateImages(mw.wand, C.MagickEvaluateOperator(op))
	return mw.GetLastError()
}

// Applys an arithmetic, relational, or logical expression to an image.
// Use these operators to lighten or darken an image, to increase or
// decrease contrast in an image, or to produce the "negative" of an image.
func (mw *MagickWand) EvaluateImageChannel(channel ChannelType, op EvaluateOperator, value float64) error {
	C.MagickEvaluateImageChannel(mw.wand, C.ChannelType(channel), C.MagickEvaluateOperator(op), C.double(value))
	return mw.GetLastError()
}

// Extracts pixel data from an image and returns it to you.
// x, y, columns, rows: These values define the perimeter of a region of pixels you want to extract.
// map: This string reflects the expected ordering of the pixel array. It can be any combination or order of
// R = red, G = green, B = blue, A = alpha (0 is transparent), O = opacity (0 is opaque), C = cyan, Y = yellow,
// M = magenta, K = black, I = intensity (for grayscale), P = pad.
// storage: Define the data type of the pixels. Float and double types are expected to be normalized [0..1]
// otherwise [0..QuantumRange]. Choose from these types: CharPixel, DoublePixel, FloatPixel, IntegerPixel,
// LongPixel, QuantumPixel, or ShortPixel.
// pixels: This array of values contain the pixel components as defined by map and type.
func (mw *MagickWand) ExportImagePixels(x, y int, cols, rows uint, pmap string, storage StorageType) (pixels []interface{}, err error) {
	cspmap := C.CString(pmap)
	defer C.free(unsafe.Pointer(cspmap))
	pixels = make([]interface{}, (cols-x)*(rows-y)*len(pmap))
	C.MagickExportImagePixels(mw.wand, C.size_t(cols), C.size_t(rows), cspmap, C.StorageType(CharPixel), unsafe.Pointer(&pixels[0]))
	return pixels, mw.GetLastError()
}

// Extends the image as defined by the geometry, gravity, and wand background color. Set the (x,y) offset of the
// geometry to move the original wand relative to the extended wand.
// width: the region width.
// height: the region height.
// x: the region x offset.
// y: the region y offset.
func (mw *MagickWand) ExtendImage(width, height uint, x, y int) error {
	C.MagickExtentImage(mw.wand, C.size_t(width), C.size_t(height), C.ssize_t(x), C.ssize_t(y))
	return mw.GetLastError()
}

// Applies a custom convolution kernel to the image.
//  kernel: An array of doubles representing the convolution kernel.
func (mw *MagickWand) FilterImage(kernel *KernelInfo) error {
	C.MagickFilterImage(mw.wand, kernel.info)
	return mw.GetLastError()
}

// Applies a custom convolution kernel to the image's channel.
//  kernel: An array of doubles representing the convolution kernel.
func (mw *MagickWand) FilterImageChannel(channel ChannelType, kernel *KernelInfo) error {
	C.MagickFilterImageChannel(mw.wand, kernel.info)
	return mw.GetLastError()
}

// Creates a vertical mirror image by reflecting the pixels around the central x-axis.
func (mw *MagickWand) FlipImage() error {
	C.MagickFlipImage(mw.wand)
	return mw.GetLastError()
}

// Changes the color value of any pixel that matches target and is an immediate neighbor.
// If the method FillToBorderMethod is specified, the color value is changed for any neighbor
// pixel that does not match the bordercolor member of image.
// fill: the floodfill color pixel wand.
// fuzz: By default target must match a particular pixel color exactly. However, in many cases
// two colors may differ by a small amount. The fuzz member of image defines how much tolerance
// is acceptable to consider two colors as the same. For example, set fuzz to 10 and the color
// red at intensities of 100 and 102 respectively are now interpreted as the same color for the
// purposes of the floodfill.
// bordercolor: the border color pixel wand.
// x, y: the starting location of the operation.
// invert: paint any pixel that does not match the target color.
func (mw *MagickWand) FloodfillPaintImage(channel ChannelType, fill *PixelWand, fuzz float64, borderColor *PixelWand, x, y int, invert bool) error {
	cinvert := C.MagickBooleanType(0)
	if invert {
		cinvert := C.MagickBooleanType(1)
	}
	C.MagickFloodfillPaintImage(mw.wand, C.ChannelType(channel), fill.pixel, C.double(fuzz), borderColor.pixel, C.ssize_t(x), C.ssize_t(y), cinvert)
	return mw.GetLastError()
}

// Creates a horizontal mirror image by reflecting the pixels around the central y-axis.
func (mw *MagickWand) UniqueImageColors() error {
	C.MagickUniqueImageColors(mw.wand)
	return mw.GetLastError()
}

// Implements the discrete Fourier transform (DFT) of the image either as a
// magnitude/phase or real/imaginary image pair.
// magnitude: if true, return as magnitude / phase pair otherwise a real / imaginary image pair.
func (mw *MagickWand) ForwardFourierTransformImage(magnitude bool) error {
	cmagnitude := C.MagickBooleanType(0)
	if magnitude {
		cmagnitude := C.MagickBooleanType(1)
	}
	C.MagickForwardFourierTransformImage(mw.wand, cmagnitude)
	return mw.GetLastError()
}

// Adds a simulated three-dimensional border around the image. The width and height specify the
// border width of the vertical and horizontal sides of the frame. The inner and outer bevels
// indicate the width of the inner and outer shadows of the frame.
// matteColor: the frame color pixel wand.
// width: the border width.
// height: the border height.
// innerBevel: the inner bevel width.
// outerBevel: the outer bevel width.
func (mw *MagickWand) FrameImage(matteColor *PixelWand, width, height uint, innerBevel, outerBevel int) error {
	C.MagickFrameImage(mw.wand, matteColor.pixel, C.size_t(width), C.size_t(height), C.ssize_t(innerBevel), C.ssize_t(outerBevel))
	return mw.GetLastError()
}

// Applys an arithmetic, relational, or logical expression to an image. Use these operators to lighten
// or darken an image, to increase or decrease contrast in an image, or to produce the "negative" of an image.
func (mw *MagickWand) FunctionImage(function MagickFunction, args []float64) error {
	C.MagickFunctionImage(mw.wand, C.MagickFunction(function), C.size_t(len(args)), (*C.double)(&args[0]))
	return mw.GetLastError()
}

// Applys an arithmetic, relational, or logical expression to an image's channel. Use these operators to lighten
// or darken an image, to increase or decrease contrast in an image, or to produce the "negative" of an image.
func (mw *MagickWand) FunctionImageChannel(channel ChannelType, function MagickFunction, number_arguments uint, arguments []float64) error {
	C.MagickFunctionImage(mw.wand, C.ChannelType(channel), C.MagickFunction(function), C.size_t(len(args)), (*C.double)(&args[0]))
	return mw.GetLastError()
}

// Evaluate expression for each pixel in the image.
func (mw *MagickWand) FxImage(expression string) *MagickWand {
	csexpression := C.CString(expression)
	C.free(unsafe.Pointer(csexpression))
	return &MagickWand{C.FxImage(csexpression)}
}

// Evaluate expression for each pixel in the image's channel
func (mw *MagickWand) FxImageChannel(channel ChannelType, expression string) *MagickWand {
	csexpression := C.CString(expression)
	C.free(unsafe.Pointer(csexpression))
	return &MagickWand{C.FxImageChannel(C.ChannelType(channel), csexpression)}
}

// Gamma-corrects an image. The same image viewed on different devices will have perceptual
// differences in the way the image's intensities are represented on the screen. Specify
// individual gamma levels for the red, green, and blue channels, or adjust all three with
// the gamma parameter. Values typically range from 0.8 to 2.3.
// You can also reduce the influence of a particular channel with a gamma value of 0.
func (mw *MagickWand) GammaImage(gamma float64) error {
	C.MagickGammaImage(mw.wand, C.double(gamma))
	return mw.GetLastError()
}

// Gamma-corrects an image's channel. The same image viewed on different devices will have perceptual
// differences in the way the image's intensities are represented on the screen. Specify
// individual gamma levels for the red, green, and blue channels, or adjust all three with
// the gamma parameter. Values typically range from 0.8 to 2.3.
// You can also reduce the influence of a particular channel with a gamma value of 0.
func (mw *MagickWand) GammaImageChannel(channel ChannelType, gamma float64) error {
	C.MagickGammaImageChannel(mw.wand, C.ChannelType(channel), C.double(gamma))
	return mw.GetLastError()
}

// Blurs an image. We convolve the image with a Gaussian operator of the given radius and standard
// deviation (sigma). For reasonable results, the radius should be larger than sigma. Use a radius
// of 0 and GaussianBlurImage() selects a suitable radius for you.
// radius: the radius of the Gaussian, in pixels, not counting the center pixel.
// sigma: the standard deviation of the Gaussian, in pixels.
func (mw *MagickWand) GaussianBlurImage(radius, sigma float64) error {
	C.MagickGaussianBlurImage(mw.wand, C.double(radius), C.double(sigma))
	return mw.GetLastError()
}

// Blurs an image's channel. We convolve the image with a Gaussian operator of the given radius and standard
// deviation (sigma). For reasonable results, the radius should be larger than sigma. Use a radius
// of 0 and GaussianBlurImage() selects a suitable radius for you.
// radius: the radius of the Gaussian, in pixels, not counting the center pixel.
// sigma: the standard deviation of the Gaussian, in pixels.
func (mw *MagickWand) GaussianBlurImageChannel(channel ChannelType, radius, sigma float64) error {
	C.MagickGaussianBlurImageChannel(mw.wand, C.ChannelType(channel), C.double(radius), C.double(sigma))
	return mw.GetLastError()
}

// Gets the image at the current image index.
func (mw *MagickWand) GetImage() *MagickWand {
	return &MagickWand{C.MagickGetImage(mw.wand)}
}

// Returns false if the image alpha channel is not activated. That is, the image is RGB rather than RGBA or CMYK rather than CMYKA.
func (mw *MagickWand) GetImageAlphaChannel() bool {
	return 1 == C.MagickGetImageAlphaChannel(mw.wand)
}

// Gets the image clip mask at the current image index.
func (mw *MagickWand) GetImageClipMask() *MagickWand {
	return &MagickWand{C.MagickGetImageClipMask(mw.wand)}
}

// Returns the image background color.
func (mw *MagickWand) GetImageBackgroundColor() (bgColor *PixelWand, err error) {
	cbgcolor := C.PixelWand{}
	C.MagickGetImageBackgroundColor(mw.wand, &cbgcolor)
	return &PixelWand{cbgcolor}, mw.GetLastError()
}

// Implements direct to memory image formats. It returns the image as a blob (a formatted "file" in memory) and its length,
// starting from the current position in the image sequence. Use SetImageFormat() to set the format to write to the blob (GIF, JPEG, PNG, etc.).
// Utilize ResetIterator() to ensure the write is from the beginning of the image sequence.
func (mw *MagickWand) GetImageBlob() []byte {
	clen := C.size_t
	csblob := C.MagickGetImageBlob(mw.wand, &clen)
	defer mw.relinquishMemory(csblob)
	return C.GoBytes(unsafe.Pointer(csblob), clen)
}

// Implements direct to memory image formats. It returns the image sequence as a blob and its length. The format of the image
// determines the format of the returned blob (GIF, JPEG, PNG, etc.). To return a different image format, use SetImageFormat().
// Note, some image formats do not permit multiple images to the same image stream (e.g. JPEG). in this instance, just the first
// image of the sequence is returned as a blob.
func (mw *MagickWand) GetImagesBlob(size_t *length) []byte {
	clen := C.size_t
	csblob := C.MagickGetImagesBlob(mw.wand, &clen)
	defer mw.relinquishMemory(csblob)
	return C.GoBytes(unsafe.Pointer(csblob), clen)
}

// Returns the chromaticy blue primary point for the image.
// x: the chromaticity blue primary x-point.
// y: the chromaticity blue primary y-point.
func (mw *MagickWand) GetImageBluePrimary() (x, y float64, err error) {
	dx, dy := C.double(0.0)
	C.MagickGetImageBluePrimary(mw.wand, &dx, &dy)
	return float64(dx), float64(dy), mw.GetLastError()
}

// Returns the image border color.
func (mw *MagickWand) GetImageBorderColor() (borderColor *PixelWand, err error) {
	cbc := C.PixelWand{}
	C.MagickGetImageBorderColor(mw.wand, &cbc)
	return PixelWand{cbc}, mw.GetLastError()
}

// Gets the depth for one or more image channels.
func (mw *MagickWand) GetImageChannelDepth(channel ChannelType) uint {
	return uint(MagickGetImageChannelDepth(mw.wand, C.ChannelType(channel)))
}

// Compares one or more image channels of an image to a reconstructed image and returns the specified distortion metrics
func (mw *MagickWand) GetImageChannelDistortion(reference *MagickWand, channel ChannelType, metric MetricType) (distortion float64, err error) {
	cdouble := C.double(0.0)
	C.MagickGetImageChannelDistortion(reference.wand, C.ChannelType(channel), C.MetricType(metric), &cdouble)
	return float64(cdouble), mw.GetLastError()
}

// Compares one or more image channels of an image to a reconstructed image and returns the specified distortion metrics.
func (mw *MagickWand) GetImageChannelDistortions(reference *MagickWand, metric MetricType) []float64 {
	cptrdbls := C.MagickGetImageChannelDistortions(mw.wand, reference.wand, C.MetricType(metric))
	defer C.relinquishMemory(cptrdbls)
	metrics := make([]float64, 0)
	q := uintptr(unsafe.Pointer(cptrdbls))
	for {
		p = (*C.double)(unsafe.Pointer(q))
		if *p == nil {
			break
		}
		metrics = append(metrics, float64(*p))
		q += unsafe.Sizeof(q)
	}
	return strings
}

// Returns features for each channel in the image in each of four directions (horizontal, vertical, left and right diagonals)
// for the specified distance. The features include the angular second moment, contrast, correlation, sum of squares: variance,
// inverse difference moment, sum average, sum varience, sum entropy, entropy, difference variance, difference entropy, information
// measures of correlation 1, information measures of correlation 2, and maximum correlation coefficient. You can access the red
// channel contrast, for example, like this:
//   channelFeatures = GetImageChannelFeatures(1);
//   contrast = channelFeatures[RedChannel].Contrast[0];
func (mw *MagickWand) GetImageChannelFeatures(distance uint) map[ChannelType]ChannelFeatures {
	ccf := C.MagickGetImageChannelFeatures(mw.wand, C.size_t(distance))
	defer mw.relinquishMemory(unsafe.Pointer(ccf))
	return []ChannelFeatures{}
	// TODO ChannelFeatures
}

// Gets the kurtosis and skewness of one or more image channels.
func (mw *MagickWand) GetImageChannelKurtosis(channel ChannelType) (kurtosis, skewness float64, err error) {
	ck, cs := C.double(0.0)
	C.MagickGetImageChannelKurtosis(mw.wand, C.ChannelType(channel), &ck, &cs)
	return float64(ck), float64(cs), mw.GetLastError()
}

// Gets the mean and standard deviation of one or more image channels.
func (mw *MagickWand) GetImageChannelMean(channel ChannelType) (mean, stdev float64, err error) {
	cm, cs := C.double(0.0)
	C.MagickGetImageChannelMean(mw.wand, C.ChannelType(channel), &cm, &cs)
	return float64(cm), float64(cs), mw.GetLastError()
}

// Gets the range for one or more image channels.
func (mw *MagickWand) GetImageChannelRange(channel ChannelType) (min, max float64, err error) {
	cmin, cmax := C.double(0.0)
	C.MagickGetImageChannelRange(mw.wand, C.ChannelType(channel), &cmin, &cmax)
	return float64(cmin), float64(cmax), mw.GetLastError()
}

// Returns statistics for each channel in the image. The statistics include the channel depth, its minima and maxima,
// the mean, the standard deviation, the kurtosis and the skewness. You can access the red channel mean, for example,
// like this:
//    channelStatistics = wand.GetImageChannelStatistics()
//    redMean = channelStatistics[RedChannel].mean
func (mw *MagickWand) GetImageChannelStatistics() map[ChannelType]ChannelStatistics {
	// TODO ChannelStatistics
}

// Returns the color of the specified colormap index.
func (mw *MagickWand) GetImageColormapColor(index uint) (color *PixelWand, err error) {
	cpw := C.PixelWand{}
	C.MagickGetImageColormapColor(mw.wand, C.size_t(index), &cpw)
	return &PixelWand{cpw}, mw.GetLastError()
}

// Gets the number of unique colors in the image.
func (mw *MagickWand) GetImageColors() uint {
	return uint(C.MagickGetImageColors(mw.wand))
}

// Gets the image colorspace.
func (mw *MagickWand) GetImageColorspace() ColorspaceType {
	return ColorspaceType(C.MagickGetImageColorspace(mw.wand))
}

// Returns the composite operator associated with the image.
func (mw *MagickWand) GetImageCompose() CompositeOperator {
	return CompositeOperator(C.MagickGetImageCompose(mw.wand))
}

// Gets the image compression.
func (mw *MagickWand) GetImageCompression() CompressionType {
	return CompressionType(C.MagickGetImageCompression(mw.wand))
}

// Gets the image compression quality.
func (mw *MagickWand) GetImageCompressionQuality() uint {
	return uint(C.MagickGetImageCompressionQuality(mw.wand))
}

// Gets the image delay.
func (mw *MagickWand) GetImageDelay() uint {
	return uint(C.MagickGetImageDelay(mw.wand))
}

// Gets the image depth.
func (mw *MagickWand) GetImageDepth() uint {
	return uint(C.MagickGetImageDepth(mw.wand))
}

// Compares an image to a reconstructed image and returns the specified distortion metric.
func (mw *MagickWand) GetImageDistortion(reference *MagickWand, metric MetricType) (distortion float64, err error) {
	cd := C.double(0.0)
	C.MagickGetImageDistortion(mw.wand, reference.wand, C.MetricType(metric), &cd)
	return float64(cd), mw.GetLastError()
}

// Gets the image disposal method.
func (mw *MagickWand) GetImageDispose() DisposeType {
	// TODO DisposeType
}

//MagickGetImageEndian() gets the image endian.
func (mw *MagickWand) GetImageEndian() EndianType {
	// TODO EndianType
}

// Returns the filename of a particular image in a sequence.
func (mw *MagickWand) GetImageFilename() string {
	return C.GoString(C.MagickGetImageFilename(mw.wand))
}

// Returns the format of a particular image in a sequence.
func (mw *MagickWand) GetImageFormat() string {
	return C.GoString(C.MagickGetImageFormat(mw.wand))
}

// Gets the image fuzz.
func (mw *MagickWand) GetImageFuzz() float64 {
	return float64(C.MagickGetImageFuzz(mw.wand))
}

// Gets the image gamma.
func (mw *MagickWand) GetImageGamma() float64 {
	return float64(C.MagickGetImageGamma(mw.wand))
}

// Gets the image gravity.
func (mw *MagickWand) GetImageGravity() GravityType {
	// TODO GravityType
}

// Returns the chromaticy green primary point.
// x: the chromaticity green primary x-point.
// y: the chromaticity green primary y-point.
func (mw *MagickWand) GetImageGreenPrimary(double *x, double *y) error {
}

// Returns the image height.
func (mw *MagickWand) GetImageHeight() uint {
	return uint(C.MagickGetImageHeight(mw.wand))
}

// Returns the image histogram as an array of PixelWand wands.
// numberColors: the number of unique colors in the image and the number of pixel wands returned.
func (mw *MagickWand) GetImageHistogram() (numberColors uint, pws []PixelWand) {
	// TODO
}

// Gets the image interlace scheme.
func (mw *MagickWand) GetImageInterlaceScheme() InterlaceType {
	// TODO InterlaceType
}

// Returns the interpolation method for the sepcified image.
func (mw *MagickWand) GetImageInterpolateMethod() InterpolatePixelMethod {
	// TODO InterpolatePixelMethod
}

// Gets the image iterations.
func (mw *MagickWand) GetImageIterations() uint {
	return uint(C.MagickGetImageIterations(mw.wand))
}

// Returns the image length in bytes.
func (mw *MagickWand) GetImageLength(MagickSizeType *length) error {
	// TODO MagickSizeType
}

// Returns the image matte color.
func (mw *MagickWand) GetImagematteColor() (matteColor *PixelWand, err error) {
	var cptrpw (*C.PixelWand)
	C.MagickGetImagematteColor(mw.wand, cptrpw)
	return &PixelWand{cptrpw}, mw.GetLastError()
}

// Returns the image orientation.
func (mw *MagickWand) GetImageOrientation() OrientationType {
	// TODO OrientationType
}

// Returns the page geometry associated with the image.
// w, h: the page width and height
// x, h: the page x-offset and y-offset.
func (mw *MagickWand) GetImagePage() (w, h uint, x, y int, err error) {
	cw, ch := C.size_t(0)
	cx, cy := C.ssize_t(0)
	C.MagickGetImagePage(mw.wand, &cw, &ch, &cx, &cy)
	return uint(cw), uint(ch), int(cx), int(cy), mw.GetLastError()
}

// Returns the color of the specified pixel.
func (mw *MagickWand) GetImagePixelColor(x, y int) (color *PixelWand, err error) {
	cpw := C.PixelWand()
	C.MagickGetImagePixelColor(C.ssize_t(x), C.ssize_t(y), &cpw)
	return &PixelWand{cpw}, mw.GetLastError()
}

// Returns the chromaticy red primary point.
// x, y: the chromaticity red primary x/y-point.
func (mw *MagickWand) GetImageRedPrimary() (x, y float64, err error) {
	var cdx, cdy C.double
	C.MagickGetImageRedPrimary(mw.wand, &cdx, &cdy)
	return float64(cdx), float64(cdy), mw.GetLastError()
}

// Extracts a region of the image and returns it as a a new wand.
func (mw *MagickWand) GetImageRegion(width uint, height uint, x int, y int) *MagickWand {
	return &MagickWand(C.MagickGetImageRegion(mw.wand, C.size_t(width), C.size_t(height), C.ssize_t(x), C.ssize_t(y)))
}

// Gets the image rendering intent.
func (mw *MagickWand) GetImageRenderingIntent() RenderingIntent {
	// TODO RenderingIntent
}

// Gets the image X and Y resolution.
func (mw *MagickWand) GetImageResolution() (x, y float64, err error) {
	var dx, dy C.double
	C.MagickGetImageResolution(mw.wand, &dx, &dy)
	return float64(dx), float64(dy), mw.GetLastError()
}

// Gets the image scene.
func (mw *MagickWand) GetImageScene() uint {
	return uint(C.MagickGetImageScene(mw.wand))
}

// Generates an SHA-256 message digest for the image pixel stream.
func (mw *MagickWand) GetImageSignature() string {
	return C.GoString(C.MagickGetImageSignature(mw.wand))
}

// Gets the image ticks-per-second.
func (mw *MagickWand) GetImageTicksPerSecond() uint {
	return uint(C.MagickGetImageTicksPerSecond(mw.wand))
}

// Gets the potential image type
// To ensure the image type matches its potential, use SetImageType(): wand.SetImageType(wand.GetImageType())
func (mw *MagickWand) GetImageType() ImageType {
	// TODO ImageType
}

// Gets the image units of resolution.
func (mw *MagickWand) GetImageUnits() ResolutionType {
	// TODO ResolutionType
}

// Returns the virtual pixel method for the specified image.
func (mw *MagickWand) GetImageVirtualPixelMethod() VirtualPixelMethod {
	// TODO VirtualPixelMethod
}

// Returns the chromaticy white point.
//x, y: the chromaticity white x/y-point.
func (mw *MagickWand) GetImageWhitePoint() (x, y float64, err error) {
	dx, dy := C.double(0.0)
	C.MagickGetImageWhitePoint(mw.wand, &dx, &dy)
	return float64(dx), float64(dy), mw.GetLastError()
}

// Returns the image width.
func (mw *MagickWand) GetImageWidth() uint {
	return uint(C.MagickGetImageWidth(mw.wand))
}

// Returns the number of images associated with a magick wand.
func (mw *MagickWand) GetNumberImages() uint {
	return uint(C.MagickGetNumberImages(mw.wand))
}

// Gets the image total ink density.
func (mw *MagickWand) GetImageTotalInkDensity() float64 {
	return float64(C.MagickGetImageTotalInkDensity(mw.wand))
}

// Replaces colors in the image from a Hald color lookup table.
// A Hald color lookup table is a 3-dimensional color cube mapped to 2 dimensions.
// Create it with the HALD coder. You can apply any color transformation to the Hald image
// and then use this method to apply the transform to the image.
func (mw *MagickWand) HaldClutImage(hald *MagickWand) error {
	C.MagickHaldClutImage(mw.wand, hald.wand)
	return mw.GetLastError()
}

// Replaces colors in the image's channel from a Hald color lookup table.
// A Hald color lookup table is a 3-dimensional color cube mapped to 2 dimensions.
// Create it with the HALD coder. You can apply any color transformation to the Hald image
// and then use this method to apply the transform to the image.
func (mw *MagickWand) HaldClutImageChannel(channel ChannelType, hald_wand *MagickWand) error {
	C.MagickHaldClutImageChannel(mw.wand, C.ChannelType(channel), hald.wand)
	return mw.GetLastError()
}

// Returns true if the wand has more images when traversing the list in the forward direction
func (mw *MagickWand) HasNextImage() bool {
	return 1 == MagickHasNextImage(mw.wand)
}

// Returns true if the wand has more images when traversing the list in the reverse direction
func (mw *MagickWand) HasPreviousImage() bool {
	return 1 == MagickHasNextImage(mw.wand)
}

// Identifies an image by printing its attributes to the file. Attributes include the image width, height, size, and others.
func (mw *MagickWand) IdentifyImage() string {
	return C.GoString(MagickIdentifyImage(mw.wand))
}

// Creates a new image that is a copy of an existing one with the image pixels "implode" by the specified percentage.
func (mw *MagickWand) ImplodeImage(radius float64) error {
	return &MagickWand{C.MagickImplodeImage(mw.wand, C.double(radius))}
}

// Accepts pixel data and stores it in the image at the location you specify.
// The pixel data can be either char, short int, int, ssize_t, float, or double in the order specified by map.
// Suppose your want to upload the first scanline of a 640x480 image from character data in red-green-blue order:
//   wand.ImportImagePixels(0, 0, 640, 1, "RGB", CharPixel, pixels)
// x, y, columns, rows: These values define the perimeter of a region of pixels you want to define.
// map: This string reflects the expected ordering of the pixel array. It can be any combination or order of R = red,
// G = green, B = blue, A = alpha (0 is transparent), O = opacity (0 is opaque), C = cyan, Y = yellow, M = magenta,
// K = black, I = intensity (for grayscale), P = pad.
// storage: Define the data type of the pixels. Float and double types are expected to be normalized [0..1] otherwise
// [0..QuantumRange]. Choose from these types: CharPixel, ShortPixel, IntegerPixel, LongPixel, FloatPixel, or DoublePixel.
// pixels: This array of values contain the pixel components as defined by map and type. You must preallocate this array
// where the expected length varies depending on the values of width, height, map, and type.
func (mw *MagickWand) ImportImageBytePixels(x, y int, columns, rows uint, pmap string, stype StorageType, pixels []interface{}) error {
	cspmap := C.CString(pmap)
	defer C.free(unsafe.Pointer(cspmap))
	C.MagickImportImagePixels(mw.wand, C.ssize_t(x), C.ssize_t(y), C.size_t(columns), C.size_t(rows), cspmap, stype, unsafe.Pointer(&pixels[0]))
	return mw.GetLastError()
}

// Implements the inverse discrete Fourier transform (DFT) of the image either as a magnitude/phase or real/imaginary image pair.
// magnitudeWand: the magnitude or real wand.
// phaseWand: the phase or imaginary wand.
// magnitude: if true, return as magnitude/phase pair otherwise a real/imaginary image pair.
func (mw *MagickWand) InverseFourierTransformImage(phaseWand *MagickWand, magnitude bool) error {
	cmagnitude := C.MagickBooleanType(0)
	if magnitude {
		cmagnitude = 1
	}
	C.MagickInverseFourierTransformImage(mw.wand, phaseWand.wand, cmagnitude)
	return mw.GetLastError()
}

// Adds a label to your image.
func (mw *MagickWand) LabelImage(label string) error {
	cslabel := C.CString(label)
	defer C.free(unsafe.Pointer(cslabel))
	C.MagickLabelImage(mw.wand, cslabel)
	return mw.GetLastError()
}

// Adjusts the levels of an image by scaling the colors falling between specified white and black points
// to the full available quantum range. The parameters provided represent the black, mid, and white points.
// The black point specifies the darkest color in the image. Colors darker than the black point are set to zero.
// Mid point specifies a gamma correction to apply to the image. White point specifies the lightest color in the
// image. Colors brighter than the white point are set to the maximum quantum value.
func (mw *MagickWand) LevelImage(blackPoint, gamma, whitePoint float64) error {
	C.MagickLevelImage(mw.wand, C.double(blackPoint), C.double(gamma), C.double(whitePoint))
	return mw.GetLastError()
}

// Adjusts the levels of an image's channel by scaling the colors falling between specified white and black points
// to the full available quantum range. The parameters provided represent the black, mid, and white points.
// The black point specifies the darkest color in the image. Colors darker than the black point are set to zero.
// Mid point specifies a gamma correction to apply to the image. White point specifies the lightest color in the
// image. Colors brighter than the white point are set to the maximum quantum value.
func (mw *MagickWand) LevelImageChannel(channel ChannelType, blackPoint, gamma, whitePoint float64) error {
	C.MagickLevelImageChannel(mw.wand, C.ChannelType(channel), C.double(blackPoint), C.double(gamma), C.double(whitePoint))
	return mw.GetLastError()
}

// Stretches with saturation the image intensity.
// You can also reduce the influence of a particular channel with a gamma value of 0.
func (mw *MagickWand) LinearStretchImage(blackPoint, whitePoint float64) error {
	C.MagickLinearStretchImage(mw.wand, C.double(blackPoint), C.double(whitePoint))
	return mw.GetLastError()
}

// Rescales image with seam carving.
// columns, rows: the number of columns and rows in the scaled image.
// deltaX: maximum seam transversal step (0 means straight seams).
// rigidity: introduce a bias for non-straight seams (typically 0).
func (mw *MagickWand) LiquidRescaleImage(columns, rows uint, deltaX, rigidity float64) error {
	C.MagickLiquidRescaleImage(mw.wand, C.size_t(columns), C.size_t(rows), C.double(deltaX), C.double(rigidity))
	return mw.GetLastError()
}

// This is a convenience method that scales an image proportionally to twice its original size.
func (mw *MagickWand) MagnifyImage() error {
	C.MagnifyImage(mw.wand)
	return mw.GetLastError()
}

// Composes all the image layers from the current given image onward to produce a single image of the merged layers.
// The inital canvas's size depends on the given ImageLayerMethod, and is initialized using the first images background color.
// The images are then compositied onto that image in sequence using the given composition that has been assigned to each individual image.
// method: the method of selecting the size of the initial canvas.
//  * MergeLayer: Merge all layers onto a canvas just large enough to hold all the actual images.
//                The virtual canvas of the first image is preserved but otherwise ignored.
//  * FlattenLayer: Use the virtual canvas size of first image. Images which fall outside this canvas is clipped.
//                  This can be used to 'fill out' a given virtual canvas.
//  * MosaicLayer: Start with the virtual canvas of the first image, enlarging left and right edges to contain all images.
//                 Images with negative offsets will be clipped.
func (mw *MagickWand) MergeImageLayers(method ImageLayerMethod) *MagickWand {
	return &MagickWand{C.MagickMergeImageLayers(C.ImageLayerMethod(method))}
}

// This is a convenience method that scales an image proportionally to one-half its original size
func (mw *MagickWand) MinifyImage() error {
	C.MinifyImage(mw.wand)
	return mw.GetLastError()
}

// Lets you control the brightness, saturation, and hue of an image. Hue is the percentage of absolute rotation from the current position.
// For example 50 results in a counter-clockwise rotation of 90 degrees, 150 results in a clockwise rotation of 90 degrees, with 0 and 200
// both resulting in a rotation of 180 degrees.
// To increase the color brightness by 20 and decrease the color saturation by 10 and leave the hue unchanged, use: 120, 90, 100.
//  * **brightness**: the percent change in brighness.
//  * **saturation**: the percent change in saturation.
//  * **hue**: the percent change in hue.
func (mw *MagickWand) ModulateImage(brightness, saturation, hue float64) error {
	C.MagickModulateImage(mw.wand, C.double(brightness), C.double(saturation), C.double(hue))
	return mw.GetLastError()
}

// Creates a composite image by combining several separate images.
// The images are tiled on the composite image with the name of the image optionally appearing just below the individual tile.
// dw: the drawing wand. The font name, size, and color are obtained from this wand.
// tileGeo: the number of tiles per row and page (e.g. 6x4+0+0).
// thumbGeo: Preferred image size and border size of each thumbnail (e.g. 120x120+4+3>).
// mode: Thumbnail framing mode: Frame, Unframe, or Concatenate.
// frame: Surround the image with an ornamental border (e.g. 15x15+3+3). The frame color is that of the thumbnail's matte color.
func (mw *MagickWand) MontageImage(dw *DrawingWand, tileGeo string, thumbGeo string, mode MontageMode, frame string) *MagickWand {
	cstile := C.CString(tileGeo)
	defer C.free(unsafe.Pointer(cstile))
	csthumb := C.CString(thumbGeo)
	defer C.free(unsafe.Pointer(csthumb))
	csframe := C.CString(frame)
	defer C.free(unsafe.Pointer(csframe))
	return &MagickWand{C.MagickMontageImage(mw.wand, dw.draw, cstile, csthumb, C.MontageMode(mode), csframe)}
}

// Method morphs a set of images. Both the image pixels and size are linearly interpolated to give the appearance of a
// meta-morphosis from one image to the next.
// numFrames: the number of in-between images to generate.
func (mw *MagickWand) MorphImages(numFrames uint) *MagickWand {
	return &MagickWand{C.MagickMorphImages(mw.wand, C.size_t(numFrames))}
}

// Applies a user supplied kernel to the image according to the given mophology method.
// channel: the image channel(s).
// method: the morphology method to be applied.
// iterations: apply the operation this many times (or no change). A value of -1 means loop until no change found. How this is applied may depend on the morphology method. Typically this is a value of 1.
// kernel: An array of doubles representing the morphology kernel.
func (mw *MagickWand) MorphologyImage(MorphologyMethod method, iterations int, KernelInfo *kernel) error {
	// TODO MorphologyMethod
}

// Applies a user supplied kernel to the image according to the given mophology method.
// channel : the image channel(s).
// method : the morphology method to be applied.
// iterations : apply the operation this many times (or no change). A value of -1 means loop until no change found. How this is applied may depend on the morphology method. Typically this is a value of 1.
// kernel : An array of doubles representing the morphology kernel.
func (mw *MagickWand) MorphologyImageChannel(ChannelType channel, MorphologyMethod method, iterations int, KernelInfo *kernel) error {
	// TODO MorphologyMethod
}

// Simulates motion blur. We convolve the image with a Gaussian operator of the given radius and standard deviation (sigma).
// For reasonable results, radius should be larger than sigma. Use a radius of 0 and MotionBlurImage() selects a suitable
// radius for you. Angle gives the angle of the blurring motion.
func (mw *MagickWand) MotionBlurImage(radius, sigma, angle float64) error {
	C.MagickMotionBlurImage(mw.wand, C.double(radius), C.double(sigma), C.double(angle))
	return mw.GetLastError()
}

// Simulates motion blur. We convolve the image with a Gaussian operator of the given radius and standard deviation (sigma).
// For reasonable results, radius should be larger than sigma. Use a radius of 0 and MotionBlurImage() selects a suitable
// radius for you. Angle gives the angle of the blurring motion.
// radius: the radius of the Gaussian, in pixels, not counting the center pixel.
// sigma: the standard deviation of the Gaussian, in pixels.
// angle: apply the effect along this angle.
func (mw *MagickWand) MotionBlurImageChannel(channel ChannelType, radius, sigma, angle float64) error {
	C.MagickMotionBlurImageChannel(mw.wand, C.ChannelType(channel), C.double(radius), C.double(sigma), C.double(angle))
	return mw.GetLastError()
}

// Negates the colors in the reference image. The Grayscale option means that only grayscale values within the image are negated.
//You can also reduce the influence of a particular channel with a gamma value of 0.
func (mw *MagickWand) NegateImage(gray bool) error {
}
func (mw *MagickWand) NegateImageChannel(channel ChannelType, gray bool) error {
}

//channel
//the image channel(s).
//gray
//If MagickTrue, only negate grayscale pixels within the image.
//MagickNewImage
//MagickNewImage() adds a blank image canvas of the specified size and background color to the wand.
func (mw *MagickWand) NewImage(columns uint, rows uint, background *PixelWand) error {
}

//width
//the image width.
//height
//the image height.
//background
//the image color.
//MagickNextImage
//MagickNextImage() sets the next image in the wand as the current image.
//It is typically used after MagickResetIterator(), after which its first use will set the first image as the current image (unless the wand is empty).
//It will return MagickFalse when no more images are left to be returned which happens when the wand is empty, or the current image is the last image.
//When the above condition (end of image list) is reached, the iterator is automaticall set so that you can start using MagickPreviousImage() to again iterate over the images in the reverse direction, starting with the last image (again). You can jump to this condition immeditally using MagickSetLastIterator().
func (mw *MagickWand) NextImage() error {
}

//MagickNormalizeImage
//MagickNormalizeImage() enhances the contrast of a color image by adjusting the pixels color to span the entire range of colors available
//You can also reduce the influence of a particular channel with a gamma value of 0.
func (mw *MagickWand) NormalizeImage() error {
}
func (mw *MagickWand) NormalizeImageChannel(channel ChannelType) error {
}

//channel
//the image channel(s).
//MagickOilPaintImage
//MagickOilPaintImage() applies a special effect filter that simulates an oil painting. Each pixel is replaced by the most frequent color occurring in a circular region defined by radius.
func (mw *MagickWand) OilPaintImage(radius double) error {
}

//radius
//the radius of the circular neighborhood.
//MagickOpaquePaintImage
//MagickOpaquePaintImage() changes any pixel that matches color with the color defined by fill.
func (mw *MagickWand) OpaquePaintImage(target *PixelWand, fill *PixelWand, fuzz double, invert bool) error {
}

func (mw *MagickWand) OpaquePaintImageChannel(channel ChannelType, target *PixelWand, fill *PixelWand, fuzz double, invert bool) error {
}

//channel
//the channel(s).
//target
//Change this target color to the fill color within the image.
//fill
//the fill pixel wand.
//fuzz
//By default target must match a particular pixel color exactly. However, in many cases two colors may differ by a small amount. The fuzz member of image defines how much tolerance is acceptable to consider two colors as the same. For example, set fuzz to 10 and the color red at intensities of 100 and 102 respectively are now interpreted as the same color for the purposes of the floodfill.
//invert
//paint any pixel that does not match the target color.
//MagickOptimizeImageLayers
//MagickOptimizeImageLayers() compares each image the GIF disposed forms of the previous image in the sequence. From this it attempts to select the smallest cropped image to replace each frame, while preserving the results of the animation.
func (mw *MagickWand) OptimizeImageLayers() *MagickWand {
}

//MagickOptimizeImageTransparency
//MagickOptimizeImageTransparency() takes a frame optimized GIF animation, and compares the overlayed pixels against the disposal image resulting from all the previous frames in the animation. Any pixel that does not change the disposal image (and thus does not effect the outcome of an overlay) is made transparent.
//WARNING: This modifies the current images directly, rather than generate a new image sequence.
func (mw *MagickWand) OptimizeImageTransparency() error {
}

//MagickOrderedPosterizeImage
//MagickOrderedPosterizeImage() performs an ordered dither based on a number of pre-defined dithering threshold maps, but over multiple intensity levels, which can be different for different channels, according to the input arguments.
func (mw *MagickWand) OrderedPosterizeImage(threshold_map string) error {
}
func (mw *MagickWand) OrderedPosterizeImageChannel(channel ChannelType, threshold_map string) error {
}

//image
//the image.
//channel
//the channel or channels to be thresholded.
//threshold_map
//A string containing the name of the threshold dither map to use, followed by zero or more numbers representing the number of color levels tho dither between.
//Any level number less than 2 is equivalent to 2, and means only binary dithering will be applied to each color channel.
//No numbers also means a 2 level (bitmap) dither will be applied to all channels, while a single number is the number of levels applied to each channel in sequence. More numbers will be applied in turn to each of the color channels.
//For example: "o3x3,6" generates a 6 level posterization of the image with a ordered 3x3 diffused pixel dither being applied between each level. While checker,8,8,4 will produce a 332 colormaped image with only a single checkerboard hash pattern (50 grey) between each color level, to basically double the number of color levels with a bare minimim of dithering.
//MagickPingImage
//MagickPingImage() is like MagickReadImage() except the only valid information returned is the image width, height, size, and format. It is designed to efficiently obtain this information from a file without reading the entire image sequence into memory.
func (mw *MagickWand) PingImage(filename string) error {
}

//filename
//the image filename.
//MagickPingImageBlob
//MagickPingImageBlob() pings an image or image sequence from a blob.
func (mw *MagickWand) PingImageBlob(blob *void, length uint) error {
}

//blob
//the blob.
//length
//the blob length.
//MagickPingImageFile
//MagickPingImageFile() pings an image or image sequence from an open file descriptor.
func (mw *MagickWand) ReadImage(filename string) error {
}

//file
//the file descriptor.
//MagickPolaroidImage
//MagickPolaroidImage() simulates a Polaroid picture.
func (mw *MagickWand) PolaroidImage(drawing_wand *DrawingWand, angle double) error {
}

//drawing_wand
//the draw wand.
//angle
//Apply the effect along this angle.
//MagickPosterizeImage
//MagickPosterizeImage() reduces the image to a limited number of color level.
func (mw *MagickWand) PosterizeImage(levels uint, dither bool) error {
}

//levels
//Number of color levels allowed in each channel. Very low values (2, 3, or 4) have the most visible effect.
//dither
//Set this integer value to something other than zero to dither the mapped image.
//MagickPreviewImages
//MagickPreviewImages() tiles 9 thumbnails of the specified image with an image processing operation applied at varying strengths. This helpful to quickly pin-point an appropriate parameter for an image processing operation.
func (mw *MagickWand) PreviewImages(preview PreviewType) *MagickWand {
}

//preview
//the preview type.
//MagickPreviousImage
//MagickPreviousImage() sets the previous image in the wand as the current image.
//It is typically used after MagickSetLastIterator(), after which its first use will set the last image as the current image (unless the wand is empty).
//It will return MagickFalse when no more images are left to be returned which happens when the wand is empty, or the current image is the first image. At that point the iterator is than reset to again process images in the forward direction, again starting with the first image in list. Images added at this point are prepended.
//Also at that point any images added to the wand using MagickAddImages() or MagickReadImages() will be prepended before the first image. In this sense the condition is not quite exactly the same as MagickResetIterator().
func (mw *MagickWand) ReadImage(filename string) error {
}

//MagickQuantizeImage
//MagickQuantizeImage() analyzes the colors within a reference image and chooses a fixed number of colors to represent the image. The goal of the algorithm is to minimize the color difference between the input and output image while minimizing the processing time.
func (mw *MagickWand) QuantizeImage(number_colors uint, colorspace ColorspaceType, treedepth uint, dither bool, measure_error bool) error {
}

//number_colors
//the number of colors.
//colorspace
//Perform color reduction in this colorspace, typically RGBColorspace.
//treedepth
//Normally, this integer value is zero or one. A zero or one tells Quantize to choose a optimal tree depth of Log4(number_colors). A tree of this depth generally allows the best representation of the reference image with the least amount of memory and the fastest computational speed. In some cases, such as an image with low color dispersion (a few number of colors), a value other than Log4(number_colors) is required. To expand the color tree completely, use a value of 8.
//dither
//A value other than zero distributes the difference between an original image and the corresponding color reduced image to neighboring pixels along a Hilbert curve.
//measure_error
//A value other than zero measures the difference between the original and quantized images. This difference is the total quantization error. The error is computed by summing over all pixels in an image the distance squared in RGB space between each reference pixel value and its quantized value.
//MagickQuantizeImages
//MagickQuantizeImages() analyzes the colors within a sequence of images and chooses a fixed number of colors to represent the image. The goal of the algorithm is to minimize the color difference between the input and output image while minimizing the processing time.
func (mw *MagickWand) QuantizeImages(number_colors uint, colorspace ColorspaceType, treedepth uint, dither bool, measure_error bool) error {
}

//number_colors
//the number of colors.
//colorspace
//Perform color reduction in this colorspace, typically RGBColorspace.
//treedepth
//Normally, this integer value is zero or one. A zero or one tells Quantize to choose a optimal tree depth of Log4(number_colors). A tree of this depth generally allows the best representation of the reference image with the least amount of memory and the fastest computational speed. In some cases, such as an image with low color dispersion (a few number of colors), a value other than Log4(number_colors) is required. To expand the color tree completely, use a value of 8.
//dither
//A value other than zero distributes the difference between an original image and the corresponding color reduced algorithm to neighboring pixels along a Hilbert curve.
//measure_error
//A value other than zero measures the difference between the original and quantized images. This difference is the total quantization error. The error is computed by summing over all pixels in an image the distance squared in RGB space between each reference pixel value and its quantized value.
//MagickRadialBlurImage
//MagickRadialBlurImage() radial blurs an image.
func (mw *MagickWand) RadialBlurImage(angle double) error {
}
func (mw *MagickWand) RadialBlurImageChannel(channel ChannelType, angle double) error {
}

//channel
//the image channel(s).
//angle
//the angle of the blur in degrees.
//MagickRaiseImage
//MagickRaiseImage() creates a simulated three-dimensional button-like effect by lightening and darkening the edges of the image. Members width and height of raise_info define the width of the vertical and horizontal edge of the effect.
func (mw *MagickWand) RaiseImage(width uint, height uint, x int, y int, raise bool) error {
}

//width,height,x,y
//Define the dimensions of the area to raise.
//raise
//A value other than zero creates a 3-D raise effect, otherwise it has a lowered effect.
//MagickRandomThresholdImage
//MagickRandomThresholdImage() changes the value of individual pixels based on the intensity of each pixel compared to threshold. The result is a high-contrast, two color image.
func (mw *MagickWand) RandomThresholdImage(low double, high double) error {
}
func (mw *MagickWand) RandomThresholdImageChannel(channel ChannelType, low double, high double) error {
}

//channel
//the image channel(s).
//low,high
//Specify the high and low thresholds. These values range from 0 to QuantumRange.
//MagickReadImage
//MagickReadImage() reads an image or image sequence. The images are inserted at the current image pointer position. Use MagickSetFirstIterator(), MagickSetLastIterator, or MagickSetImageIndex() to specify the current image pointer position at the beginning of the image list, the end, or anywhere in-between respectively.
func (mw *MagickWand) ReadImage(filename string) error {
}

//filename
//the image filename.
//MagickReadImageBlob
//MagickReadImageBlob() reads an image or image sequence from a blob.
func (mw *MagickWand) ReadImageBlob(blob *void, length uint) error {
}

//blob
//the blob.
//length
//the blob length.
//MagickReadImageFile
//MagickReadImageFile() reads an image or image sequence from an open file descriptor.
func (mw *MagickWand) ReadImageFile(FILE *file) error {
}

//file
//the file descriptor.
//MagickRemapImage
//MagickRemapImage() replaces the colors of an image with the closest color from a reference image.
func (mw *MagickWand) RemapImage(remap_wand *MagickWand, method DitherMethod) error {
}

//affinity
//the affinity wand.
//method
//choose from these dither methods: NoDitherMethod, RiemersmaDitherMethod, or FloydSteinbergDitherMethod.
//MagickRemoveImage
//MagickRemoveImage() removes an image from the image list.
func (mw *MagickWand) RemoveImage() error {
}

//insert
//the splice wand.
//MagickResampleImage
//MagickResampleImage() resample image to desired resolution.
//Bessel Blackman Box Catrom Cubic Gaussian Hanning Hermite Lanczos Mitchell Point Quandratic Sinc Triangle
//Most of the filters are FIR (finite impulse response), however, Bessel, Gaussian, and Sinc are IIR (infinite impulse response). Bessel and Sinc are windowed (brought down to zero) with the Blackman filter.
func (mw *MagickWand) ResampleImage(x_resolution double, y_resolution double, filter FilterTypes, blur double) error {
}

//x_resolution
//the new image x resolution.
//y_resolution
//the new image y resolution.
//filter
//Image filter to use.
//blur
//the blur factor where > 1 is blurry, < 1 is sharp.
//MagickResetImagePage
//MagickResetImagePage() resets the Wand page canvas and position.
func (mw *MagickWand) ResetImagePage(page string) error {
}

//page
//the relative page specification.
//MagickResizeImage
//MagickResizeImage() scales an image to the desired dimensions with one of these filters:
//Bessel   Blackman   Box
//Catrom   CubicGaussian
//Hanning  Hermite    Lanczos
//Mitchell PointQuandratic
//Sinc     Triangle
//Most of the filters are FIR (finite impulse response), however, Bessel, Gaussian, and Sinc are IIR (infinite impulse response). Bessel and Sinc are windowed (brought down to zero) with the Blackman filter.
func (mw *MagickWand) ResizeImage(columns uint, rows uint, filter FilterTypes, blur double) error {
}

//columns
//the number of columns in the scaled image.
//rows
//the number of rows in the scaled image.
//filter
//Image filter to use.
//blur
//the blur factor where > 1 is blurry, < 1 is sharp.
//MagickRollImage
//MagickRollImage() offsets an image as defined by x and y.
func (mw *MagickWand) RollImage(x int, y uint) error {
}

//x
//the x offset.
//y
//the y offset.
//MagickRotateImage
//MagickRotateImage() rotates an image the specified number of degrees. Empty triangles left over from rotating the image are filled with the background color.
func (mw *MagickWand) RotateImage(background *PixelWand, degrees double) error {
}

//background
//the background pixel wand.
//degrees
//the number of degrees to rotate the image.
//MagickSampleImage
//MagickSampleImage() scales an image to the desired dimensions with pixel sampling. Unlike other scaling methods, this method does not introduce any additional color into the scaled image.
func (mw *MagickWand) SampleImage(columns uint, rows uint) error {
}

//columns
//the number of columns in the scaled image.
//rows
//the number of rows in the scaled image.
//MagickScaleImage
//MagickScaleImage() scales the size of an image to the given dimensions.
func (mw *MagickWand) ScaleImage(columns uint, rows uint) error {
}

//columns
//the number of columns in the scaled image.
//rows
//the number of rows in the scaled image.
//MagickSegmentImage
//MagickSegmentImage() segments an image by analyzing the histograms of the color components and identifying units that are homogeneous with the fuzzy C-means technique.
func (mw *MagickWand) SegmentImage(colorspace ColorspaceType, invert bool, cluster_threshold double, smooth_threshold double) error {
}

//wand
//the wand.
//colorspace
//the image colorspace.
//verbose
//Set to MagickTrue to print detailed information about the identified classes.
//cluster_threshold
//This represents the minimum number of pixels contained in a hexahedra before it can be considered valid (expressed as a percentage).
//smooth_threshold
//the smoothing threshold eliminates noise in the second derivative of the histogram. As the value is increased, you can expect a smoother second derivative.
//MagickSelectiveBlurImage
//MagickSelectiveBlurImage() selectively blur an image within a contrast threshold. It is similar to the unsharpen mask that sharpens everything with contrast above a certain threshold.
func (mw *MagickWand) SelectiveBlurImage(radius double, sigma double, threshold double) error {
}
func (mw *MagickWand) SelectiveBlurImageChannel(channel ChannelType, radius double, sigma double, threshold double) error {
}

//channel
//the image channel(s).
//radius
//the radius of the gaussian, in pixels, not counting the center pixel.
//sigma
//the standard deviation of the gaussian, in pixels.
//threshold
//only pixels within this contrast threshold are included in the blur operation.
//MagickSeparateImageChannel
//MagickSeparateImageChannel() separates a channel from the image and returns a grayscale image. A channel is a particular color component of each pixel in the image.
func (mw *MagickWand) SeparateImageChannel(channel ChannelType) error {
}

//channel
//the image channel(s).
//MagickSepiaToneImage
//MagickSepiaToneImage() applies a special effect to the image, similar to the effect achieved in a photo darkroom by sepia toning. Threshold ranges from 0 to QuantumRange and is a measure of the extent of the sepia toning. A threshold of 80 is a good starting point for a reasonable tone.
func (mw *MagickWand) SepiaToneImage(threshold double) error {
}

//threshold
//Define the extent of the sepia toning.
//MagickSetImage
//MagickSetImage() replaces the last image returned by MagickSetImageIndex(), MagickNextImage(), MagickPreviousImage() with the images from the specified wand.
func (mw *MagickWand) SetImage(set_wand *MagickWand) error {
}

//set_wand
//the set_wand wand.
//MagickSetImageAlphaChannel
//MagickSetImageAlphaChannel() activates, deactivates, resets, or sets the alpha channel.
func (mw *MagickWand) SetImageAlphaChannel(alpha_type AlphaChannelType) error {
}

//alpha_type
//the alpha channel type: ActivateAlphaChannel, DeactivateAlphaChannel, OpaqueAlphaChannel, or SetAlphaChannel.
//MagickSetImageBackgroundColor
//MagickSetImageBackgroundColor() sets the image background color.
func (mw *MagickWand) SetImageBackgroundColor(background *PixelWand) error {
}

//background
//the background pixel wand.
//MagickSetImageBias
//MagickSetImageBias() sets the image bias for any method that convolves an image (e.g. MagickConvolveImage()).
func (mw *MagickWand) SetImageBias(bias double) error {
}

//bias
//the image bias.
//MagickSetImageBluePrimary
//MagickSetImageBluePrimary() sets the image chromaticity blue primary point.
func (mw *MagickWand) SetImageBluePrimary(x double, y double) error {
}

//x
//the blue primary x-point.
//y
//the blue primary y-point.
//MagickSetImageBorderColor
//MagickSetImageBorderColor() sets the image border color.
func (mw *MagickWand) SetImageBorderColor(border *PixelWand) error {
}

//border
//the border pixel wand.
//MagickSetImageChannelDepth
//MagickSetImageChannelDepth() sets the depth of a particular image channel.
func (mw *MagickWand) SetImageChannelDepth(channel ChannelType, depth uint) error {
}

//channel
//the image channel(s).
//depth
//the image depth in bits.
//MagickSetImageClipMask
//MagickSetImageClipMask() sets image clip mask.
func (mw *MagickWand) SetImageClipMask(clip_mask *MagickWand) error {
}

//clip_mask
//the clip_mask wand.
//MagickSetImageColor
//MagickSetImageColor() set the entire wand canvas to the specified color.
func (mw *MagickWand) SetImageColor(color *PixelWand) error {
}

//background
//the image color.
//MagickSetImageColormapColor
//MagickSetImageColormapColor() sets the color of the specified colormap index.
func (mw *MagickWand) SetImageColormapColor(index uint, color *PixelWand) error {
}

//index
//the offset into the image colormap.
//color
//Return the colormap color in this wand.
//MagickSetImageColorspace
//MagickSetImageColorspace() sets the image colorspace.
func (mw *MagickWand) SetImageColorspace(colorspace ColorspaceType) error {
}

//colorspace
//the image colorspace: UndefinedColorspace, RGBColorspace, GRAYColorspace, TransparentColorspace, OHTAColorspace, XYZColorspace, YCbCrColorspace, YCCColorspace, YIQColorspace, YPbPrColorspace, YPbPrColorspace, YUVColorspace, CMYKColorspace, sRGBColorspace, HSLColorspace, or HWBColorspace.
//MagickSetImageCompose
//MagickSetImageCompose() sets the image composite operator, useful for specifying how to composite the image thumbnail when using the MagickMontageImage() method.
func (mw *MagickWand) SetImageCompose(compose CompositeOperator) error {
}

//compose
//the image composite operator.
//MagickSetImageCompression
//MagickSetImageCompression() sets the image compression.
func (mw *MagickWand) SetImageCompression(compression CompressionType) error {
}

//compression
//the image compression type.
//MagickSetImageCompressionQuality
//MagickSetImageCompressionQuality() sets the image compression quality.
func (mw *MagickWand) SetImageCompressionQuality(quality uint) error {
}

//quality
//the image compression tlityype.
//MagickSetImageDelay
//MagickSetImageDelay() sets the image delay.
func (mw *MagickWand) SetImageDelay(delay uint) error {
}

//delay
//the image delay in ticks-per-second units.
//MagickSetImageDepth
//MagickSetImageDepth() sets the image depth.
func (mw *MagickWand) SetImageDepth(depth uint) error {
}

//depth
//the image depth in bits: 8, 16, or 32.
//MagickSetImageDispose
//MagickSetImageDispose() sets the image disposal method.
func (mw *MagickWand) SetImageDispose(dispose DisposeType) error {
}

//dispose
//the image disposeal type.
//MagickSetImageEndian
//MagickSetImageEndian() sets the image endian method.
func (mw *MagickWand) SetImageEndian(endian EndianType) error {
}

//endian
//the image endian type.
//MagickSetImageExtent
//MagickSetImageExtent() sets the image size (i.e. columns & rows).
func (mw *MagickWand) SetImageExtent(columns uint, rows unsigned) error {
}

//columns
//The image width in pixels.
//rows
//The image height in pixels.
//MagickSetImageFilename
//MagickSetImageFilename() sets the filename of a particular image in a sequence.
func (mw *MagickWand) SetImageFilename(filename string) error {
}

//filename
//the image filename.
//MagickSetImageFormat
//MagickSetImageFormat() sets the format of a particular image in a sequence.
func (mw *MagickWand) SetImageFormat(format string) error {
}

//format
//the image format.
//MagickSetImageFuzz
//MagickSetImageFuzz() sets the image fuzz.
func (mw *MagickWand) SetImageFuzz(fuzz double) error {
}

//fuzz
//the image fuzz.
//MagickSetImageGamma
//MagickSetImageGamma() sets the image gamma.
func (mw *MagickWand) SetImageGamma(gamma double) error {
}

//gamma
//the image gamma.
//MagickSetImageGravity
//MagickSetImageGravity() sets the image gravity type.
func (mw *MagickWand) SetImageGravity(gravity GravityType) error {
}

//gravity
//the image interlace scheme: NoInterlace, LineInterlace, PlaneInterlace, PartitionInterlace.
//MagickSetImageGreenPrimary
//MagickSetImageGreenPrimary() sets the image chromaticity green primary point.
func (mw *MagickWand) SetImageGreenPrimary(x double, y double) error {
}

//x
//the green primary x-point.
//y
//the green primary y-point.
//MagickSetImageInterlaceScheme
//MagickSetImageInterlaceScheme() sets the image interlace scheme.
func (mw *MagickWand) SetImageInterlaceScheme(interlace InterlaceType) error {
}

//interlace
//the image interlace scheme: NoInterlace, LineInterlace, PlaneInterlace, PartitionInterlace.
//MagickSetImageInterpolateMethod
//MagickSetImageInterpolateMethod() sets the image interpolate pixel method.
func (mw *MagickWand) SetImageInterpolateMethod(method InterpolatePixelMethod) error {
}

//method
//the image interpole pixel methods: choose from Undefined, Average, Bicubic, Bilinear, Filter, Integer, Mesh, NearestNeighbor.
//MagickSetImageIterations
//MagickSetImageIterations() sets the image iterations.
func (mw *MagickWand) SetImageIterations(iterations uint) error {
}

//delay
//the image delay in 1/100th of a second.
//MagickSetImageMatte
//MagickSetImageMatte() sets the image matte channel.
func (mw *MagickWand) SetImageMatteColor(matte *MagickBooleanType) error {
}

//matte
//Set to MagickTrue to enable the image matte channel otherwise MagickFalse.
//MagickSetImageMatteColor
//MagickSetImageMatteColor() sets the image matte color.
func (mw *MagickWand) SetImageMatteColor(matte *PixelWand) error {
}

//matte
//the matte pixel wand.
//MagickSetImageOpacity
//MagickSetImageOpacity() sets the image to the specified opacity level.
func (mw *MagickWand) SetImageOpacity(alpha double) error {
}

//alpha
//the level of transparency: 1.0 is fully opaque and 0.0 is fully transparent.
//MagickSetImageOrientation
//MagickSetImageOrientation() sets the image orientation.
func (mw *MagickWand) SetImageOrientation(orientation OrientationType) error {
}

//orientation
//the image orientation type.
//MagickSetImagePage
//MagickSetImagePage() sets the page geometry of the image.
func (mw *MagickWand) SetImagePage(width uint, height uint, x int, y int) error {
}

//width
//the page width.
//height
//the page height.
//x
//the page x-offset.
//y
//the page y-offset.
//MagickSetImageProgressMonitor
//MagickSetImageProgressMonitor() sets the wand image progress monitor to the specified method and returns the previous progress monitor if any. The progress monitor method looks like this:
//MagickBooleanType MagickProgressMonitor(text string,offset MagickOffsetType,span MagickSizeType,void *client_data)
//If the progress monitor returns MagickFalse, the current operation is interrupted.
//MagickProgressMonitor MagickSetImageProgressMonitor(mw.wand,progress_monitor MagickProgressMonitor,void *client_data)
//progress_monitor
//Specifies a pointer to a method to monitor progress of an image operation.
//client_data
//Specifies a pointer to any client data.
//MagickSetImageRedPrimary
//MagickSetImageRedPrimary() sets the image chromaticity red primary point.
func (mw *MagickWand) SetImageRedPrimary(x double, y double) error {
}

//x
//the red primary x-point.
//y
//the red primary y-point.
//MagickSetImageRenderingIntent
//MagickSetImageRenderingIntent() sets the image rendering intent.
func (mw *MagickWand) SetImageRenderingIntent(rendering_intent RenderingIntent) error {
}

//rendering_intent
//the image rendering intent: UndefinedIntent, SaturationIntent, PerceptualIntent, AbsoluteIntent, or RelativeIntent.
//MagickSetImageResolution
//MagickSetImageResolution() sets the image resolution.
func (mw *MagickWand) SetImageResolution(x_resolution double, y_resolution double) error {
}

//x_resolution
//the image x resolution.
//y_resolution
//the image y resolution.
//MagickSetImageScene
//MagickSetImageScene() sets the image scene.
func (mw *MagickWand) SetImageScene(scene uint) error {
}

//delay
//the image scene number.
//MagickSetImageTicksPerSecond
//MagickSetImageTicksPerSecond() sets the image ticks-per-second.
func (mw *MagickWand) SetImageTicksPerSecond(ticks_per_second int) error {
}

//ticks_per_second
//the units to use for the image delay.
//MagickSetImageType
//MagickSetImageType() sets the image type.
func (mw *MagickWand) SetImageType(image_type ImageType) error {
}

//image_type
//the image type: UndefinedType, BilevelType, GrayscaleType, GrayscaleMatteType, PaletteType, PaletteMatteType, TrueColorType, TrueColorMatteType, ColorSeparationType, ColorSeparationMatteType, or OptimizeType.
//MagickSetImageUnits
//MagickSetImageUnits() sets the image units of resolution.
func (mw *MagickWand) SetImageUnits(units ResolutionType) error {
}

//units
//the image units of resolution : UndefinedResolution, PixelsPerInchResolution, or PixelsPerCentimeterResolution.
//MagickSetImageVirtualPixelMethod
//MagickSetImageVirtualPixelMethod() sets the image virtual pixel method.
//VirtualPixelMethod MagickSetImageVirtualPixelMethod(mw.wand,method VirtualPixelMethod)
//method
//the image virtual pixel method : UndefinedVirtualPixelMethod, ConstantVirtualPixelMethod, EdgeVirtualPixelMethod, MirrorVirtualPixelMethod, or TileVirtualPixelMethod.
//MagickSetImageWhitePoint
//MagickSetImageWhitePoint() sets the image chromaticity white point.
func (mw *MagickWand) SetImageWhitePoint(x double, y double) error {
}

//x
//the white x-point.
//y
//the white y-point.
//MagickShadeImage
//MagickShadeImage() shines a distant light on an image to create a three-dimensional effect. You control the positioning of the light with azimuth and elevation; azimuth is measured in degrees off the x axis and elevation is measured in pixels above the Z axis.
func (mw *MagickWand) ShadeImage(invert bool, azimuth double, elevation double) error {
}

//gray
//A value other than zero shades the intensity of each pixel.
//azimuth, elevation
//Define the light source direction.
//MagickShadowImage
//MagickShadowImage() simulates an image shadow.
func (mw *MagickWand) ShadowImage(opacity double, sigma double, x int, y int) error {
}

//opacity
//percentage transparency.
//sigma
//the standard deviation of the Gaussian, in pixels.
//x
//the shadow x-offset.
//y
//the shadow y-offset.
//MagickSharpenImage
//MagickSharpenImage() sharpens an image. We convolve the image with a Gaussian operator of the given radius and standard deviation (sigma). For reasonable results, the radius should be larger than sigma. Use a radius of 0 and MagickSharpenImage() selects a suitable radius for you.
func (mw *MagickWand) SharpenImage(radius double, sigma double) error {
}
func (mw *MagickWand) SharpenImageChannel(channel ChannelType, radius double, sigma double) error {
}

//channel
//the image channel(s).
//radius
//the radius of the Gaussian, in pixels, not counting the center pixel.
//sigma
//the standard deviation of the Gaussian, in pixels.
//MagickShaveImage
//MagickShaveImage() shaves pixels from the image edges. It allocates the memory necessary for the new Image structure and returns a pointer to the new image.
func (mw *MagickWand) ShaveImage(columns uint, rows uint) error {
}

//columns
//the number of columns in the scaled image.
//rows
//the number of rows in the scaled image.
//MagickShearImage
//MagickShearImage() slides one edge of an image along the X or Y axis, creating a parallelogram. An X direction shear slides an edge along the X axis, while a Y direction shear slides an edge along the Y axis. The amount of the shear is controlled by a shear angle. For X direction shears, x_shear is measured relative to the Y axis, and similarly, for Y direction shears y_shear is measured relative to the X axis. Empty triangles left over from shearing the image are filled with the background color.
func (mw *MagickWand) ShearImage(background *PixelWand, x_shear double, y_shear double) error {
}

//background
//the background pixel wand.
//x_shear
//the number of degrees to shear the image.
//y_shear
//the number of degrees to shear the image.
//MagickSigmoidalContrastImage
//MagickSigmoidalContrastImage() adjusts the contrast of an image with a non-linear sigmoidal contrast algorithm. Increase the contrast of the image using a sigmoidal transfer function without saturating highlights or shadows. Contrast indicates how much to increase the contrast (0 is none; 3 is typical; 20 is pushing it); mid-point indicates where midtones fall in the resultant image (0 is white; 50 is middle-gray; 100 is black). Set sharpen to MagickTrue to increase the image contrast otherwise the contrast is reduced.
func (mw *MagickWand) SigmoidalContrastImage(invert bool, alpha double, beta double) error {
}
func (mw *MagickWand) SigmoidalContrastImageChannel(channel ChannelType, sharpen bool, alpha double, beta double) error {
}

//channel
//Identify which channel to level: RedChannel, GreenChannel,
//sharpen
//Increase or decrease image contrast.
//alpha
//strength of the contrast, the larger the number the more 'threshold-like' it becomes.
//beta
//midpoint of the function as a color value 0 to QuantumRange.
//MagickSimilarityImage
//MagickSimilarityImage() compares the reference image of the image and returns the best match offset. In addition, it returns a similarity image such that an exact match location is completely white and if none of the pixels match, black, otherwise some gray level in-between.
func (mw *MagickWand) SimilarityImage(reference *MagickWand, RectangeInfo *offset, double *similarity) *MagickWand {
}

//reference
//the reference wand.
//offset
//the best match offset of the reference image within the image.
//similarity
//the computed similarity between the images.
//MagickSketchImage
//MagickSketchImage() simulates a pencil sketch. We convolve the image with a Gaussian operator of the given radius and standard deviation (sigma). For reasonable results, radius should be larger than sigma. Use a radius of 0 and SketchImage() selects a suitable radius for you. Angle gives the angle of the blurring motion.
func (mw *MagickWand) SketchImage(radius double, sigma double, angle double) error {
}

//radius
//the radius of the Gaussian, in pixels, not counting the center pixel.
//sigma
//the standard deviation of the Gaussian, in pixels.
//angle
//Apply the effect along this angle.
//MagickSmushImages
//MagickSmushImages() takes all images from the current image pointer to the end of the image list and smushs them to each other top-to-bottom if the stack parameter is true, otherwise left-to-right.
func (mw *MagickWand) SmushImages(stack bool, offset int) *MagickWand {
}

//stack
//By default, images are stacked left-to-right. Set stack to MagickTrue to stack them top-to-bottom.
//offset
//minimum distance in pixels between images.
//MagickSolarizeImage
//MagickSolarizeImage() applies a special effect to the image, similar to the effect achieved in a photo darkroom by selectively exposing areas of photo sensitive paper to light. Threshold ranges from 0 to QuantumRange and is a measure of the extent of the solarization.
func (mw *MagickWand) SolarizeImage(threshold double) error {
}
func (mw *MagickWand) SolarizeImageChannel(channel ChannelType, threshold double) error {
}

//channel
//the image channel(s).
//threshold
//Define the extent of the solarization.
//MagickSparseColorImage
//MagickSparseColorImage(), given a set of coordinates, interpolates the colors found at those coordinates, across the whole image, using various methods.
func (mw *MagickWand) SparseColorImage(channel ChannelType, method SparseColorMethod, number_arguments uint, arguments []double) error {
}

//image
//the image to be sparseed.
//method
//the method of image sparseion.
//ArcSparseColorion will always ignore source image offset, and always 'bestfit' the destination image with the top left corner offset relative to the polar mapping center.
//Bilinear has no simple inverse mapping so will not allow 'bestfit' style of image sparseion.
//Affine, Perspective, and Bilinear, will do least squares fitting of the distrotion when more than the minimum number of control point pairs are provided.
//Perspective, and Bilinear, will fall back to a Affine sparseion when less than 4 control point pairs are provided. While Affine sparseions will let you use any number of control point pairs, that is Zero pairs is a No-Op (viewport only) distrotion, one pair is a translation and two pairs of control points will do a scale-rotate-translate, without any shearing.
//number_arguments
//the number of arguments given for this sparseion method.
//arguments
//the arguments for this sparseion method.
//MagickSpliceImage
//MagickSpliceImage() splices a solid color into the image.
func (mw *MagickWand) SpliceImage(width uint, height uint, x int, y int) error {
}

//width
//the region width.
//height
//the region height.
//x
//the region x offset.
//y
//the region y offset.
//MagickSpreadImage
//MagickSpreadImage() is a special effects method that randomly displaces each pixel in a block defined by the radius parameter.
func (mw *MagickWand) SpreadImage(radius double) error {
}

//radius
//Choose a random pixel in a neighborhood of this extent.
//MagickStatisticImage
//MagickStatisticImage() replace each pixel with corresponding statistic from the neighborhood of the specified width and height.
func (mw *MagickWand) StatisticImage(stype StatisticType, width double, height uint) error {
}
func (mw *MagickWand) StatisticImageChannel(channel ChannelType, stype StatisticType, width double, height uint) error {
}

//channel
//the image channel(s).
//type
//the statistic type (e.g. median, mode, etc.).
//width
//the width of the pixel neighborhood.
//height
//the height of the pixel neighborhood.
//MagickSteganoImage
//MagickSteganoImage() hides a digital watermark within the image. Recover the hidden watermark later to prove that the authenticity of an image. Offset defines the start position within the image to hide the watermark.
func (mw *MagickWand) SteganoImage(watermark_wand *MagickWand, offset int) *MagickWand {
}

//watermark_wand
//the watermark wand.
//offset
//Start hiding at this offset into the image.
//MagickStereoImage
//MagickStereoImage() composites two images and produces a single image that is the composite of a left and right image of a stereo pair
func (mw *MagickWand) StereoImage(offset_wand *MagickWand) *MagickWand {
}

//offset_wand
//Another image wand.
//MagickStripImage
//MagickStripImage() strips an image of all profiles and comments.
func (mw *MagickWand) StripImage() error {
}

//MagickSwirlImage
//MagickSwirlImage() swirls the pixels about the center of the image, where degrees indicates the sweep of the arc through which each pixel is moved. You get a more dramatic effect as the degrees move from 1 to 360.
func (mw *MagickWand) SwirlImage(degrees double) error {
}

//degrees
//Define the tightness of the swirling effect.
//MagickTextureImage
//MagickTextureImage() repeatedly tiles the texture image across and down the image canvas.
//MagickWand *MagickTextureImage(mw.wand, texture_wand *MagickWand)
//texture_wand
//the texture wand
//MagickThresholdImage
//MagickThresholdImage() changes the value of individual pixels based on the intensity of each pixel compared to threshold. The result is a high-contrast, two color image.
func (mw *MagickWand) ThresholdImage(threshold double) error {
}
func (mw *MagickWand) ThresholdImageChannel(channel ChannelType, threshold double) error {
}

//channel
//the image channel(s).
//threshold
//Define the threshold value.
//MagickThumbnailImage
//MagickThumbnailImage() changes the size of an image to the given dimensions and removes any associated profiles. The goal is to produce small low cost thumbnail images suited for display on the Web.
func (mw *MagickWand) ThumbnailImage(columns uint, rows uint) error {
}

//columns
//the number of columns in the scaled image.
//rows
//the number of rows in the scaled image.
//MagickTintImage
//MagickTintImage() applies a color vector to each pixel in the image. The length of the vector is 0 for black and white and at its maximum for the midtones. The vector weighting function is f(x)=(1-(4.0*((x-0.5)*(x-0.5)))).
func (mw *MagickWand) TintImage(tint *PixelWand, opacity *PixelWand) error {
}

//tint
//the tint pixel wand.
//opacity
//the opacity pixel wand.
//MagickTransformImage
//MagickTransformImage() is a convenience method that behaves like MagickResizeImage() or MagickCropImage() but accepts scaling and/or cropping information as a region geometry specification. If the operation fails, a NULL image handle is returned.
func (mw *MagickWand) TransformImage(crop string, geometry string) *MagickWand {
}

//crop
//A crop geometry string. This geometry defines a subregion of the image to crop.
//geometry
//An image geometry string. This geometry defines the final size of the image.
//MagickTransformImageColorspace
//MagickTransformImageColorspace() transform the image colorspace, setting the images colorspace while transforming the images data to that colorspace.
func (mw *MagickWand) TransformImageColorspace(colorspace ColorspaceType) error {
}

//colorspace
//the image colorspace: UndefinedColorspace, sRGBColorspace, RGBColorspace, GRAYColorspace, OHTAColorspace, XYZColorspace, YCbCrColorspace, YCCColorspace, YIQColorspace, YPbPrColorspace, YPbPrColorspace, YUVColorspace, CMYKColorspace, HSLColorspace, HWBColorspace.
//MagickTransparentPaintImage
//MagickTransparentPaintImage() changes any pixel that matches color with the color defined by fill.
func (mw *MagickWand) TransparentPaintImage(target *PixelWand, alpha double, fuzz double, invert bool) error {
}

//target
//Change this target color to specified opacity value within the image.
//alpha
//the level of transparency: 1.0 is fully opaque and 0.0 is fully transparent.
//fuzz
//By default target must match a particular pixel color exactly. However, in many cases two colors may differ by a small amount. The fuzz member of image defines how much tolerance is acceptable to consider two colors as the same. For example, set fuzz to 10 and the color red at intensities of 100 and 102 respectively are now interpreted as the same color for the purposes of the floodfill.
//invert
//paint any pixel that does not match the target color.
//MagickTransposeImage
//MagickTransposeImage() creates a vertical mirror image by reflecting the pixels around the central x-axis while rotating them 90-degrees.
func (mw *MagickWand) TransposeImage() error {
}

//MagickTransverseImage
//MagickTransverseImage() creates a horizontal mirror image by reflecting the pixels around the central y-axis while rotating them 270-degrees.
func (mw *MagickWand) TransverseImage() error {
}

//MagickTrimImage
//MagickTrimImage() remove edges that are the background color from the image.
func (mw *MagickWand) TrimImage(fuzz double) error {
}

//fuzz
//By default target must match a particular pixel color exactly. However, in many cases two colors may differ by a small amount. The fuzz member of image defines how much tolerance is acceptable to consider two colors as the same. For example, set fuzz to 10 and the color red at intensities of 100 and 102 respectively are now interpreted as the same color for the purposes of the floodfill.
//MagickUniqueImageColors
//MagickUniqueImageColors() discards all but one of any pixel color.
func (mw *MagickWand) UniqueImageColors() error {
}

//MagickUnsharpMaskImage
//MagickUnsharpMaskImage() sharpens an image. We convolve the image with a Gaussian operator of the given radius and standard deviation (sigma). For reasonable results, radius should be larger than sigma. Use a radius of 0 and UnsharpMaskImage() selects a suitable radius for you.
func (mw *MagickWand) UnsharpMaskImage(radius double, sigma double, amount double, threshold double) error {
}
func (mw *MagickWand) UnsharpMaskImageChannel(channel ChannelType, radius double, sigma double, amount double, threshold double) error {
}

//channel
//the image channel(s).
//radius
//the radius of the Gaussian, in pixels, not counting the center pixel.
//sigma
//the standard deviation of the Gaussian, in pixels.
//amount
//the percentage of the difference between the original and the blur image that is added back into the original.
//threshold
//the threshold in pixels needed to apply the diffence amount.
//MagickVignetteImage
//MagickVignetteImage() softens the edges of the image in vignette style.
func (mw *MagickWand) VignetteImage(black_point double, white_point double, x int, y int) error {
}

//black_point
//the black point.
//white_point
//the white point.
//x, y
//Define the x and y ellipse offset.
//MagickWaveImage
//MagickWaveImage() creates a "ripple" effect in the image by shifting the pixels vertically along a sine wave whose amplitude and wavelength is specified by the given parameters.
func (mw *MagickWand) WaveImage(amplitude double, wave_length double) error {
}

//amplitude, wave_length
//Define the amplitude and wave length of the sine wave.
//MagickWhiteThresholdImage
//MagickWhiteThresholdImage() is like ThresholdImage() but force all pixels above the threshold into white while leaving all pixels below the threshold unchanged.
func (mw *MagickWand) WhiteThresholdImage(threshold *PixelWand) error {
}

//threshold
//the pixel wand.
// Writes an image to the specified filename.
func (mw *MagickWand) WriteImage(filename string) error {
	csfilename := C.CString(filename)
	defer C.free(unsafe.Pointer(csfilename))
	C.MagickWriteImage(mw.wand, csfilename)
	return mw.GetLastError()
}

// Writes an image to an open file descriptor.
func (mw *MagickWand) WriteImageImageFile(out *os.File) error {
	cmode := C.CString("w+")
	defer C.free(unsafe.Pointer(cmode))
	file := C.fdopen(C.int(out.Fd()), cmode)
	C.MagickWriteImageFile(mw.wand, file)
	return mw.GetLastError()
}

// Writes an image or image sequence.
func (mw *MagickWand) WriteImages(filename string, adjoin bool) error {
	cbool := C.MagickBooleanType(0)
	if adjoin {
		cbool = 1
	}
	csfilename := C.CString(filename)
	defer C.free(unsafe.Pointer(csfilename))
	C.MagickWriteImages(mw.wand, csfilename, cbool)
	return mw.GetLastError()
}

// Writes an image sequence to an open file descriptor.
func (mw *MagickWand) WriteImagesFile(out *os.File) error {
	cmode := C.CString("w+")
	defer C.free(unsafe.Pointer(cmode))
	file := C.fdopen(C.int(out.Fd()), cmode)
	C.MagickWriteImagesFile(mw.wand, file)
	return mw.GetLastError()
}