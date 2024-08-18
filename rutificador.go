package rutificador

import (
	"strconv"
	"strings"

	customerrors "github.com/grayjacketstudios/rutificador/customErrors"
	"github.com/grayjacketstudios/rutificador/utils"
)

var multiplicador = [9]int{2, 3, 4, 5, 6, 7, 2, 3, 4}

// Genera el digito verificador de un rut, a partir de un string con los numeros requeridos, los cuales pueden estar o no separados por puntos.
func obtenerDV(rut string) (dv rune, err error) {
	rut = strings.ReplaceAll(rut, ".", "")
	if len(rut) == 0 {
		return dv, &customerrors.EmptyInputError{}
	}
	rut = utils.Reverse(rut)
	_, err = strconv.Atoi(rut)
	if err != nil {
		return dv, &customerrors.InvalidInputError{Input: rut}
	}

	suma := 0
	for i, dig := range rut {
		suma += (int(dig) - '0') * multiplicador[i]
	}
	res := 11 - (suma - (int(suma/11) * 11))
	switch res {
	case 11:
		return 0, nil
	case 10:
		return 'K', nil
	default:
		return rune(res), nil
	}

}
