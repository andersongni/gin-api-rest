package controllers

import (
	"net/http"

	"github.com/andersongni/gin-api-rest/database"
	"github.com/andersongni/gin-api-rest/models"
	"github.com/gin-gonic/gin"
	_ "github.com/swaggo/swag/example/celler/httputil"
)

// @BasePath /api/v1

// ExibeTodosAlunos godoc
// @Summary      Exibe Todos Os Alunos
// @Description  get string
// @Tags         alunos
// @Accept       json
// @Produce      json
// @Success      200  {object}  models.Aluno
// @Failure      400  {object}  httputil.HTTPError
// @Failure      404  {object}  httputil.HTTPError
// @Failure      500  {object}  httputil.HTTPError
// @Router       /alunos [get]
func ExibeTodosAlunos(c *gin.Context) {
	var alunos []models.Aluno
	database.DB.Find(&alunos)
	c.JSON(http.StatusOK, alunos)
}

// Saudacao      godoc
// @Summary      Realiza uma saudacao
// @Description  get string
// @Tags         nome
// @Accept       json
// @Produce      json
// @Param        nome path      string              true "Nome da pessoa"
// @Success      200  {object}  models.Aluno
// @Failure      400  {object}  httputil.HTTPError
// @Failure      404  {object}  httputil.HTTPError
// @Failure      500  {object}  httputil.HTTPError
// @Router       /{nome} [get]
func Saudacao(c *gin.Context) {
	nome := c.Params.ByName("nome")
	c.JSON(http.StatusOK, gin.H{
		"API diz:": "E ai " + nome + ", tudo beleza?",
	})
}

// CriaNovoAluno godoc
// @Summary      Cria Novo Aluno
// @Description  Rota para criar um novo aluno
// @Tags         alunos
// @Accept       json
// @Produce      json
// @Param        aluno body      models.Aluno         true "Objeto Aluno"
// @Success      200   {object}  models.Aluno
// @Failure      400   {object}  httputil.HTTPError
// @Failure      404   {object}  httputil.HTTPError
// @Failure      500   {object}  httputil.HTTPError
// @Router       /alunos [post]
func CriaNovoAluno(c *gin.Context) {
	var aluno models.Aluno
	err := c.ShouldBindJSON(&aluno)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	if err := models.ValidaAluno(&aluno); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	// print(fmt.Sprintf("aluno lido: %+v", aluno))
	database.DB.Create(&aluno)
	c.JSON(http.StatusOK, aluno)
}

// BuscaAlunoPorID godoc
// @Summary      Busca Aluno Por ID
// @Description  Localiza um unico aluno
// @Tags         alunos
// @Accept       json
// @Produce      json
// @Param        id    path      string         true "ID do Aluno"
// @Success      200   {object}  models.Aluno
// @Failure      400   {object}  httputil.HTTPError
// @Failure      404   {object}  httputil.HTTPError
// @Failure      500   {object}  httputil.HTTPError
// @Router       /alunos/{id} [get]
func BuscaAlunoPorID(c *gin.Context) {
	var aluno models.Aluno
	id := c.Params.ByName("id")
	database.DB.First(&aluno, id)

	if aluno.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"Not Found": "Aluno de ID " + id + " não encontrado",
		})
		return
	}

	c.JSON(http.StatusOK, aluno)
}

// DeletaAlunoPorId godoc
// @Summary      Deleta Aluno Por ID
// @Description  Deleta um unico aluno
// @Tags         alunos
// @Accept       json
// @Produce      json
// @Param        id    path      string         true "ID do Aluno"
// @Success      200   {object}  models.Aluno
// @Failure      400   {object}  httputil.HTTPError
// @Failure      404   {object}  httputil.HTTPError
// @Failure      500   {object}  httputil.HTTPError
// @Router       /alunos/{id} [delete]
func DeletaAlunoPorId(c *gin.Context) {
	var aluno models.Aluno
	id := c.Params.ByName("id")
	database.DB.Delete(&aluno, id)

	c.JSON(http.StatusOK, gin.H{
		"Delete": "Aluno de ID " + id + " removido com sucesso!",
	})
}

// AtualizaAlunoPorId godoc
// @Summary      Atualiza Aluno Por ID
// @Description  Atualiza um unico aluno
// @Tags         alunos
// @Accept       json
// @Produce      json
// @Param        id    path      string         true "ID do Aluno"
// @Param        aluno body      models.Aluno         true "Objeto Aluno"
// @Success      200   {object}  models.Aluno
// @Failure      400   {object}  httputil.HTTPError
// @Failure      404   {object}  httputil.HTTPError
// @Failure      500   {object}  httputil.HTTPError
// @Router       /alunos/{id} [patch]
func AtualizaAlunoPorId(c *gin.Context) {
	var aluno models.Aluno
	id := c.Params.ByName("id")
	database.DB.First(&aluno, id)

	err := c.ShouldBindJSON(&aluno)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	if err := models.ValidaAluno(&aluno); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	database.DB.Model(&aluno).UpdateColumns(aluno)

	c.JSON(http.StatusOK, aluno)
}

// BuscaAlunoPorCPF godoc
// @Summary      Busca Aluno Por CPF
// @Description  Localiza um unico aluno
// @Tags         alunos
// @Accept       json
// @Produce      json
// @Param        cpf   path      string         true "CPF do Aluno"
// @Success      200   {object}  models.Aluno
// @Failure      400   {object}  httputil.HTTPError
// @Failure      404   {object}  httputil.HTTPError
// @Failure      500   {object}  httputil.HTTPError
// @Router       /alunos/cpf/{cpf} [get]
func BuscaAlunoPorCPF(c *gin.Context) {
	var aluno models.Aluno
	cpf := c.Param("cpf")
	database.DB.Where("cpf = ?", cpf).First(&aluno)

	if aluno.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"Not Found": "Aluno de CPF " + cpf + " não encontrado",
		})
		return
	}

	c.JSON(http.StatusOK, aluno)
}

func ExibePaginaIndex(c *gin.Context) {
	var alunos []models.Aluno
	database.DB.Find(&alunos)
	c.HTML(http.StatusOK, "index.html", gin.H{
		"alunos": alunos,
	})
}

func RotaNaoEncontrada(c *gin.Context) {
	c.HTML(http.StatusNotFound, "404.html", nil)
}
