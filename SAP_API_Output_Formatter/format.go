package sap_api_output_formatter

import (
	"encoding/json"
	"sap-api-integrations-product-master-reads/SAP_API_Caller/responses"

	"github.com/latonaio/golang-logging-library/logger"
	"golang.org/x/xerrors"
)

func ConvertToGeneral(raw []byte, l *logger.Logger) ([]General, error) {
	pm := &responses.General{}
	err := json.Unmarshal(raw, pm)
	if err != nil {
		return nil, xerrors.Errorf("cannot convert to General. unmarshal error: %w", err)
	}
	if len(pm.D.Results) == 0 {
		return nil, xerrors.New("Result data is not exist")
	}
	if len(pm.D.Results) > 10 {
		l.Info("raw data has too many Results. %d Results exist. show the first 10 of Results array", len(pm.D.Results))
	}
	general := make([]General, 0, 10)
	for i := 0; i < 10 && i < len(pm.D.Results); i++ {
		data := pm.D.Results[i]
		general = append(general, General{
		Product:             data.Product,
		ProductDescription:  data.ToDescription.Deferred.URI,
		BaseUnit:            data.BaseUnit,
		ValidityStartDate:   data.ValidityStartDate,
		ProductGroup:        data.ProductGroup,
		Division:            data.Division,
		GrossWeight:         data.GrossWeight,
		WeightUnit:          data.WeightUnit,
		SizeOrDimensionText: data.SizeOrDimensionText,
		ProductStandardID:   data.ProductStandardID,
		})
	}

	return general, nil
}

func ConvertToPlant(raw []byte, l *logger.Logger) ([]Plant, error) {
	pm := &responses.Plant{}
	err := json.Unmarshal(raw, pm)
	if err != nil {
		return nil, xerrors.Errorf("cannot convert to Plant. unmarshal error: %w", err)
	}
	if len(pm.D.Results) == 0 {
		return nil, xerrors.New("Result data is not exist")
	}
	if len(pm.D.Results) > 10 {
		l.Info("raw data has too many Results. %d Results exist. show the first 10 of Results array", len(pm.D.Results))
	}
	plant := make([]Plant, 0, 10)
	for i := 0; i < 10 && i < len(pm.D.Results); i++ {
		data := pm.D.Results[i]
		plant = append(plant, Plant{
		Product:                       data.Product,
		Plant:                         data.Plant,
		PurchasingGroup:               data.PurchasingGroup,
		ProductionInvtryManagedLoc:    data.ProductionInvtryManagedLoc,
		AvailabilityCheckType:         data.AvailabilityCheckType,
		ProfitCenter:                  data.ProfitCenter,
		GoodsReceiptDuration:          data.GoodsReceiptDuration,
		MRPType:                       data.MRPType,
		MRPResponsible:                data.MRPResponsible,
		MinimumLotSizeQuantity:        data.MinimumLotSizeQuantity,
		MaximumLotSizeQuantity:        data.MaximumLotSizeQuantity,
		FixedLotSizeQuantity:          data.FixedLotSizeQuantity,
		IsBatchManagementRequired:     data.IsBatchManagementRequired,
		ProcurementType:               data.ProcurementType,
		IsInternalBatchManaged:        data.IsInternalBatchManaged,
		GoodsIssueUnit:                data.GoodsIssueUnit,
		MaterialFreightGroup:          data.MaterialFreightGroup,
		ProductLogisticsHandlingGroup: data.ProductLogisticsHandlingGroup,
		IsMarkedForDeletion:           data.IsMarkedForDeletion,
		})
	}

	return plant, nil
}

