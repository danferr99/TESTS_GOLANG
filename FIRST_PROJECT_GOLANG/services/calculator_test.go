package services_test

import (
	"projetoTeste/services"
	"testing"
)

func TestTMB(t *testing.T) {

	if services.TMB(80, 1.89, 17) != 0 {
		t.Error("ERRO - NÃO PODE SER IGUAL A 0")
	}
}

func TestDiv(t *testing.T) {

	if services.Div(10, 3) != 2 {
		t.Error("ERRO - RESULTADO DA DIVISÃO DIFERENTE DE 2")
	}

}

func TestValidarSenha(t *testing.T) {

	if services.ValidarSenha("Daniel123") == "Daniel" {
		t.Error("ERRO - ESCOLHA OUTRA SENHA")

	}

	if services.ValidarSenha("Daniel123") == "123" {
		t.Error("ERRO - ESCOLHA OUTRA SENHA")

	}
	if services.ValidarSenha("Daniel123") == "" {
		t.Error("ERRO - ESCOLHA OUTRA SENHA")

	} else {
		println("SENHA VÁLIDA")

	}

}
