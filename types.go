package fda

// A DeviceSafetyReport reflects a report filed about one or more devices
type DeviceSafetyReport struct {
	AdverseEvent                 string `json:"adverse_event_flag"`
	ProductProblem               string `json:"product_problem_flag"`
	EventDate                    string `json:"date_of_event"`
	DateReport                   string `json:"date_report"`
	ReportReceivedDate           string `json:"date_received"`
	NumberOfDevices              string `json:"number_devices_in_event"`
	NumberOfPatients             string `json:"number_patients_in_event"`
	ReportNumber                 string `json:"report_number"`
	ReportSourceCode             string `json:"report_source_code"`
	ReportedByHealthProfessional string `json:"health_professional"`
	ReporterOccupationCode       string `json:"reporter_occupation_code"`
	WasInitialReportToFDA        string `json:"initial_report_to_fda"`
	DeviceWasReused              string `json:"reprocessed_and_reused_flag"`

	Devices  []Device `json:"device"`
	Patients []struct {
		SequenceNumber string   `json:"patient_sequence_number"`
		ReceivedDate   string   `json:"date_received"`
		Treatment      string   `json:"sequence_number_treatment"`
		Outcomes       []string `json:"sequence_number_outcome"`
	} `json:"patient"`
	Reports []struct {
		SequenceNumber string `json:"patient_sequence_number"`
		Description    string `json:"text_type_code"`
		Text           string `json:"text"`
		TextKey        string `json:"mdr_text_key"`
		ReceivedDate   string `json:"date_received"`
	} `json:"mdr_text"`

	ReportTypes                           []string `json:"type_of_report"`
	FacilityAwareDate                     string   `json:"date_facility_aware"`
	ReportDate                            string   `json:"report_date"`
	ReportedToFDA                         string   `json:"report_to_fda"`
	ReportedToFDADate                     string   `json:"date_report_to_fda"`
	ReportedToManufacturer                string   `json:"report_to_manufacturer"`
	ReportedToManufacturerDate            string   `json:"date_report_to_manufacturer"`
	EventLocation                         string   `json:"event_location"`
	DistributorName                       string   `json:"distributor_name"`
	DistributorAddress1                   string   `json:"distributor_address_1"`
	DistributorAddress2                   string   `json:"distributor_address_2"`
	DistributorCity                       string   `json:"distributor_city"`
	DistributorState                      string   `json:"distributor_state"`
	DistributorZip                        string   `json:"distributor_zip_code"`
	DistributorZipExt                     string   `json:"distributor_zip_code_ext"`
	ManufacturerName                      string   `json:"manufacturer_name"`
	ManufacturerAddress1                  string   `json:"manufacturer_address_1"`
	ManufacturerAddress2                  string   `json:"manufacturer_address_2"`
	ManufacturerCity                      string   `json:"manufacturer_city"`
	ManufacturerState                     string   `json:"manufacturer_state"`
	ManufacturerZip                       string   `json:"manufacturer_zip_code"`
	ManufacturerZipExt                    string   `json:"manufacturer_zip_code_ext"`
	ManufacturerCountry                   string   `json:"manufacturer_country"`
	ManufacturerPostalCode                string   `json:"manufacturer_postal_code"`
	EventTypes                            string   `json:"event_type"`
	ManufactureDate                       string   `json:"device_date_of_manufacture"`
	SingleUse                             string   `json:"single_use_flag"`
	PreviousUse                           string   `json:"previous_use_code"`
	RemedialActions                       []string `json:"remedial_action"`
	RemovalCorrectionNumber               string   `json:"removal_correction_number"`
	ManufacturerContactTitle              string   `json:"manufacturer_contact_t_name"`
	ManufacturerContactFirstName          string   `json:"manufacturer_contact_f_name"`
	ManufacturerContactLastName           string   `json:"manufacturer_contact_l_name"`
	ManufacturerContactStreet1            string   `json:"manufacturer_contact_street_1"`
	ManufacturerContactStreet2            string   `json:"manufacturer_contact_street_2"`
	ManufacturerContactCity               string   `json:"manufacturer_contact_city"`
	ManufacturerContactState              string   `json:"manufacturer_contact_state"`
	ManufacturerContactZip                string   `json:"manufacturer_contact_zip_code"`
	ManufacturerContactZipExt             string   `json:"manufacturer_contact_zip_ext"`
	ManufacturerContactPostal             string   `json:"manufacturer_contact_postal"`
	ManufacturerContactCountry            string   `json:"manufacturer_contact_country"`
	ManufacturerContactPhoneCountry       string   `json:"manufacturer_contact_pcountry"`
	ManufacturerContactPhoneAreaCode      string   `json:"manufacturer_contact_area_code"`
	ManufacturerContactPhoneExchange      string   `json:"manufacturer_contact_exchange"`
	ManufacturerContactPhoneExtension     string   `json:"manufacturer_contact_extension"`
	ManufacturerContactPhoneCity          string   `json:"manufacturer_contact_pcity"`
	ManufacturerContactPhoneContactNumber string   `json:"manufacturer_contact_phone_number"`
	ManufacturerContactPhoneLocal         string   `json:"manufacturer_contact_plocal"`
	ManufacturerG1Name                    string   `json:"manufacturer_g1_name"`
	ManufacturerG1Street1                 string   `json:"manufacturer_g1_street_1"`
	ManufacturerG1Street2                 string   `json:"manufacturer_g1_street_2"`
	ManufacturerG1City                    string   `json:"manufacturer_g1_city"`
	ManufacturerG1State                   string   `json:"manufacturer_g1_state"`
	ManufacturerG1Zip                     string   `json:"manufacturer_g1_zip_code"`
	ManufacturerG1ZipExt                  string   `json:"manufacturer_g1_zip_ext"`
	ManufacturerG1Postal                  string   `json:"manufacturer_g1_postal_code"`
	ManufacturerG1Country                 string   `json:"manufacturer_g1_country"`
	ManufacturerG1ManufacturerReceived    string   `json:"date_manufacturer_received"`
	SourceTypes                           []string `json:"source_type"`
	EventKey                              string   `json:"event_key"`
	MDRReportKey                          string   `json:"mdr_report_key"`
	LinkFlag                              string   `json:"manufacturer_link_flag_"`
}

