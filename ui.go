package main

import (
	"fmt"

	ui "github.com/gizak/termui"
	"github.com/gizak/termui/extra"
)

var (
	field int
)

type UiService struct {
	Tabpane *extra.Tabpane
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
	ConfigStatus   *ui.Par
}

func NewUiService() UiService {
	err := ui.Init()
	if err != nil {
		panic(err)
	}

	tab1 := extra.NewTab("Статистика")
	tab2 := extra.NewTab("Конфигурация")
	tabpane := extra.NewTabpane()
	tabpane.Y = 1
	tabpane.Width = 20
	tabpane.SetTabs(*tab1, *tab2)

	listCommands := ui.NewList()
	strs := []string{
		"[q] [Выход](fg-red)",
		"[u] [Обновление состояния](fg-blue)",
		"[t] [Тестовое письмо](fg-red)",
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
	configList.Height = 10
	configList.Width = 25
	configList.Y = 0

	configStatus := ui.NewPar("[Конфигурация прочитана успешно](fg-green)")
	configStatus.Height = 3
	configStatus.Width = 37
	configStatus.Y = 4
	configStatus.BorderFg = ui.ColorGreen
	field = 1

	return UiService{
		Tabpane:        tabpane,
		Commands:       listCommands,
		Status:         bc,
		Logs:           lg,
		ConfigCommands: configCommands,
		Config:         configList,
		ConfigStatus:   configStatus,
		logs: []string{
			"[TMP] [FIRE! FIRE!](fg-red)",
		},
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

	// checking keys
	u.CheckKeys()
	ui.Loop()
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

func (u UiService) CheckKeys() {
	ui.Handle("q", func(ui.Event) {
		ui.StopLoop()
	})
	ui.Handle("u", func(ui.Event) {
		// обновление состояния
		if field == 1 {
			parse(&u, config)
		}
	})
	ui.Handle("r", func(ui.Event) {
		if field == 2 {
			u.RefreshConfig()
		}
	})
	ui.Handle("n", func(ui.Event) {
		if field == 2 {
			CreateConfig(&u)
		}
	})
	ui.Handle("t", func(ui.Event) {
		// тестовое  письмо
		if field == 1 {
			u.LogError("[Email] [шлю тестовый email](fg-green)")
			sendMessageByEmail(&u ,"Тестовое письмо")
			u.LogError("[Email] [Отправлено](fg-green)")
		}
	})

	// switch layout
	ui.Handle("1", func(ui.Event) {
		u.SetMainLayout()
		field = 1
	})
	ui.Handle("2", func(ui.Event) {
		u.SetConfigLayout()
		field = 2
	})
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
		u.ConfigStatus.Text = "[Ошибка чтения конфигурации. Создайте новую нажав [n]](fg-red)"
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
