# Teste Técnico Weni.

## Aplicação API Rest de uma To Do List.
---

### Requisito para Rodar a API:
- Banco de dados MySQL.


### Como integrar a API na sua aplicação:
- No arquivo .env altere o usuario e senha para os mesmos que estão configurados no seu banco de dados.
- O diretório sql possuí um arquivo sql.sql que contém todos os comandos para criar um database e duas tabelas que são usadas pela API.(Esse é apenas um modelo que pode ser alterado de acordo com o gosto do usuário, desde que satisfaça o funcionamento da aplicação).

### Funcionamento da aplicação:

Essa API recebe requisições com o protocolo HTTP e envia respostas de acordo com as requisições recebidas.

Os endpoints criados nessa aplicação são:
- POST: /tarefas > Para criar uma nova tarefa. Devolve um JSON da tarefa criada.
- GET: /tarefas" > Para retornar todas as tarefas.
- GET: /tarefas/abertas > Para retornar todas as tarefas não concluidas.
- GET: /tarefas/fechadas > Para retornar todas as tarefas concluidas.
- PUT: /tarefas/{tarefaId} > Para alterar o titulo de uma tarefa. Devolve um JSON da tarefa alterada.
- PUT: /tarefas/abrir/{tarefaId} > Para tirar uma tarefa da lista de concluidas. Decolve um JSON da tarefa aberta.
- PUT: /tarefas/fechar/{tarefaId} > Para concluir uma tarefa. Decolve um JSON da tarefa fechada.
- DELETE: /tarefas/abertas/{tarefaId} > Para apagar uma tarefa da lista de não concluidas.
- DELETE: /tarefas/fechadas/{tarefaId} > Para apagar uma tarefa da lista de concluidas.

Utilize no formato: localhost:5000/\<endpoint\>

Para testar e ter uma visão melhor do funcionamento dessa aplicação use o Postman. No diretório principal contem o arquivo "weni-api.postman_collection.json" para importação.

---

### Para mais detalhes do código veja a documentação.
