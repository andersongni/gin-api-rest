package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"

	"github.com/andersongni/gin-api-rest/controllers"
	"github.com/andersongni/gin-api-rest/database"
	"github.com/andersongni/gin-api-rest/models"
	"github.com/gin-gonic/gin"

	"github.com/stretchr/testify/assert"
)

var ID uint

func ConfiguraRotas() *gin.Engine {
	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()
	return r
}

func RequisitaSaudacao(nome string) *httptest.ResponseRecorder {
	// Define a rota
	r := ConfiguraRotas()
	r.GET("/:nome", controllers.Saudacao)

	// Realizando requisicao
	url := fmt.Sprintf("/%s", nome)
	requisicao, _ := http.NewRequest("GET", url, nil)
	resposta := httptest.NewRecorder()

	r.ServeHTTP(resposta, requisicao)

	return resposta
}

func CriaAlunoMock() {
	database.ConectaComBancoDeDados()

	aluno := models.Aluno{Nome: "AlunoMock", CPF: "12345678901", RG: "123456789"}
	database.DB.Create(&aluno)
	ID = aluno.ID
}

func DeletaAlunoMock() {
	database.ConectaComBancoDeDados()
	var aluno models.Aluno
	database.DB.Delete(&aluno, ID)
}

func TestStatusCodeSaudacao(t *testing.T) {
	resposta := RequisitaSaudacao("anderson")

	// Compara Staus Code com Testify
	erro := fmt.Sprintf("Status Code esperado era %d, mas foi recebido %d", http.StatusOK, resposta.Code)
	assert.Equal(t, http.StatusOK, resposta.Code, erro)
}

func TestPayloadSaudacao(t *testing.T) {
	resposta := RequisitaSaudacao("anderson")

	respostaEsperada := `{"API diz:":"E ai anderson, tudo beleza?"}`
	assert.Equal(t, respostaEsperada, resposta.Body.String(), "Payload diferente")
}

func TestStatusCodeAlunos(t *testing.T) {

	database.ConectaComBancoDeDados()

	CriaAlunoMock()
	defer DeletaAlunoMock()

	// Define a rota
	r := ConfiguraRotas()
	r.GET("/:nome", controllers.ExibeTodosAlunos)

	// Realizando requisicao
	requisicao, _ := http.NewRequest("GET", "/alunos", nil)
	resposta := httptest.NewRecorder()

	r.ServeHTTP(resposta, requisicao)

	// Compara Status Code
	erro := fmt.Sprintf("Status Code esperado era %d, mas foi recebido %d", http.StatusOK, resposta.Code)
	assert.Equal(t, http.StatusOK, resposta.Code, erro)

}

func TestBuscaAlunoPorID(t *testing.T) {
	database.ConectaComBancoDeDados()
	CriaAlunoMock()
	defer DeletaAlunoMock()

	r := ConfiguraRotas()
	r.GET("/alunos/:id", controllers.BuscaAlunoPorID)

	url := fmt.Sprintf("/alunos/%d", ID)
	requisicao, _ := http.NewRequest("GET", url, nil)
	resposta := httptest.NewRecorder()
	r.ServeHTTP(resposta, requisicao)

	var alunoRecebido models.Aluno
	json.Unmarshal(resposta.Body.Bytes(), &alunoRecebido)

	assert.Equal(t, "AlunoMock", alunoRecebido.Nome, "Nome do aluno diferente do esperado")
	assert.Equal(t, "12345678901", alunoRecebido.CPF)
	assert.Equal(t, "123456789", alunoRecebido.RG)

}

func TestDeletaAluno(t *testing.T) {
	database.ConectaComBancoDeDados()

	r := ConfiguraRotas()
	r.DELETE("/alunos/:id", controllers.DeletaAlunoPorId)
	r.GET("/alunos/:id", controllers.BuscaAlunoPorID)

	CriaAlunoMock()
	defer DeletaAlunoMock()
	url := "/alunos/" + strconv.Itoa(int(ID))

	requisicaoDelete, _ := http.NewRequest("DELETE", url, nil)
	respostaDelete := httptest.NewRecorder()
	r.ServeHTTP(respostaDelete, requisicaoDelete)

	assert.Equal(t, http.StatusOK, respostaDelete.Code, "Status Code inesperado no endpoint de Delete")

	requisicaoGet, _ := http.NewRequest("GET", url, nil)
	respostaGet := httptest.NewRecorder()
	r.ServeHTTP(respostaGet, requisicaoGet)

	assert.Equal(t, http.StatusNotFound, respostaGet.Code)

}

func TestEditaAluno(t *testing.T) {
	database.ConectaComBancoDeDados()

	r := ConfiguraRotas()
	r.PATCH("/alunos/:id", controllers.AtualizaAlunoPorId)

	CriaAlunoMock()
	defer DeletaAlunoMock()
	url := fmt.Sprintf("/alunos/%d", ID)
	alunoMockBytes := models.Aluno{Nome: "AlunoMock", CPF: "47123456789", RG: "123456700"}
	alunoMockJson, _ := json.Marshal(alunoMockBytes)

	requisicao, _ := http.NewRequest("PATCH", url, bytes.NewBuffer(alunoMockJson))
	resposta := httptest.NewRecorder()
	r.ServeHTTP(resposta, requisicao)

	var alunoRecebido models.Aluno
	json.Unmarshal(resposta.Body.Bytes(), &alunoRecebido)

	assert.Equal(t, "AlunoMock", alunoRecebido.Nome)
	assert.Equal(t, "47123456789", alunoRecebido.CPF)
	assert.Equal(t, "123456700", alunoRecebido.RG)

}
