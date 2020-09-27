package video

import "godoom/configuration"

func BindVideoVariables() {
	configuration.BindBoolVariable("use_mouse", &UseMouse)
	configuration.BindBoolVariable("fullscreen", &Fullscreen)
	configuration.BindIntVariable("video_display", &VideoDisplay)
	configuration.BindBoolVariable("aspect_ratio_correct", &AspectRatioCorrect)
	configuration.BindBoolVariable("integer_scaling", &IntegerScaling)
	configuration.BindBoolVariable("vga_porch_flash", &VgaPorchFlash)
	configuration.BindIntVariable("startup_delay", &StartupDelay)
	configuration.BindIntVariable("fullscreen_width", &FullscreenWidth)
	configuration.BindIntVariable("fullscreen_height", &FullscreenHeight)
	configuration.BindBoolVariable("force_software_renderer", &ForceSoftwareRenderer)
	configuration.BindIntVariable("max_scaling_buffer_pixels", &MaxScalingBufferPixels)
	configuration.BindIntVariable("window_width", &WindowWidth)
	configuration.BindIntVariable("window_height", &WindowHeight)
	configuration.BindBoolVariable("grabmouse", &GrabMouse)
	configuration.BindStringVariable("video_driver", &VideoDriver)
	configuration.BindStringVariable("window_position", &WindowPosition)
	configuration.BindIntVariable("usegamma", &UseGamma)
	configuration.BindBoolVariable("png_screenshots", &PngScreenshots)
}
