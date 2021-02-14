package resources

import (
	"flag"
	"github.com/kelseyhightower/envconfig"
	"log"
	"regexp"
)

var (
	timeout = flag.Int("timeout", 30, "Timeout of connection in seconds")
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
			log.Println("Could not connect", r)
		}
	}()

	flag.Parse()

	err := envconfig.Process("", conf)
	if err != nil {
		panic(err.Error())
	}

	format := "Debug: %v\nPort: %d\nTimeout: %d\n"
	log.Printf(format, conf.Debug, conf.Port, *timeout)

	err = validateConfiguration(*conf)
	if err != nil {
		conf.CloseConnection()
		panic(err.Error())
	}
	conf.connected = true
	log.Println("Connected")
}

func (conf *Configuration) CloseConnection() {
	conf.connected = false

	log.Println("Connection is closed")
}

//validateConfiguration valid configuration
func validateConfiguration(conf Configuration) error {
	err := validDBURL(conf.DBURL)

	if err != nil {
		log.Println(err.Error())
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
