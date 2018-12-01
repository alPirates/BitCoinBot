package main

var (
	config *Config
)

func main() {
	config = &Config{}
	notifyService := NewNotifyService(
		NewEmailService(),
		NewVkService(),
	)
	ui := NewUiService(&notifyService)
	go loopParse(&ui, config)
	ui.Init()
}
