package main

import (
	"fmt"
	"os"

	ui "github.com/gizak/termui"
)

var (
	field int
)

type UiService struct {
	// NotifyService
	NotifyServ *NotifyService
	// tabpane
	Tabpane *ui.TabPane
	// main layout
	Layout1  []*ui.Row
	Commands *ui.List
	Logs     *ui.List
	Status   *ui.BarChart
	logs     []string
	// config layout
	Layout2        []*ui.Row
	ConfigCommands *ui.List
	Config         *ui.List
	ConfigStatus   *ui.Paragraph
}

func NewUiService(notify *NotifyService) UiService {
	err := ui.Init()
	if err != nil {
		panic(err)
	}

	tab1 := ui.NewTab("Статистика")
	tab2 := ui.NewTab("Конфигурация")
	tabpane := ui.NewTabPane()
	tabpane.Y = 1
	tabpane.Width = 20
	tabpane.SetTabs(*tab1, *tab2)

	listCommands := ui.NewList()
	strs := []string{
		"[q] [Выход](fg-red)",
		"[u] [Обновление состояния](fg-blue)",
		"[t] [Тестовое письмо](fg-red)",
		"[v] [Тестовое сообщение vk](fg-red)",
		"---------------------",
		"[1] [Главная панель](fg-green)",
		"[2] [Конфигурация](fg-yellow)",
	}
	listCommands.Items = strs
	listCommands.ItemFgColor = ui.ColorYellow
	listCommands.BorderLabel = "Возможные комманды"
	listCommands.Height = 10
	listCommands.Width = 25
	listCommands.Y = 0

	bc := ui.NewBarChart()
	bc.BorderLabel = "Состояние"
	bc.Width = 26
	bc.Height = 10
	bc.TextColor = ui.ColorGreen
	bc.BarColor = ui.ColorRed
	bc.NumColor = ui.ColorYellow

	lg := ui.NewList()

	lg.ItemFgColor = ui.ColorYellow
	lg.BorderLabel = "Логи"
	lg.Height = 15
	lg.Width = 25
	lg.Y = 0

	configCommands := ui.NewList()
	commandsList := []string{
		"[r] [Обновить конфигурацию](fg-red)",
		"[n] [Создать дефолтный файл](fg-red)",
		"---------------------",
		"[1] [Главная панель](fg-green)",
		"[2] [Конфигурация](fg-yellow)",
	}
	configCommands.Items = commandsList
	configCommands.ItemFgColor = ui.ColorYellow
	configCommands.BorderLabel = "Меню"
	configCommands.Height = 10
	configCommands.Width = 25
	configCommands.Y = 0

	configList := ui.NewList()
	configListList := config.toStringMas()
	configList.Items = configListList
	configList.ItemFgColor = ui.ColorCyan
	configList.BorderLabel = "Конфигурация"
	configList.Height = 11
	configList.Width = 25
	configList.Y = 0

	configStatus := ui.NewParagraph("[Конфигурация прочитана успешно](fg-green)")
	configStatus.Height = 3
	configStatus.Width = 37
	configStatus.Y = 4
	configStatus.BorderFg = ui.ColorGreen
	field = 1

	return UiService{
		NotifyServ:     notify,
		Tabpane:        tabpane,
		Commands:       listCommands,
		Status:         bc,
		Logs:           lg,
		ConfigCommands: configCommands,
		Config:         configList,
		ConfigStatus:   configStatus,
		logs:           []string{},
	}
}

func (u UiService) Init() {
	layout1 := []*ui.Row{
		ui.NewRow(
			ui.NewCol(12, 0, u.Tabpane),
		),
		ui.NewRow(
			ui.NewCol(6, 0, u.Commands),
			ui.NewCol(6, 0, u.Status),
		),
		ui.NewRow(
			ui.NewCol(12, 0, u.Logs),
		),
	}

	layout2 := []*ui.Row{
		ui.NewRow(
			ui.NewCol(12, 0, u.Tabpane),
		),
		ui.NewRow(
			ui.NewCol(12, 0, u.ConfigCommands),
		),
		ui.NewRow(
			ui.NewCol(12, 0, u.Config),
		),
		ui.NewRow(
			ui.NewCol(12, 0, u.ConfigStatus),
		),
	}

	u.SetCharts([]int{80, 54, 100, 50, 88})

	u.RefreshConfig()

	u.Layout1 = layout1
	u.Layout2 = layout2

	u.SetMainLayout()

	uiEvents := ui.PollEvents()
	for {
		select {
		case e := <-uiEvents:
			switch e.ID {
			// press 'q' or 'C-c' to quit
			case "q", "<C-c>":
				ui.Close()
				os.Exit(0)
				// quit
			case "u":
				// обновление состояния
				if field == 1 {
					parse(&u, config)
				}
			case "r":
				if field == 2 {
					u.RefreshConfig()
				}
			case "n":
				if field == 2 {
					CreateConfig(&u)
				}
			case "t":
				// тестовое  письмо
				if field == 1 {
					u.LogError("[Email] [шлю тестовый email](fg-green)")
					sendMessageByEmail(&u, "Тестовое письмо")
					u.LogError("[Email] [Отправлено](fg-green)")
				}
			case "v":
				// тестовое  письмо
				if field == 1 {
					vk := NewVkService()
					u.LogError("[VK] [шлю тестовое сообщение vk](fg-green)")
					vk.Notify(&u, "Тестовое письмо")
					u.LogError("[VK] [Отправлено](fg-green)")
				}
			case "1":
				u.SetMainLayout()
				field = 1
			case "2":
				u.SetConfigLayout()
				field = 2
			case "<Resize>":
				payload := e.Payload.(ui.Resize)
				width := payload.Width
				ui.Body.Width = width
				ui.Body.Align()
				ui.Render(ui.Body)
			}
		}
	}
}

func (u UiService) SetMainLayout() {
	ui.Clear()
	ui.Body.Rows = u.Layout1
	ui.Body.Align()
	u.Tabpane.SetActiveLeft()
	ui.Render(ui.Body, u.Tabpane)
}

func (u UiService) SetConfigLayout() {
	ui.Clear()
	ui.Body.Rows = u.Layout2
	ui.Body.Align()
	u.Tabpane.SetActiveRight()
	ui.Render(ui.Body, u.Tabpane)
}

func (u UiService) LogError(message string) {
	if len(u.Logs.Items) == 13 {
		u.Logs.Items = u.Logs.Items[1:]
	}
	u.Logs.Items = append(u.Logs.Items, message)
	if field == 1 {
		ui.Render(u.Logs)
	}
}

func (u UiService) RefreshConfig() {
	config.getConfig(&u)
	u.Config.Items = config.toStringMas()
	ui.Render(u.Config)
}

func (u UiService) SetStatus(status bool) {
	if status {
		u.ConfigStatus.BorderFg = ui.ColorGreen
		u.ConfigStatus.Text = "[Конфигурация прочитана успешно](fg-green)"

	} else {
		u.ConfigStatus.BorderFg = ui.ColorRed
		u.ConfigStatus.Text = "[Ошибка чтения конфигурации. Создана новая](fg-red)"
		CreateConfig(&u)
		config.getConfig(&u)
	}
	ui.Render(u.ConfigStatus)
}

func (u UiService) SetCharts(data []int) {
	u.Status.Data = data
	lables := make([]string, 0)
	for i, _ := range data {
		lables = append(lables, fmt.Sprintf("V%d", i+1))
	}
	u.Status.DataLabels = lables
	ui.Render(u.Status)
}
