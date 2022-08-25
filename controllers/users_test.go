package controllers

import (
	"bytes"
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/google/uuid"
	"github.com/lengebretsen/go-practice/models"
	"github.com/lengebretsen/go-practice/util/assert"
)

type mockUserRepository struct {
	users []models.User
	err   error
}

func (m *mockUserRepository) SelectAllUsers() ([]models.User, error) {
	return m.users, m.err
}
func (m *mockUserRepository) SelectOneUser(id uuid.UUID) (models.User, error) {
	if len(m.users) > 0 {
		return m.users[0], m.err
	} else {
		return models.User{}, m.err
	}
}
func (m *mockUserRepository) InsertUser(usr models.User) (models.User, error) {
	if usr.Id == uuid.Nil {
		log.Fatalln("UUID value for new user was nil")
	}
	if len(m.users) > 0 {
		//overwrite Id val so we can match to a known val in the test
		usr.Id = m.users[0].Id
		return usr, m.err
	} else {
		return models.User{}, m.err
	}
}
func (m *mockUserRepository) UpdateUser(usr models.User) (models.User, error) {
	if len(m.users) > 0 {
		return usr, m.err
	} else {
		return models.User{}, m.err
	}
}
func (m *mockUserRepository) DeleteUser(id uuid.UUID) error {
	return m.err
}

func TestFetchUsersRoute(t *testing.T) {
	type test struct {
		mockResult  mockUserRepository
		wantedCode  int
		wantedBody  []models.User
		wantedError ApiError
	}

	tests := []test{
		{
			mockResult: mockUserRepository{
				users: []models.User{
					{Id: uuid.MustParse("493adb28-9da1-4db8-893d-73cc2d7bd4ee"), FirstName: "Test", LastName: "User"},
					{Id: uuid.MustParse("ddcfdd51-9715-4d4d-bea3-317cccea16ea"), FirstName: "Some", LastName: "Guy"},
				},
			},
			wantedCode: 200,
			wantedBody: []models.User{
				{Id: uuid.MustParse("493adb28-9da1-4db8-893d-73cc2d7bd4ee"), FirstName: "Test", LastName: "User"},
				{Id: uuid.MustParse("ddcfdd51-9715-4d4d-bea3-317cccea16ea"), FirstName: "Some", LastName: "Guy"},
			},
		},
		{
			mockResult: mockUserRepository{users: []models.User{}, err: nil},
			wantedCode: 200,
			wantedBody: []models.User{},
		},
		{
			mockResult:  mockUserRepository{users: nil, err: errors.New("Kaboom!")},
			wantedCode:  500,
			wantedError: ApiError{Message: "Error fetching user records"},
		},
	}

	for _, testCase := range tests {
		router := SetupRouter()
		RegisterRoutes(router, &testCase.mockResult, nil)

		w := httptest.NewRecorder()
		req, _ := http.NewRequest("GET", "/users/", nil)
		router.ServeHTTP(w, req)

		assert.Equal(t, testCase.wantedCode, w.Code)

		if testCase.wantedBody != nil {
			//Unmarshal json resp into slice of Users
			parsedResp := []models.User{}
			json.Unmarshal(w.Body.Bytes(), &parsedResp)

			//compare expected slice w/ unmarshaled response
			assert.Equal(t, parsedResp, testCase.wantedBody)
		} else {
			//Unmarshal json resp into ApiError response
			parsedResp := ApiError{}
			json.Unmarshal(w.Body.Bytes(), &parsedResp)
			assert.Equal(t, parsedResp, testCase.wantedError)
		}
	}
}

