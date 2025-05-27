package controllers

import (
	"net/http"

	"github.com/andersongni/gin-api-rest/database"
	"github.com/andersongni/gin-api-rest/models"
	"github.com/gin-gonic/gin"
)

func ExibeTodosAlunos(c *gin.Context) {
	var alunos []models.Aluno
	database.DB.Find(&alunos)
	c.JSON(http.StatusOK, alunos)
}

func Saudacao(c *gin.Context) {
	nome := c.Params.ByName("nome")
	c.JSON(http.StatusOK, gin.H{
		"API diz:": "E ai " + nome + ", tudo beleza?",
	})
}

func CriaNovoAluno(c *gin.Context) {
	var aluno models.Aluno
	err := c.ShouldBindJSON(&aluno)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	// if database.DB == nil {
	// 	fmt.Println("database.DB is nil!")
	// 	c.JSON(http.StatusInternalServerError, gin.H{
	// 		"error": "Database connection not initialized",
	// 	})
	// 	return
	// }

	// print(fmt.Sprintf("aluno lido: %+v", aluno))
	database.DB.Create(&aluno)
	c.JSON(http.StatusOK, aluno)
}

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

func DeletaAlunoPorId(c *gin.Context) {
	var aluno models.Aluno
	id := c.Params.ByName("id")
	database.DB.Delete(&aluno, id)

	c.JSON(http.StatusOK, gin.H{
		"Delete": "Aluno de ID " + id + " removido com sucesso!",
	})
}

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
	database.DB.Model(&aluno).UpdateColumns(aluno)

	c.JSON(http.StatusOK, aluno)
}

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
