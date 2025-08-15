package excel

import (
	"context"
	"fmt"
	"scanner-backend/barcode"

	"github.com/xuri/excelize/v2"
)

func Create(ctx context.Context) error {
	filename := "planilha.xlsx"
	f := excelize.NewFile()
	f.NewSheet("Plan1")
	f.SetActiveSheet(0)

	data, err := barcode.GetAll(ctx)
	if err != nil {
		return err
	}
	f.SetCellValue("Plan1", "A1", "Descrição")
	f.SetCellValue("Plan1", "B1", "Código")
	linha := 2
	for _, item := range data {
		f.SetCellValue("Plan1", fmt.Sprintf("A%d", linha), item.Produto)
		f.SetCellValue("Plan1", fmt.Sprintf("B%d", linha), item.Code)
		linha++
	}

	if err := f.SaveAs("data/" + filename); err != nil {
		fmt.Println("Erro ao salvar:", err)
	} else {
		fmt.Println("Arquivo Excel sobrescrito com sucesso!")
	}
	return nil
}
