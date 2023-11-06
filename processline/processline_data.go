package processline

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"strings"
)

// Function to extract data fields from the line
func extractData(line string, start, end int) string {
	fields := strings.Fields(line[start:end])
	result := strings.Join(fields, " ")
	return result
}

// InsertJSONData inserts JSON data into the database.
func InsertJSONData(db *sql.DB, tableName, columnName, jsonData string) error {
	id := 1 // Insert a row into hot_file table and get the id
	// Prepare the SQL statement with a placeholder for JSON data
	insertSQL := fmt.Sprintf("INSERT INTO %s (hot_file_id, %s) VALUES ($1, $2)", tableName, columnName)
	fmt.Println(insertSQL)
	fmt.Println(jsonData)
	// Execute the SQL statement
	_, err := db.Exec(insertSQL, id, jsonData)

	fmt.Println("JSON data inserted successfully.")

	return err
}

// ProcessLine processes a line of data based on its record type.
func ProcessLine(db *sql.DB, id int, line string) {
	// Determine the record type based on the first three characters (Standard Message Identifier)
	recordType := line[:3]

	switch recordType {
	case "BFH":
		// File Header Record
		data := FileHeaderBFH01{
			SequenceNumber:             extractData(line, 3, 11),
			StandardNumericQualifier:   extractData(line, 11, 13),
			BSPIdentifier:              extractData(line, 13, 16),
			TicketingAirlineCodeNumber: extractData(line, 16, 19),
			HandbookRevisionNumber:     extractData(line, 19, 22),
			TestorProductionStatus:     extractData(line, 22, 26),
			ProcessingDate:             extractData(line, 26, 32),
			ProcessingTime:             extractData(line, 32, 36),
			ISOCountryCode:             extractData(line, 36, 38),
			FileSequenceNumber:         extractData(line, 38, 44),
		}
		// Marshal the struct to JSON
		jsonData, err := json.Marshal(data)
		if err != nil {
			log.Fatal(err)
		}
		InsertJSONData(db, "hot_file_header", "file_header_record", string(jsonData))
		fmt.Println("%%%%%%%%%%%%%%%%%%")
		fmt.Println("BFH:", string(jsonData))

	case "BCH":
		data := BillingCycleHeaderBCH02{
			SequenceNumber:            extractData(line, 3, 11),
			StandardNumericQualifier:  extractData(line, 11, 13),
			ProcessingDateIdentifier:  extractData(line, 13, 16),
			ProcessingCycleIdentifier: extractData(line, 16, 17),
			BillingAnalysisEndingDate: extractData(line, 17, 23),
			DynamicRunIdentifier:      extractData(line, 23, 24),
			HotReportingEndDate:       extractData(line, 24, 30),
		}
		jsonData, err := json.Marshal(data)
		if err != nil {
			log.Fatal(err)
		}
		InsertJSONData(db, "hot_file_header", "billing_analysis_header", string(jsonData))
		fmt.Println("BCH:", string(jsonData))
	case "BOH":
		//BOH03 (Reporting Agent) Office Header Record
		data := OfficeHeaderRecordBOH03{
			SequenceNumber:             extractData(line, 3, 11),
			StandardNumericQualifier:   extractData(line, 11, 13),
			AgentNumericCode:           extractData(line, 13, 21),
			RemittancePeriodEndingDate: extractData(line, 21, 27),
			CurrencyType:               extractData(line, 27, 31),
			MultiLocationIdentifier:    extractData(line, 31, 34),
		}

		jsonData, err := json.Marshal(data)
		if err != nil {
			log.Fatal(err)
		}
		InsertJSONData(db, "hot_file_header", "office_header", string(jsonData))
		fmt.Println("BOH:", string(jsonData))
		//TRANSACTION RECORDS
	case "BKT":
		// BKT06 Transaction Header Record
		data := TransactionHeaderRecordBKT06{
			SequenceNumber:                    extractData(line, 3, 11),
			StandardNumericQualifier:          extractData(line, 11, 13),
			TransactionNumber:                 extractData(line, 13, 19),
			NetReportingIndicator:             extractData(line, 19, 21),
			TransactionRecordCounter:          extractData(line, 21, 24),
			TicketingAirlineCodeNumber:        extractData(line, 24, 27),
			CommercialAgreementReference:      extractData(line, 27, 37),
			CustomerFileReference:             extractData(line, 37, 64),
			ReportingSystemIdentifier:         extractData(line, 66, 68),
			SettlementAuthorisationCode:       extractData(line, 68, 82),
			DataInputStatusIndicator:          extractData(line, 82, 83),
			NetReportingMethodIndicator:       extractData(line, 83, 84),
			NetReportingCalculationType:       extractData(line, 84, 85),
			AutomatedRepricingEngineIndicator: extractData(line, 85, 86),
		}

		jsonData, err := json.Marshal(data)
		if err != nil {
			log.Fatal(err)
		}
		InsertJSONData(db, "hot_file_transaction_records", "transaction_header", string(jsonData))
		fmt.Println("BOH:", string(jsonData))
	case "BKS":
		//BKS24 Ticket/Document Identification
		//BKS30 STD/Document Amounts
		//BKS31 Coupon Tax Information
		//BKS39 Commission
		//BKS42 Tax on Commission
		//BKS45 Related Ticket/Document Information
		//BKS46 Qualifying Issue Information for Sales Transactions
		//BKS47 Netting Values

		ProcessLineBKS(db, line)

	case "BKI":
		//BKI61 Unticketed Point Information
		//BKI62 Additional Itinerary Data
		//BKI63 Itinerary Data Segment

		ProcessLineBKI(db, line)

	case "BAR":

		//BAR64 Document Amounts
		//BAR65 Additional Information–Passenger
		//BAR66 Additional Information–Form of Payment
		//BAR67 Additional Tax Information
		ProcessLineBAR(db, line)

	case "BMD":
		//BMD75 Electronic Miscellaneous Document Coupon Detail
		//BMD76 Electronic Miscellaneous Document Remarks
		ProcessLineBMD(db, line)

	case "BKF":
		// BKF81 Fare Calculation Record
		data := FareCalculationRecordBKF81{
			SequenceNumber:                extractData(line, 3, 11),
			StandardNumericQualifier:      extractData(line, 11, 13),
			DateofIssue:                   extractData(line, 13, 19),
			TransactionNumber:             extractData(line, 19, 25),
			TicketDocumentNumber:          extractData(line, 25, 39),
			CheckDigit:                    extractData(line, 39, 40),
			FareCalculationSequenceNumber: extractData(line, 40, 41),
			FareCalculationArea:           extractData(line, 41, 128),
		}

		jsonData, err := json.Marshal(data)
		if err != nil {
			log.Fatal(err)
		}
		InsertJSONData(db, "hot_file_transaction_records", "fare_calculation", string(jsonData))
		fmt.Println("BOH:", string(jsonData))
	case "BCC":
		// BCC82 Additional Card Information Record
		data := AdditionalCardInformationRecorBCC82{
			SequenceNumber:                     extractData(line, 3, 11),
			StandardNumericQualifier:           extractData(line, 11, 13),
			DateofIssue:                        extractData(line, 13, 19),
			TransactionNumber:                  extractData(line, 19, 25),
			FormOfPaymentType:                  extractData(line, 25, 35),
			FormOfPaymentTransactionIdentifier: extractData(line, 35, 60),
		}

		jsonData, err := json.Marshal(data)
		if err != nil {
			log.Fatal(err)
		}
		InsertJSONData(db, "hot_file_transaction_records", "additional_card_information", string(jsonData))
		fmt.Println("BOH:", string(jsonData))
	case "BCX":
		// BCX83 3DS Authentication and Additional Card Payment Information Record
		data := THREEDSAuthenticationandAdditionalCardPaymentInformationRecordBCX83{
			SequenceNumber:                   extractData(line, 3, 11),
			StandardNumericQualifier:         extractData(line, 11, 13),
			DateofIssue:                      extractData(line, 13, 19),
			TransactionNumber:                extractData(line, 19, 25),
			FormOfPaymentType:                extractData(line, 25, 35),
			CardAuthenticationSequenceNumber: extractData(line, 35, 37),
			ThreeDSecureAuthenticationAndAdditionalCardPaymentInformation: extractData(line, 13, 136),
		}

		jsonData, err := json.Marshal(data)
		if err != nil {
			log.Fatal(err)
		}
		InsertJSONData(db, "hot_file_transaction_records", "3ds_authentication_card_payment_information", string(jsonData))
		fmt.Println("BOH:", string(jsonData))
	case "BKP":
		// BKP84 Form of Payment Record
		data := FormOfPaymentRecordBKP84{
			SequenceNumber:              extractData(line, 3, 11),
			StandardNumericQualifier:    extractData(line, 11, 13),
			DateofIssue:                 extractData(line, 13, 19),
			TransactionNumber:           extractData(line, 19, 25),
			FormOfPaymentType:           extractData(line, 25, 35),
			FormOfPaymentAmount:         extractData(line, 35, 46),
			FormOfPaymentAccountNumber:  extractData(line, 46, 65),
			ExpiryDate:                  extractData(line, 65, 69),
			ExtendedPaymentCode:         extractData(line, 69, 71),
			ApprovalCode:                extractData(line, 71, 77),
			InvoiceNumber:               extractData(line, 77, 91),
			InvoiceDate:                 extractData(line, 91, 97),
			RemittanceAmount:            extractData(line, 97, 108),
			CardVerificationValueResult: extractData(line, 108, 109),
			ReservedSpace:               extractData(line, 109, 132),
			CurrencyType:                extractData(line, 132, 136),
		}

		jsonData, err := json.Marshal(data)
		if err != nil {
			log.Fatal(err)
		}
		InsertJSONData(db, "hot_file_transaction_records", "form_of_payment", string(jsonData))
		fmt.Println("BOH:", string(jsonData))
		//TOTAL RECORDS
	case "BOT":

		//BOT93 Office Subtotals per Transaction Code and Currency Type Record
		//BOT94 Office Totals per Currency Type Record
		ProcessLineBOT(db, line)
	case "BCT":
		// BCT95 Billing Analysis (Cycle) Totals per Currency Type Record
		data := BillingAnalysisTotalsPerCurrencyTypeRecordBCT95{
			SequenceNumber:                   extractData(line, 3, 11),
			StandardNumericQualifier:         extractData(line, 11, 13),
			ProcessingDateIdentifier:         extractData(line, 13, 16),
			ProcessingCycleIdentifier:        extractData(line, 16, 17),
			OfficeCount:                      extractData(line, 17, 22),
			GrossValueAmount:                 extractData(line, 22, 37),
			TotalRemittanceAmount:            extractData(line, 37, 52),
			TotalCommissionValueAmount:       extractData(line, 52, 67),
			TotalTaxOrMiscellaneousFeeAmount: extractData(line, 67, 82),
			TotalTaxOnCommissionAmount:       extractData(line, 82, 97),
			ReservedSpace:                    extractData(line, 97, 132),
			CurrencyType:                     extractData(line, 132, 136),
		}

		jsonData, err := json.Marshal(data)
		if err != nil {
			log.Fatal(err)
		}
		InsertJSONData(db, "hot_file_total_records", "billing_analysis_totals_per_currency_type", string(jsonData))
		fmt.Println("BOH:", string(jsonData))

	case "BFT":

		//BFT99 File Totals per Currency Type Record
		data := FileTotalsPerCurrencyTypeRecordBFT99{
			SequenceNumber:                   extractData(line, 3, 11),
			StandardNumericQualifier:         extractData(line, 11, 13),
			BSPIdentifier:                    extractData(line, 13, 16),
			OfficeCount:                      extractData(line, 16, 21),
			GrossValueAmount:                 extractData(line, 21, 36),
			TotalRemittanceAmount:            extractData(line, 36, 51),
			TotalCommissionValueAmount:       extractData(line, 51, 66),
			TotalTaxOrMiscellaneousFeeAmount: extractData(line, 66, 81),
			TotalTaxOnCommissionAmount:       extractData(line, 81, 96),
			ReservedSpace:                    extractData(line, 96, 132),
			CurrencyType:                     extractData(line, 132, 136),
		}

		jsonData, err := json.Marshal(data)
		if err != nil {
			log.Fatal(err)
		}
		InsertJSONData(db, "hot_file_total_records", "file_totals_per_currency_type", string(jsonData))
		fmt.Println("BOH:", string(jsonData))

	default:
		fmt.Println("Unknown record type:", recordType)
	}
}

