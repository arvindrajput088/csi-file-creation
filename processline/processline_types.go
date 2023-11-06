package processline

//HEADER RECORDS
type FileHeaderBFH01 struct {
	SequenceNumber             string `json:"sequenceNumber"`
	StandardNumericQualifier   string `json:"standardNumericQualifier"`
	BSPIdentifier              string `json:"bspIdentifier"`
	TicketingAirlineCodeNumber string `json:"ticketingAirlineCodeNumber"`
	HandbookRevisionNumber     string `json:"handbookRevisionNumber"`
	TestorProductionStatus     string `json:"testorProductionStatus"`
	ProcessingDate             string `json:"processingDate"`
	ProcessingTime             string `json:"processingTime"`
	ISOCountryCode             string `json:"isoCountryCode"`
	FileSequenceNumber         string `json:"fileSequenceNumber"`
}

type BillingCycleHeaderBCH02 struct {
	SequenceNumber            string `json:"sequenceNumber"`
	StandardNumericQualifier  string `json:"standardNumericQualifier"`
	ProcessingDateIdentifier  string `json:"processingDateIdentifier"`
	ProcessingCycleIdentifier string `json:"processingCycleIdentifier"`
	BillingAnalysisEndingDate string `json:"billingAnalysisEndingDate"`
	DynamicRunIdentifier      string `json:"dynamicRunIdentifier"`
	HotReportingEndDate       string `json:"hotReportingEndDate"`
}

type OfficeHeaderRecordBOH03 struct {
	SequenceNumber             string `json:"sequenceNumber"`
	StandardNumericQualifier   string `json:"standardNumericQualifier"`
	AgentNumericCode           string `json:"agentNumericCode"`
	RemittancePeriodEndingDate string `json:"remittancePeriodEndingDate"`
	CurrencyType               string `json:"CurrencyType"`
	MultiLocationIdentifier    string `json:"multiLocationIdentifier"`
}

//TRANSACTION RECORDS
type TransactionHeaderRecordBKT06 struct {
	SequenceNumber                    string `json:"sequenceNumber"`
	StandardNumericQualifier          string `json:"standardNumericQualifier"`
	TransactionNumber                 string `json:"transactionNumber"`
	NetReportingIndicator             string `json:"netReportingIndicator"`
	TransactionRecordCounter          string `json:"transactionRecordCounter"`
	TicketingAirlineCodeNumber        string `json:"ticketingAirlineCodeNumber"`
	CommercialAgreementReference      string `json:"commercialAgreementReference"`
	CustomerFileReference             string `json:"customerFileReference"`
	ReportingSystemIdentifier         string `json:"reportingSystemIdentifier"`
	SettlementAuthorisationCode       string `json:"settlementAuthorisationCode"`
	DataInputStatusIndicator          string `json:"dataInputStatusIndicator"`
	NetReportingMethodIndicator       string `json:"netReportingMethodIndicator"`
	NetReportingCalculationType       string `json:"netReportingCalculationType"`
	AutomatedRepricingEngineIndicator string `json:"automatedRepricingEngineIndicator"`
}

type TicketIdentificationRecordBKS24 struct {
	SequenceNumber                    string `json:"sequenceNumber"`
	StandardNumericQualifier          string `json:"standardNumericQualifier"`
	DateofIssue                       string `json:"dateofIssue"`
	TransactionNumber                 string `json:"transactionNumber"`
	TicketDocumentNumber              string `json:"ticketDocumentNumber"`
	CheckDigit                        string `json:"check-Digit"`
	CouponUseIndicator                string `json:"couponUseIndicator"`
	ConjunctionTicketIndicator        string `json:"conjunctionTicketIndicator"`
	AgentNumericCode                  string `json:"agentNumericCode"`
	ReasonForIssuanceCode             string `json:"reasonForIssuanceCode"`
	TourCode                          string `json:"tourCode"`
	TransactionCode                   string `json:"transactionCode"`
	TrueOriginDestinationCityCodes    string `json:"trueOriginDestinationCityCodes"`
	PNRReferenceAndAirlineData        string `json:"pnrReferenceAndAirlineData"`
	TimeOfIssue                       string `json:"timeOfIssue"`
	JourneyTournaroundAirportCityCode string `json:"journeyTournaroundAirportCityCode"`
}

