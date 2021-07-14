// Package controllers é responsável por abrir uma conexão com o banco de dados, e por devolver respostas para cada requisição HTTP recebida.
package controllers

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"
	"weni/api/src/database"
	"weni/api/src/models"
	"weni/api/src/repositories"
	"weni/api/src/responses"

	"github.com/gorilla/mux"
)

// CriarTarefa cria uma nova tarefa na lista de tarefas.
func CriarTarefa(w http.ResponseWriter, r *http.Request) {
	corpoRequest, erro := ioutil.ReadAll(r.Body)
	if erro != nil {
		responses.Erro(w, http.StatusUnprocessableEntity, erro)
		return
	}

	var tarefa models.Tarefa
	if erro = json.Unmarshal(corpoRequest, &tarefa); erro != nil {
		responses.Erro(w, http.StatusBadRequest, erro)
		return
	}

	if erro := tarefa.Preparar(); erro != nil {
		responses.Erro(w, http.StatusBadRequest, erro)
		return
	}

	db, erro := database.Conectar()
	if erro != nil {
		responses.Erro(w, http.StatusInternalServerError, erro)
		return
	}
	defer db.Close()

	repositorio := repositories.NovoRepositorioDeTarefas(db)
	tarefa.ID, erro = repositorio.Criar(tarefa)
	if erro != nil {
		responses.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	responses.JSON(w, http.StatusCreated, tarefa)
}

// MostrarTarefas mostra todas as tarefas da lista de tarefas.
func MostrarTarefas(w http.ResponseWriter, r *http.Request) {
	db, erro := database.Conectar()
	if erro != nil {
		responses.Erro(w, http.StatusInternalServerError, erro)
		return
	}
	defer db.Close()

	repositorio := repositories.NovoRepositorioDeTarefas(db)
	tarefas, erro := repositorio.BuscarTodos()
	if erro != nil {
		responses.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	responses.JSON(w, http.StatusOK, tarefas)
}

// MostrarTarefasAbertas mostra todas as tarefas abertas da lista de tarefas.
func MostrarTarefasAbertas(w http.ResponseWriter, r *http.Request) {
	db, erro := database.Conectar()
	if erro != nil {
		responses.Erro(w, http.StatusInternalServerError, erro)
		return
	}
	defer db.Close()

	repositorio := repositories.NovoRepositorioDeTarefas(db)
	tarefasAbertos, erro := repositorio.BuscarAbertos()
	if erro != nil {
		responses.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	responses.JSON(w, http.StatusOK, tarefasAbertos)
}

// MostrarTarefasFechadas mostra todas as tarefas fechadas da lista de tarefas.
func MostrarTarefasFechadas(w http.ResponseWriter, r *http.Request) {
	db, erro := database.Conectar()
	if erro != nil {
		responses.Erro(w, http.StatusInternalServerError, erro)
		return
	}
	defer db.Close()

	repositorio := repositories.NovoRepositorioDeTarefas(db)
	tarefasFechados, erro := repositorio.BuscarFechados()
	if erro != nil {
		responses.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	responses.JSON(w, http.StatusOK, tarefasFechados)
}

// AlterarTarefa altera o conteudo de uma tarefa da lista de tarefas.
func AlterarTarefa(w http.ResponseWriter, r *http.Request) {
	parametros := mux.Vars(r)
	tarefaID, erro := strconv.ParseUint(parametros["tarefaId"], 10, 64)
	if erro != nil {
		responses.Erro(w, http.StatusBadRequest, erro)
		return
	}

	corpoRequest, erro := ioutil.ReadAll(r.Body)
	if erro != nil {
		responses.Erro(w, http.StatusUnprocessableEntity, erro)
		return
	}

	var tarefa models.Tarefa
	if erro := json.Unmarshal(corpoRequest, &tarefa); erro != nil {
		responses.Erro(w, http.StatusBadRequest, erro)
		return
	}

	if erro := tarefa.Preparar(); erro != nil {
		responses.Erro(w, http.StatusBadRequest, erro)
		return
	}

	db, erro := database.Conectar()
	if erro != nil {
		responses.Erro(w, http.StatusInternalServerError, erro)
		return
	}
	defer db.Close()

	repositorio := repositories.NovoRepositorioDeTarefas(db)
	tarefa.ID, erro = repositorio.Atualizar(tarefaID, tarefa)
	if erro != nil {
		responses.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	responses.JSON(w, http.StatusOK, tarefa)
}

// AbrirTarefa abre ou completa uma tarefa e manda para a lista de tarefas concluidas.
func AbrirTarefa(w http.ResponseWriter, r *http.Request) {
	parametros := mux.Vars(r)
	tarefaID, erro := strconv.ParseUint(parametros["tarefaId"], 10, 64)
	if erro != nil {
		responses.Erro(w, http.StatusBadRequest, erro)
		return
	}

	db, erro := database.Conectar()
	if erro != nil {
		responses.Erro(w, http.StatusInternalServerError, erro)
		return
	}
	defer db.Close()

	repositorio := repositories.NovoRepositorioDeTarefas(db)
	tarefa, erro := repositorio.Abrir(tarefaID)
	if erro != nil {
		responses.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	responses.JSON(w, http.StatusOK, tarefa)
}

// FecharTarefa fecha ou restaura uma tarefa e manda para a lista principal.
func FecharTarefa(w http.ResponseWriter, r *http.Request) {
	parametros := mux.Vars(r)
	tarefaID, erro := strconv.ParseUint(parametros["tarefaId"], 10, 64)
	if erro != nil {
		responses.Erro(w, http.StatusBadRequest, erro)
		return
	}

	db, erro := database.Conectar()
	if erro != nil {
		responses.Erro(w, http.StatusInternalServerError, erro)
		return
	}
	defer db.Close()

	repositorio := repositories.NovoRepositorioDeTarefas(db)
	tarefa, erro := repositorio.Fechar(tarefaID)
	if erro != nil {
		responses.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	responses.JSON(w, http.StatusOK, tarefa)
}

// DeletarTarefa apaga uma tarefa da lista de tarefas abertas(concluidas).
func DeletarTarefaAberta(w http.ResponseWriter, r *http.Request) {
	parametros := mux.Vars(r)
	tarefaID, erro := strconv.ParseUint(parametros["tarefaId"], 10, 64)
	if erro != nil {
		responses.Erro(w, http.StatusBadRequest, erro)
		return
	}

	db, erro := database.Conectar()
	if erro != nil {
		responses.Erro(w, http.StatusInternalServerError, erro)
		return
	}
	defer db.Close()

	repositorio := repositories.NovoRepositorioDeTarefas(db)
	if erro = repositorio.DeletarAbertos(tarefaID); erro != nil {
		responses.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	responses.JSON(w, http.StatusNoContent, nil)

}

// DeletarTarefa apaga uma tarefa da lista de tarefas fechadas(não concluidas).
func DeletarTarefaFechada(w http.ResponseWriter, r *http.Request) {
	parametros := mux.Vars(r)
	tarefaID, erro := strconv.ParseUint(parametros["tarefaId"], 10, 64)
	if erro != nil {
		responses.Erro(w, http.StatusBadRequest, erro)
		return
	}

	db, erro := database.Conectar()
	if erro != nil {
		responses.Erro(w, http.StatusInternalServerError, erro)
		return
	}
	defer db.Close()

	repositorio := repositories.NovoRepositorioDeTarefas(db)
	if erro = repositorio.DeletarAbertos(tarefaID); erro != nil {
		responses.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	responses.JSON(w, http.StatusNoContent, nil)

}
