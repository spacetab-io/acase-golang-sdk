package test

import (
	"context"
	"testing"

	"github.com/nbio/st"
)

func TestStatusListRequest_Ok(t *testing.T) {
	testRequest("statuslist_response_example.xml", false)
	defer gockOff()

	data, err := acApi.StatusListRequest(context.Background())
	er := getCustomErrorType()
	st.Expect(t, err, er)
	st.Expect(t, len(data.Status), 29)
	st.Expect(t, data.Status[0].Code, 1)
	st.Expect(t, data.Status[0].Name, "Запрос цены")
	st.Expect(t, data.Status[1].Code, 5)
	st.Expect(t, data.Status[1].Name, "Предложение цены")
	st.Expect(t, data.Status[2].Code, 11)
	st.Expect(t, data.Status[2].Name, "Получен")
	st.Expect(t, data.Status[3].Code, 14)
	st.Expect(t, data.Status[3].Name, "Получен")
	st.Expect(t, data.Status[4].Code, 18)
	st.Expect(t, data.Status[4].Name, "В работе")
	st.Expect(t, data.Status[5].Code, 21)
	st.Expect(t, data.Status[5].Name, "В работе")
	st.Expect(t, data.Status[6].Code, 25)
	st.Expect(t, data.Status[6].Name, "Отказан")
	st.Expect(t, data.Status[7].Code, 29)
	st.Expect(t, data.Status[7].Name, "Отказан")
	st.Expect(t, data.Status[8].Code, 32)
	st.Expect(t, data.Status[8].Name, "Лист ожидания")
	st.Expect(t, data.Status[9].Code, 35)
	st.Expect(t, data.Status[9].Name, "Подтвержден")
}

func TestStatusListRequest_Error(t *testing.T) {
	testRequest("statuslist_error_example.xml", true)
	defer gockOff()

	_, err := acApi.StatusListRequest(context.Background())

	st.Expect(t, err.Code, "9998")
	st.Expect(t, err.Message, "Доступ запрещен !")
}
