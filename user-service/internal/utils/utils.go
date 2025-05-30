package utils

import "fmt"

func GenerateEmailFromBarcode(barcode string) string {
	return fmt.Sprintf("%s@astanait.edu.kz", barcode)
}
