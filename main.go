package main

import (
	sap_api_caller "sap-api-integrations-product-master-reads/SAP_API_Caller"
	sap_api_input_reader "sap-api-integrations-product-master-reads/SAP_API_Input_Reader"

	"github.com/latonaio/golang-logging-library/logger"
)

func main() {
	l := logger.NewLogger()
	fr := sap_api_input_reader.NewFileReader()
	inoutSDC := fr.ReadSDC("./Inputs/SDC_Product_Master_Work_Scheduling_sample.json")
	caller := sap_api_caller.NewSAPAPICaller(
		"https://sandbox.api.sap.com/s4hanacloud/sap/opu/odata/sap/", l,
	)

	accepter := inoutSDC.Accepter
	if len(accepter) == 0 || accepter[0] == "All" {
		accepter = []string{
			"Product", "Plant", "MRPArea", "Procurement",
			"WorkScheduling", "SalesPlant",
			"Accounting", "SalesOrganization", "ProductDesc",
		}
	}

	caller.AsyncGetProductMaster(
		inoutSDC.Product.Product,
		inoutSDC.Product.Plant.Plant,
		inoutSDC.Product.Plant.MRPArea.MRPArea,
		inoutSDC.Product.Accounting.ValuationArea,
		inoutSDC.Product.SalesOrganization.ProductSalesOrg,
		inoutSDC.Product.SalesOrganization.ProductDistributionChnl,
		accepter,
	)
}
