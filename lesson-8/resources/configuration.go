package resources

import (
	"flag"
	"fmt"
	"github.com/kelseyhightower/envconfig"
	"regexp"
)

var (
	timeout = flag.Int("timeout", 30, "Timeout of connection")
)

type Configuration struct {
	Debug     bool   `envconfig:"DEBUG" default:"false" required:"true"`
	Port      int    `envconfig:"PORT" default:"8080" required:"true"`
	DBURL     string `envconfig:"DB_URL" default:"postgres://user:password@localhost:5432/petstore?sslmode=disable" required:"true"`
	connected bool
}

//NewConnection set new connection
func (conf *Configuration) NewConnection() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Could not connect", r)
		}
	}()

	flag.Parse()

	err := envconfig.Process("", conf)
	if err != nil {
		panic(err.Error())
	}

	format := "Debug: %v\nPort: %d\nTimeout: %d\n"
	_, err = fmt.Printf(format, conf.Debug, conf.Port, timeout)
	if err != nil {
		panic(err.Error())
	}

	err = validateConfiguration(*conf)
	if err != nil {
		conf.CloseConnection()
		panic(err.Error())
	}
	conf.connected = true
	fmt.Println("Connected")
}

func (conf *Configuration) CloseConnection() {
	conf.connected = false

	fmt.Println("Connection is closed")
}

//validateConfiguration valid configuration
func validateConfiguration(conf Configuration) error {
	err := validDBURL(conf.DBURL)

	if err != nil {
		fmt.Println(err.Error())
	}

	return nil
}

func validDBURL(dbUrl string) error {
	var valid = regexp.MustCompile(`postgres:\/\/(?:[^:]+):(?:[^:]+):(?:\d{4})`)

	if ok := valid.MatchString(dbUrl); ok != true {
		panic("DB_URL is not set right")
	}

	return nil
}