// ProcessLineBKS processes a line of data based on its record type.
func ProcessLineBKS(db *sql.DB, line string) {
	// Determine the BKS record type based on the characters(12 TO 13) (Standard Numeric Qualifier)
	recordBKSType := extractData(line, 11, 13)
	fmt.Println("***********************")

	fmt.Println("recordBKSType:", recordBKSType)
	switch recordBKSType {
	case "24":
		//BKS24 Ticket/Document Identification Record
		data := TicketIdentificationRecordBKS24{
			SequenceNumber:                    extractData(line, 3, 11),
			StandardNumericQualifier:          extractData(line, 11, 13),
			DateofIssue:                       extractData(line, 13, 19),
			TransactionNumber:                 extractData(line, 19, 25),
			TicketDocumentNumber:              extractData(line, 25, 39),
			CheckDigit:                        extractData(line, 39, 40),
			CouponUseIndicator:                extractData(line, 40, 44),
			ConjunctionTicketIndicator:        extractData(line, 44, 47),
			AgentNumericCode:                  extractData(line, 47, 55),
			ReasonForIssuanceCode:             extractData(line, 55, 56),
			TourCode:                          extractData(line, 56, 71),
			TransactionCode:                   extractData(line, 71, 75),
			TrueOriginDestinationCityCodes:    extractData(line, 75, 85),
			PNRReferenceAndAirlineData:        extractData(line, 85, 98),
			TimeOfIssue:                       extractData(line, 98, 102),
			JourneyTournaroundAirportCityCode: extractData(line, 102, 107),
		}
		// Marshal the struct to JSON
		jsonData, err := json.Marshal(data)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("@@@@@@@@@@@@@@@@@")

		InsertJSONData(db, "hot_file_transaction_records", "ticket_identification", string(jsonData))
		fmt.Println("BKS24:", string(jsonData))

	case "30":
		//BKS30 STD/Document Amounts Record

		data := STDOrDocumentAmountsRecordBKS30{
			SequenceNumber:               extractData(line, 3, 11),
			StandardNumericQualifier:     extractData(line, 11, 13),
			DateofIssue:                  extractData(line, 13, 19),
			TransactionNumber:            extractData(line, 19, 25),
			TicketDocumentNumber:         extractData(line, 25, 39),
			CheckDigit:                   extractData(line, 39, 40),
			CommissionableAmount:         extractData(line, 40, 51),
			NetFareAmount:                extractData(line, 51, 62),
			TaxOrMiscellaneousFeeType1:   extractData(line, 62, 70),
			TaxOrMiscellaneousFeeAmount1: extractData(line, 70, 81),
			TaxOrMiscellaneousFeeType2:   extractData(line, 81, 89),
			TaxOrMiscellaneousFeeAmount2: extractData(line, 89, 100),
			TaxOrMiscellaneousFeeType3:   extractData(line, 100, 108),
			TaxOrMiscellaneousFeeAmount3: extractData(line, 108, 119),
			TicketOrDocumentAmount:       extractData(line, 119, 130),
			ReservedSpace:                extractData(line, 130, 132),
			CurrencyType:                 extractData(line, 132, 136),
		}
		jsonData, err := json.Marshal(data)
		if err != nil {
			log.Fatal(err)
		}
		InsertJSONData(db, "hot_file_transaction_records", "std_amounts", string(jsonData))
		fmt.Println("BCH:", string(jsonData))
	case "31":
		//BKS31 Coupon Tax Information Record
		data := CouponTaxInformationRecordBKS31{
			SequenceNumber:             extractData(line, 3, 11),
			StandardNumericQualifier:   extractData(line, 11, 13),
			DateofIssue:                extractData(line, 13, 19),
			TransactionNumber:          extractData(line, 19, 25),
			TicketDocumentNumber:       extractData(line, 25, 39),
			CheckDigit:                 extractData(line, 39, 40),
			SegmentIdentifier1:         extractData(line, 40, 41),
			CouponTaxAirportCode1:      extractData(line, 41, 46),
			SegmentTaxAirportCode1:     extractData(line, 46, 52),
			CouponTaxCode1:             extractData(line, 52, 54),
			CouponTaxType1:             extractData(line, 54, 57),
			CouponTaxReportedAmount1:   extractData(line, 57, 68),
			CouponTaxCurrencyType1:     extractData(line, 68, 72),
			CouponTaxApplicableAmount1: extractData(line, 72, 83),
			SegmentIdentifier2:         extractData(line, 83, 84),
			CouponTaxAirportCode2:      extractData(line, 84, 89),
			SegmentTaxAirportCode2:     extractData(line, 89, 95),
			CouponTaxCode2:             extractData(line, 95, 97),
			CouponTaxType2:             extractData(line, 97, 100),
			CouponTaxReportedAmount2:   extractData(line, 100, 111),
			CouponTaxCurrencyType2:     extractData(line, 111, 115),
			CouponTaxApplicableAmount2: extractData(line, 115, 126),
			ReservedSpace:              extractData(line, 126, 132),
			CurrencyType:               extractData(line, 132, 136),
		}

		jsonData, err := json.Marshal(data)
		if err != nil {
			log.Fatal(err)
		}
		InsertJSONData(db, "hot_file_transaction_records", "coupon_tax_information", string(jsonData))
		fmt.Println("BOH:", string(jsonData))

	case "39":
		//BKS39 Commission Record
		data := CommissionRecordBKS39{
			SequenceNumber:                          extractData(line, 3, 11),
			StandardNumericQualifier:                extractData(line, 11, 13),
			DateofIssue:                             extractData(line, 13, 20),
			TransactionNumber:                       extractData(line, 20, 25),
			TicketDocumentNumber:                    extractData(line, 25, 39),
			CheckDigit:                              extractData(line, 39, 40),
			StatisticalCode:                         extractData(line, 40, 43),
			CommissionType:                          extractData(line, 43, 49),
			CommissionRate:                          extractData(line, 49, 54),
			CommissionAmount:                        extractData(line, 54, 65),
			SupplementaryType:                       extractData(line, 65, 71),
			SupplementaryRate:                       extractData(line, 71, 76),
			SupplementaryAmount:                     extractData(line, 76, 87),
			EffectiveCommissionRate:                 extractData(line, 87, 92),
			EffectiveCommissionAmount:               extractData(line, 92, 103),
			AmountPaidbyCustomer:                    extractData(line, 103, 114),
			RoutingDomesticOrInternationalIndicator: extractData(line, 114, 115),
			CommissionControlAdjustmentIndicator:    extractData(line, 115, 116),
			ReservedSpace:                           extractData(line, 116, 132),
			CurrencyType:                            extractData(line, 132, 136),
		}

		jsonData, err := json.Marshal(data)
		if err != nil {
			log.Fatal(err)
		}
		InsertJSONData(db, "hot_file_transaction_records", "commission", string(jsonData))
		fmt.Println("BOH:", string(jsonData))
	case "42":
		//BKS42 Tax on Commission Record
		data := TaxOnCommissionRecordBKS42{
			SequenceNumber:           extractData(line, 3, 11),
			StandardNumericQualifier: extractData(line, 11, 13),
			DateofIssue:              extractData(line, 13, 19),
			TransactionNumber:        extractData(line, 19, 25),
			TicketDocumentNumber:     extractData(line, 25, 39),
			CheckDigit:               extractData(line, 39, 40),
			TaxonCommissionType1:     extractData(line, 40, 46),
			TaxonCommissionAmount1:   extractData(line, 46, 57),
			TaxonCommissionType2:     extractData(line, 57, 63),
			TaxonCommissionAmount2:   extractData(line, 63, 74),
			TaxonCommissionType3:     extractData(line, 74, 80),
			TaxonCommissionAmount3:   extractData(line, 80, 91),
			TaxonCommissionType4:     extractData(line, 91, 97),
			TaxonCommissionAmount4:   extractData(line, 97, 108),
			ReservedSpace:            extractData(line, 108, 132),
			CurrencyType:             extractData(line, 132, 136),
		}

		jsonData, err := json.Marshal(data)
		if err != nil {
			log.Fatal(err)
		}
		InsertJSONData(db, "hot_file_transaction_records", "tax_on_commission", string(jsonData))
		fmt.Println("BOH:", string(jsonData))

	case "45":
		//BKS45 Related Ticket/Document Information Record
		data := RelatedTicketOrDocumentInformationRecordBKS45{
			SequenceNumber:                extractData(line, 3, 11),
			StandardNumericQualifier:      extractData(line, 11, 13),
			RemittancePeriodEndingDate:    extractData(line, 13, 19),
			TransactionNumber:             extractData(line, 19, 25),
			RelatedTicketOrDocumentNumber: extractData(line, 25, 39),
			CheckDigit:                    extractData(line, 37, 40),
			WaiverCode:                    extractData(line, 40, 54),
			ReasonForMemoIssuanceCode:     extractData(line, 53, 59),
			RelatedTicketOrDocumentCouponNumberIdentifier: extractData(line, 59, 63),
			DateOfIssueRelatedDocument:                    extractData(line, 63, 69),
		}

		jsonData, err := json.Marshal(data)
		if err != nil {
			log.Fatal(err)
		}
		InsertJSONData(db, "hot_file_transaction_records", "related_ticket_information", string(jsonData))
		fmt.Println("BOH:", string(jsonData))
	case "46":
		//BKS46 Qualifying Issue Information for Sales Transactions Record
		data := QualifyingIssueInformationForSalesTransactionsRecordBKS46{
			SequenceNumber:                          extractData(line, 3, 11),
			StandardNumericQualifier:                extractData(line, 11, 13),
			DateofIssue:                             extractData(line, 13, 19),
			TransactionNumber:                       extractData(line, 19, 25),
			TicketDocumentNumber:                    extractData(line, 25, 39),
			CheckDigit:                              extractData(line, 39, 40),
			OriginalIssueTicketOrDocumentNumber:     extractData(line, 40, 54),
			OriginalIssueLocationCityCode:           extractData(line, 54, 57),
			OriginalIssueDateDDMMMYY:                extractData(line, 57, 64),
			OriginalIssueAgentNumericCodeIATANumber: extractData(line, 64, 72),
			EndorsementsOrRestrictions:              extractData(line, 72, 121),
		}

		jsonData, err := json.Marshal(data)
		if err != nil {
			log.Fatal(err)
		}
		InsertJSONData(db, "hot_file_transaction_records", "qualifying_issue_information_for_sales_transactions", string(jsonData))
		fmt.Println("BOH:", string(jsonData))
	case "47":
		//BKS47 Netting Values Record
		data := NettingValuesRecordBKS47{
			SequenceNumber:           extractData(line, 3, 11),
			StandardNumericQualifier: extractData(line, 11, 13),
			DateofIssue:              extractData(line, 13, 19),
			TransactionNumber:        extractData(line, 19, 25),
			TicketDocumentNumber:     extractData(line, 25, 39),
			CheckDigit:               extractData(line, 39, 40),
			NettingType1:             extractData(line, 40, 41),
			NettingCode1:             extractData(line, 41, 49),
			NettingAmount1:           extractData(line, 49, 60),
			NettingType2:             extractData(line, 60, 61),
			NettingCode2:             extractData(line, 61, 69),
			NettingAmount2:           extractData(line, 69, 80),
			NettingType3:             extractData(line, 80, 81),
			NettingCode3:             extractData(line, 81, 89),
			NettingAmount3:           extractData(line, 89, 100),
			NettingType4:             extractData(line, 100, 101),
			NettingCode4:             extractData(line, 101, 109),
			NettingAmount4:           extractData(line, 109, 120),
			ReservedSpace:            extractData(line, 120, 132),
			CurrencyType:             extractData(line, 132, 136),
		}

		jsonData, err := json.Marshal(data)
		if err != nil {
			log.Fatal(err)
		}
		InsertJSONData(db, "hot_file_transaction_records", "netting_values", string(jsonData))
		fmt.Println("BOH:", string(jsonData))

	default:
		fmt.Println("Unknown record type:", recordBKSType)
	}
}

