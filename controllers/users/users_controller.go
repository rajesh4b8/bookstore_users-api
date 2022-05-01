package users

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rajesh4b8/bookstore_users-api/domain/users"
	usersService "github.com/rajesh4b8/bookstore_users-api/services/users"
	"github.com/rajesh4b8/bookstore_users-api/utils/errors"
)

func CreateUser(c *gin.Context) {
	var user users.User
	// bytes, err := ioutil.ReadAll(c.Request.Body)
	// if err != nil {
	// 	fmt.Println(err)
	// 	// TODO Handle error
	// 	return
	// }

	// if err := json.Unmarshal(bytes, &user); err != nil {
	// 	fmt.Println(err)
	// 	// TODO handle err
	// 	return
	// }

	if err := c.ShouldBindJSON(&user); err != nil {
		restErr := errors.NewBadRequestError(err.Error())
		c.JSON(restErr.Status, restErr)
		return
	}

	result, saveErr := usersService.CreateUser(user)
	if saveErr != nil {
		c.JSON(saveErr.Status, saveErr)
		return
	}
	c.JSON(http.StatusCreated, result)
}

func GetUser(c *gin.Context) {
	c.String(http.StatusNotImplemented, "Implement me!")
}

func FindUser(c *gin.Context) {
	c.String(http.StatusNotImplemented, "Implement me!")
}
