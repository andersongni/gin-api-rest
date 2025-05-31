package routes

import (
	"github.com/andersongni/gin-api-rest/controllers"
	"github.com/gin-gonic/gin"

	docs "github.com/andersongni/gin-api-rest/docs"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func HandleRequest() {
	r := gin.Default()
	basePath := "/api/v1"
	docs.SwaggerInfo.BasePath = basePath
	r.LoadHTMLGlob("templates/*")
	r.Static("/assets", "./assets")

	api := r.Group(basePath)
	{
		api.GET("/alunos", controllers.ExibeTodosAlunos)
		api.GET("/:nome", controllers.Saudacao)
		api.POST("/alunos", controllers.CriaNovoAluno)
		api.GET("/alunos/:id", controllers.BuscaAlunoPorID)
		api.DELETE("/alunos/:id", controllers.DeletaAlunoPorId)
		api.PATCH("/alunos/:id", controllers.AtualizaAlunoPorId)
		api.GET("/alunos/cpf/:cpf", controllers.BuscaAlunoPorCPF)
		api.GET("/index", controllers.ExibePaginaIndex)
	}

	r.NoRoute(controllers.RotaNaoEncontrada)
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	r.Run()
}
