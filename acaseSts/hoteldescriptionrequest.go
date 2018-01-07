package acaseSts

import "encoding/xml"

type HotelDescriptionRequestType struct {
	XMLName			xml.Name			`xml:"HotelDescription"`
	BuyerId			string				`xml:"BuyerId,attr"`
	UserId			string				`xml:"UserId,attr"`
	Password		string				`xml:"Password,attr"`
	Language		LanguageTypeEnum	`xml:"Language,attr,omitempty"`
	HotelCode		int					`xml:"HotelCode,attr"`
	CurrencyCode	int					`xml:"CurrencyCode,attr,omitempty"`
	Options			string				`xml:"Opt,attr,omitempty"`
}

type HotelDescriptionResponseType struct {
	XMLName					xml.Name					`xml:"HotelDescription"`
	Success					string						`xml:"Success"`
	Error					ErrorType					`xml:"Error,omitempty"`
	Code					int							`xml:"Code,attr"`
	Name					string						`xml:"Name,attr"`
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
	AdmUnit1				AdmUnit1Type				`xml:"AdmUnit1"`
	AdmUnit2				AdmUnit2Type				`xml:"AdmUnit2"`
	TypeOfPlace				TypeOfPlaceType				`xml:"TypeOfPlace"`
	Country					CountryType					`xml:"Country"`
	Position				PositionType				`xml:"Position"`
	Rating					RatingType					`xml:"Rating"`
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
	XMLName			xml.Name			`xml:"Object"`
	Code			int					`xml:"Code,attr"`
	Name			string				`xml:"Name,attr"`
	Distance		float64				`xml:"Distance,attr"`
	ObjectType		ObjectTypeType		`xml:"ObjectType"`
	ObjectSubType	ObjectSubTypeType	`xml:"ObjectSubType"`
	City			CityType			`xml:"City"`
}

type ObjectSubTypeType struct {
	XMLName	xml.Name	`xml:"ObjectSubType"`
	Code	int			`xml:"Code,attr"`
	Name	string		`xml:"Name,attr"`
}

type ObjectTypeType struct {
	XMLName	xml.Name	`xml:"ObjectType"`
	Code	int			`xml:"Code,attr"`
	Name	string		`xml:"Name,attr"`
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
	XMLName						xml.Name	`xml:"Room"`
	Code						int			`xml:"Code,attr"`
	Name						string		`xml:"Name,attr"`
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
	XMLName	xml.Name	`xml:"Amenity"`
	Code	int			`xml:"Code,attr"`
	Name	string		`xml:"Name,attr"`
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
	XMLName				xml.Name				`xml:"Stars"`
	Code				int						`xml:"Code,attr"`
	Name				string					`xml:"Name,attr"`
	Value				string					`xml:"Value,attr"`
	Url					string					`xml:"Url,attr"`
	OfficialCertificate	OfficialCertificateType	`xml:"OfficialCertificate"`
}

type OfficialCertificateType struct {
	XMLName		xml.Name		`xml:"OfficialCertificate"`
	Code		int				`xml:"Code,attr"`
	Name		string			`xml:"Name,attr"`
	Certificate	CertificateType	`xml:"Certificate"`
}

type CertificateType struct {
	XMLName		xml.Name	`xml:"Certificate"`
	Code		int			`xml:"Code,attr"`
	Name		string		`xml:"Name,attr"`
	ValidBefore	string		`xml:"ValidBefore,attr"`
	Url			string		`xml:"Url,attr"`
}

type RatingType struct {
	XMLName	xml.Name	`xml:"Rating"`
	Code	int			`xml:"Code,attr"`
	Name	string		`xml:"Name,attr"`
}

type AlternativeCitiesType struct {
	XMLName			xml.Name				`xml:"AlternativeCities"`
	AlternativeCity	[]AlternativeCityType	`xml:"AlternativeCity"`
}

type AlternativeCityType struct {
	XMLName		xml.Name		`xml:"AlternativeCity"`
	Code		int				`xml:"Code,attr"`
	Name		string			`xml:"Name,attr"`
	Country		CountryType		`xml:"Country"`
	AdmUnit1	AdmUnit1Type	`xml:"AdmUnit1"`
	AdmUnit2	AdmUnit2Type	`xml:"AdmUnit2"`
	TypeOfPlace	TypeOfPlaceType	`xml:"TypeOfPlace"`
}

type ObjTypeType struct {
	XMLName		xml.Name	`xml:"ObjType"`
	Code		int			`xml:"Code,attr"`
	Name		string		`xml:"Name,attr"`
	NameWhere	string		`xml:"NameWhere,attr,omitempty"`
}
