// Package repositories é responsavel por interagir com o banco. Ex: Criar uma nova tarefa, listar todas as tarefas e etc.
package repositories

import (
	"database/sql"
	"weni/api/src/models"
)

// tarefas representa um repositório de tarefas.
type Tarefas struct {
	db *sql.DB
}

// NovoRepositorioDeTarefas cria um novo repositorios de usuários.
func NovoRepositorioDeTarefas(db *sql.DB) *Tarefas {
	return &Tarefas{db}
}

// Criar insere uma tarefa no banco de dados.
func (repositorio Tarefas) Criar(tarefa models.Tarefa) (uint64, error) {
	statement, erro := repositorio.db.Prepare(
		"insert into abertos (title, checked) values (?, false)",
	)
	if erro != nil {
		return 0, erro
	}
	defer statement.Close()

	resultado, erro := statement.Exec(tarefa.Title)
	if erro != nil {
		return 0, erro
	}

	ultimoIDInserido, erro := resultado.LastInsertId()
	if erro != nil {
		return 0, erro
	}

	return uint64(ultimoIDInserido), nil
}

// BuscarTodos retorna todos as tarefas que estão no banco de dados.
func (repositorio Tarefas) BuscarTodos() ([]models.Tarefa, error) {
	linhasAbertos, erro := repositorio.db.Query("select * from abertos")
	if erro != nil {
		return nil, erro
	}
	defer linhasAbertos.Close()

	linhasFechados, erro := repositorio.db.Query("select * from fechados")
	if erro != nil {
		return nil, erro
	}
	defer linhasFechados.Close()

	var tasks []models.Tarefa

	for linhasAbertos.Next() {
		var task models.Tarefa
		if erro := linhasAbertos.Scan(&task.ID, &task.Title, &task.Checked); erro != nil {
			return nil, erro
		}

		tasks = append(tasks, task)
	}

	for linhasFechados.Next() {
		var task models.Tarefa
		if erro := linhasFechados.Scan(&task.ID, &task.Title, &task.Checked); erro != nil {
			return nil, erro
		}

		tasks = append(tasks, task)
	}

	return tasks, nil
}

// BuscarAbertos retorna todas as tarefas abertas(ainda não concluidas).
func (repositorio Tarefas) BuscarAbertos() ([]models.Tarefa, error) {
	linhasAbertos, erro := repositorio.db.Query("select * from abertos")
	if erro != nil {
		return nil, erro
	}
	defer linhasAbertos.Close()

	var tasks []models.Tarefa

	for linhasAbertos.Next() {
		var task models.Tarefa
		if erro := linhasAbertos.Scan(&task.ID, &task.Title, &task.Checked); erro != nil {
			return nil, erro
		}

		tasks = append(tasks, task)
	}

	return tasks, nil
}

// BuscarFechados retorna todas as tarefas fechadas(concluidas).
func (repositorio Tarefas) BuscarFechados() ([]models.Tarefa, error) {
	linhasFechados, erro := repositorio.db.Query("select * from fechados")
	if erro != nil {
		return nil, erro
	}
	defer linhasFechados.Close()

	var tasks []models.Tarefa

	for linhasFechados.Next() {
		var task models.Tarefa
		if erro := linhasFechados.Scan(&task.ID, &task.Title, &task.Checked); erro != nil {
			return nil, erro
		}

		tasks = append(tasks, task)
	}

	return tasks, nil
}

// Atualizar altera as informações de uma tarefa no banco de dados.
func (repositorio Tarefas) Atualizar(ID uint64, task models.Tarefa) (uint64, error) {
	statement, erro := repositorio.db.Prepare("update abertos set title = ? where id = ?")
	if erro != nil {
		return 0, erro
	}
	defer statement.Close()

	if _, erro := statement.Exec(task.Title, ID); erro != nil {
		return 0, erro
	}

	return ID, erro
}

