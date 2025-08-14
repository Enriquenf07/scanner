package barcode

type BarcodeSimples struct {
	Code     string
	Datahora string
	Produto  string
}

type BarcodeRequest struct {
	Barcode string
	Produto string
}

type Barcodes struct {
	Barcodes []BarcodeSimples
}