type STDOrDocumentAmountsRecordBKS30 struct {
	SequenceNumber               string `json:"sequenceNumber"`
	StandardNumericQualifier     string `json:"standardNumericQualifier"`
	DateofIssue                  string `json:"dateofIssue"`
	TransactionNumber            string `json:"transactionNumber"`
	TicketDocumentNumber         string `json:"ticketDocumentNumber"`
	CheckDigit                   string `json:"check-Digit"`
	CommissionableAmount         string `json:"commissionableAmount"`
	NetFareAmount                string `json:"netFareAmount"`
	TaxOrMiscellaneousFeeType1   string `json:"taxOrMiscellaneousFeeType1"`
	TaxOrMiscellaneousFeeAmount1 string `json:"taxOrMiscellaneousFeeAmount1"`
	TaxOrMiscellaneousFeeType2   string `json:"taxOrMiscellaneousFeeType2"`
	TaxOrMiscellaneousFeeAmount2 string `json:"taxOrMiscellaneousFeeAmount2"`
	TaxOrMiscellaneousFeeType3   string `json:"taxOrMiscellaneousFeeType3"`
	TaxOrMiscellaneousFeeAmount3 string `json:"taxOrMiscellaneousFeeAmount3"`
	TicketOrDocumentAmount       string `json:"ticketOrDocumentAmount"`
	ReservedSpace                string `json:"reservedSpace"`
	CurrencyType                 string `json:"currencyType"`
}

type CouponTaxInformationRecordBKS31 struct {
	SequenceNumber           string `json:"sequenceNumber"`
	StandardNumericQualifier string `json:"standardNumericQualifier"`
	DateofIssue              string `json:"dateofIssue"`
	TransactionNumber        string `json:"transactionNumber"`
	TicketDocumentNumber     string `json:"ticketDocumentNumber"`
	CheckDigit               string `json:"check-Digit"`
	//First Coupon Tax Code
	SegmentIdentifier1         string `json:"segmentIdentifier1"`
	CouponTaxAirportCode1      string `json:"couponTaxAirportCode1"`
	SegmentTaxAirportCode1     string `json:"segmentTaxAirportCode1"`
	CouponTaxCode1             string `json:"couponTaxCode1"`
	CouponTaxType1             string `json:"couponTaxType1"`
	CouponTaxReportedAmount1   string `json:"couponTaxReportedAmount1"`
	CouponTaxCurrencyType1     string `json:"couponTaxCurrencyType1"`
	CouponTaxApplicableAmount1 string `json:"couponTaxApplicableAmount1"`
	//Second Coupon Tax Code
	SegmentIdentifier2         string `json:"segmentIdentifier2"`
	CouponTaxAirportCode2      string `json:"couponTaxAirportCode2"`
	SegmentTaxAirportCode2     string `json:"segmentTaxAirportCode2"`
	CouponTaxCode2             string `json:"couponTaxCode2"`
	CouponTaxType2             string `json:"couponTaxType2"`
	CouponTaxReportedAmount2   string `json:"couponTaxReportedAmount2"`
	CouponTaxCurrencyType2     string `json:"couponTaxCurrencyType2"`
	CouponTaxApplicableAmount2 string `json:"couponTaxApplicableAmount2"`
	ReservedSpace              string `json:"reservedSpace"`
	CurrencyType               string `json:"currencyType"`
}

type CommissionRecordBKS39 struct {
	SequenceNumber                          string `json:"sequenceNumber"`
	StandardNumericQualifier                string `json:"standardNumericQualifier"`
	DateofIssue                             string `json:"dateofIssue"`
	TransactionNumber                       string `json:"transactionNumber"`
	TicketDocumentNumber                    string `json:"ticketDocumentNumber"`
	CheckDigit                              string `json:"check-Digit"`
	StatisticalCode                         string `json:"statisticalCode"`
	CommissionType                          string `json:"commissionType"`
	CommissionRate                          string `json:"commissionRate"`
	CommissionAmount                        string `json:"commissionAmount"`
	SupplementaryType                       string `json:"supplementaryType"`
	SupplementaryRate                       string `json:"supplementaryRate"`
	SupplementaryAmount                     string `json:"supplementaryAmount"`
	EffectiveCommissionRate                 string `json:"effectiveCommissionRate"`
	EffectiveCommissionAmount               string `json:"effectiveCommissionAmount"`
	AmountPaidbyCustomer                    string `json:"amountPaidbyCustomer"`
	RoutingDomesticOrInternationalIndicator string `json:"routingDomesticOrInternationalIndicator"`
	CommissionControlAdjustmentIndicator    string `json:"commissionControlAdjustmentIndicator"`
	ReservedSpace                           string `json:"reservedSpace"`
	CurrencyType                            string `json:"currencyType"`
}

