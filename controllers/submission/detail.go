package submission

import (
	SubmissionService "Rabbit-OJ-Backend/services/submission"
	"github.com/gin-gonic/gin"
)

func Detail(c *gin.Context) {
	sid := c.Param("sid")

	submission, err := SubmissionService.Detail(sid)
	if err != nil {
		c.JSON(400, gin.H{
			"code":    400,
			"message": err.Error(),
		})
	} else {
		c.JSON(200, gin.H{
			"code":    200,
			"message": submission,
		})
	}
}