func ConvertToMRPArea(raw []byte, l *logger.Logger) ([]MRPArea, error) {
	pm := &responses.MRPArea{}
	err := json.Unmarshal(raw, pm)
	if err != nil {
		return nil, xerrors.Errorf("cannot convert to MRPArea. unmarshal error: %w", err)
	}
	if len(pm.D.Results) == 0 {
		return nil, xerrors.New("Result data is not exist")
	}
	if len(pm.D.Results) > 10 {
		l.Info("raw data has too many Results. %d Results exist. show the first 10 of Results array", len(pm.D.Results))
	}
	mRPArea := make([]MRPArea, 0, 10)
	for i := 0; i < 10 && i < len(pm.D.Results); i++ {
		data := pm.D.Results[i]
		mRPArea = append(mRPArea, MRPArea{
		Product:                       data.Product,
		Plant:                         data.Plant,
		MRPArea:                       data.MRPArea,
		MRPType:                       data.MRPType,
		MRPResponsible:                data.MRPResponsible,
		MRPGroup:                      data.MRPGroup,
		ReorderThresholdQuantity:      data.ReorderThresholdQuantity,
		PlanningTimeFence:             data.PlanningTimeFence,
		LotSizingProcedure:            data.LotSizingProcedure,
		LotSizeRoundingQuantity:       data.LotSizeRoundingQuantity,
		MinimumLotSizeQuantity:        data.MinimumLotSizeQuantity,
		MaximumLotSizeQuantity:        data.MaximumLotSizeQuantity,
		MaximumStockQuantity:          data.MaximumStockQuantity,
		ProcurementSubType:            data.ProcurementSubType,
		DfltStorageLocationExtProcmt:  data.DfltStorageLocationExtProcmt,
		MRPPlanningCalendar:           data.MRPPlanningCalendar,
		SafetyStockQuantity:           data.SafetyStockQuantity,
		SafetyDuration:                data.SafetyDuration,
		FixedLotSizeQuantity:          data.FixedLotSizeQuantity,
		PlannedDeliveryDurationInDays: data.PlannedDeliveryDurationInDays,
		StorageLocation:               data.StorageLocation,
		IsMarkedForDeletion:           data.IsMarkedForDeletion,
		})
	}

	return mRPArea, nil
}

func ConvertToProcurement(raw []byte, l *logger.Logger) ([]Procurement, error) {
	pm := &responses.Procurement{}
	err := json.Unmarshal(raw, pm)
	if err != nil {
		return nil, xerrors.Errorf("cannot convert to Procurement. unmarshal error: %w", err)
	}
	if len(pm.D.Results) == 0 {
		return nil, xerrors.New("Result data is not exist")
	}
	if len(pm.D.Results) > 10 {
		l.Info("raw data has too many Results. %d Results exist. show the first 10 of Results array", len(pm.D.Results))
	}
	procurement := make([]Procurement, 0, 10)
	for i := 0; i < 10 && i < len(pm.D.Results); i++ {
		data := pm.D.Results[i]
		procurement = append(procurement, Procurement{
		Product:                     data.Product,
		Plant:                       data.Plant,
		IsAutoPurOrdCreationAllowed: data.IsAutoPurOrdCreationAllowed,
		IsSourceListRequired:        data.IsSourceListRequired,
		})
	}

	return procurement, nil
}

func ConvertToWorkScheduling(raw []byte, l *logger.Logger) ([]WorkScheduling, error) {
	pm := &responses.WorkScheduling{}
	err := json.Unmarshal(raw, pm)
	if err != nil {
		return nil, xerrors.Errorf("cannot convert to WorkScheduling. unmarshal error: %w", err)
	}
	if len(pm.D.Results) == 0 {
		return nil, xerrors.New("Result data is not exist")
	}
	if len(pm.D.Results) > 10 {
		l.Info("raw data has too many Results. %d Results exist. show the first 10 of Results array", len(pm.D.Results))
	}
	workScheduling := make([]WorkScheduling, 0, 10)
	for i := 0; i < 10 && i < len(pm.D.Results); i++ {
		data := pm.D.Results[i]
		workScheduling = append(workScheduling, WorkScheduling{
		Product:                       data.Product,
		Plant:                         data.Plant,
		ProductionInvtryManagedLoc:    data.ProductionInvtryManagedLoc,
		ProductProcessingTime:         data.ProductProcessingTime,
		ProductionSupervisor:          data.ProductionSupervisor,
		ProductProductionQuantityUnit: data.ProductProductionQuantityUnit,
		ProdnOrderIsBatchRequired:     data.ProdnOrderIsBatchRequired,
		MatlCompIsMarkedForBackflush:  data.MatlCompIsMarkedForBackflush,
		ProductionSchedulingProfile:   data.ProductionSchedulingProfile,
		})
	}

	return workScheduling, nil
}

