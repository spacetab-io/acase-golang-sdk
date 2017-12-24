package test

import (
	"testing"

	"github.com/nbio/st"
	"gopkg.in/h2non/gock.v1"
)

func TestCityDescriptionRequest_Ok(t *testing.T)  {
	testRequest("cdresp_example.xml")
	defer gock.Off()

	data, err := acApi.CityDescriptionRequest("2")
	er := getCustomErrorType()
	st.Expect(t, err, er)
	st.Expect(t, data.Code, "2")
	st.Expect(t, data.Name, "Москва")
	st.Expect(t, data.Description, "Город Москва - столица Российской Федерации - был основан в 1147 г. князем Юрием Долгоруким на берегу р.Москва как крепость. Более чем за 850 лет город стал одним из самых крупных городов мира с площадью почти 1000 кв.км и населением около 9 млн.чел. Современная Москва - крупнейший экономический, политический и культурный центр России, где находятся правительство Российской Федерации, многие российские и иностранные компании, образовательные учреждения. Москва славится своими архитектурными и культурными памятниками мирового значения, среди которых: один из самых красивых архитектурных ансамблей мира, бывшая резиденция российских царей - Кремль (15-17 вв.), Собор Василия Блаженного (1551-1561 гг.) - шедевр русской архитектуры, расположенный на Красной Площади - уникальном архитектурном ансамбле. Москва известна своими театрами и музеями, среди которых Большой Театр (1776 г.) - один из ведущих мировых театров оперы и балета, Музей изобразительных искусств им. Пушкина, Третьяковская галерея, Исторический музей, усадьба Коломенское (14 в.) - одна из резиденций русских князей и царей, усадьба Кусково (18 в.). Московский метрополитен считается одним из самых красивых в мире.")
	st.Expect(t, data.Url, "http://xml.acase.ru/cities_images/ru_mow00.jpg")
	st.Expect(t, data.Country.Code, "9")
	st.Expect(t, data.Country.Name, "Россия")
	st.Expect(t, data.Country.Position.Latitude, "61.7731")
	st.Expect(t, data.Country.Position.Longitude, "93.8672")
	st.Expect(t, data.AdmUnit1.Code, 1)
	st.Expect(t, data.AdmUnit1.Name, "Не определено")
	st.Expect(t, data.AdmUnit2.Code, 1)
	st.Expect(t, data.AdmUnit2.Name, "Не определено")
	st.Expect(t, data.TypeOfPlace.Code, 2)
	st.Expect(t, data.TypeOfPlace.Name, "город")
	st.Expect(t, data.Position.Latitude, "55.7731")
	st.Expect(t, data.Position.Longitude, "37.8672")
}

func TestCityDescriptionRequest_Error(t *testing.T) {
	testRequest("cderror_example.xml")
	defer gock.Off()

	_, err := acApi.CityDescriptionRequest("2")

	st.Expect(t, err.Code, "9998")
	st.Expect(t, err.Message, "Доступ запрещен !")
}

