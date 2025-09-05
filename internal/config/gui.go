package config

type GUIConfig struct {
	Width     float32
	Height    float32
	Resizable bool
}

func BuildGUIConfig(width, height float32, resizable bool) *GUIConfig {
	return &GUIConfig{
		Width:     width,
		Height:    height,
		Resizable: resizable,
	}
}