// ProcessLineBKI processes a line of data based on its record type.
func ProcessLineBKI(db *sql.DB, line string) {
	// Determine the BKI record type based on the characters(12 TO 13) (Standard Numeric Qualifier)
	recordBKIType := extractData(line, 11, 13)
	switch recordBKIType {
	case "61":
		//BKS61 Unticketed Point Information Record
		data := UnticketedPointInformationRecordBKI61{
			SequenceNumber:                        extractData(line, 3, 11),
			StandardNumericQualifier:              extractData(line, 11, 13),
			DateofIssue:                           extractData(line, 13, 19),
			TransactionNumber:                     extractData(line, 19, 25),
			TicketDocumentNumber:                  extractData(line, 25, 39),
			CheckDigit:                            extractData(line, 39, 40),
			SegmentIdentifier:                     extractData(line, 40, 41),
			UnticketedPointAirportOrCityCode:      extractData(line, 41, 46),
			UnticketedPointDateOfArrival:          extractData(line, 46, 53),
			UnticketedPointLocalTimeOfArrival:     extractData(line, 53, 58),
			UnticketedPointDateOfDeparture:        extractData(line, 58, 65),
			UnticketedPointLocalTimeOfDeparture:   extractData(line, 65, 70),
			UnticketedPointDepartureEquipmentCode: extractData(line, 70, 73),
		}
		// Marshal the struct to JSON
		jsonData, err := json.Marshal(data)
		if err != nil {
			log.Fatal(err)
		}
		InsertJSONData(db, "hot_file_transaction_records", "unticketed_point_information", string(jsonData))
		fmt.Println("BFH:", string(jsonData))

	case "62":
		//BKI62 Additional Itinerary Data Record
		data := AdditionalItineraryDataRecordBKI62{
			SequenceNumber:               extractData(line, 3, 11),
			StandardNumericQualifier:     extractData(line, 11, 13),
			DateofIssue:                  extractData(line, 13, 19),
			TransactionNumber:            extractData(line, 19, 25),
			TicketDocumentNumber:         extractData(line, 25, 39),
			CheckDigit:                   extractData(line, 39, 40),
			SegmentIdentifier:            extractData(line, 40, 41),
			OriginAirportOrCityCode:      extractData(line, 41, 46),
			FlightDepartureDate:          extractData(line, 46, 53),
			FlightDepartureTime:          extractData(line, 53, 58),
			FlightDepartureTerminal:      extractData(line, 58, 63),
			DestinationAirportOrCityCode: extractData(line, 63, 68),
			FlightArrivalDate:            extractData(line, 68, 75),
			FlightArrivalTime:            extractData(line, 75, 80),
			FlightArrivalTerminal:        extractData(line, 80, 85),
		}
		// Marshal the struct to JSON
		jsonData, err := json.Marshal(data)
		if err != nil {
			log.Fatal(err)
		}
		InsertJSONData(db, "hot_file_transaction_records", "additional_itinerary_data", string(jsonData))
		fmt.Println("BFH:", string(jsonData))

	case "63":
		//BKI63 Itinerary Data Segment Record
		data := ItineraryDataSegmentRecordBKI63{
			SequenceNumber:                       extractData(line, 3, 11),
			StandardNumericQualifier:             extractData(line, 11, 13),
			DateofIssue:                          extractData(line, 13, 19),
			TransactionNumber:                    extractData(line, 19, 25),
			TicketDocumentNumber:                 extractData(line, 25, 39),
			CheckDigit:                           extractData(line, 39, 40),
			SegmentIdentifier:                    extractData(line, 40, 41),
			StopoverCode:                         extractData(line, 41, 42),
			NotValidBeforeDate:                   extractData(line, 42, 47),
			NotValidAfterDate:                    extractData(line, 47, 52),
			OriginAirportOrCityCode:              extractData(line, 52, 57),
			DestinationAirportOrCityCode:         extractData(line, 57, 62),
			Carrier:                              extractData(line, 62, 65),
			SoldPassengerCabin:                   extractData(line, 65, 66),
			FlightNumber:                         extractData(line, 66, 71),
			ReservationBookingDesignator:         extractData(line, 71, 73),
			FlightDepartureDate:                  extractData(line, 73, 80),
			FlightDepartureTime:                  extractData(line, 80, 85),
			FlightBookingStatus:                  extractData(line, 85, 87),
			BaggageAllowance:                     extractData(line, 87, 90),
			FareBasisOrTicketDesignator:          extractData(line, 90, 105),
			FrequentFlyerReference:               extractData(line, 105, 125),
			FareComponentPricedPassengerTypeCode: extractData(line, 125, 128),
			ThroughOrChangeOfGaugeIndicator:      extractData(line, 128, 129),
			EquipmentCode:                        extractData(line, 129, 132),
		}
		// Marshal the struct to JSON
		jsonData, err := json.Marshal(data)
		if err != nil {
			log.Fatal(err)
		}
		InsertJSONData(db, "hot_file_transaction_records", "itinerary_data_segment", string(jsonData))
		fmt.Println("BFH:", string(jsonData))

	default:
		fmt.Println("Unknown record type:", recordBKIType)
	}
}