type TaxOnCommissionRecordBKS42 struct {
	SequenceNumber           string `json:"sequenceNumber"`
	StandardNumericQualifier string `json:"standardNumericQualifier"`
	DateofIssue              string `json:"dateofIssue"`
	TransactionNumber        string `json:"transactionNumber"`
	TicketDocumentNumber     string `json:"ticketDocumentNumber"`
	CheckDigit               string `json:"check-Digit"`
	TaxonCommissionType1     string `json:"TaxonCommissionType1"`
	TaxonCommissionAmount1   string `json:"TaxonCommissionAmount1"`
	TaxonCommissionType2     string `json:"TaxonCommissionType2"`
	TaxonCommissionAmount2   string `json:"TaxonCommissionAmount2"`
	TaxonCommissionType3     string `json:"TaxonCommissionType3"`
	TaxonCommissionAmount3   string `json:"TaxonCommissionAmount3"`
	TaxonCommissionType4     string `json:"TaxonCommissionType4"`
	TaxonCommissionAmount4   string `json:"TaxonCommissionAmount4"`
	ReservedSpace            string `json:"reservedSpace"`
	CurrencyType             string `json:"currencyType"`
}

type RelatedTicketOrDocumentInformationRecordBKS45 struct {
	SequenceNumber                                string `json:"sequenceNumber"`
	StandardNumericQualifier                      string `json:"standardNumericQualifier"`
	RemittancePeriodEndingDate                    string `json:"remittancePeriodEndingDate"`
	TransactionNumber                             string `json:"transactionNumber"`
	RelatedTicketOrDocumentNumber                 string `json:"RelatedTicketOrDocumentNumber"`
	CheckDigit                                    string `json:"check-Digit"`
	WaiverCode                                    string `json:"waiverCode"`
	ReasonForMemoIssuanceCode                     string `json:"reasonForMemoIssuanceCode"`
	RelatedTicketOrDocumentCouponNumberIdentifier string `json:"relatedTicketOrDocumentCouponNumberIdentifier"`
	DateOfIssueRelatedDocument                    string `json:"dateOfIssueRelatedDocument"`
}

type QualifyingIssueInformationForSalesTransactionsRecordBKS46 struct {
	SequenceNumber                          string `json:"sequenceNumber"`
	StandardNumericQualifier                string `json:"standardNumericQualifier"`
	DateofIssue                             string `json:"dateofIssue"`
	TransactionNumber                       string `json:"transactionNumber"`
	TicketDocumentNumber                    string `json:"ticketDocumentNumber"`
	CheckDigit                              string `json:"check-Digit"`
	OriginalIssueTicketOrDocumentNumber     string `json:"OriginalIssueTicketOrDocumentNumber"`
	OriginalIssueLocationCityCode           string `json:"OriginalIssueLocationCityCode"`
	OriginalIssueDateDDMMMYY                string `json:"OriginalIssueDateDDMMMYY"`
	OriginalIssueAgentNumericCodeIATANumber string `json:"OriginalIssueAgentNumericCodeIATANumber"`
	EndorsementsOrRestrictions              string `json:"EndorsementsOrRestrictions"`
}

type NettingValuesRecordBKS47 struct {
	SequenceNumber           string `json:"sequenceNumber"`
	StandardNumericQualifier string `json:"standardNumericQualifier"`
	DateofIssue              string `json:"dateofIssue"`
	TransactionNumber        string `json:"transactionNumber"`
	TicketDocumentNumber     string `json:"ticketDocumentNumber"`
	CheckDigit               string `json:"check-Digit"`
	NettingType1             string `json:"nettingType1"`
	NettingCode1             string `json:"nettingCode1"`
	NettingAmount1           string `json:"nettingAmount1"`
	NettingType2             string `json:"nettingType2"`
	NettingCode2             string `json:"nettingCode2"`
	NettingAmount2           string `json:"nettingAmount2"`
	NettingType3             string `json:"nettingType3"`
	NettingCode3             string `json:"nettingCode3"`
	NettingAmount3           string `json:"nettingAmount3"`
	NettingType4             string `json:"nettingType4"`
	NettingCode4             string `json:"nettingCode4"`
	NettingAmount4           string `json:"nettingAmount4"`
	ReservedSpace            string `json:"reservedSpace"`
	CurrencyType             string `json:"currencyType"`
}

