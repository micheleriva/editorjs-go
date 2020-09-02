package main

type Options struct {
	Image ImageOptions
}

type ImageOptions struct {
	Classes ImageClasses
	Caption string
}

type ImageClasses struct {
	WithBorder     string
	Stretched      string
	WithBackground string
}

func main() {

}
