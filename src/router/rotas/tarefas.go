package rotas

import (
	"net/http"
	"weni/api/src/controllers"
)

// rotasTarefas é um slice de structs do tipo Rota
var rotasTarefas = []Rota{
	{
		URI:    "/tarefas",
		Metodo: http.MethodPost,
		Funcao: controllers.CriarTarefa,
	},
	{
		URI:    "/tarefas",
		Metodo: http.MethodGet,
		Funcao: controllers.MostrarTarefas,
	},
	{
		URI:    "/tarefas/abertas",
		Metodo: http.MethodGet,
		Funcao: controllers.MostrarTarefasAbertas,
	},
	{
		URI:    "/tarefas/fechadas",
		Metodo: http.MethodGet,
		Funcao: controllers.MostrarTarefasFechadas,
	},
	{
		URI:    "/tarefas/{tarefaId}",
		Metodo: http.MethodPut,
		Funcao: controllers.AlterarTarefa,
	},
	{
		URI:    "/tarefas/abrir/{tarefaId}",
		Metodo: http.MethodPut,
		Funcao: controllers.AbrirTarefa,
	},
	{
		URI:    "/tarefas/fechar/{tarefaId}",
		Metodo: http.MethodPut,
		Funcao: controllers.FecharTarefa,
	},
	{
		URI:    "/tarefas/abertas/{tarefaId}",
		Metodo: http.MethodDelete,
		Funcao: controllers.DeletarTarefaAberta,
	},
	{
		URI:    "/tarefas/fechadas/{tarefaId}",
		Metodo: http.MethodDelete,
		Funcao: controllers.DeletarTarefaFechada,
	},
}
