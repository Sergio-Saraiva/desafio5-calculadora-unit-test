package main

import "testing"

func TestSoma(t *testing.T) {
	teste := soma(1, 2, 3, 4, 5)
	resultado := 15
	if teste != resultado {
		t.Errorf("Soma esperada: %d, obtida: %d", resultado, teste)
	}
}

func TestSoma2(t *testing.T) {
	teste := soma(1, 2, 3, 4, 5)
	resultado := 14
	if teste != resultado {
		t.Errorf("Soma esperada: %d, obtida: %d", resultado, teste)
	}
}

func TestMultiplica(t *testing.T) {
	teste := multiplica(1, 2, 3, 4, 5)
	resultado := 120
	if teste != resultado {
		t.Errorf("Multiplicação esperada: %d, obtida: %d", resultado, teste)
	}
}

func TestMultiplica2(t *testing.T) {
	teste := multiplica(1, 2, 3, 4, 5)
	resultado := 121
	if teste != resultado {
		t.Errorf("Multiplicação esperada: %d, obtida: %d", resultado, teste)
	}
}

func TestSubtrai(t *testing.T) {
	teste := subtrai(1, 2, 3, 4, 5)
	resultado := -15
	if teste != resultado {
		t.Errorf("Subtração esperada: %d, obtida: %d", resultado, teste)
	}
}

func TestSubtrai2(t *testing.T) {
	teste := subtrai(1, 2, 3, 4, 5)
	resultado := -14
	if teste != resultado {
		t.Errorf("Subtração esperada: %d, obtida: %d", resultado, teste)
	}
}

func TestDivide(t *testing.T) {
	teste := divide(1, 2, 3, 4, 5)
	resultado := 0
	if teste != resultado {
		t.Errorf("Divisão esperada: %d, obtida: %d", resultado, teste)
	}
}

func TestDivide2(t *testing.T) {
	teste := divide(1, 2, 3, 4, 5)
	resultado := 1
	if teste != resultado {
		t.Errorf("Divisão esperada: %d, obtida: %d", resultado, teste)
	}
}