// Abrir abre uma tarefa retirando ela da lista de tarefas fechadas(concluidas).
func (repositorio Tarefas) Abrir(ID uint64) (models.Tarefa, error) {
	statement, erro := repositorio.db.Prepare("update fechados set checked = false where id = ?")
	if erro != nil {
		return models.Tarefa{}, erro
	}
	defer statement.Close()

	if _, erro := statement.Exec(ID); erro != nil {
		return models.Tarefa{}, erro
	}

	linhas, erro := repositorio.db.Query("select id, title, checked from fechados where id = ?", ID)
	if erro != nil {
		return models.Tarefa{}, nil
	}
	defer linhas.Close()

	var tarefa models.Tarefa

	if linhas.Next() {
		if erro := linhas.Scan(&tarefa.ID, &tarefa.Title, &tarefa.Checked); erro != nil {
			return models.Tarefa{}, erro
		}
	}

	repositorio.DeletarFechados(ID)

	novoID, erro := repositorio.inserirAbertos(tarefa)
	if erro != nil {
		return models.Tarefa{}, erro
	}

	tarefa.ID = novoID

	return tarefa, nil
}

// Fechar fecha uma tarefa retirando ela da lista de tarefas abertas(não concluidas).
func (repositorio Tarefas) Fechar(ID uint64) (models.Tarefa, error) {
	statement, erro := repositorio.db.Prepare("update abertos set checked = true where id = ?")
	if erro != nil {
		return models.Tarefa{}, erro
	}
	defer statement.Close()

	if _, erro := statement.Exec(ID); erro != nil {
		return models.Tarefa{}, erro
	}

	linhas, erro := repositorio.db.Query("select id, title, checked from abertos where id = ?", ID)
	if erro != nil {
		return models.Tarefa{}, nil
	}
	defer linhas.Close()

	var tarefa models.Tarefa

	if linhas.Next() {
		if erro := linhas.Scan(&tarefa.ID, &tarefa.Title, &tarefa.Checked); erro != nil {
			return models.Tarefa{}, erro
		}
	}

	repositorio.DeletarAbertos(ID)

	novoID, erro := repositorio.inserirFechados(tarefa)
	if erro != nil {
		return models.Tarefa{}, erro
	}

	tarefa.ID = novoID

	return tarefa, nil
}

// DeletarAbertos exclui uma tarefa banco de tarefas abertas(não concluidas).
func (repositorio Tarefas) DeletarAbertos(ID uint64) error {
	statement, erro := repositorio.db.Prepare("delete from abertos where id = ?")
	if erro != nil {
		return erro
	}
	defer statement.Close()

	if _, erro := statement.Exec(ID); erro != nil {
		return erro
	}

	return nil
}

// Deletar exclui uma tarefa banco de tarefas fechadas(concluidas).
func (repositorio Tarefas) DeletarFechados(ID uint64) error {
	statement, erro := repositorio.db.Prepare("delete from fechados where id = ?")
	if erro != nil {
		return erro
	}
	defer statement.Close()

	if _, erro := statement.Exec(ID); erro != nil {
		return erro
	}

	return nil
}

// inserirAbertos insere uma tarefa na lista de tarefas abertas(não concluidas). É usada na função Abrir.
func (repositorio Tarefas) inserirAbertos(tarefa models.Tarefa) (uint64, error) {
	statement, erro := repositorio.db.Prepare(
		"insert into abertos (title, checked) values (?, ?)",
	)
	if erro != nil {
		return 0, erro
	}
	defer statement.Close()

	resultado, erro := statement.Exec(tarefa.Title, tarefa.Checked)
	if erro != nil {
		return 0, erro
	}

	ultimoIDInserido, erro := resultado.LastInsertId()
	if erro != nil {
		return 0, erro
	}

	return uint64(ultimoIDInserido), nil
}

// inserirAbertos insere uma tarefa na lista de tarefas fechadas(concluidas). É usada na função Fechar.
func (repositorio Tarefas) inserirFechados(tarefa models.Tarefa) (uint64, error) {
	statement, erro := repositorio.db.Prepare(
		"insert into fechados (title, checked) values (?, ?)",
	)
	if erro != nil {
		return 0, erro
	}
	defer statement.Close()
	resultado, erro := statement.Exec(tarefa.Title, tarefa.Checked)
	if erro != nil {
		return 0, erro
	}

	ultimoIDInserido, erro := resultado.LastInsertId()
	if erro != nil {
		return 0, erro
	}

	return uint64(ultimoIDInserido), nil
}