func TestFetchUserRoute(t *testing.T) {
	type test struct {
		userId      string
		mockResult  mockUserRepository
		wantedCode  int
		wantedBody  models.User
		wantedError ApiError
	}

	tests := []test{
		{
			userId:     "493adb28-9da1-4db8-893d-73cc2d7bd4ee",
			mockResult: mockUserRepository{users: []models.User{{Id: uuid.MustParse("493adb28-9da1-4db8-893d-73cc2d7bd4ee"), FirstName: "Test", LastName: "User"}}},
			wantedCode: 200,
			wantedBody: models.User{Id: uuid.MustParse("493adb28-9da1-4db8-893d-73cc2d7bd4ee"), FirstName: "Test", LastName: "User"},
		},
		{
			userId:      "493adb28-9da1-4db8-893d-73cc2d7bd4ee",
			mockResult:  mockUserRepository{users: []models.User{}, err: models.ErrModelNotFound},
			wantedCode:  404,
			wantedError: ApiError{Message: "No user exists with Id [493adb28-9da1-4db8-893d-73cc2d7bd4ee]"},
		},
		{
			userId:      "bob",
			mockResult:  mockUserRepository{users: []models.User{}, err: nil},
			wantedCode:  400,
			wantedError: ApiError{Message: "Id [bob] is not a valid UUID"},
		},
		{
			userId:      "493adb28-9da1-4db8-893d-73cc2d7bd4ee",
			mockResult:  mockUserRepository{users: []models.User{}, err: errors.New("Kaboom!")},
			wantedCode:  500,
			wantedError: ApiError{Message: "Error fetching user record with Id [493adb28-9da1-4db8-893d-73cc2d7bd4ee]"},
		},
	}

	for _, testCase := range tests {
		router := SetupRouter()
		RegisterRoutes(router, &testCase.mockResult, nil)

		w := httptest.NewRecorder()
		req, _ := http.NewRequest("GET", "/users/"+testCase.userId, nil)
		router.ServeHTTP(w, req)

		assert.Equal(t, testCase.wantedCode, w.Code)

		if testCase.wantedBody != (models.User{}) {
			//Unmarshal json resp into User
			parsedResp := models.User{}
			json.Unmarshal(w.Body.Bytes(), &parsedResp)

			//compare expected User w/ unmarshaled response
			assert.Equal(t, parsedResp, testCase.wantedBody)
		} else {
			//Unmarshal json resp into ApiError response
			parsedResp := ApiError{}
			json.Unmarshal(w.Body.Bytes(), &parsedResp)
			assert.Equal(t, parsedResp, testCase.wantedError)
		}
	}
}

func TestAddUserRoute(t *testing.T) {
	type test struct {
		requestBody string
		mockResult  mockUserRepository
		wantedCode  int
		wantedBody  models.User
		wantedError ApiError
	}

	tests := []test{
		{
			mockResult:  mockUserRepository{users: []models.User{{Id: uuid.MustParse("493adb28-9da1-4db8-893d-73cc2d7bd4ee")}}},
			requestBody: `{"firstName":"New", "lastName":"User"}`,
			wantedCode:  201,
			wantedBody:  models.User{Id: uuid.MustParse("493adb28-9da1-4db8-893d-73cc2d7bd4ee"), FirstName: "New", LastName: "User"},
		},
		{
			mockResult:  mockUserRepository{users: []models.User{{Id: uuid.MustParse("493adb28-9da1-4db8-893d-73cc2d7bd4ee")}}},
			requestBody: `{"lastName":42}`,
			wantedCode:  400,
			wantedError: ApiError{Message: "Invalid request body."},
		},
		{
			requestBody: `{"firstName":"New", "lastName":"Guy"}`,
			mockResult:  mockUserRepository{users: []models.User{}, err: errors.New("Kaboom!")},
			wantedCode:  500,
			wantedError: ApiError{Message: "Error creating new user"},
		},
	}

	for _, testCase := range tests {
		router := SetupRouter()
		RegisterRoutes(router, &testCase.mockResult, nil)

		w := httptest.NewRecorder()
		req, _ := http.NewRequest("POST", "/users/", bytes.NewBuffer([]byte(testCase.requestBody)))
		req.Header.Set("Content-Type", "application/json")
		router.ServeHTTP(w, req)

		assert.Equal(t, w.Code, testCase.wantedCode)

		if testCase.wantedBody != (models.User{}) {
			//Unmarshal json resp into User
			parsedResp := models.User{}
			json.Unmarshal(w.Body.Bytes(), &parsedResp)

			//compare expected User w/ unmarshaled response
			assert.Equal(t, parsedResp, testCase.wantedBody)
		} else {
			//Unmarshal json resp into ApiError response
			parsedResp := ApiError{}
			json.Unmarshal(w.Body.Bytes(), &parsedResp)
			assert.Equal(t, parsedResp, testCase.wantedError)
		}
	}
}

