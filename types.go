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

// OpenFDA reflects OpenFDA-enhanced fields. These are not guaranteed
// to be present due to a variety of reasons (typos in the original
// safety report, etc.)
type OpenFDA struct {
	ApplicationNumbers   []string `json:"application_number"`
	BrandNames           []string `json:"brand_name"`
	DosageForms          []string `json:"dosage_form"`
	GenericNames         []string `json:"generic_name"`
	ManufacturerNames    []string `json:"manufacturer_name"`
	IsOriginalPackager   bool     `json:"is_original_packager"`
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
