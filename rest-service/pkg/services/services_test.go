package services

import (
	"github.com/labstack/echo/v4"
	uuid "github.com/satori/go.uuid"
	"github.com/stackpath/backend-developer-tests/rest-service/pkg/models"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"
)

func TestGetPeople(t *testing.T) {
	var sb strings.Builder
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/people", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	// Get /people
	if assert.NoError(t, GetPeople(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)

		for _, person := range models.AllPeople() {
			personStr, err := person.ToJSON()

			if err != nil {
				c.Logger().Error(err)
				continue
			}
			sb.WriteString(personStr)
		}

		assert.Equal(t, sb.String(), rec.Body.String())
	}

	//Get Jane Doe
	q := make(url.Values)
	q.Set("first_name", "Jane")
	q.Set("last_name", "Doe")

	req = httptest.NewRequest(http.MethodGet, "/people?"+q.Encode(), nil)
	rec = httptest.NewRecorder()
	c = e.NewContext(req, rec)
	sb.Reset()
	if assert.NoError(t, GetPeople(c)) {
		for _, person := range models.FindPeopleByName("Jane", "Doe") {
			personStr, err := person.ToJSON()

			if err != nil {
				c.Logger().Error(err)
				continue
			}
			sb.WriteString(personStr)
		}
		assert.Equal(t, sb.String(), rec.Body.String())
	}

	// Get Jim Doe
	// Doesn't exist
	q = make(url.Values)
	q.Set("first_name", "Jim")
	q.Set("last_name", "Doe")

	req = httptest.NewRequest(http.MethodGet, "/people?"+q.Encode(), nil)
	rec = httptest.NewRecorder()
	c = e.NewContext(req, rec)
	sb.Reset()
	if assert.NoError(t, GetPeople(c)) {
		for _, person := range models.FindPeopleByName("Jim", "Doe") {
			personStr, err := person.ToJSON()

			if err != nil {
				c.Logger().Error(err)
				continue
			}
			sb.WriteString(personStr)
		}
		assert.Equal(t, "Not Found", rec.Body.String())
	}

	// Get Brian Smith by phone number
	q = make(url.Values)
	q.Set("phone_number", "+44 7700 900077")

	req = httptest.NewRequest(http.MethodGet, "/people?"+q.Encode(), nil)
	rec = httptest.NewRecorder()
	c = e.NewContext(req, rec)
	sb.Reset()
	if assert.NoError(t, GetPeople(c)) {
		for _, person := range models.FindPeopleByPhoneNumber("+44 7700 900077") {
			personStr, err := person.ToJSON()

			if err != nil {
				c.Logger().Error(err)
				continue
			}
			sb.WriteString(personStr)
		}
		assert.Equal(t, sb.String(), rec.Body.String())
	}
}

func TestGetPersonByID(t *testing.T) {

	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	// Retrieve person by id
	c.SetPath("/people/:id")
	c.SetParamNames("id")
	c.SetParamValues("df12ce76-767b-4bf0-bccb-816745df9e70")

	if assert.NoError(t, GetPersonByID(c)) {
		uuid, _ := uuid.FromString("df12ce76-767b-4bf0-bccb-816745df9e70")

		person, _ := models.FindPersonByID(uuid)

		personStr, _ := person.ToJSON()
		assert.Equal(t, personStr, rec.Body.String())
	}

	// Test invalid UUID
	rec = httptest.NewRecorder()
	c = e.NewContext(req, rec)
	c.SetPath("/people/:id")
	c.SetParamNames("id")
	c.SetParamValues("Not A UUID")

	if assert.NoError(t, GetPersonByID(c)){
		assert.Equal(t, rec.Code, 400)
		assert.Equal(t, "Invalid UUID", rec.Body.String())
	}

	// Test valid ID but not found
	rec = httptest.NewRecorder()
	c = e.NewContext(req, rec)
	c.SetPath("/people/:id")
	c.SetParamNames("id")
	c.SetParamValues("ef12ce76-767b-4bf0-bccb-816745df9e70")

	if assert.NoError(t, GetPersonByID(c)){
		assert.Equal(t, rec.Code, 404)
		assert.Equal(t, "user ID ef12ce76-767b-4bf0-bccb-816745df9e70 not found", rec.Body.String())
	}
}
