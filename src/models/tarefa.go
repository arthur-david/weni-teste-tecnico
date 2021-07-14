// Package models contém o modelo de tarefa usado na API, além de preparar os dados desse modelo, caso haja algum erro de escrita.
package models

import (
	"errors"
	"strings"
)

// Tarefa representa a estrutura de uma tarefa.
type Tarefa struct {
	ID      uint64 `json:"id,omitempty"`
	Title   string `json:"title"`
	Checked bool   `json:"checked"`
}

// Preparar valida e formata a tarefa recebida através dos métodos validar e formatar.
func (t *Tarefa) Preparar() error {
	if erro := t.validar(); erro != nil {
		return erro
	}
	t.formatar()

	return nil
}

// validar verifica se o campo Title do modelo está em branco.
func (t *Tarefa) validar() error {
	if t.Title == "" {
		return errors.New("o campo Title é obrigatório e não pode estar em branco")
	}

	return nil
}

// formatar retira todos os espaços vazios no final e no começo do Title.
func (t *Tarefa) formatar() {
	t.Title = strings.TrimSpace(t.Title)
}
