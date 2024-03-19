package main

import "testing"

func TestShouldSumCorrectTest(t *testing.T) {
	teste := soma(3, 2, 1)
	resultado := 6

	if teste != resultado {
		t.Error("Valor esperado:", resultado, "Valor retornado:", teste)
	}
}

func TestShouldSumCorrectTest2(t *testing.T) {
	teste := soma(5, 4, 3)
	resultado := 12

	if teste != resultado {
		t.Error("Valor esperado:", resultado, "Valor retornado:", teste)
	}
}

func TestShouldMultCorrectTest(t *testing.T) {
	teste := multiplica(10, 10)
	resultado := 100
	if teste != resultado {
		t.Error("Valor esperado:", resultado, "Valor retornado:", teste)
	}
}

func TestShouldMultCorrectTest2(t *testing.T) {
	teste := multiplica(2, 5)
	resultado := 10
	if teste != resultado {
		t.Error("Valor esperado:", resultado, "Valor retornado:", teste)
	}
}

func TestShouldSubCorrectTest(t *testing.T) {
	teste := substrai(50, 12)
	resultado := 38
	if teste != resultado {
		t.Error("Valor esperado:", resultado, "Valor retornado:", teste)
	}
}

func TestShouldSubCorrectTest2(t *testing.T) {
	teste := substrai()
	resultado := 0
	if teste != resultado {
		t.Error("Valor esperado:", resultado, "Valor retornado:", teste)
	}
}

func TestShouldDivCorrectTest(t *testing.T) {
	teste := divide(10, 2)
	resultado := 5
	if teste != resultado {
		t.Error("Valor esperado:", resultado, "Valor retornado:", teste)
	}
}

func TestShouldDivCorrectTest2(t *testing.T) {
	teste := divide()
	resultado := 0
	if teste != resultado {
		t.Error("Valor esperado:", resultado, "Valor retornado:", teste)
	}
}
