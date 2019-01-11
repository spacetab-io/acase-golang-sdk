package test

import (
	"context"
	"testing"

	"github.com/nbio/st"
)

func TestStarListRequest_Ok(t *testing.T) {
	testRequest("starlist_response_example.xml", false)
	defer gockOff()

	data, err := acApi.StarListRequest(context.Background(), 0, "", "")
	er := getCustomErrorType()
	st.Expect(t, err, er)
	st.Expect(t, len(data.Star), 7)
	st.Expect(t, data.Star[0].Code, 1)
	st.Expect(t, data.Star[0].Name, "Уровень 5 из 5")
	st.Expect(t, data.Star[0].Value, "5")
	st.Expect(t, data.Star[1].Code, 2)
	st.Expect(t, data.Star[1].Name, "Уровень 4 из 5")
	st.Expect(t, data.Star[1].Value, "4")
	st.Expect(t, data.Star[2].Code, 3)
	st.Expect(t, data.Star[2].Name, "Уровень 3 из 5")
	st.Expect(t, data.Star[2].Value, "3")
	st.Expect(t, data.Star[3].Code, 4)
	st.Expect(t, data.Star[3].Name, "Уровень 2 из 5")
	st.Expect(t, data.Star[3].Value, "2")
	st.Expect(t, data.Star[4].Code, 5)
	st.Expect(t, data.Star[4].Name, "Уровень 1 из 5")
	st.Expect(t, data.Star[4].Value, "1")
	st.Expect(t, data.Star[5].Code, 6)
	st.Expect(t, data.Star[5].Name, "-")
	st.Expect(t, data.Star[5].Value, "-")
	st.Expect(t, data.Star[6].Code, 99)
	st.Expect(t, data.Star[6].Name, "-")
	st.Expect(t, data.Star[6].Value, "-")
}

func TestStarListRequest_Error(t *testing.T) {
	testRequest("starlist_error_example.xml", true)
	defer gockOff()

	_, err := acApi.StarListRequest(context.Background(), 0, "", "")

	st.Expect(t, err.Code, "9998")
	st.Expect(t, err.Message, "Доступ запрещен !")
}