func TestUpdateUserRoute(t *testing.T) {
	type test struct {
		userId      string
		requestBody string
		mockResult  mockUserRepository
		wantedCode  int
		wantedBody  models.User
		wantedError ApiError
	}

	tests := []test{
		{
			userId:      "493adb28-9da1-4db8-893d-73cc2d7bd4ee",
			mockResult:  mockUserRepository{users: []models.User{{Id: uuid.MustParse("493adb28-9da1-4db8-893d-73cc2d7bd4ee")}}},
			requestBody: `{"firstName":"Updated", "lastName":"Name"}`,
			wantedCode:  200,
			wantedBody:  models.User{Id: uuid.MustParse("493adb28-9da1-4db8-893d-73cc2d7bd4ee"), FirstName: "Updated", LastName: "Name"},
		},
		{
			userId:      "493adb28-9da1-4db8-893d-73cc2d7bd4ee",
			mockResult:  mockUserRepository{users: []models.User{}},
			requestBody: `{"lastName":42}`,
			wantedCode:  400,
			wantedError: ApiError{Message: "Invalid request body."},
		},
		{
			userId:      "493adb28-9da1-4db8-893d-73cc2d7bd4ee",
			requestBody: `{"firstName":"Updated", "lastName":"Name"}`,
			mockResult:  mockUserRepository{users: []models.User{}, err: models.ErrModelNotFound},
			wantedCode:  404,
			wantedError: ApiError{Message: "No user exists with Id [493adb28-9da1-4db8-893d-73cc2d7bd4ee]"},
		},
		{
			userId:      "bob",
			requestBody: `{"firstName":"Updated", "lastName":"Name"}`,
			mockResult:  mockUserRepository{users: []models.User{}, err: nil},
			wantedCode:  400,
			wantedError: ApiError{Message: "Id [bob] is not a valid UUID"},
		},
		{
			userId:      "493adb28-9da1-4db8-893d-73cc2d7bd4ee",
			requestBody: `{"firstName":"Updated", "lastName":"Name"}`,
			mockResult:  mockUserRepository{users: []models.User{}, err: errors.New("Kaboom!")},
			wantedCode:  500,
			wantedError: ApiError{Message: "Error updating user record with Id [493adb28-9da1-4db8-893d-73cc2d7bd4ee]"},
		},
	}

	for _, testCase := range tests {
		router := SetupRouter()
		RegisterRoutes(router, &testCase.mockResult, nil)

		w := httptest.NewRecorder()
		req, _ := http.NewRequest("PUT", "/users/"+testCase.userId, bytes.NewBuffer([]byte(testCase.requestBody)))
		req.Header.Set("Content-Type", "application/json")
		router.ServeHTTP(w, req)

		assert.Equal(t, w.Code, testCase.wantedCode)

		if testCase.wantedBody != (models.User{}) {
			//Unmarshal json resp into User
			parsedResp := models.User{}
			json.Unmarshal(w.Body.Bytes(), &parsedResp)

			//compare expected User w/ unmarshaled response
			assert.Equal(t, parsedResp, testCase.wantedBody)
		} else {
			//Unmarshal json resp into ApiError response
			parsedResp := ApiError{}
			json.Unmarshal(w.Body.Bytes(), &parsedResp)
			assert.Equal(t, parsedResp, testCase.wantedError)
		}
	}
}

func TestDeleteUserRoute(t *testing.T) {
	type test struct {
		userId      string
		mockResult  mockUserRepository
		wantedCode  int
		wantedError ApiError
	}

	tests := []test{
		{
			userId:     "493adb28-9da1-4db8-893d-73cc2d7bd4ee",
			mockResult: mockUserRepository{},
			wantedCode: 204,
		},
		{
			userId:      "493adb28-9da1-4db8-893d-73cc2d7bd4ee",
			mockResult:  mockUserRepository{err: models.ErrModelNotFound},
			wantedCode:  404,
			wantedError: ApiError{Message: "No user exists with Id [493adb28-9da1-4db8-893d-73cc2d7bd4ee]"},
		},
		{
			userId:      "bob",
			mockResult:  mockUserRepository{users: []models.User{}, err: nil},
			wantedCode:  400,
			wantedError: ApiError{Message: "Id [bob] is not a valid UUID"},
		},
		{
			userId:      "493adb28-9da1-4db8-893d-73cc2d7bd4ee",
			mockResult:  mockUserRepository{users: []models.User{}, err: errors.New("Kaboom!")},
			wantedCode:  500,
			wantedError: ApiError{Message: "Error deleting user record with Id [493adb28-9da1-4db8-893d-73cc2d7bd4ee]"},
		},
	}

	for _, testCase := range tests {
		router := SetupRouter()
		RegisterRoutes(router, &testCase.mockResult, nil)

		w := httptest.NewRecorder()
		req, _ := http.NewRequest("DELETE", "/users/"+testCase.userId, nil)
		req.Header.Set("Content-Type", "application/json")
		router.ServeHTTP(w, req)

		assert.Equal(t, w.Code, testCase.wantedCode)

		if testCase.wantedCode != 204 {
			//Unmarshal json resp into ApiError response
			parsedResp := ApiError{}
			json.Unmarshal(w.Body.Bytes(), &parsedResp)
			assert.Equal(t, parsedResp, testCase.wantedError)
		}
	}
}
