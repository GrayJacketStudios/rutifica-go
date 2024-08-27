package rutificagor

import (
	"errors"
	"strconv"
	"strings"
	"testing"

	customerrors "github.com/grayjacketstudios/rutificagor/customErrors"
)

func TestObtenerDV(t *testing.T) {
	testCases := []struct {
		input    string
		expected rune
		err      error
	}{
		{"18622178", '8', nil},
		{"11111111", '1', nil},
		{"70360100", '6', nil},
		{"22222222", '2', nil},
		{"15304340", 'K', nil},
		{"8367720", '1', nil},
		{"18.622.178", '8', nil},
		{"1111111f", '0', &customerrors.InvalidInputError{}},
		{"", '0', &customerrors.EmptyInputError{}},
	}

	for _, tc := range testCases {
		if result, err := ObtenerDV(tc.input); tc.expected != result || errors.Is(err, &customerrors.InvalidInputError{}) {
			t.Errorf("ObtenerDV(%s) = %v; want: value: %v expected error: %v, received error: %v", tc.input, result, tc.expected, tc.err, err)
		}
	}
}

func TestValidarRut(t *testing.T) {
	testCases := []struct {
		input    string
		expected bool
		err      error
	}{
		{"186221788", true, nil},
		{"18622178-8", true, nil},
		{"18.622.178-8", true, nil},
		{"18..622.1788", true, nil},
		{"11111111-1", true, nil},
		{"11111111-3", false, nil},
		{"15304340-k", true, nil},
		{"15304340-K", true, nil},
		{"15g#4z52-4", false, &customerrors.InvalidInputError{}},
		{"", false, &customerrors.EmptyInputError{}},
	}

	for _, tc := range testCases {
		if result, err := ValidarRut(tc.input); tc.expected != result || errors.Is(err, &customerrors.InvalidInputError{}) {
			t.Errorf("ValidarRut(%s) = %v; want: value: %v | expected error: %v, received error: %v", tc.input, result, tc.expected, tc.err, err)
		}
	}
}

func TestGenerarRutRandom(t *testing.T) {
	for range 10 {
		if result := GenerarRutRandom(); !(len(result) >= 9 && len(result) <= 10) || !strings.Contains(result, "-") {
			t.Errorf("GenerarRutRandom() = %s; value should be between 9 and 10 characters, with the DV split by '-'", result)
		}
	}
}

func TestGenerarRut(t *testing.T) {
	testCases := []struct {
		min int
		max int
	}{
		{
			min: 17000000,
			max: 25000000,
		},
		{
			min: 4000000,
			max: 25000000,
		},
		{
			min: 60000000,
			max: 75000000,
		},
		{
			min: 5000000,
			max: 9000000,
		},
	}
	for _, tc := range testCases {
		result := GenerarRut(tc.min, tc.max)
		split := strings.Split(result, "-")
		rut, _ := strconv.Atoi(split[0])
		if rut < tc.min || rut > tc.max {
			t.Errorf("GenerarRut(%d, %d) = %s; value is out of min max range", tc.min, tc.max, result)
		}
		if len(split[1]) > 1 {
			t.Errorf("GenerarRut(%d, %d) = %s; DV should be only 1 character long", tc.min, tc.max, result)
		}
	}
}

func TestFormatPuntosGuion(t *testing.T) {
	testCases := []struct {
		rut      string
		formated string
	}{
		{rut: "18622178-8", formated: "18.622.178-8"},
		{rut: "11111111-1", formated: "11.111.111-1"},
		{rut: "8000000-8", formated: "8.000.000-8"},
	}
	for _, tc := range testCases {
		if result := formatPuntosGuion(tc.rut); result != tc.formated {
			t.Errorf("formatPuntosGuion(%s) = %s; want: %s", tc.rut, result, tc.formated)
		}
	}
}

func TestFormatearRut(t *testing.T) {
	testCases := []struct {
		rut      string
		option   int
		formated string
		err      error
	}{
		{rut: "18622178-8", option: 1, formated: "18.622.178-8", err: nil},
		{rut: "11111111-1", option: 1, formated: "11.111.111-1", err: nil},
		{rut: "8000000-8", option: 1, formated: "8.000.000-8", err: nil},
		{rut: "18.622.178-8", option: 2, formated: "18622178-8", err: nil},
		{rut: "11.111.111-1", option: 2, formated: "11111111-1", err: nil},
		{rut: "8.000.000-8", option: 2, formated: "8000000-8", err: nil},
		{rut: "18.622.178-8", option: 3, formated: "186221788", err: nil},
		{rut: "11.111.111-1", option: 3, formated: "111111111", err: nil},
		{rut: "8.000.000-8", option: 3, formated: "80000008", err: nil},
		{rut: "18622178-8", option: 23, formated: "18622178-8", err: &customerrors.InvalidOptionError{}},
	}
	for _, tc := range testCases {
		if result, err := FormatearRut(tc.rut, tc.option); result != tc.formated || tc.err != nil && err == tc.err {
			t.Errorf("formatPuntosGuion(%s, %d) = %s; want: %s; error wanted: %s", tc.rut, tc.option, result, tc.formated, tc.err)
		}
	}
}