type UnticketedPointInformationRecordBKI61 struct {
	SequenceNumber                        string `json:"sequenceNumber"`
	StandardNumericQualifier              string `json:"standardNumericQualifier"`
	DateofIssue                           string `json:"dateofIssue"`
	TransactionNumber                     string `json:"transactionNumber"`
	TicketDocumentNumber                  string `json:"ticketDocumentNumber"`
	CheckDigit                            string `json:"check-Digit"`
	SegmentIdentifier                     string `json:"segmentIdentifier"`
	UnticketedPointAirportOrCityCode      string `json:"unticketedPointAirportOrCityCode"`
	UnticketedPointDateOfArrival          string `json:"unticketedPointDateOfArrival"`
	UnticketedPointLocalTimeOfArrival     string `json:"unticketedPointLocalTimeOfArrival"`
	UnticketedPointDateOfDeparture        string `json:"unticketedPointDateOfDeparture"`
	UnticketedPointLocalTimeOfDeparture   string `json:"unticketedPointLocalTimeOfDeparture"`
	UnticketedPointDepartureEquipmentCode string `json:"unticketedPointDepartureEquipmentCode"`
}

type AdditionalItineraryDataRecordBKI62 struct {
	SequenceNumber               string `json:"sequenceNumber"`
	StandardNumericQualifier     string `json:"standardNumericQualifier"`
	DateofIssue                  string `json:"dateofIssue"`
	TransactionNumber            string `json:"transactionNumber"`
	TicketDocumentNumber         string `json:"ticketDocumentNumber"`
	CheckDigit                   string `json:"check-Digit"`
	SegmentIdentifier            string `json:"segmentIdentifier"`
	OriginAirportOrCityCode      string `json:"originAirportOrCityCode"`
	FlightDepartureDate          string `json:"flightDepartureDate"`
	FlightDepartureTime          string `json:"flightDepartureTime"`
	FlightDepartureTerminal      string `json:"flightDepartureTerminal"`
	DestinationAirportOrCityCode string `json:"destinationAirportOrCityCode"`
	FlightArrivalDate            string `json:"flightArrivalDate"`
	FlightArrivalTime            string `json:"flightArrivalTime"`
	FlightArrivalTerminal        string `json:"flightArrivalTerminal"`
}

type ItineraryDataSegmentRecordBKI63 struct {
	SequenceNumber                       string `json:"sequenceNumber"`
	StandardNumericQualifier             string `json:"standardNumericQualifier"`
	DateofIssue                          string `json:"dateofIssue"`
	TransactionNumber                    string `json:"transactionNumber"`
	TicketDocumentNumber                 string `json:"ticketDocumentNumber"`
	CheckDigit                           string `json:"check-Digit"`
	SegmentIdentifier                    string `json:"segmentIdentifier"`
	StopoverCode                         string `json:"stopoverCode"`
	NotValidBeforeDate                   string `json:"notValidBeforeDate"`
	NotValidAfterDate                    string `json:"notValidAfterDate"`
	OriginAirportOrCityCode              string `json:"originAirportOrCityCode"`
	DestinationAirportOrCityCode         string `json:"destinationAirportOrCityCode"`
	Carrier                              string `json:"carrier"`
	SoldPassengerCabin                   string `json:"soldPassengerCabin"`
	FlightNumber                         string `json:"flightNumber"`
	ReservationBookingDesignator         string `json:"reservationBookingDesignator"`
	FlightDepartureDate                  string `json:"flightDepartureDate"`
	FlightDepartureTime                  string `json:"flightDepartureTime"`
	FlightBookingStatus                  string `json:"flightBookingStatus"`
	BaggageAllowance                     string `json:"baggageAllowance"`
	FareBasisOrTicketDesignator          string `json:"fareBasisOrTicketDesignator"`
	FrequentFlyerReference               string `json:"frequentFlyerReference"`
	FareComponentPricedPassengerTypeCode string `json:"fareComponentPricedPassengerTypeCode"`
	ThroughOrChangeOfGaugeIndicator      string `json:"throughOrChangeOfGaugeIndicator"`
	EquipmentCode                        string `json:"equipmentCode"`
}

