package main

import "log"

type Notifier interface {
	Notify(ui *UiService, message string) error
}

type NotifyService struct {
	Services  []Notifier
	UiService *UiService
}

func NewNotifyService(services ...Notifier) NotifyService {
	return NotifyService{
		Services: services,
	}
}

func (n NotifyService) Notify(message string) {
	for _, notifier := range n.Services {
		err := notifier.Notify(n.UiService, message)
		if err != nil {
			// handle error
			log.Println(err)
		}
	}
}