// A Device is a specific device mention in a DeviceSafetyReport
type Device struct {
	SequenceNumber          string `json:"device_sequence_number"`
	EventKey                string `json:"device_event_key"`
	ReceivedDate            string `json:"date_received"`
	BrandName               string `json:"brand_name"`
	GenericName             string `json:"generic_name"`
	ProductCode             string `json:"device_report_product_code"`
	ModelNumber             string `json:"model_number"`
	CatalogNumber           string `json:"catalog_number"`
	LotNumber               string `json:"lot_number"`
	OtherIDNumber           string `json:"other_id_number"`
	ExpirationDate          string `json:"expiration_date_of_device"`
	Age                     string `json:"device_age_text"`
	Availability            string `json:"device_availability"`
	ReturnedDate            string `json:"date_returned_to_manufacturer"`
	EvaluatedByManufacturer string `json:"device_evaluated_by_manufacturer"`
	Operator                string `json:"device_operator"`
	Implanted               string `json:"implant_flag"`
	RemovedDate             string `json:"date_removed_flag"`
	ManufacturerName        string `json:"manufacturer_d_name"`
	ManufacturerAddress1    string `json:"manufacturer_d_address_1"`
	ManufacturerAddress2    string `json:"manufacturer_d_address_2"`
	ManufacturerCity        string `json:"manufacturer_d_city"`
	ManufacturerState       string `json:"manufacturer_d_state"`
	ManufacturerCountry     string `json:"manufacturer_d_country"`
	ManufacturerZip         string `json:"manufacturer_d_zip_code"`
	ManufacturerZipExt      string `json:"manufacturer_d_zip_code_ext"`
	ManufacturerPostalCode  string `json:"manufacturer_d_postal_code"`
}

