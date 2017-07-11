package hello

import (
	"expvar"
	"log"
	"net/http"

	logging "gopkg.in/tokopedia/logging.v1"
)

type ServerConfig struct {
	Name string
}

type Config struct {
	Server ServerConfig
}

type HelloWorldModule struct {
	cfg       *Config
	something string
	stats     *expvar.Int
}

func NewHelloWorldModule() *HelloWorldModule {

	var cfg Config

	ok := logging.ReadModuleConfig(&cfg, "config", "hello") || logging.ReadModuleConfig(&cfg, "files/etc/gosample", "hello")
	if !ok {
		// when the app is run with -e switch, this message will automatically be redirected to the log file specified
		log.Fatalln("failed to read config")
	}

	// this message only shows up if app is run with -debug option, so its great for debugging
	logging.Debug.Println("hello init called", cfg.Server.Name)

	return &HelloWorldModule{
		cfg:       &cfg,
		something: "John Doe",
		stats:     expvar.NewInt("rpsStats"),
	}

}

func (hlm *HelloWorldModule) SayHelloWorld(w http.ResponseWriter, r *http.Request) {
	hlm.stats.Add(1)

	// r.ParseForm()
	// r.Form -> map[string][]string
	// for key, val := range r.Form {
	// 	for _, val2 := range val {
	// 		fmt.Println(key, val2)
	// 	}
	// }

	// mapGw := map[string]string{
	// 	"foo": "bar",
	// }
	// var mapGw map[string]string

	// mapGw := make(map[string]interface{})
	// mapGw["foo1"] = 1
	// mapGw["foo2"] = "bar"
	// mapGw["foo"] = "bar"

	name := r.FormValue("name")
	w.Write([]byte("Hello " + name))
}
