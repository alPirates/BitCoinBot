package main

var (
    config *Config
)

func main() {
	config = getConfig()
    ui := NewUiService()
    go loopParse(&ui, config)
    ui.Init()
}