// ProcessLineBAR processes a line of data based on its record type.
func ProcessLineBAR(db *sql.DB, line string) {
	// Determine the BAR record type based on the characters(12 TO 13) (Standard Numeric Qualifier)
	recordBARType := extractData(line, 11, 13)
	switch recordBARType {
	case "64":
		//BAR64 Document Amounts Record
		data := DocumentAmountsRecordBAR64{
			SequenceNumber:           extractData(line, 3, 11),
			StandardNumericQualifier: extractData(line, 11, 13),
			DateofIssue:              extractData(line, 13, 19),
			TransactionNumber:        extractData(line, 19, 25),
			TicketDocumentNumber:     extractData(line, 25, 39),
			CheckDigit:               extractData(line, 39, 40),
			Fare:                     extractData(line, 40, 52),
			TicketingModeIndicator:   extractData(line, 52, 53),
			EquivalentFarePaid:       extractData(line, 53, 65),
			Total:                    extractData(line, 65, 77),
			ServicingAirlineOrSystemProviderIdentifier:         extractData(line, 77, 81),
			FareCalculationModeIndicator:                       extractData(line, 81, 82),
			BookingAgentIdentification:                         extractData(line, 82, 88),
			BookingEntityOutletType:                            extractData(line, 88, 89),
			AirlineIssFareCalculationPricingIndicatoruingAgent: extractData(line, 89, 90),
			AirlineIssuingAgent:                                extractData(line, 90, 98),
		}
		// Marshal the struct to JSON
		jsonData, err := json.Marshal(data)
		if err != nil {
			log.Fatal(err)
		}
		InsertJSONData(db, "hot_file_transaction_records", "document_amounts", string(jsonData))
		fmt.Println("BFH:", string(jsonData))
	case "65":
		//BAR65 Additional Information–Passenger Record
		data := AdditionalInformationPassengerRecordBAR65{
			SequenceNumber:           extractData(line, 3, 11),
			StandardNumericQualifier: extractData(line, 11, 13),
			DateofIssue:              extractData(line, 13, 19),
			TransactionNumber:        extractData(line, 19, 25),
			TicketDocumentNumber:     extractData(line, 25, 39),
			CheckDigit:               extractData(line, 39, 40),
			PassengerName:            extractData(line, 40, 89),
			PassengerSpecificData:    extractData(line, 89, 118),
			DateOfBirth:              extractData(line, 118, 125),
			PassengerTypeCode:        extractData(line, 125, 128),
		}
		// Marshal the struct to JSON
		jsonData, err := json.Marshal(data)
		if err != nil {
			log.Fatal(err)
		}
		InsertJSONData(db, "hot_file_transaction_records", "additional_information_passenger", string(jsonData))
		fmt.Println("BFH:", string(jsonData))

	case "66":
		//BAR66 Additional Information–Form of Payment Record
		data := AdditionalInformationFormOfPaymentRecordBAR66{
			SequenceNumber:              extractData(line, 3, 11),
			StandardNumericQualifier:    extractData(line, 11, 13),
			DateofIssue:                 extractData(line, 13, 19),
			TransactionNumber:           extractData(line, 19, 25),
			TicketDocumentNumber:        extractData(line, 25, 39),
			CheckDigit:                  extractData(line, 39, 40),
			FormOfPaymentSequenceNumber: extractData(line, 40, 41),
			FormOfPaymentInformation:    extractData(line, 41, 91),
		}
		// Marshal the struct to JSON
		jsonData, err := json.Marshal(data)
		if err != nil {
			log.Fatal(err)
		}
		InsertJSONData(db, "hot_file_transaction_records", "additional_information_form_of_payment", string(jsonData))
		fmt.Println("BFH:", string(jsonData))

	case "67":
		//BAR67 Additional Information–Taxes
		data := AdditionalInformationTaxesBAR67{
			SequenceNumber:               extractData(line, 3, 11),
			StandardNumericQualifier:     extractData(line, 11, 13),
			DateofIssue:                  extractData(line, 13, 19),
			TransactionNumber:            extractData(line, 19, 25),
			TicketDocumentNumber:         extractData(line, 25, 39),
			CheckDigit:                   extractData(line, 39, 40),
			TaxInformationSequenceNumber: extractData(line, 40, 42),
			TaxInformationIdentifier:     extractData(line, 42, 46),
			AdditionalTaxInformation:     extractData(line, 46, 116),
		}
		// Marshal the struct to JSON
		jsonData, err := json.Marshal(data)
		if err != nil {
			log.Fatal(err)
		}
		InsertJSONData(db, "hot_file_transaction_records", "additional_tax_information", string(jsonData))
		fmt.Println("BFH:", string(jsonData))

	default:
		fmt.Println("Unknown record type:", recordBARType)
	}
}