// An EnforcementReport contains information about an enforcement action
// taken by the FDA
type EnforcementReport struct {
	ReportDate              string `json:"report_date"`
	RecallInitiationDate    string `json:"recall_initiation_date"`
	RecallNumber            string `json:"recall_number"`
	ReasonForRecall         string `json:"reason_for_recall"`
	Status                  string `json:"status"`
	Classification          string `json:"classification"`
	ProductType             string `json:"product_type"`
	ProductDescription      string `json:"product_description"`
	ProductQuantity         string `json:"product_quantity"`
	InitialFirmNotification string `json:"initial_firm_notification"`
	DistributionPattern     string `json:"distribution_pattern"`
	Country                 string `json:"country"`
	State                   string `json:"state"`
	City                    string `json:"city"`
	EventID                 string `json:"event_id"`
	RecallingFirm           string `json:"recalling_firm"`
	Voluntary               string `json:"voluntary_mandated"`
	CodeInfo                string `json:"code_info"`

	OpenFDA OpenFDA
}

// A DrugSafetyReport is one report filed regarding one individual who is on
// one or more medications. Subjectively, while there are many fields in
// each DrugSafetyReport, often the majority of fields are left empty.
type DrugSafetyReport struct {
	ID                               string  `json:"safetyreportid"`
	FulfillExpediteCriteria          string  `json:"fulfillexpeditecriteria"`
	Version                          string  `json:"safetyreportversion"`
	ReceiveDate                      string  `json:"receivedate"`
	ReceiveDateFormat                string  `json:"receivedateformat"`
	ReceiptDate                      string  `json:"receiptdate"`
	ReceiptDateFormat                string  `json:"receiptdateformat"`
	Epoch                            float64 `json:"@epoch"`
	Serious                          string
	SeriousnessCongentialAbnormality string `json:"seriousnesscongenitalanomali"`
	SeriousnessDeath                 string `json:"seriousnessdeath"`
	SeriousnessDisabling             string `json:"seriousnessdisabling"`
	SeriousnessHospitalization       string `json:"seriousnesshospitalization"`
	SeriousnessLifeThreatening       string `json:"seriousnesslifethreatening"`
	SeriouesnessOther                string `json:"seriousnessother"`
	TransmissionDate                 string `json:"transmissiondate"`
	TransmissionDateFormat           string `json:"transmissiondateformat"`
	Duplicate                        string `json:"duplicate"`
	CompanyNumber                    string `json:"companynumb"`
	OccurCountry                     string `json:"occurcountry"`
	PrimarySourceCountry             string `json:"primarysourcecountry"`
	PrimarySource                    struct {
		Qualification   string
		ReporterCountry string `json:"reportercountry"`
	} `json:"primarysource"`
	ReportDuplicate struct {
		Source string `json:"duplicatesource"`
		Number string `json:"duplicatenumb"`
	} `json:"reportduplicate"`
	Sender struct {
		Type         string `json:"sendertype"`
		Organization string `json:"senderorganization"`
	}
	Receiver struct {
		Type         string `json:"receivertype"`
		Organization string `json:"receiverorganization"`
	}

	Patient Patient
}

// A Patient is a representation of the individual in the SafetyReport.
// Many of these fields are often left blank.
type Patient struct {
	Age     string `json:"patientonsetage"`
	AgeUnit string `json:"patientonsetageunit"`
	Sex     string `json:"patientsex"`
	Weight  string `json:"patientweight"` //In kilograms
	Death   struct {
		Date       string `json:"patientdeathdate"`
		DateFormat string `json:"patientdeathdateformat"`
	} `json:"patientdeath"`

	Reactions []DrugReaction `json:"reaction"`
	Drugs     []Drug         `json:"drug"`
}

// DrugReaction represents the reaction that caused the DrugSafetyReport to be filed
// for a particular Patient, as well as the outcome of the event
type DrugReaction struct {
	MedDRA        string `json:"reactionmeddrapt"`
	MedDRAVersion string `json:"reactionmeddraversionpt"`
	Outcome       string `json:"reactionoutcome"`
}

