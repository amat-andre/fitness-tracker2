package actioninfo

import (
	"fmt"
	"log"

	//"github.com/Yandex-Practicum/tracker/internal/daysteps"
	// "github.com/Yandex-Practicum/tracker/internal/trainings"
)

type DataParser interface {
	Parse(string) error
	ActionInfo() (string, error)
}

func Info(dataset []string, dp DataParser) {
	for _, data := range dataset {
		err := dp.Parse(data)
		if err != nil {
			log.Println(err)
			continue
		}

		line, err := dp.ActionInfo()
		if err != nil {
			log.Println(err)
		}
		fmt.Println(line)
	}
}
