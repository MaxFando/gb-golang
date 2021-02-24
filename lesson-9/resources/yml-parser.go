package resources

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"os"
)

type Employee struct {
	Name  string `yaml:"name"`
	Job   string `yaml:"job"`
	Skill string `yaml:"skill"`
}

func YmlDecode(filename string) string {
	file, err := os.Open(filename)
	if err != nil {
		panic(fmt.Sprintln("Cannot open the file: ", filename, err.Error()))
	}

	defer func() {
		err := file.Close()
		if err != nil {
			panic(fmt.Sprintln("Cannot close the file:", err.Error()))
		}
	}()

	employee := &Employee{}
	err = yaml.NewDecoder(file).Decode(employee)
	if err != nil {
		panic(fmt.Sprintln("Cannot decode file:", filename, err.Error()))
	}

	return employee.Name
}