// Drug provides information on a drug reported in a SafetyReport.
// Many of these fields are often left blank.
type Drug struct {
	ActionDrug                 string  `json:"actiondrug"`
	Additional                 string  `json:"drugadditional"`
	CumulativeDose             string  `json:"drugcumulativedosagenumb"`
	CumulativeDoseUnit         string  `json:"drugcumulativedosageunit"`
	DosageForm                 string  `json:"drugdosageform"`
	IntervalDosageDefinition   string  `json:"drugintervaldosagedefinition"`
	IntervalDosageUnits        string  `json:"drugintervaldosageunitnumb"`
	OccurredOnReadministration string  `json:"drugrecurreadministration"`
	NumSeparateDosages         string  `json:"drugseparatedosagenumb"`
	NumDoses                   string  `json:"drugstructuredosagenumb"`
	DosageUnit                 string  `json:"drugstructuredosageunit"`
	AdminRoute                 string  `json:"drugadministrationroute"`
	AuthorizationNum           string  `json:"drugauthorizationnumb"`
	BatchNum                   string  `json:"drugbatchnumb"`
	RoleCharacterization       string  `json:"drugcharacterization"`
	DosageText                 string  `json:"drugdosagetext"`
	StartDate                  string  `json:"drugstartdate"`
	StartDateFormat            string  `json:"drugstartdateformat"`
	EndDate                    string  `json:"drugenddate"`
	EndDateFormat              string  `json:"drugenddateformat"`
	TreatmentDuration          string  `json:"drugtreatmentduration"`
	TreatmentDurationUnit      string  `json:"drugtreatmentdurationunit"`
	Indication                 string  `json:"drugindication"`
	Name                       string  `json:"medicinalproduct"`
	OpenFDA                    OpenFDA `json:"openfda"`
}

