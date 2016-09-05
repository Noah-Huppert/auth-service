package resp

import "github.com/gin-gonic/gin"

const (
	AcceptJson = "application/json"
	AcceptXML = "application/xml"
)

func Resp (data interface{}, status int, c *gin.Context) {
	accept := c.Request.Header.Get("Accept")

	if accept == "application/xml" {
		c.Header("Content-Type", "application/xml")

		c.XML(status, data)
	} else {
		c.Header("Content-Type", "application/json")

		if c.Query("json_expanded") == "true" {
			c.IndentedJSON(status, data)
		} else {
			c.JSON(status, data)
		}
	}
}
