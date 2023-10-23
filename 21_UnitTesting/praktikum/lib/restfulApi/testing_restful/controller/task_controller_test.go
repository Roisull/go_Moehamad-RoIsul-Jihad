package controller

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/labstack/echo"
	"github.com/stretchr/testify/assert"
	"github.com/user/go_MoehamadRoisulJihad/21_UnitTesting/praktikum/lib/restfulApi/testing_restful/config"
	"github.com/user/go_MoehamadRoisulJihad/21_UnitTesting/praktikum/lib/restfulApi/testing_restful/model"
)

func InitEchoTestAPI() *echo.Echo {
	config.InitDBTest()
	e := echo.New()
	return e
}

func InsertDataUserForGetUsers() error {
	user := model.User{
		Name:     "Rois",
		Password: "123",
		Email:    "Rois@gmail.com",
	}

	var err error
	if err = config.DB.Save(&user).Error; err != nil {
		return err
	}
	return nil
}

func TestGetUserController(t *testing.T) {
	var testCases = []struct {
		name       string
		path       string
		expectCode int
		sizeData   int
	}{
		{
			name:       "get user normal",
			path:       "/users",
			expectCode: http.StatusOK,
			sizeData:   1,
		},
	}

	e := InitEchoTestAPI()
	InsertDataUserForGetUsers()
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	for _, testCase := range testCases {
		c.SetPath(testCase.path)

		if assert.NoError(t, GetUsersController(c)) {
			assert.Equal(t, testCase.expectCode, rec.Code)
			body := rec.Body.String()

			// open file
			// convert struct
			var user UserResponse
			err := json.Unmarshal([]byte(body), &user)

			if err != nil {
				assert.Error(t, err, "error")
			}
			assert.Equal(t, testCase.sizeData, len(user.Data))
		}

	}

}