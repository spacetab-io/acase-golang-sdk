package test

import (
	"testing"

	"github.com/nbio/st"
	"gopkg.in/h2non/gock.v1"
)

func TestHotelAmenityListRequest_Ok(t *testing.T)  {
	testRequest("hotelamenitylist_response_example.xml")
	defer gock.Off()

	data, err := acApi.HotelAmenityListRequest(0, "")
	er := getCustomErrorType()
	st.Expect(t, err, er)
	st.Expect(t, len(data.HotelAmenity), 90)
	st.Expect(t, data.HotelAmenity[0].Code, 1)
	st.Expect(t, data.HotelAmenity[0].Name, "Отель в зоне отдыха")
	st.Expect(t, data.HotelAmenity[0].Url, "http://images.acase.ru/signs_images/hotel_in_recreation_area.gif")
	st.Expect(t, data.HotelAmenity[1].Code, 2)
	st.Expect(t, data.HotelAmenity[1].Name, "Парковка охраняемая")
	st.Expect(t, data.HotelAmenity[1].Url, "http://images.acase.ru/signs_images/guarded_parking.gif")
	st.Expect(t, data.HotelAmenity[2].Code, 3)
	st.Expect(t, data.HotelAmenity[2].Name, "Парковка неохраняемая")
	st.Expect(t, data.HotelAmenity[2].Url, "http://images.acase.ru/signs_images/nonguarded_parking.gif")
	st.Expect(t, data.HotelAmenity[3].Code, 4)
	st.Expect(t, data.HotelAmenity[3].Name, "Поднос багажа")
	st.Expect(t, data.HotelAmenity[3].Url, "http://images.acase.ru/signs_images/luggage_service.gif")
	st.Expect(t, data.HotelAmenity[4].Code, 6)
	st.Expect(t, data.HotelAmenity[4].Name, "Сейф в отеле")
	st.Expect(t, data.HotelAmenity[4].Url, "http://images.acase.ru/signs_images/safe_deposit_box_at_reception.gif")
	st.Expect(t, data.HotelAmenity[5].Code, 7)
	st.Expect(t, data.HotelAmenity[5].Name, "Пункт обмена валюты")
	st.Expect(t, data.HotelAmenity[5].Url, "http://images.acase.ru/signs_images/currency_exchange.gif")
	st.Expect(t, data.HotelAmenity[6].Code, 69)
	st.Expect(t, data.HotelAmenity[6].Name, "Банкомат")
	st.Expect(t, data.HotelAmenity[6].Url, "http://images.acase.ru/signs_images/atm.gif")
	st.Expect(t, data.HotelAmenity[7].Code, 8)
	st.Expect(t, data.HotelAmenity[7].Name, "Лифт")
	st.Expect(t, data.HotelAmenity[7].Url, "http://images.acase.ru/signs_images/elevator.gif")
	st.Expect(t, data.HotelAmenity[8].Code, 9)
	st.Expect(t, data.HotelAmenity[8].Name, "Ресторан")
	st.Expect(t, data.HotelAmenity[8].Url, "http://images.acase.ru/signs_images/restaurant.gif")
	st.Expect(t, data.HotelAmenity[9].Code, 10)
	st.Expect(t, data.HotelAmenity[9].Name, "Бар")
	st.Expect(t, data.HotelAmenity[9].Url, "http://images.acase.ru/signs_images/bar.gif")
}

func TestHotelAmenityListRequest_Error(t *testing.T) {
	testRequest("hotelamenitylist_error_example.xml")
	defer gock.Off()

	_, err := acApi.HotelAmenityListRequest(0, "")

	st.Expect(t, err.Code, "9998")
	st.Expect(t, err.Message, "Доступ запрещен !")
}

