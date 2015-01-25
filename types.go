package fda

// A SafetyReport is one report filed regarding one individual who is on
// one or more medications. Subjectively, while there are many fields in
// each SafetyReport, often the majority of fields are left empty.
type SafetyReport struct {
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

	Reactions []Reaction `json:"reaction"`
	Drugs     []Drug     `json:"drug"`
}

// Reaction represents the reaction that caused the SafetyReport to be filed,
// and the outcome of the event
type Reaction struct {
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