type DocumentAmountsRecordBAR64 struct {
	SequenceNumber                                     string `json:"sequenceNumber"`
	StandardNumericQualifier                           string `json:"standardNumericQualifier"`
	DateofIssue                                        string `json:"dateofIssue"`
	TransactionNumber                                  string `json:"transactionNumber"`
	TicketDocumentNumber                               string `json:"ticketDocumentNumber"`
	CheckDigit                                         string `json:"check-Digit"`
	Fare                                               string `json:"fare"`
	TicketingModeIndicator                             string `json:"ticketingModeIndicator"`
	EquivalentFarePaid                                 string `json:"equivalentFarePaid"`
	Total                                              string `json:"total"`
	ServicingAirlineOrSystemProviderIdentifier         string `json:"servicingAirlineOrSystemProviderIdentifier"`
	FareCalculationModeIndicator                       string `json:"fareCalculationModeIndicator"`
	BookingAgentIdentification                         string `json:"bookingAgentIdentification"`
	BookingEntityOutletType                            string `json:"bookingEntityOutletType"`
	AirlineIssFareCalculationPricingIndicatoruingAgent string `json:"AirlineIssFareCalculationPricingIndicatoruingAgent"`
	AirlineIssuingAgent                                string `json:"airlineIssuingAgent"`
}

type AdditionalInformationPassengerRecordBAR65 struct {
	SequenceNumber           string `json:"sequenceNumber"`
	StandardNumericQualifier string `json:"standardNumericQualifier"`
	DateofIssue              string `json:"dateofIssue"`
	TransactionNumber        string `json:"transactionNumber"`
	TicketDocumentNumber     string `json:"ticketDocumentNumber"`
	CheckDigit               string `json:"check-Digit"`
	PassengerName            string `json:"passengerName"`
	PassengerSpecificData    string `json:"passengerSpecificData"`
	DateOfBirth              string `json:"dateOfBirth"`
	PassengerTypeCode        string `json:"passengerTypeCode"`
}

type AdditionalInformationFormOfPaymentRecordBAR66 struct {
	SequenceNumber              string `json:"sequenceNumber"`
	StandardNumericQualifier    string `json:"standardNumericQualifier"`
	DateofIssue                 string `json:"dateofIssue"`
	TransactionNumber           string `json:"transactionNumber"`
	TicketDocumentNumber        string `json:"ticketDocumentNumber"`
	CheckDigit                  string `json:"check-Digit"`
	FormOfPaymentSequenceNumber string `json:"FormOfPaymentSequenceNumber"`
	FormOfPaymentInformation    string `json:"FormOfPaymentInformation"`
}

type AdditionalInformationTaxesBAR67 struct {
	SequenceNumber               string `json:"sequenceNumber"`
	StandardNumericQualifier     string `json:"standardNumericQualifier"`
	DateofIssue                  string `json:"dateofIssue"`
	TransactionNumber            string `json:"transactionNumber"`
	TicketDocumentNumber         string `json:"ticketDocumentNumber"`
	CheckDigit                   string `json:"check-Digit"`
	TaxInformationSequenceNumber string `json:"taxInformationSequenceNumber"`
	TaxInformationIdentifier     string `json:"taxInformationIdentifier"`
	AdditionalTaxInformation     string `json:"additionalTaxInformation"`
}

