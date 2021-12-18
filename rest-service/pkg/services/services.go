package services

import (
	"github.com/labstack/echo/v4"
	uuid "github.com/satori/go.uuid"
	"github.com/stackpath/backend-developer-tests/rest-service/pkg/models"
	"net/http"
	"strings"
)

func GetPeople(c echo.Context) error {
	var sb strings.Builder
	var people []*models.Person

	firstName := c.QueryParam("first_name")
	lastName := c.QueryParam("last_name")
	phoneNumber := c.QueryParam("phone_number")

	if firstName == "" && lastName == "" && phoneNumber == ""{
		people = models.AllPeople()
	} else if firstName != "" && lastName != ""{
		people = models.FindPeopleByName(firstName, lastName)
	} else if phoneNumber != "" {
		people = models.FindPeopleByPhoneNumber(phoneNumber)
	}

	for _, person := range people {
		personStr, err := person.ToJSON()

		if err != nil {
			c.Logger().Error(err)
			continue
		}
		sb.WriteString(personStr)
	}
	returnString := sb.String()

	if returnString == "" {
		returnString = "Not Found"
	}
	return c.String(http.StatusOK, returnString)
}

func GetPersonByID(c echo.Context) error {
	id := c.Param("id")
	uuid, err := uuid.FromString(id)

	if err != nil {
		c.Logger().Error(err)
		return c.String(http.StatusBadRequest, "Invalid UUID")
	}

	person, err := models.FindPersonByID(uuid)

	if err != nil {
		c.Logger().Error(err)
		return c.String(http.StatusNotFound, err.Error())
	}

	personStr, err := person.ToJSON()

	if err != nil {
		c.Logger().Error(err)
		return c.String(http.StatusInternalServerError, err.Error())
	}

	return c.String(http.StatusOK, personStr)
}
