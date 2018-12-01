package main

var (
	config *Config
)

func main() {
	config = &Config{}
	ui := NewUiService()
	go loopParse(&ui, config)
	ui.Init()
}
