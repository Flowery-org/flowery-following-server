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

			var rel dto.CreateRelation
			err = json.Unmarshal(val, &rel)
			rel.CreatedAt = strconv.FormatInt(time.Now().UnixMilli(), 10)

			if err != nil {
				c.JSON(400, err)
			}

			client := lib.GetNeo4jClientInstance()

			err = client.Instance.FollowUser(c, rel)
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

			var rel dto.DeleteRelation
			err = json.Unmarshal(val, &rel)

			if err != nil {
				c.JSON(400, err)
			}

			client := lib.GetNeo4jClientInstance()

			err = client.Instance.UnfollowUser(c, rel)

			if err != nil {
				c.JSON(400, err)
			}

			c.JSON(200, gin.H{
				"ok": true,
			})

		})
	}
}
