package main

import (
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Tarefa struct { // estrutura do JSON
	ID           uint   `json:"id"`
	TituloTarefa string `json:"titulo_tarefa"`
}

var db *gorm.DB //

func main() {
	dsn := "root:@tcp(localhost:3306)/test?charset=utf8mb4&parseTime=True" // configura as autenticações do banco de dados
	var err error
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{}) // faz a conexão com o banco dados
	if err != nil {
		panic("Erro ao conectar no banco!") // captura o erro e mostra uma mensagem erro na tela
	}

	r := gin.Default() //cria conexão com o servidor

	r.GET("/tarefas", buscarTarefas) //captura a requisição do JSON
	r.POST("/tarefas", criarTarefas)
	r.PUT("/tarefas", attTarefas)
	r.Run(":8081")
}

func buscarTarefas(c *gin.Context) {
	var tarefas []Tarefa
	db.Find(&tarefas)
	c.JSON(200, tarefas)
}

func criarTarefas(c *gin.Context) {
	var tarefa Tarefa
	c.ShouldBindJSON(&tarefa)
	db.Create(&Tarefa{
		ID:           tarefa.ID,
		TituloTarefa: tarefa.TituloTarefa,
	})
}

func attTarefas(c *gin.Context) {
	var tarefa Tarefa
	c.ShouldBindJSON(&tarefa)
	db.Save(&Tarefa{
		ID:           tarefa.ID,
		TituloTarefa: tarefa.TituloTarefa,
	})
}
