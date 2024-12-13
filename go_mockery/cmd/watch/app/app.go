package app

import "fmt"

type IWatch interface {
	GetCurrentTime() (string, error)
}

type App struct {
	Watch IWatch
}

func NewApp(w IWatch) *App {
	return &App{Watch: w}
}

func (a *App) ShowCurrentTime() string {
	current, err := a.Watch.GetCurrentTime()
	if err != nil {
		return "Cannot Get Current Time"
	}
	return fmt.Sprintf("This time is : %v", current)
}

func (a *App) Run() {
	fmt.Println(a.ShowCurrentTime())
}
