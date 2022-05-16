package users

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/rajesh4b8/bookstore_users-api/domain/users"
	"github.com/rajesh4b8/bookstore_users-api/services"
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

	result, saveErr := services.UsersService.CreateUser(user)
	if saveErr != nil {
		c.JSON(saveErr.Status, saveErr)
		return
	}
	c.JSON(http.StatusCreated, result)
}

func GetUser(c *gin.Context) {
	userId, err := strconv.ParseInt(c.Param("user_id"), 10, 64)
	if err != nil {
		restErr := errors.NewBadRequestError("invalid user_id")
		c.JSON(restErr.Status, restErr)
		return
	}

	user, getErr := services.UsersService.GetUser(userId)
	if getErr != nil {
		c.JSON(getErr.Status, getErr)
		return
	}

	c.JSON(http.StatusOK, user)
}

func FindUser(c *gin.Context) {
	c.String(http.StatusNotImplemented, "Implement me!")
}