// ProcessLineBMD processes a line of data based on its record type.
func ProcessLineBMD(db *sql.DB, line string) {
	// Determine the BMD record type based on the characters(12 TO 13) (Standard Numeric Qualifier)
	recordBMDType := extractData(line, 11, 13)
	switch recordBMDType {
	case "75":
		//BMD75 Electronic Miscellaneous Document Coupon Detail Record
		data := ElectronicMiscellaneousDocumentCouponDetailRecordBMD75{
			SequenceNumber:                         extractData(line, 3, 11),
			StandardNumericQualifier:               extractData(line, 11, 13),
			DateofIssue:                            extractData(line, 13, 19),
			TransactionNumber:                      extractData(line, 19, 25),
			TicketDocumentNumber:                   extractData(line, 25, 39),
			CheckDigit:                             extractData(line, 39, 40),
			EMDCouponNumber:                        extractData(line, 40, 41),
			EMDCouponValue:                         extractData(line, 41, 52),
			EMDRelatedTicketOrDocumentNumber:       extractData(line, 52, 66),
			EMDRelatedCouponNumber:                 extractData(line, 66, 67),
			EMDServiceType:                         extractData(line, 67, 68),
			EMDReasonForIssuanceSubCode:            extractData(line, 68, 71),
			EMDFeeOwnerAirlineDesignator:           extractData(line, 71, 74),
			EMDExcessBaggageOverAllowanceQualifier: extractData(line, 74, 75),
			EMDExcessBaggageCurrencyCode:           extractData(line, 75, 78),
			EMDExcessBaggageRatePerUnit:            extractData(line, 78, 90),
			EMDExcessBaggageTotalNumberInExcess:    extractData(line, 90, 102),
			EMDConsumedAtIssuanceIndicator:         extractData(line, 102, 103),
			EMDNumberOfServices:                    extractData(line, 103, 106),
			EMDOperatingCarrier:                    extractData(line, 106, 109),
			EMDAttributeGroup:                      extractData(line, 109, 112),
			EMDAttributeSubGroup:                   extractData(line, 112, 115),
			EMDIndustryCarrierIndicator:            extractData(line, 115, 116),
			ReservedSpace:                          extractData(line, 116, 132),
			CurrencyType:                           extractData(line, 132, 136),
		}
		// Marshal the struct to JSON
		jsonData, err := json.Marshal(data)
		if err != nil {
			log.Fatal(err)
		}
		InsertJSONData(db, "hot_file_transaction_records", "electronic_miscellaneous_document_coupon_detail", string(jsonData))
		fmt.Println("BFH:", string(jsonData))

	case "76":
		//BMD76 Electronic Miscellaneous Document Coupon Remarks Record
		data := ElectronicMiscellaneousDocumentCouponRemarksRecordBMD76{
			SequenceNumber:           extractData(line, 3, 11),
			StandardNumericQualifier: extractData(line, 11, 13),
			DateofIssue:              extractData(line, 13, 19),
			TransactionNumber:        extractData(line, 19, 25),
			TicketDocumentNumber:     extractData(line, 25, 39),
			CheckDigit:               extractData(line, 39, 40),
			CouponNumber:             extractData(line, 40, 41),
			EMDRemarks:               extractData(line, 41, 111),
		}
		// Marshal the struct to JSON
		jsonData, err := json.Marshal(data)
		if err != nil {
			log.Fatal(err)
		}
		InsertJSONData(db, "hot_file_transaction_records", "electronic_miscellaneous_document_remarks", string(jsonData))
		fmt.Println("BFH:", string(jsonData))
	default:
		fmt.Println("Unknown record type:", recordBMDType)
	}
}

