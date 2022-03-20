package users

import (
	"github.com/gin-gonic/gin"
	"github.com/mustafakocatepe/go-bookstore-users-api/domain/users"
	"github.com/mustafakocatepe/go-bookstore-users-api/services"
	"github.com/mustafakocatepe/go-bookstore-users-api/utils/errors"
	"net/http"
	"strconv"
)

func getUserId(userIdParam string) (int64, *errors.RestErr) {
	userId, userErr := strconv.ParseInt(userIdParam, 10, 64)
	if userErr != nil {
		return 0, errors.NewBadRequestError("user id should be a number")
	}
	return userId, nil

}

func Create(c *gin.Context) {
	var user users.User

	/*
		bytes, err := ioutil.ReadAll(c.Request.Body)
		if err != nil {
			//TODO: handle error
			return
		}
		if err := json.Unmarshal(bytes, &user); err != nil {
			//TODO: handle JSON error
			fmt.Println(err.Error())
			return
		}
	*/

	if err := c.ShouldBindJSON(&user); err != nil {
		restErr := errors.NewBadRequestError("invalid json body")
		c.JSON(restErr.Status, restErr)
		return
	}

	result, saveErr := services.CreateUser(user)
	if saveErr != nil {
		c.JSON(saveErr.Status, saveErr)
		return
	}

	c.JSON(http.StatusCreated, result)
}

func Get(c *gin.Context) {
	userId, idErr := getUserId(c.Param("user_id"))
	if idErr != nil {
		c.JSON(idErr.Status, idErr)
		return
	}

	user, getErr := services.GetUser(userId)
	if getErr != nil {
		c.JSON(getErr.Status, getErr)
		return
	}
	c.JSON(http.StatusOK, user)
}

func Update(c *gin.Context) {
	userId, idErr := getUserId(c.Param("user_id"))
	if idErr != nil {
		c.JSON(idErr.Status, idErr)
		return
	}

	var user users.User
	if err := c.ShouldBindJSON(&user); err != nil {
		restErr := errors.NewBadRequestError("invalid json body")
		c.JSON(restErr.Status, restErr)
		return
	}
	user.Id = userId

	result, err := services.UpdateUSer(user)
	if err != nil {
		c.JSON(err.Status, err)
		return
	}
	c.JSON(http.StatusOK, result)

}

func Delete(c *gin.Context) {
	userId, idErr := getUserId(c.Param("user_id"))
	if idErr != nil {
		c.JSON(idErr.Status, idErr)
		return
	}
	if err := services.DeleteUser(userId); err != nil {
		c.JSON(idErr.Status, idErr)
		return
	}
	c.JSON(http.StatusOK, map[string]string{"status": "deleted"})
}

//GET : /users/search?status=active - status is a query param
func Search(c *gin.Context) {
	status := c.Query("status")

	users, err := services.Search(status)
	if err != nil {
		c.JSON(err.Status, err)
		return
	}
	c.JSON(http.StatusOK, users)
}

/*
func Search(c *gin.Context) {
	c.String(http.StatusNotImplemented, "implement me!")
}
*/
