package routes

import (
	"encoding/json"
	"flowery-following-server/dto"
	"flowery-following-server/lib"
	"github.com/gin-gonic/gin"
	"io"
	"strconv"
	"time"
)

func Relation(router *gin.RouterGroup) {
	// /rel
	rel := router.Group("/rel")
	{
		// PUT /
		rel.PUT("", func(c *gin.Context) {
			body := c.Request.Body
			val, err := io.ReadAll(body)

			if err != nil {
				c.JSON(500, err)
			}

			var createRel dto.CreateRelation
			err = json.Unmarshal(val, &createRel)
			createRel.CreatedAt = strconv.FormatInt(time.Now().UnixMilli(), 10)

			if err != nil {
				c.JSON(500, err)
			}

			client := lib.GetNeo4jClientInstance()

			err = client.Instance.FollowUser(c, createRel)
			if err != nil {
				c.JSON(400, err)
			}

			c.JSON(200, gin.H{
				"ok": true,
			})
		})

		rel.DELETE("", func(c *gin.Context) {
			body := c.Request.Body
			val, err := io.ReadAll(body)
			if err != nil {
				c.JSON(500, err)
			}

			var delRel dto.DeleteRelation
			err = json.Unmarshal(val, &delRel)

			if err != nil {
				c.JSON(400, err)
			}

			client := lib.GetNeo4jClientInstance()

			err = client.Instance.UnfollowUser(c, delRel)

			if err != nil {
				c.JSON(400, err)
			}

			c.JSON(200, gin.H{
				"ok": true,
			})

		})

		rel.GET("/followers", func(c *gin.Context) {
			userId := c.Query("userId")
			client := lib.GetNeo4jClientInstance()

			followers, err := client.Instance.GetAllFollowers(c, userId)
			if err != nil {
				c.JSON(500, err)
			}
			c.JSON(200, followers)
		})

		rel.GET("/followings", func(c *gin.Context) {
			userId := c.Query("userId")
			client := lib.GetNeo4jClientInstance()

			followers, err := client.Instance.GetAllFollowings(c, userId)
			if err != nil {
				c.JSON(500, err)
			}
			c.JSON(200, followers)
		})
	}
}
