package app

type App struct {
	name string
}

func New() *App {
	return &App{name: "English Platform"}
}

func (a *App) Name() string {
	return a.name
}