type ElectronicMiscellaneousDocumentCouponDetailRecordBMD75 struct {
	SequenceNumber                         string `json:"sequenceNumber"`
	StandardNumericQualifier               string `json:"standardNumericQualifier"`
	DateofIssue                            string `json:"dateofIssue"`
	TransactionNumber                      string `json:"transactionNumber"`
	TicketDocumentNumber                   string `json:"ticketDocumentNumber"`
	CheckDigit                             string `json:"check-Digit"`
	EMDCouponNumber                        string `json:"EMDCouponNumber"`
	EMDCouponValue                         string `json:"EMDCouponValue"`
	EMDRelatedTicketOrDocumentNumber       string `json:"additionalTaxInformation"`
	EMDRelatedCouponNumber                 string `json:"emdRelatedCouponNumber"`
	EMDServiceType                         string `json:"emdServiceType"`
	EMDReasonForIssuanceSubCode            string `json:"emdReasonForIssuanceSubCode"`
	EMDFeeOwnerAirlineDesignator           string `json:"emdFeeOwnerAirlineDesignator"`
	EMDExcessBaggageOverAllowanceQualifier string `json:"emdExcessBaggageOverAllowanceQualifier"`
	EMDExcessBaggageCurrencyCode           string `json:"emdExcessBaggageCurrencyCode"`
	EMDExcessBaggageRatePerUnit            string `json:"emdExcessBaggageRatePerUnit"`
	EMDExcessBaggageTotalNumberInExcess    string `json:"emdExcessBaggageTotalNumberInExcess"`
	EMDConsumedAtIssuanceIndicator         string `json:"emdConsumedAtIssuanceIndicator"`
	EMDNumberOfServices                    string `json:"emdNumberOfServices"`
	EMDOperatingCarrier                    string `json:"emdOperatingCarrier"`
	EMDAttributeGroup                      string `json:"emdAttributeGroup"`
	EMDAttributeSubGroup                   string `json:"emdAttributeSubGroup"`
	EMDIndustryCarrierIndicator            string `json:"emdIndustryCarrierIndicator"`
	ReservedSpace                          string `json:"reservedSpace"`
	CurrencyType                           string `json:"currencyType"`
}

type ElectronicMiscellaneousDocumentCouponRemarksRecordBMD76 struct {
	SequenceNumber           string `json:"sequenceNumber"`
	StandardNumericQualifier string `json:"standardNumericQualifier"`
	DateofIssue              string `json:"dateofIssue"`
	TransactionNumber        string `json:"transactionNumber"`
	TicketDocumentNumber     string `json:"ticketDocumentNumber"`
	CheckDigit               string `json:"check-Digit"`
	CouponNumber             string `json:"couponNumber"`
	EMDRemarks               string `json:"emdRemarks"`
}

type FareCalculationRecordBKF81 struct {
	SequenceNumber                string `json:"sequenceNumber"`
	StandardNumericQualifier      string `json:"standardNumericQualifier"`
	DateofIssue                   string `json:"dateofIssue"`
	TransactionNumber             string `json:"transactionNumber"`
	TicketDocumentNumber          string `json:"ticketDocumentNumber"`
	CheckDigit                    string `json:"check-Digit"`
	FareCalculationSequenceNumber string `json:"FareCalculationSequenceNumber"`
	FareCalculationArea           string `json:"FareCalculationArea"`
}

type AdditionalCardInformationRecorBCC82 struct {
	SequenceNumber                     string `json:"sequenceNumber"`
	StandardNumericQualifier           string `json:"standardNumericQualifier"`
	DateofIssue                        string `json:"dateofIssue"`
	TransactionNumber                  string `json:"transactionNumber"`
	FormOfPaymentType                  string `json:"formOfPaymentType"`
	FormOfPaymentTransactionIdentifier string `json:"formOfPaymentTransactionIdentifier"`
}

type THREEDSAuthenticationandAdditionalCardPaymentInformationRecordBCX83 struct {
	SequenceNumber                                                string `json:"sequenceNumber"`
	StandardNumericQualifier                                      string `json:"standardNumericQualifier"`
	DateofIssue                                                   string `json:"dateofIssue"`
	TransactionNumber                                             string `json:"transactionNumber"`
	FormOfPaymentType                                             string `json:"formOfPaymentType"`
	CardAuthenticationSequenceNumber                              string `json:"cardAuthenticationSequenceNumber"`
	ThreeDSecureAuthenticationAndAdditionalCardPaymentInformation string `json:"threeDSecureAuthenticationAndAdditionalCardPaymentInformation"`
}

type FormOfPaymentRecordBKP84 struct {
	SequenceNumber              string `json:"sequenceNumber"`
	StandardNumericQualifier    string `json:"standardNumericQualifier"`
	DateofIssue                 string `json:"dateofIssue"`
	TransactionNumber           string `json:"transactionNumber"`
	FormOfPaymentType           string `json:"formOfPaymentType"`
	FormOfPaymentAmount         string `json:"formOfPaymentAmount"`
	FormOfPaymentAccountNumber  string `json:"formOfPaymentAccountNumber"`
	ExpiryDate                  string `json:"expiryDate"`
	ExtendedPaymentCode         string `json:"extendedPaymentCode"`
	ApprovalCode                string `json:"approvalCode"`
	InvoiceNumber               string `json:"invoiceNumber"`
	InvoiceDate                 string `json:"invoiceDate"`
	RemittanceAmount            string `json:"remittanceAmount"`
	CardVerificationValueResult string `json:"cardVerificationValueResult"`
	ReservedSpace               string `json:"reservedSpace"`
	CurrencyType                string `json:"currencyType"`
}

