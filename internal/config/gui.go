package config

type GUIConfig struct {
	Width     int
	Height    int
	Resizable bool
}

func BuildGUIConfig(width, height int, resizable bool) *GUIConfig {
	return &GUIConfig{
		Width:     width,
		Height:    height,
		Resizable: resizable,
	}
}
