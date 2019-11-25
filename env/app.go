package env

type envApp struct {
	Username string `default:"name"`
	Password string `default:"password"`
	Dbname   string `default:"dbname"`
}

var app *envApp

func App() *envApp {
	if app == nil {
		app = process("", &envApp{}).(*envApp)
	}
	return app
}
