package rutificagor

import (
	"strconv"
	"strings"

	customerrors "github.com/grayjacketstudios/rutificagor/customErrors"
	"github.com/grayjacketstudios/rutificagor/utils"
)

var multiplicador = [9]int{2, 3, 4, 5, 6, 7, 2, 3, 4}

// Genera el digito verificador de un rut, a partir de un string con los numeros requeridos, los cuales pueden estar o no separados por puntos.
func ObtenerDV(rut string) (dv rune, err error) {
	rut = strings.ReplaceAll(strings.ReplaceAll(rut, ".", ""), "-", "")
	if len(rut) == 0 {
		return '0', &customerrors.EmptyInputError{}
	}
	rut = utils.Reverse(rut)
	_, err = strconv.Atoi(rut)
	if err != nil {
		return '0', &customerrors.InvalidInputError{Input: rut}
	}

	suma := 0
	for i, dig := range rut {
		suma += (int(dig) - '0') * multiplicador[i]
	}
	res := 11 - (suma - (int(suma/11) * 11))
	switch res {
	case 11:
		return '0', nil
	case 10:
		return 'K', nil
	default:
		return rune(res + '0'), nil
	}
}

/*
	Se valida el rut entregado, separando el digito verificador y constrastando con la funcion ObtenerDV, obteniendo un boolean y error como respuesta

Se puede enviar el rut en los siguientes formatos:
  - "11.111.111-1" (separados con punto y guion)
  - "11111111-1" (separacion del dv con guion)
  - "111111111" (sin puntos ni guion, incluyendo el DV)
*/
func ValidarRut(rut string) (bool, error) {
	if len(rut) <= 1 {
		return false, &customerrors.EmptyInputError{}
	}
	rut = strings.ToUpper(rut)
	lastDigit := rut[len(rut)-1]
	rut = rut[0 : len(rut)-1]
	dv, err := ObtenerDV(rut)
	if err != nil {
		return false, err
	}

	if dv != rune(lastDigit) {
		return false, nil
	}
	return true, nil
}
