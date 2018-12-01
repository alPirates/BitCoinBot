package main

import "log"

type Notifier interface {
	Notify(ui *UiService, message string) error
}

type NotifyService struct {
	Services []Notifier
}

func NewNotifyService(services ...Notifier) NotifyService {
	return NotifyService{
		Services: services,
	}
}

func (n NotifyService) Notify(ui *UiService, message string) {
	for _, notifier := range n.Services {
		err := notifier.Notify(ui, message)
		if err != nil {
			log.Println(err)
		}
	}
}
