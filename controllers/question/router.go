package question

import (
	"Rabbit-OJ-Backend/controllers/middlewares"
	"Rabbit-OJ-Backend/controllers/validators"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
)

func Router(baseRouter *gin.Engine) {
	questionRouter := baseRouter.Group("/question")

	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		if err := v.RegisterValidation("language", validators.Language); err != nil {
			fmt.Println(err)
		}
		if err := v.RegisterValidation("code", validators.Code); err != nil {
			fmt.Println(err)
		}
	}

	questionRouter.GET("/list/:page", middlewares.TryAuthJWT(), List)
	questionRouter.GET("/record/:tid/:page", middlewares.AuthJWT(true), Record)
	questionRouter.GET("/item/:tid", middlewares.TryAuthJWT(), Detail)
	questionRouter.POST("/item", middlewares.AuthJWT(true), Create)
	questionRouter.PUT("/item/:tid", middlewares.AuthJWT(true), Edit)
	questionRouter.DELETE("/item/:tid", middlewares.AuthJWT(true), Delete)
	questionRouter.POST("/submit/:tid", middlewares.AuthJWT(true), Submit)
	questionRouter.GET("/judge/:tid", middlewares.AuthJWT(true), GetJudge)
	questionRouter.POST("/judge/:tid", middlewares.AuthJWT(true), SetJudge)
}
