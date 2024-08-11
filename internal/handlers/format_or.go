package handlers

import (
	"fmt"
	"strconv"
	"time"

	log "github.com/sirupsen/logrus"
	"github.com/xuri/excelize/v2"
)

func HandleOR() {

	// Leer excel a modificar
	f, err := excelize.OpenFile("or-base.xlsx")
	if err != nil {
		log.Error(err)
		return
	}

	// cerrar archivo cuando termina el programa
	defer func() {
		if err := f.Close(); err != nil {
			log.Error(err)
		}
	}()

	//Leer filas

	var sheet string = "Hoja1"
	rows, err := f.GetRows(sheet)
	if err != nil {
		log.Error(err)
		return
	}

	//iterar sobre las filas de precio

	// formula, err := f.GetCellFormula(sheet, "D2")
	// fmt.Println(formula)
	// baby, err := f.GetCellValue(sheet, "E2")
	// fmt.Println(baby == "#N/A")

	for i := range rows {
		if i > 1 {
			//Coord. donde se escribe res BUSCARV
			vlupWriteCoord, err := excelize.CoordinatesToCellName(4, i+1)
			if err != nil {
				log.Error(err)
				return
			}

			//coor. de criterio BUSCARV
			critCoord, err := excelize.CoordinatesToCellName(1, i+1)
			if err != nil {
				log.Error(err)
				return
			}

			//ejecutar BUSCARV
			if err = f.SetCellFormula(sheet, vlupWriteCoord, fmt.Sprintf("=VLOOKUP(%s,Hoja2!A1:C860,3,0)", critCoord)); err != nil {
				log.Error(err)
			}

			//obtener valor calculado de celda BUSCARV

			vlupCells, err := f.CalcCellValue(sheet, vlupWriteCoord)
			if err != nil {
				log.Error(err)
				continue
			}

			//convertir valor a int

			val, err := strconv.Atoi(vlupCells)
			if err != nil {
				log.Error(err)
			}

			//sumar porcentaje de gananacia

			percentage := (val / 100) * 10
			newPrice := val + percentage

			//obtener coordenadas de celdas a aplicar porcentaje de ganancias

			newPriceCoords, err := excelize.CoordinatesToCellName(5, i+1)
			if err != nil {
				log.Error(err)
				return
			}

			//escribir nuevo valor

			if err = f.SetCellFormula(sheet, newPriceCoords, fmt.Sprintf("=DOLLAR(%d, 2)", newPrice)); err != nil {
				log.Error(err)
				return
			}
		}

	}

	now := time.Now()

	if err = f.SetCellValue(sheet, "E2", now.Format("02/01/2006")); err != nil {
		log.Error(err)
	}

	if err := f.RemoveCol(sheet, "D"); err != nil {
		log.Error(err)
		return
	}

	if err = f.SaveAs("NewList.xlsx"); err != nil {
		log.Error(err)
		return
	}
}
