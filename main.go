package main

var (
	config *Config
)

func main() {
	config = &Config{}
	ui := NewUiService()
	config.getConfig(&ui)
	go loopParse(&ui, config)
	ui.Init()
}
