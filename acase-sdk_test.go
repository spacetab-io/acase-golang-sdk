package acaseSdk

import (
	"context"
	"encoding/xml"
	"math/rand"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strconv"
	"testing"

	"github.com/nbio/st"
	"github.com/tmconsulting/acase-golang-sdk/acaseSts"
)

func TestApi_OrderListRequest(t *testing.T) {
	auth := Auth{
		BuyerId:  "B1",
		UserId:   "U1",
		Password: "P1",
		Language: "L1",
	}

	server := httptest.NewServer(http.HandlerFunc(func(resp http.ResponseWriter, req *http.Request) {
		req.ParseForm()

		body, _ := xml.Marshal(&acaseSts.CityDescriptionType{})

		resp.Header().Set("Content-Type", "text/xml")
		resp.Write(body)
	}))
	defer server.Close()

	var (
		reqMethodName string
		reqMimeType   string
		reqData       []byte
		reqCtxValue   int

		respMethodName string
		respMimeType   string
		respData       []byte
		respCtxValue   int
	)

	api := NewApi(auth, server.URL)

	api.RegisterEventHandler(BeforeRequestSend, func(ctx context.Context, methodName, mimeType string, data []byte) {
		reqMethodName = methodName
		reqMimeType = mimeType

		params, _ := url.ParseQuery(string(data))

		reqData = []byte(params.Get("CityCode"))

		reqCtxValue = ctx.Value("value").(int)

	}).RegisterEventHandler(AfterResponseReceive, func(ctx context.Context, methodName, mimeType string, data []byte) {
		respMethodName = methodName
		respMimeType = mimeType

		respData = data

		respCtxValue = ctx.Value("value").(int)

	})

	value := rand.Intn(99999) + 1
	cityCode := int64(10)

	ctx := context.WithValue(context.Background(), "value", value)

	_, err := api.CityDescriptionRequest(ctx, cityCode)

	st.Expect(t, err, (*AcaseResponseError)(nil))

	st.Expect(t, reqMethodName, "CityDescriptionRequest")
	st.Expect(t, reqMimeType, "application/x-www-form-urlencoded")
	st.Expect(t, reqCtxValue, value)
	st.Expect(t, reqData, []byte(strconv.Itoa(int(cityCode))))

	st.Expect(t, respMethodName, "CityDescriptionRequest")
	st.Expect(t, respMimeType, "text/xml")
	st.Expect(t, respCtxValue, value)
	st.Expect(t, string(respData), `<CityDescription Code="" Name="" Description="" Url=""><Success></Success><Country Code="0" Name=""></Country><AdmUnit1 Code="0" Name=""></AdmUnit1><AdmUnit2 Code="0" Name=""><AdmUnit1 Code="0" Name=""></AdmUnit1></AdmUnit2><TypeOfPlace Code="0" Name=""></TypeOfPlace><Position></Position></CityDescription>`)
}
