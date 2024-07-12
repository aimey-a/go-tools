package gexcel

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/xuri/excelize/v2"
)

var (
	ExcelFileFormat = map[string]string{}
	FileFormat      = []string{".xlsx", ".xls", ".xlsm", ".xlsb", ".xml", ".xlam", ".xltm", ".xla", ".xlam"}
)

func init() {
	for _, v := range FileFormat {
		ExcelFileFormat[v] = v
	}
}

func ReadFolderExcel(str ...string) (xlsxFiles []string) {
	var dir string
	if len(str) > 0 {
		dir = str[0]
	} else {
		exePath, err := os.Executable()
		if err != nil {
			// 处理错误
			fmt.Println("Error:", err)
			return
		}
		dir = filepath.Dir(exePath)
	}
	fmt.Println("Executable Directory:", dir)
	files, _ := os.ReadDir(dir)
	for _, file := range files {
		if !file.IsDir() && ExcelFileFormat[filepath.Ext(file.Name())] != "" {
			xlsxFiles = append(xlsxFiles, dir+"/"+file.Name())
		}
	}
	return
}

func ReadExcelContent(str string) []string {
	var data []string
	f, err := excelize.OpenFile(str)
	if err != nil {
		fmt.Println(err)
		return data
	}
	sheetList := f.GetSheetList()
	if len(sheetList) == 0 {
		return data
	}
	sheet := sheetList[0]
	// 获取 Sheet1 上所有单元格
	rows, err := f.GetRows(sheet)
	if err != nil {
		fmt.Println(err)
		return data
	}
	idx := -1
	for i, row := range rows {
		for j, colCell := range row {
			if i == 0 && colCell == "SENSITIVEWORDS" {
				idx = j
			} else {
				if idx == j && colCell != "" {
					data = append(data, colCell)
				}
			}
		}
	}
	return data
}

func NewExcelFile(name string, data []string) {
	f := excelize.NewFile()
	index, err := f.NewSheet("Sheet1")
	if err != nil {
		fmt.Println(err)
		return
	}
	f.SetCellValue("Sheet1", "A1", "word")
	for i, v := range data {
		f.SetCellValue("Sheet1", fmt.Sprintf("A%v", i+2), v)
	}
	f.SetActiveSheet(index)
	if err := f.SaveAs(name + ".xlsx"); err != nil {
		fmt.Println(err)
	}
}
