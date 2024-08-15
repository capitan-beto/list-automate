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
	f, err := excelize.OpenFile("list-base/or-base.xlsx")
	if err != nil {
		log.Error(err)
		return
	}

	// cerrar archivo cuando termina el programa
	defer func() {
		if err := f.Close(); err != nil {
			log.Error(err)
		}
		fmt.Println("Archivo guardado")
	}()

	//Leer filas

	var sheet string = "precios-julio"
	rows, err := f.GetRows(sheet)
	if err != nil {
		log.Error(err)
		return
	}

	// formula, err := f.GetCellFormula(sheet, "D2")  //Con este metodo podemos obtener funcion a utilizar
	// fmt.Println(formula)

	//iterar sobre las filas de precio

	for i := range rows {
		if i > 1 {

			// BUSCARV recibe tres argumentos (VB, RB, C, Coincidencia)
			// VB: Valor de referencia (El codigo de producto al que le queremos actualizar precio).
			// RB: Rango de busqueda, donde se encuentra el valor de busqueda (celdas de la Hoja2)
			// C: Columna dentro del rango de busqueda (columna de precio)
			// Coincidencia: VERDADERO o FALSO (resultado aprox o exacto)

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
			err = f.SetCellFormula(sheet, vlupWriteCoord,
				fmt.Sprintf("=VLOOKUP(%s,Hoja2!A1:C860,3,0)", critCoord))
			if err != nil {
				log.Error(err)
			}

			//obtener valor calculado de celda BUSCARV

			vlupCells, err := f.CalcCellValue(sheet, vlupWriteCoord)
			if err != nil {
				log.Error("error: '", err, "' in row: ", i, "\n")
				continue
			}

			//convertir valor a int

			val, err := strconv.Atoi(vlupCells)
			if err != nil {
				log.Error(err)
				return
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

			if err = f.SetCellFormula(sheet, newPriceCoords,
				fmt.Sprintf("=DOLLAR(%d, 2)", newPrice)); err != nil {
				log.Error(err)
				return
			}
		}

		fmt.Printf("Row %d processed \n", i)

	}

	fmt.Println("All rows processed!")

	now := time.Now()

	if err = f.SetCellValue(sheet, "E2", now.Format("02/01/2006")); err != nil {
		log.Error(err)
	}

	if err := f.RemoveCol(sheet, "D"); err != nil {
		log.Error(err)
		return
	}

	if err = f.SaveAs(fmt.Sprintf("lists/lista-or-%s.xlsx", now.Format("02-01"))); err != nil {
		log.Error(err)
		return
	}
}