// A DrugLabel is labeling information for a medication, including
// OTC and prescription medications. SetID, ID, Version, and EffectiveTime
// are supposed to be guaranteed to be present, while other fields are optional.
type DrugLabel struct {
	SetID         string `json:"set_id"`
	ID            string `json:"id"`
	Version       string `json:"version"`
	EffectiveTime string `json:"effective_time"`

	AbuseAndDependence           []string `json:"drug_abuse_and_dependence"`
	ControlledSubstance          []string `json:"controlled_substance"`
	Abuse                        []string `json:"abuse"`
	Dependence                   []string `json:"dependence"`
	Overdosage                   []string `json:"overdosage"`
	AdverseReactions             []string `json:"adverse_reactions"`
	DrugInteractions             []string `json:"drug_interactions"`
	LaboratoryInterference       []string `json:"drug_and_or_laboratory_test_interactions"`
	ClinicalPharmacology         []string `json:"clinical_pharmacology"`
	MechanismOfAction            []string `json:"mechanism_of_action"`
	Pharmacodynamics             []string `json:"pharmacodynamics"`
	Pharmacokinetics             []string `json:"pharmacokinetics"`
	IndicationsAndUsage          []string `json:"indications_and_usage"`
	Contraindications            []string `json:"contraindications"`
	DosageAndAdministration      []string `json:"dosage_and_administration"`
	DosageFormsAndStrengths      []string `json:"dosage_forms_and_strengths"`
	Purpose                      []string `json:"purpose"`
	Description                  []string `json:"description"`
	ActiveIngredient             []string `json:"active_ingredient"`
	InactiveIngredient           []string `json:"inactive_ingredient"`
	ProductIngredients           []string `json:"spl_product_data_elements"`
	PackageInsert                []string `json:"spl_patient_package_insert"`
	InformationForPatients       []string `json:"information_for_patients"`
	InformationForCaregivers     []string `json:"information_for_owners_or_caregivers"`
	InstructionsForUse           []string `json:"instructions_for_use"`
	WhenToAskADoctor             []string `json:"ask_doctor"`
	WhenToAskADoctorOrPharmacist []string `json:"ask_doctor_or_pharmacist"`
	AbsoluteContraindications    []string `json:"do_not_use"`
	InformationRegardingChildren []string `json:"keep_out_of_reach_of_children"`
	OtherSafetyInformation       []string `json:"other_safety_information"`
	TelephoneNumberForQuestions  []string `json:"questions"`
	WhenToDiscontinueImmediately []string `json:"stop_use"`
	SideEffectsAndWhatToAvoid    []string `json:"when_using"`
	PatientMedicationGuide       []string `json:"patient_medication_information"`
	MedicationGuide              []string `json:"spl_medguide"`
	SpecificPopulations          []string `json:"use_in_specific_populations"`
	EffectsOnPregnancy           []string `json:"pregnancy"`
	TeratogenicEffects           []string `json:"teratogenic_effects"`
	NonteratogenicEffects        []string `json:"nonteratogenic_effects"`
	LaborAndDelivery             []string `json:"labor_and_delivery"`
	NursingMothers               []string `json:"nursing_mothers"`
	PregnancyOrBreastfeeding     []string `json:"pregnancy_or_breast_feeding"`
	PediatricUse                 []string `json:"pediatric_use"`
	GeriatricUse                 []string `json:"geriatric_use"`
	NonclinicalToxicology        []string `json:"nonclinical_toxicology"`
	Carcinogenesis               []string `json:"carcinogenesis_and_mutagenesis_and_impairment_of_fertility"`
	AnimalPharmacology           []string `json:"animal_pharmacology_and_or_toxicology"`
	ClinicalStudies              []string `json:"clinical_studies"`
	References                   []string `json:"references"`
	HowSupplied                  []string `json:"how_supplied"`
	StorageAndHandling           []string `json:"storage_and_handling"`
	SafeHandlingWarning          []string `json:"safe_handling_warning"`
	BoxedWarning                 []string `json:"boxed_warning"`
	WarningsAndPrecautions       []string `json:"warnings_and_precautions"`
	UserSafetyWarnings           []string `json:"user_safety_warnings"`
	Precautions                  []string `json:"precautions"`
	Warnings                     []string `json:"warnings"`
	GeneralPrecautions           []string `json:"general_precautions"`
	LaboratoryTests              []string `json:"laboratory_tests"`
	RecentMajorChanges           []string `json:"recent_major_changes"`
	Microbiology                 []string `json:"microbiology"`
	PrincipalDisplayPanel        []string `json:"package_label_principal_display_panel"`
	UnclassifiedSection          []string `json:"spl_unclassified_section"`

	OpenFDA OpenFDA `json:"openfda"`
}

// OpenFDA reflects OpenFDA-enhanced fields. These are not guaranteed
// to be present due to a variety of reasons (typos in the original
// safety report, etc.)
type OpenFDA struct {
	ApplicationNumbers   []string `json:"application_number"`
	BrandNames           []string `json:"brand_name"`
	DosageForms          []string `json:"dosage_form"`
	GenericNames         []string `json:"generic_name"`
	ManufacturerNames    []string `json:"manufacturer_name"`
	IsOriginalPackager   []bool   `json:"is_original_packager"`
	ProductNDCs          []string `json:"product_ndc"`
	ProductTypes         []string `json:"product_type"`
	Routes               []string `json:"route"`
	SubstanceNames       []string `json:"substance_name"`
	SPLIDs               []string `json:"spl_id"`
	SPLSetIDs            []string `json:"spl_set_id"`
	MechanismsOfAction   []string `json:"pharm_class_moa"`
	ChemicalStructures   []string `json:"pharm_class_cs"`
	PhysiologicEffects   []string `json:"pharm_class_pe"`
	PharmacologicClasses []string `json:"pharm_class_epc"`
	UPCs                 []string `json:"upc"`
	UNIIs                []string `json:"unii"`
	RXCUIs               []string `json:"rxcui"`
}
