package controllers

import (
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/google/uuid"
	"github.com/lengebretsen/go-practice/models"
	"github.com/lengebretsen/go-practice/testing/assert"
)

type mockAddressRepository struct {
	addrs []models.Address
	err   error
}

func (m *mockAddressRepository) FetchAddresses() ([]models.Address, error) {
	if m.addrs != nil {
		return m.addrs, nil
	} else {
		return nil, m.err
	}
}
func (m *mockAddressRepository) FetchOneAddress(id uuid.UUID) (models.Address, error) {
	if len(m.addrs) > 0 {
		return m.addrs[0], nil
	} else {
		return models.Address{}, m.err
	}
}
func (m *mockAddressRepository) InsertAddress(models.Address) (models.Address, error) {
	panic("IMPLEMENT ME")
}
func (m *mockAddressRepository) UpdateAddress(models.Address) (models.Address, error) {
	panic("IMPLEMENT ME")
}
func (m *mockAddressRepository) DeleteAddress(id uuid.UUID) error {
	panic("IMPLEMENT ME")
}
func (m *mockAddressRepository) FindAddressesByUserId(userId uuid.UUID) ([]models.Address, error) {
	panic("IMPLEMENT ME")
}

func TestFetchAddressesRoute(t *testing.T) {
	type test struct {
		mockResult mockAddressRepository
		wantedCode int
		wantedBody []models.Address
		wantedErr  ApiError
	}

	tests := []test{
		{
			mockResult: mockAddressRepository{
				addrs: []models.Address{
					{
						Id:     uuid.MustParse("34ecb0a8-7184-42fa-8840-6fa5c496d161"),
						UserId: uuid.MustParse("80e4de8a-91c4-46cc-a66d-23d3cf364036"),
						Street: "123 A St.",
						City:   "Anytown",
						State:  "GA",
						Zip:    "30033",
						Type:   "HOME",
					},
					{
						Id:     uuid.MustParse("160ded2d-3074-417f-9bc1-a0d44a403cf2"),
						UserId: uuid.MustParse("80e4de8a-91c4-46cc-a66d-23d3cf364036"),
						Street: "456 B St.",
						City:   "Anothertown",
						State:  "TN",
						Zip:    "38028",
						Type:   "WORK",
					},
				},
			},
			wantedCode: 200,
			wantedBody: []models.Address{
				{
					Id:     uuid.MustParse("34ecb0a8-7184-42fa-8840-6fa5c496d161"),
					UserId: uuid.MustParse("80e4de8a-91c4-46cc-a66d-23d3cf364036"),
					Street: "123 A St.",
					City:   "Anytown",
					State:  "GA",
					Zip:    "30033",
					Type:   "HOME",
				},
				{
					Id:     uuid.MustParse("160ded2d-3074-417f-9bc1-a0d44a403cf2"),
					UserId: uuid.MustParse("80e4de8a-91c4-46cc-a66d-23d3cf364036"),
					Street: "456 B St.",
					City:   "Anothertown",
					State:  "TN",
					Zip:    "38028",
					Type:   "WORK",
				},
			},
		},
		{
			mockResult: mockAddressRepository{err: errors.New("Kaboom!!")},
			wantedCode: 500,
			wantedErr:  ApiError{Message: "Error fetching address records"},
		},
	}

	for _, testCase := range tests {
		router := SetupRouter()
		RegisterRoutes(router, nil, &testCase.mockResult)

		w := httptest.NewRecorder()
		req, _ := http.NewRequest("GET", "/addresses/", nil)
		router.ServeHTTP(w, req)

		assert.Equal(t, testCase.wantedCode, w.Code)

		if testCase.wantedBody != nil {
			//Unmarshal json resp
			parsedResp := []models.Address{}
			json.Unmarshal(w.Body.Bytes(), &parsedResp)

			//compare expected slice w/ unmarshaled response
			assert.Equal(t, parsedResp, testCase.wantedBody)
		} else {
			//Unmarshal json resp into ApiError response
			parsedResp := ApiError{}
			json.Unmarshal(w.Body.Bytes(), &parsedResp)
			assert.Equal(t, parsedResp, testCase.wantedErr)
		}
	}
}

func TestFetchSingleAddressRoute(t *testing.T) {
	type test struct {
		mockResult mockAddressRepository
		addrId     string
		wantedCode int
		wantedBody models.Address
		wantedErr  ApiError
	}

	tests := []test{
		{
			mockResult: mockAddressRepository{
				addrs: []models.Address{
					{
						Id:     uuid.MustParse("34ecb0a8-7184-42fa-8840-6fa5c496d161"),
						UserId: uuid.MustParse("80e4de8a-91c4-46cc-a66d-23d3cf364036"),
						Street: "123 A St.",
						City:   "Anytown",
						State:  "GA",
						Zip:    "30033",
						Type:   "HOME",
					},
				},
			},
			addrId:     "34ecb0a8-7184-42fa-8840-6fa5c496d161",
			wantedCode: 200,
			wantedBody: models.Address{
				Id:     uuid.MustParse("34ecb0a8-7184-42fa-8840-6fa5c496d161"),
				UserId: uuid.MustParse("80e4de8a-91c4-46cc-a66d-23d3cf364036"),
				Street: "123 A St.",
				City:   "Anytown",
				State:  "GA",
				Zip:    "30033",
				Type:   "HOME",
			},
		},
		{
			addrId:     "bob",
			wantedCode: 400,
			wantedErr:  ApiError{Message: "Id [bob] is not a valid UUID"},
		},
		{
			addrId:     "34ecb0a8-7184-42fa-8840-6fa5c496d161",
			mockResult: mockAddressRepository{err: models.ErrModelNotFound},
			wantedCode: 404,
			wantedErr:  ApiError{Message: "No address exists with Id [34ecb0a8-7184-42fa-8840-6fa5c496d161]"},
		},
		{
			addrId:     "34ecb0a8-7184-42fa-8840-6fa5c496d161",
			mockResult: mockAddressRepository{err: errors.New("Kaboom!!")},
			wantedCode: 500,
			wantedErr:  ApiError{Message: "Error fetching address record with Id [34ecb0a8-7184-42fa-8840-6fa5c496d161]"},
		},
	}

	for _, testCase := range tests {
		router := SetupRouter()
		RegisterRoutes(router, nil, &testCase.mockResult)

		w := httptest.NewRecorder()
		req, _ := http.NewRequest("GET", "/addresses/"+testCase.addrId, nil)
		router.ServeHTTP(w, req)

		assert.Equal(t, testCase.wantedCode, w.Code)

		if testCase.wantedBody != (models.Address{}) {
			//Unmarshal json resp
			parsedResp := models.Address{}
			json.Unmarshal(w.Body.Bytes(), &parsedResp)

			//compare expected w/ unmarshaled response
			assert.Equal(t, parsedResp, testCase.wantedBody)
		} else {
			//Unmarshal json resp into ApiError response
			parsedResp := ApiError{}
			json.Unmarshal(w.Body.Bytes(), &parsedResp)
			assert.Equal(t, parsedResp, testCase.wantedErr)
		}
	}
}
