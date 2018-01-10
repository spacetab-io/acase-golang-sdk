package acaseSts

import "encoding/xml"

type HotelDescriptionRequestType struct {
	Credentials
	XMLName			xml.Name			`xml:"HotelDescription"`
	HotelCode		int					`xml:"HotelCode,attr"`
	CurrencyCode	int					`xml:"CurrencyCode,attr,omitempty"`
	Options			string				`xml:"Opt,attr,omitempty"`
}

type HotelDescriptionResponseType struct {
	BaseResponse
	SimpleCodeNameType
	XMLName					xml.Name					`xml:"HotelDescription"`
	Zip						string						`xml:"Zip,attr"`
	Address					string						`xml:"Address,attr"`
	Story					string						`xml:"Story,attr"`
	Built					string						`xml:"Built,attr"`
	Reconstructed			string						`xml:"Reconstructed,attr"`
	NumberOfBlocks			string						`xml:"NumberOfBlocks,attr"`
	NumberOfRooms			int							`xml:"NumberOfRooms,attr"`
	NumberOfFloors			int							`xml:"NumberOfFloors,attr"`
	CityCentre				string						`xml:"CityCentre,attr"`
	Underground				string						`xml:"Underground,attr"`
	Description				string						`xml:"Description,attr"`
	Airport					string						`xml:"Airport,attr"`
	RailwayStation			string						`xml:"RailwayStation,attr"`
	RiverPort				string						`xml:"RiverPort,attr"`
	SeaPort					string						`xml:"SeaPort,attr"`
	Map						string						`xml:"Map,attr"`
	SpecialRequirements		string						`xml:"SpecialRequirements,attr"`
	ObjType					ObjTypeType					`xml:"ObjType"`
	City					CityType					`xml:"City"`
	AlternativeCities		AlternativeCitiesType		`xml:"AlternativeCities"`
	AdmUnit1				SimpleCodeNameType			`xml:"AdmUnit1"`
	AdmUnit2				AdmUnit2Type				`xml:"AdmUnit2"`
	TypeOfPlace				SimpleCodeNameType			`xml:"TypeOfPlace"`
	Country					CountryType					`xml:"Country"`
	Position				PositionType				`xml:"Position"`
	Rating					SimpleCodeNameType			`xml:"Rating"`
	Stars					StarsType					`xml:"Stars"`
	TimeCheck				TimeCheckType				`xml:"TimeCheck"`
	Images					ImagesType					`xml:"Images"`
	Amenities				AmenitiesType				`xml:"Amenities"`
	Product					string						`xml:"Product"`
	RoomsList				RoomsListType				`xml:"RoomsList"`
	SpecialRequirementList	SpecialRequirementListType	`xml:"SpecialRequirementList"`
	Objects					ObjectsType					`xml:"Objects"`
}

type ObjectsType struct {
	XMLName	xml.Name		`xml:"Objects"`
	Object	[]ObjectType	`xml:"Object"`
}

type ObjectType struct {
	SimpleCodeNameType
	XMLName			xml.Name			`xml:"Object"`
	Distance		float64				`xml:"Distance,attr"`
	ObjectType		SimpleCodeNameType	`xml:"ObjectType"`
	ObjectSubType	SimpleCodeNameType	`xml:"ObjectSubType"`
	City			CityType			`xml:"City"`
}

type SpecialRequirementListType struct {
	XMLName				xml.Name					`xml:"SpecialRequirementList"`
	SpecialRequirement	[]SpecialRequirementType	`xml:"SpecialRequirement"`
}

type SpecialRequirementType struct {
	XMLName	xml.Name	`xml:"SpecialRequirement"`
	Code	int			`xml:"Code,attr"`
	Text	string		`xml:"Text,attr"`
}

type RoomsListType struct {
	XMLName	xml.Name	`xml:"RoomsList"`
	Room	[]RoomType	`xml:"Room"`
}

type RoomType struct {
	SimpleCodeNameType
	XMLName						xml.Name	`xml:"Room"`
	Description					string		`xml:"Description,attr"`
	MaxNumberOfGuests			int			`xml:"MaxNumberOfGuests,attr"`
	MaxNumberOfExtraBeds		int			`xml:"MaxNumberOfExtraBeds,attr"`
	MaxNumberOfExtraBedsAdult	int			`xml:"MaxNumberOfExtraBedsAdult,attr"`
	MaxNumberOfExtraBedsChild	int			`xml:"MaxNumberOfExtraBedsChild,attr"`
	MaxNumberOfExtraBedsInfant	int			`xml:"MaxNumberOfExtraBedsInfant,attr"`
}

type AmenitiesType struct {
	XMLName	xml.Name		`xml:"Amenities"`
	Amenity	[]AmenityType	`xml:"Amenity"`
}

type AmenityType struct {
	SimpleCodeNameType
	XMLName	xml.Name	`xml:"Amenity"`
	Url		string		`xml:"Url,attr,omitempty"`
}

type ImagesType struct {
	XMLName	xml.Name	`xml:"Images"`
	Image	[]ImageType	`xml:"Image"`
}

type ImageType struct {
	XMLName	xml.Name	`xml:"Image"`
	Url		string		`xml:"Url,attr"`
}

type TimeCheckType struct {
	XMLName	xml.Name	`xml:"TimeCheck"`
	In		string		`xml:"In,attr"`
	Out		string		`xml:"Out,attr"`
}

type StarsType struct {
	SimpleCodeNameType
	XMLName				xml.Name				`xml:"Stars"`
	Value				string					`xml:"Value,attr"`
	Url					string					`xml:"Url,attr"`
	OfficialCertificate	OfficialCertificateType	`xml:"OfficialCertificate"`
}

type OfficialCertificateType struct {
	SimpleCodeNameType
	XMLName		xml.Name		`xml:"OfficialCertificate"`
	Certificate	CertificateType	`xml:"Certificate"`
}

type CertificateType struct {
	SimpleCodeNameType
	XMLName		xml.Name	`xml:"Certificate"`
	ValidBefore	string		`xml:"ValidBefore,attr"`
	Url			string		`xml:"Url,attr"`
}

type AlternativeCitiesType struct {
	XMLName			xml.Name				`xml:"AlternativeCities"`
	AlternativeCity	[]AlternativeCityType	`xml:"AlternativeCity"`
}

type AlternativeCityType struct {
	SimpleCodeNameType
	XMLName		xml.Name			`xml:"AlternativeCity"`
	Country		CountryType			`xml:"Country"`
	AdmUnit1	SimpleCodeNameType	`xml:"AdmUnit1"`
	AdmUnit2	AdmUnit2Type		`xml:"AdmUnit2"`
	TypeOfPlace	SimpleCodeNameType	`xml:"TypeOfPlace"`
}

type ObjTypeType struct {
	SimpleCodeNameType
	XMLName		xml.Name	`xml:"ObjType"`
	NameWhere	string		`xml:"NameWhere,attr,omitempty"`
}
