package main

import (
	"reflect"
	"strconv"

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
	bc.Data = []int{87, 54, 100, 50, 88}
	bc.DataLabels = []string{"S1", "S2", "S3", "S4", "S5"}
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

	config := ui.NewList()
	configList := []string{
		"[email] [test@mail.ru](fg-green)",
		"[password] [test](fg-yellow)",
	}
	config.Items = configList
	config.ItemFgColor = ui.ColorCyan
	config.BorderLabel = "Конфигурация"
	config.Height = 10
	config.Width = 25
	config.Y = 0

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
		Config:         config,
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
			config.getConfig(&u)
			mas := make([]string, 0)
			v := reflect.ValueOf(*config)
			n := v.Type().NumField()
			for i := 0; i < n; i++ {
				g := v.Type().Field(i)
				switch v.Field(i).Kind() {
				case reflect.String:
					mas = append(mas, "["+g.Name+"] ["+v.Field(i).String()+"](fg-yellow)")
					break
				case reflect.Int:
					mas = append(mas, "["+g.Name+"] ["+strconv.Itoa((int)(v.Field(i).Int()))+"](fg-yellow)")
					break
				}
			}
			u.Config.Items = mas
			ui.Render(u.Config)
		}
	})
	ui.Handle("t", func(ui.Event) {
		// тестовое  письмо
		if field == 1 {
			u.LogError("[Email] [шлю тестовый email](fg-red)")
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
