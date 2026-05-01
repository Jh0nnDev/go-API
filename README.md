# API REST com Go Lang

API REST criada com Go para operações de CRUD, usando MySQL como banco de dados.

## 🛠️ Tecnologias
- Gin Web Framework
- MySQL
- GORM Library

## 🚀 Como rodar

1.Ter o GO instalado
2.Fazer o Download do GORM e GIN
3.Ter o MySQL instalado 
4.Instalar o Postman para testes
5.Configurar o banco de dados na variavel "dsn"

## 📋 Endpoints

| Método | Rota | Descrição |
|--------|------|-----------|
| GET | /tarefas | Recupera lista de tarefas |
| POST | /tarefas | Cria uma nova tarefa |
| PUT | /tarefas | Atualiza uma tarefa pelo id |
| DELETE | /tarefas | Exclui uma tarefa |
## ⚠️ Configuração
As credenciais do banco de dados estão na variável `dsn` dentro do `main.go`. 
Substitua com suas próprias credenciais antes de rodar.