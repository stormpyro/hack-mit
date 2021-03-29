package models

// Solicitud de licencia de edificación
type LicenceApplication struct {
	ID               string   `json:"_id" bson:"_id"`
	DateInit         int64    `json:"dateInit"`
	TransactionType  string   `json:"transactionType"`
	ConstructionType string   `json:"constructionType"`
	ApprovalModality []string `json:"approvalModality"`
	Annexes          []string `json:"annexes"`
}

// Del solicitante
type Applicant struct {
	ID                string            `json:"_id" bson:"_id"`
	Type              string            `json:"type"`
	NaturalPerson     NaturalPerson     `json:"naturalPerson"`
	LegalPerson       LegalPerson       `json:"legalPerson"`
	LegalRepresentant LegalRepresentant `json:"legalRepresentant"`
}

// Persona Natural
type NaturalPerson struct {
	LastName        string  `json:"lastName"`
	MothersLastName string  `json:"mothersLastName"`
	Names           string  `json:"names"`
	DNI             int64   `json:"dni"`
	Address         Address `json:"address"`
	CivilState      string  `json:"civilState"`
	Spouse          Spouse  `json:"spouse"`
}

type Address struct {
	Apartment    string `json:"apartment"`
	Province     string `json:"province"`
	District     string `json:"district"`
	Urbanization string `json:"urbanization"`
	Apple        int64  `json:"apple"`
	Portion      int64  `json:"portion"`
	SubPortion   int64  `json:"subPortion"`
	Street       string `json:"street"`
	StreetNumber int64  `json:"streetNumber"`
	Interior     int64  `json:"interior"`
}

// Cónyuge
type Spouse struct {
	LastName        string `json:"lastName"`
	MothersLastName string `json:"MothersLastName"`
	Names           string `json:"names"`
	DNI             int64  `json:"dni"`
}

// Persona juridica
type LegalPerson struct {
	BusinessName string  `json:"businessName"`
	RUC          int64   `json:"ruc"`
	Address      Address `json:"address"`
}

// Respresentante legal
type LegalRepresentant struct {
	Type              string  `json:"type"`
	LastName          string  `json:"lastName"`
	MothersLastName   string  `json:"mothersLastName"`
	Names             string  `json:"names"`
	DNI               int64   `json:"dni"`
	Address           Address `json:"address"`
	PowerInscribedIn  string  `json:"powerInscribedIn"`
	MandatRegister    bool    `json:"mandatRegister"`
	MercantilRegister bool    `json:"mercantilRegister"`
	OfficeRegister    string  `json:"officeRegister"`
}

// Del terreno
type Ground struct {
	ID                    string                `json:"_id" bson:"_id"`
	Ubication             Address               `json:"ubication"`
	PerimeterMeasurements PerimeterMeasurements `json:"perimeterMeasurements"`
}

type PerimeterMeasurements struct {
	TotalArea         float64 `json:"totalArea"`
	FromTheFront      float64 `json:"fromTheFront"`
	FromTheLeft       float64 `json:"fromTheLeft"`
	FromTheBackground float64 `json:"fromTheBackground"`
	FromTheRight      float64 `json:"fromTheRight"`
}

// De la titulación del predio
type LandTitling struct {
	ID               string           `json:"_id" bson:"_id"`
	Ground           LandGround       `json:"ground"`
	ExistingBuilding ExistingBuilding `json:"existingBuilding"`
}

type LandGround struct {
	PropertyRegime     string `json:"propertyRegime"`
	CondominiumsNumber int64  `json:"condominiumsNumber"`
	LandRegistry       string `json:"landRegistry"`
	LandCode           int64  `json:"landCode"`
	PowerInscribedIn   string `json:"powerInscribedIn"`
}

type ExistingBuilding struct {
	Type             string `json:"type"`
	PreviousBuilding int64  `json:"previousBuilding"`
	LandRegistry     string `json:"landRegistry"`
	LandCode         int64  `json:"landCode"`
	PowerInscribedIn string `json:"powerInscribedIn"`
}