func ConvertToSalesPlant(raw []byte, l *logger.Logger) ([]SalesPlant, error) {
	pm := &responses.SalesPlant{}
	err := json.Unmarshal(raw, pm)
	if err != nil {
		return nil, xerrors.Errorf("cannot convert to SalesPlant. unmarshal error: %w", err)
	}
	if len(pm.D.Results) == 0 {
		return nil, xerrors.New("Result data is not exist")
	}
	if len(pm.D.Results) > 10 {
		l.Info("raw data has too many Results. %d Results exist. show the first 10 of Results array", len(pm.D.Results))
	}
	salesPlant := make([]SalesPlant, 0, 10)
	for i := 0; i < 10 && i < len(pm.D.Results); i++ {
		data := pm.D.Results[i]
		salesPlant = append(salesPlant, SalesPlant{
		Product:               data.Product,
		Plant:                 data.Plant,
		LoadingGroup:          data.LoadingGroup,
		AvailabilityCheckType: data.AvailabilityCheckType,
		})
	}

	return salesPlant, nil
}

func ConvertToAccounting(raw []byte, l *logger.Logger) ([]Accounting, error) {
	pm := &responses.Accounting{}
	err := json.Unmarshal(raw, pm)
	if err != nil {
		return nil, xerrors.Errorf("cannot convert to Accounting. unmarshal error: %w", err)
	}
	if len(pm.D.Results) == 0 {
		return nil, xerrors.New("Result data is not exist")
	}
	if len(pm.D.Results) > 10 {
		l.Info("raw data has too many Results. %d Results exist. show the first 10 of Results array", len(pm.D.Results))
	}
	accounting := make([]Accounting, 0, 10)
	for i := 0; i < 10 && i < len(pm.D.Results); i++ {
		data := pm.D.Results[i]
		accounting = append(accounting, Accounting{
		Product:             data.Product,
		ValuationArea:       data.ValuationArea,
		ValuationClass:      data.ValuationClass,
		StandardPrice:       data.StandardPrice,
		PriceUnitQty:        data.PriceUnitQty,
		MovingAveragePrice:  data.MovingAveragePrice,
		PriceLastChangeDate: data.PriceLastChangeDate,
		PlannedPrice:        data.PlannedPrice,
		IsMarkedForDeletion: data.IsMarkedForDeletion,
		})
	}

	return accounting, nil
}

func ConvertToSalesOrganization(raw []byte, l *logger.Logger) ([]SalesOrganization, error) {
	pm := &responses.SalesOrganization{}
	err := json.Unmarshal(raw, pm)
	if err != nil {
		return nil, xerrors.Errorf("cannot convert to SalesOrganization. unmarshal error: %w", err)
	}
	if len(pm.D.Results) == 0 {
		return nil, xerrors.New("Result data is not exist")
	}
	if len(pm.D.Results) > 10 {
		l.Info("raw data has too many Results. %d Results exist. show the first 10 of Results array", len(pm.D.Results))
	}
	salesOrganization := make([]SalesOrganization, 0, 10)
	for i := 0; i < 10 && i < len(pm.D.Results); i++ {
		data := pm.D.Results[i]
		salesOrganization = append(salesOrganization, SalesOrganization{
		Product:                        data.Product,
		ProductSalesOrg:                data.ProductSalesOrg,
		ProductDistributionChnl:        data.ProductDistributionChnl,
		SupplyingPlant:                 data.SupplyingPlant,
		PriceSpecificationProductGroup: data.PriceSpecificationProductGroup,
		AccountDetnProductGroup:        data.AccountDetnProductGroup,
		ItemCategoryGroup:              data.ItemCategoryGroup,
		SalesMeasureUnit:               data.SalesMeasureUnit,
		ProductHierarchy:               data.ProductHierarchy,
		IsMarkedForDeletion:            data.IsMarkedForDeletion,
		})
	}

	return salesOrganization, nil
}

func ConvertToProductDesc(raw []byte, l *logger.Logger) ([]ProductDesc, error) {
	pm := &responses.ProductDesc{}
	err := json.Unmarshal(raw, pm)
	if err != nil {
		return nil, xerrors.Errorf("cannot convert to ProductDesc. unmarshal error: %w", err)
	}
	if len(pm.D.Results) == 0 {
		return nil, xerrors.New("Result data is not exist")
	}
	if len(pm.D.Results) > 10 {
		l.Info("raw data has too many Results. %d Results exist. show the first 10 of Results array", len(pm.D.Results))
	}
	productDesc := make([]ProductDesc, 0, 10)
	for i := 0; i < 10 && i < len(pm.D.Results); i++ {
		data := pm.D.Results[i]
		productDesc = append(productDesc, ProductDesc{
		Product:            data.Product,
		Language:           data.Language,
		ProductDescription: data.ProductDescription,
		})
	}

	return productDesc, nil
}
