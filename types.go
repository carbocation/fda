package fda

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

type Reaction struct {
	MedDRA        string `json:"reactionmeddrapt"`
	MedDRAVersion string `json:"reactionmeddraversionpt"`
	Outcome       string `json:"reactionoutcome"`
}

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

type OpenFDA map[string]interface{}
