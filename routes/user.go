package routes

import (
	"encoding/json"
	"flowery-following-server/dto"
	"flowery-following-server/lib"
	"github.com/gin-gonic/gin"
	"io"
)

func User(router *gin.RouterGroup) {

	user := router.Group("/user")
	{
		// PUT /
		user.PUT("", func(c *gin.Context) {
			body := c.Request.Body
			val, err := io.ReadAll(body)

			if err != nil { //* TODO: Better Error Handling
				c.JSON(500, err)
			}

			var user dto.User

			err = json.Unmarshal(val, &user)
			if err != nil {
				c.JSON(500, err)
			}

			client := lib.GetNeo4jClientInstance()

			// Create User
			err = client.Instance.CreateUser(c, user)
			if err != nil {
				c.JSON(500, err)
			}

			c.JSON(200, gin.H{
				"ok": true,
			})
		})

		user.DELETE("", func(c *gin.Context) {
			body := c.Request.Body
			val, err := io.ReadAll(body)

			if err != nil { //* TODO: Better Error Handling
				c.JSON(500, err)
			}

			var user dto.User

			err = json.Unmarshal(val, &user)
			if err != nil {
				c.JSON(500, err)
			}

			client := lib.GetNeo4jClientInstance()

			// Create User
			err = client.Instance.DeleteUser(c, user.Id)
			if err != nil {
				c.JSON(500, err)
			}

			c.JSON(200, gin.H{
				"ok": true,
			})
		})
	}
}
