package test

import (
	"context"
	"testing"

	"github.com/nbio/st"
)

func TestHotelProductRequest_Ok(t *testing.T) {
	testRequest("hotelproduct_response_example.xml", false)
	defer gockOff()

	data, err := acApi.HotelProductRequest(context.Background(), 0, 0, 0, 0, 0,
		0, 0, "", "", "", "")
	er := getCustomErrorType()
	st.Expect(t, err, er)
	st.Expect(t, data.Code, 800300)
	st.Expect(t, data.Name, "Аквариум Отель (Крокус Экспо)")
	st.Expect(t, data.StartDate, "12.06.2009")
	st.Expect(t, data.EndDate, "13.06.2009")
	st.Expect(t, data.NumberOfGuests, 2)
	st.Expect(t, data.NumberOfExtraBedsAdult, 1)
	st.Expect(t, data.NumberOfExtraBedsChild, 0)
	st.Expect(t, data.NumberOfExtraBedsInfant, 0)
	st.Expect(t, data.SpecialRequirements, "1. Гостиница имеет право взимать регистрационный сбор или депозит за дополнительные услуги при размещении. Размер регистрационного сбора гостиница устанавливает самостоятельно..")
	st.Expect(t, data.BuyerId, "MyCompanyId")
	st.Expect(t, data.UserId, "MyUserId")
	st.Expect(t, data.Password, "MyPassword")
	st.Expect(t, string(data.Language), "ru")
	st.Expect(t, data.Country.Code, 9)
	st.Expect(t, data.Country.Name, "Россия")
	st.Expect(t, data.Country.Position.Latitude, "57.5158")
	st.Expect(t, data.Country.Position.Longitude, "82.6172")
	st.Expect(t, data.AdmUnit1.Code, 1)
	st.Expect(t, data.AdmUnit1.Name, "Не определено")
	st.Expect(t, data.AdmUnit2.Code, 1)
	st.Expect(t, data.AdmUnit2.Name, "Не определено")
	st.Expect(t, data.TypeOfPlace.Code, 2)
	st.Expect(t, data.TypeOfPlace.Name, "город")
	st.Expect(t, data.City.Code, 2)
	st.Expect(t, data.City.Name, "Москва")
	st.Expect(t, data.City.Position.Latitude, "55.724")
	st.Expect(t, data.City.Position.Longitude, "37.7628")
	st.Expect(t, data.Position.Latitude, "55.828133")
	st.Expect(t, data.Position.Longitude, "37.389736")
	st.Expect(t, data.ObjType.Code, 2)
	st.Expect(t, data.ObjType.Name, "Гостиница")
	st.Expect(t, data.Currency.Code, 3)
	st.Expect(t, data.Currency.Name, "USD")
	st.Expect(t, data.Infants.AgeTo, 2)
	st.Expect(t, data.Infants.AgeFrom, 0)
	st.Expect(t, data.Infants.UseThisAge.Code, 1)
	st.Expect(t, data.Infants.UseThisAge.Name, "Да")
	st.Expect(t, data.Children.UseThisAge.Code, 2)
	st.Expect(t, data.Children.UseThisAge.Name, "Нет")
	st.Expect(t, len(data.SpecialRequirementList.SpecialRequirement), 1)
	st.Expect(t, data.SpecialRequirementList.SpecialRequirement[0].Code, 595)
	st.Expect(t, data.SpecialRequirementList.SpecialRequirement[0].Text, "Гостиница имеет право взимать регистрационный сбор или депозит за дополнительные услуги при размещении. Размер регистрационного сбора гостиница устанавливает самостоятельно.")
	st.Expect(t, len(data.HotelProductList.Items), 7)
	st.Expect(t, data.HotelProductList.Items[0].Code, "900457")
	st.Expect(t, data.HotelProductList.Items[0].RoomName, "Studio /Double")
	st.Expect(t, data.HotelProductList.Items[0].RoomCode, "900252")
	st.Expect(t, data.HotelProductList.Items[0].RateCode, "900092")
	st.Expect(t, data.HotelProductList.Items[0].RateName, "")
	st.Expect(t, data.HotelProductList.Items[0].RateDescription, "")
	st.Expect(t, data.HotelProductList.Items[0].IsHourRate, "2")
	st.Expect(t, data.HotelProductList.Items[0].MaxHours, 0)
	st.Expect(t, data.HotelProductList.Items[0].StartTime, "")
	st.Expect(t, data.HotelProductList.Items[0].EndTime, "")
	st.Expect(t, data.HotelProductList.Items[0].RateGroupCode, "4")
	st.Expect(t, data.HotelProductList.Items[0].RateGroupName, "FIT")
	st.Expect(t, data.HotelProductList.Items[0].MinimumRetailPrice, 250.00)
	st.Expect(t, data.HotelProductList.Items[0].Price, 216.00)
	st.Expect(t, data.HotelProductList.Items[0].TravelAgencyCommission, 0.00)
	st.Expect(t, data.HotelProductList.Items[0].FullVATInPrice, "1")
	st.Expect(t, data.HotelProductList.Items[0].VATIncludedInPrice, 32.95)
	st.Expect(t, data.HotelProductList.Items[0].DeadlineDate, "11.06.2009")
	st.Expect(t, data.HotelProductList.Items[0].DeadlineTimeLoc, "11.06.2009 13:00")
	st.Expect(t, data.HotelProductList.Items[0].DeadlineTimeSys, "11.06.2009 13:00")
	st.Expect(t, data.HotelProductList.Items[0].DeadlineTimeUTC, "11.06.2009 10:00")
	st.Expect(t, data.HotelProductList.Items[0].PenaltySize, 216.00)
	st.Expect(t, data.HotelProductList.Items[0].MaxNumberOfGuests, 2)
	st.Expect(t, data.HotelProductList.Items[0].MaxNumberOfExtraBeds, 1)
	st.Expect(t, data.HotelProductList.Items[0].MaxNumberOfExtraBedsAdult, 1)
	st.Expect(t, data.HotelProductList.Items[0].MaxNumberOfExtraBedsChild, 0)
	st.Expect(t, data.HotelProductList.Items[0].MaxNumberOfExtraBedsInfant, 0)
	st.Expect(t, len(data.HotelProductList.Items[0].SpecialOfferList.Items), 1)
	st.Expect(t, data.HotelProductList.Items[0].SpecialOfferList.Items[0].StayNights, 5)
	st.Expect(t, data.HotelProductList.Items[0].SpecialOfferList.Items[0].PayNights, 4)
	st.Expect(t, data.HotelProductList.Items[0].SpecialOfferList.Items[0].SpecialOfferType.Code, 17)
	st.Expect(t, data.HotelProductList.Items[0].SpecialOfferList.Items[0].SpecialOfferType.Name, "Бесплатные ночи")
	st.Expect(t, data.HotelProductList.Items[0].SpecialOfferList.Items[0].SpecialOfferType.Id, 3)
	st.Expect(t, data.HotelProductList.Items[0].SpecialOfferList.Items[0].IsMultiple.Code, 2)
	st.Expect(t, data.HotelProductList.Items[0].SpecialOfferList.Items[0].IsMultiple.Name, "Нет")
	st.Expect(t, len(data.HotelProductList.Items[0].Meals.Items), 1)
	st.Expect(t, data.HotelProductList.Items[0].Meals.Items[0].Code, 0)
	st.Expect(t, data.HotelProductList.Items[0].Meals.Items[0].Name, "Завтрак \"Шведский стол\"")
	st.Expect(t, data.HotelProductList.Items[0].Meals.Items[0].Price, 0.00)
	st.Expect(t, data.HotelProductList.Items[0].Meals.Items[0].IncludedInPrice.Code, 1)
	st.Expect(t, data.HotelProductList.Items[0].Meals.Items[0].IncludedInPrice.Name, "Да")
	st.Expect(t, data.HotelProductList.Items[0].Availability.Code, 2)
	st.Expect(t, data.HotelProductList.Items[0].Availability.Name, "Свободная продажа")
	st.Expect(t, data.HotelProductList.Items[0].Availability.Allotment.Code, "9500001+900392+9500001+2")
	st.Expect(t, data.HotelProductList.Items[0].Availability.Allotment.Name, "")
	st.Expect(t, data.HotelProductList.Items[0].Availability.Allotment.Rooms, 0)
	st.Expect(t, len(data.HotelProductList.Items[0].WhereToPayList.Items), 1)
	st.Expect(t, data.HotelProductList.Items[0].WhereToPayList.Items[0].Code, 3)
	st.Expect(t, data.HotelProductList.Items[0].WhereToPayList.Items[0].Name, "Оплата Академсервису согласно договору")
}

func TestHotelProductRequest_Error(t *testing.T) {
	testRequest("hotelproduct_error_example.xml", true)
	defer gockOff()

	_, err := acApi.HotelProductRequest(context.Background(), 0, 0, 0, 0, 0,
		0, 0, "", "", "", "")

	st.Expect(t, err.Code, "9998")
	st.Expect(t, err.Message, "Доступ запрещен !")
}
