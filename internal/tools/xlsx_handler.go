package tools

import (
	"cmd/api/main.go/models"
	"database/sql"
	"fmt"

	"github.com/shopspring/decimal"
	log "github.com/sirupsen/logrus"
	"github.com/xuri/excelize/v2"
)

func XlsxHandler(db *sql.DB) error {

	f, err := excelize.OpenFile("list-base/led-smd.xlsx")
	if err != nil {
		log.Error(err)
		return err
	}

	defer func() {
		if err := f.Close(); err != nil {
			log.Error(err)
		}
	}()

	var sheet string = "Hoja1"
	rows, err := f.GetRows(sheet)
	if err != nil {
		log.Error(err)
		return err
	}

	for i := range rows {
		if i > 0 {
			price, err := decimal.NewFromString(rows[i][2])
			if err != nil {
				log.Error("error: '", err, "' in row: ", i, "\n")
				continue
			}
			product := models.Product{
				Id:     rows[i][0],
				Desc:   rows[i][1],
				Price:  price,
				Subcat: rows[i][3],
				Cat:    rows[i][4],
				Src:    rows[i][5],
				Date:   rows[i][6],
			}

			fmt.Println(product)
		}
	}

	return nil
}
