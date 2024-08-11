package handlers

import (
	"fmt"
	"strconv"

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

			vlupCells, err := f.CalcCellValue(sheet, vlupWriteCoord)
			if err != nil {
				log.Error(err)
			}

			val, err := strconv.Atoi(vlupCells)
			if err != nil {
				log.Error(err)
			}

			fmt.Println(val * 2)
		}

	}

	if err = f.SaveAs("NewList.xlsx"); err != nil {
		log.Error(err)
		return
	}
}
