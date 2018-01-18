package test

import (
	"testing"
	"github.com/nbio/st"
)

func TestObjTypeListRequest_Ok(t *testing.T) {
	testRequest("objtypelist_response_example.xml", false)
	defer gockOff()

	data, err := acApi.ObjTypeListRequest("", "")
	er := getCustomErrorType()
	st.Expect(t, err, er)
	st.Expect(t, len(data.ObjType), 6)
	st.Expect(t, data.ObjType[0].Code, 1)
	st.Expect(t, data.ObjType[0].Name, "-")
	st.Expect(t, data.ObjType[1].Code, 2)
	st.Expect(t, data.ObjType[1].Name, "Гостиница")
	st.Expect(t, data.ObjType[2].Code, 3)
	st.Expect(t, data.ObjType[2].Name, "Санаторий")
	st.Expect(t, data.ObjType[3].Code, 4)
	st.Expect(t, data.ObjType[3].Name, "Дом отдыха")
	st.Expect(t, data.ObjType[4].Code, 5)
	st.Expect(t, data.ObjType[4].Name, "Хостел")
	st.Expect(t, data.ObjType[5].Code, 6)
	st.Expect(t, data.ObjType[5].Name, "Пансионат")
}

func TestObjTypeListRequest_Error(t *testing.T) {
	testRequest("objtypelist_error_example.xml", true)
	defer gockOff()

	_, err := acApi.ObjTypeListRequest("", "")

	st.Expect(t, err.Code, "9998")
	st.Expect(t, err.Message, "Доступ запрещен !")
}