// ProcessLineBOT processes a line of data based on its record type.
func ProcessLineBOT(db *sql.DB, line string) {
	// Determine the BMD record type based on the characters(12 TO 13) (Standard Numeric Qualifier)
	recordBOTType := extractData(line, 11, 13)
	switch recordBOTType {
	case "93":
		//BOT93 Office Subtotals per Transaction Code and Currency Type Record
		data := OfficeSubtotalsPerTransactionCodeAndCurrencyTypeRecordBOT93{
			SequenceNumber:                   extractData(line, 3, 11),
			StandardNumericQualifier:         extractData(line, 11, 13),
			AgentNumericCode:                 extractData(line, 13, 21),
			RemittancePeriodEndingDate:       extractData(line, 21, 27),
			GrossValueAmount:                 extractData(line, 27, 42),
			TotalRemittanceAmount:            extractData(line, 42, 57),
			TotalCommissionValueAmount:       extractData(line, 57, 72),
			TotalTaxOrMiscellaneousFeeAmount: extractData(line, 72, 87),
			TransactionCode:                  extractData(line, 87, 91),
			TotalTaxOnCommissionAmount:       extractData(line, 91, 106),
			ReservedSpace:                    extractData(line, 106, 132),
			CurrencyType:                     extractData(line, 132, 136),
		}
		// Marshal the struct to JSON
		jsonData, err := json.Marshal(data)
		if err != nil {
			log.Fatal(err)
		}
		InsertJSONData(db, "hot_file_total_records", "office_subtotals_per_transaction_code_currency_type", string(jsonData))
		fmt.Println("BFH:", string(jsonData))
	case "94":
		//BOT94 Office Totals per Currency Type Record
		data := OfficeTotalsPerCurrencyTypeRecordBOT94{
			SequenceNumber:                   extractData(line, 3, 11),
			StandardNumericQualifier:         extractData(line, 11, 13),
			AgentNumericCode:                 extractData(line, 13, 21),
			RemittancePeriodEndingDate:       extractData(line, 21, 27),
			GrossValueAmount:                 extractData(line, 27, 42),
			TotalRemittanceAmount:            extractData(line, 42, 57),
			TotalCommissionValueAmount:       extractData(line, 57, 72),
			TotalTaxOrMiscellaneousFeeAmount: extractData(line, 72, 87),
			TotalTaxOnCommissionAmount:       extractData(line, 87, 102),
			ReservedSpace:                    extractData(line, 102, 132),
			CurrencyType:                     extractData(line, 132, 136),
		}
		// Marshal the struct to JSON
		jsonData, err := json.Marshal(data)
		if err != nil {
			log.Fatal(err)
		}
		InsertJSONData(db, "hot_file_total_records", "office_totals_per_currency_type", string(jsonData))
		fmt.Println("BFH:", string(jsonData))

	default:
		fmt.Println("Unknown record type:", recordBOTType)
	}
}
