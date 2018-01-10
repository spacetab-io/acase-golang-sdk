Unofficial golang SDK for integration with hotels reservation API, powered by "Academ service" provider.

[![develop](https://circleci.com/gh/tmconsulting/acase-golang-sdk/tree/develop.svg?style=shield)](https://circleci.com/gh/tmconsulting/acase-golang-sdk/tree/develop)

---------------

This SDK is used to connect to the Academservice and use it methods for a 3-rd party API.

* [Installation](#installation)
* [Get started](#get-started)
* [Api Reference](#api-reference)
* [Contact us](#contact-us)
* [License](#license)

### Installation
```
go get -u github.com/tmconsulting/acase-golang-sdk
```

### Get started

To start you will need to have credentials for the Academservice. <br>
Import package
```golang
import "github.com/tmconsulting/acase-golang-sdk"
```

Api initialization
```golang
credentials := acaseSdk.Auth{
	BuyerId:"BuyerId",
	UserId:"UserId",
	Password:"Password",
	Language:"RU",
}
api := acaseSdk.NewApi(credentials, "http://test-www.acase.ru/xml/form.jsp")
```

Now you can call methods, for example list of citizenships:
```golang
result, err := api.CitizenshipListRequest()
```

As a result, the pointer to structure is returned.
In the example above: *acaseSts.CitizenshipListType.
Every method the second value return an error object of type *AcaseResponseError or nil.

### API Reference

Implemented methods:

| Method name | Description |
|-------------|-------------|
| AdmUnit1Request | Returns the list of administrative units of the first level |
| AdmUnit2Request | Returns the list of administrative units of the second level |
| CitizenshipListRequest | Returns the list of citizenships |
| CityDescriptionRequest | Returns the detailed description of the selected city |
| CityListRequest | Returns the list of cities where Academservice sells hotels |
| CountryDescriptionRequest | Returns the detailed description of the selected country |
| ClientCategoryListRequest | Returns the list of client categories |
| CountryListRequest | Returns the list of countries where Academservice sells hotels |
| CurrencyListRequest | Returns the list of available currencies |
| CustomerRequestCreate | Allows to create end customer |
| CustomerRequestUpdate | Allows to update end customer |
| CustomerRequestDelete | Allows to delete end customer |
| CustomerRequestList | Allows to get list of end customer |
| CustomerRequestInfo | Allows to get information about end customer |
| HotelAmenityListRequest | Returns the list of hotel amenities |
| HotelDescriptionRequest | Returns the detailed description of the selected hotel |
| HotelListRequest | Returns the list of hotels to sell |
| HotelPricingRequest2 | Returns the calculated price of the accommodation for the specified accommodation plan |
| HotelProductRequest | Returns the calculated price of the accommodation |
| HotelSearchRequest | Returns the list of tour products which meet the specified criteria |
| MealRequest | Returns the list of meal |
| MealTypeRequest | Returns the list of meal types |
| ObjectRequest | Returns the list of objects |
| ObjectSubTypeRequest | Returns the list of object's subtypes |
| ObjectTypeRequest | Returns the list of object's types |
| ObjTypeListRequest | List of types accommodation object |
| OrderDocRequest | Allows to get reservation documents on request |
| OrderInfoNotifyRequest | Notify buyer's computer system about changes in orders |
| OrderInfoAwocNotifyRequest | Notify buyer's computer system about changes in orders (Hotels with prices on request) |
| OrderListRequest | Returns the list of orders which meet the specified criteria |
| OrderRequest | Create, amend or cancel an order |
| OrderAwocRequest | Create, amend or cancel an order (Hotels with prices on request) |
| RateGroupRequest | Returns the list of group of rates |
| RouteRequest | Returns the list of route |
| PenaltyReasonRequest | Returns the list of penalty reasons |
| SpecialOfferTypeRequest | Returns the list of type of action |
| StarListRequest | Returns the list of hotel categories |
| StatusListRequest | Returns the list of technological statuses services |
| TypeOfPlaceRequest | Returns the list of settlements type |

All methos of Academservice describe [here](http://www.acase.ru/xml/en-xml.jsp)

### Contact us

If you have any issues or questions regarding the API or the SDK it self, you are welcome to create an issue, or
You can write an Email to `atkachev@gmail.com` or `roquie0@gmail.com`

### License.

SDK is released under the [MIT License](./LICENSE).
