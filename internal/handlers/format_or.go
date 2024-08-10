package handlers

import (
	"fmt"

	log "github.com/sirupsen/logrus"
	"github.com/xuri/excelize/v2"
)

func HandleOR() {
	f, err := excelize.OpenFile("or-base.xlsx")
	if err != nil {
		log.Error(err)
		return
	}

	defer func() {
		if err := f.Close(); err != nil {
			log.Error(err)
		}
	}()

	rows, err := f.GetRows("Hoja1")
	if err != nil {
		log.Error(err)
		return
	}

	for i, _ := range rows {
		fmt.Println(i)
	}
}