//TOTAL RECORDS

type OfficeSubtotalsPerTransactionCodeAndCurrencyTypeRecordBOT93 struct {
	SequenceNumber                   string `json:"sequenceNumber"`
	StandardNumericQualifier         string `json:"standardNumericQualifier"`
	AgentNumericCode                 string `json:"agentNumericCode"`
	RemittancePeriodEndingDate       string `json:"remittancePeriodEndingDate"`
	GrossValueAmount                 string `json:"grossValueAmount"`
	TotalRemittanceAmount            string `json:"totalRemittanceAmount"`
	TotalCommissionValueAmount       string `json:"totalCommissionValueAmount"`
	TotalTaxOrMiscellaneousFeeAmount string `json:"totalTaxOrMiscellaneousFeeAmount"`
	TransactionCode                  string `json:"transactionCode"`
	TotalTaxOnCommissionAmount       string `json:"totalTaxOnCommissionAmount"`
	ReservedSpace                    string `json:"reservedSpace"`
	CurrencyType                     string `json:"currencyType"`
}

type OfficeTotalsPerCurrencyTypeRecordBOT94 struct {
	SequenceNumber                   string `json:"sequenceNumber"`
	StandardNumericQualifier         string `json:"standardNumericQualifier"`
	AgentNumericCode                 string `json:"agentNumericCode"`
	RemittancePeriodEndingDate       string `json:"remittancePeriodEndingDate"`
	GrossValueAmount                 string `json:"grossValueAmount"`
	TotalRemittanceAmount            string `json:"totalRemittanceAmount"`
	TotalCommissionValueAmount       string `json:"totalCommissionValueAmount"`
	TotalTaxOrMiscellaneousFeeAmount string `json:"totalTaxOrMiscellaneousFeeAmount"`
	TotalTaxOnCommissionAmount       string `json:"totalTaxOnCommissionAmount"`
	ReservedSpace                    string `json:"reservedSpace"`
	CurrencyType                     string `json:"currencyType"`
}

type BillingAnalysisTotalsPerCurrencyTypeRecordBCT95 struct {
	SequenceNumber                   string `json:"sequenceNumber"`
	StandardNumericQualifier         string `json:"standardNumericQualifier"`
	ProcessingDateIdentifier         string `json:"processingDateIdentifier"`
	ProcessingCycleIdentifier        string `json:"processingCycleIdentifier"`
	OfficeCount                      string `json:"officeCount"`
	GrossValueAmount                 string `json:"grossValueAmount"`
	TotalRemittanceAmount            string `json:"totalRemittanceAmount"`
	TotalCommissionValueAmount       string `json:"totalCommissionValueAmount"`
	TotalTaxOrMiscellaneousFeeAmount string `json:"totalTaxOrMiscellaneousFeeAmount"`
	TotalTaxOnCommissionAmount       string `json:"totalTaxOnCommissionAmount"`
	ReservedSpace                    string `json:"reservedSpace"`
	CurrencyType                     string `json:"currencyType"`
}

type FileTotalsPerCurrencyTypeRecordBFT99 struct {
	SequenceNumber                   string `json:"sequenceNumber"`
	StandardNumericQualifier         string `json:"standardNumericQualifier"`
	BSPIdentifier                    string `json:"bspIdentifier"`
	OfficeCount                      string `json:"officeCount"`
	GrossValueAmount                 string `json:"grossValueAmount"`
	TotalRemittanceAmount            string `json:"totalRemittanceAmount"`
	TotalCommissionValueAmount       string `json:"totalCommissionValueAmount"`
	TotalTaxOrMiscellaneousFeeAmount string `json:"totalTaxOrMiscellaneousFeeAmount"`
	TotalTaxOnCommissionAmount       string `json:"totalTaxOnCommissionAmount"`
	ReservedSpace                    string `json:"reservedSpace"`
	CurrencyType                     string `json:"currencyType"`
}
