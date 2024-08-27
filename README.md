# Rutificagor

**Rutificagor** es una biblioteca en Go (Golang) diseñada para manipular y validar RUTs (Rol Único Tributario) chilenos. La biblioteca permite generar, validar y formatear RUTs con facilidad.

## Instalación

Para usar **rutificagor** en tu proyecto, puedes instalarlo mediante `go get`:

```bash
go get github.com/grayjacketstudios/rutificagor
```

## Uso

### Importar el paquete

```go
import (
    "github.com/grayjacketstudios/rutificagor"
)
```

### Funciones Disponibles
1. `ObtenerDV`
Genera (calcula) el dígito verificador (DV) de un RUT.

```go
dv, err := rutificagor.ObtenerDV("11111111")
if err != nil {
    // Manejo del error
}
fmt.Println(dv) // Output: '1'
```

2. `ValidarRut`
Valida un RUT dado en diferentes formatos.

```go
isValid, err := rutificagor.ValidarRut("11.111.111-1")
if err != nil {
    // Manejo del error
}
fmt.Println(isValid) // Output: true
```

3. `GenerarRut`
Genera un RUT aleatorio dentro de un rango específico, con el DV correspondiente.
```go
rut := rutificagor.GenerarRut(8000000, 24000000)
fmt.Println(rut) // Output: "12345678-9"
```

4. `GenerarRutRandom`
Genera un RUT aleatorio en un rango predefinido (4 millones a 99 millones).
```go
rut := rutificagor.GenerarRutRandom()
fmt.Println(rut) // Output: "87654321-K"
```

5. `FormatearRut`
Formatea un RUT en diferentes estilos, según la opción especificada.

- Opción 1: Formato con puntos y guion (ej. "11.111.111-1")
- Opción 2: Sin puntos, con guion (ej. "11111111-1")
- Opción 3: Sin puntos ni guion (ej. "111111111")

```go
rutFormated, err := rutificagor.FormatearRut("11111111-1", 1)
if err != nil {
    // Manejo del error
}
fmt.Println(rutFormated) // Output: "11.111.111-1"
```

## Manejo de Errores
La biblioteca utiliza errores personalizados para manejar casos específicos como entradas vacías, entradas inválidas o opciones no válidas.

Ejemplo:
```go
if err != nil {
    switch err.(type) {
    case *customerrors.EmptyInputError:
        fmt.Println("La entrada está vacía.")
    case *customerrors.InvalidInputError:
        fmt.Println("La entrada es inválida.")
    case *customerrors.InvalidOptionError:
        fmt.Println("La opción de formato es inválida.")
    default:
        fmt.Println("Error desconocido.")
    }
}
```

## Contribuir

Las contribuciones son bienvenidas. Por favor, abre un "issue" o envía un "pull request" en GitHub.
## Licencia

Este proyecto está licenciado bajo la Licencia MIT. Consulta el archivo LICENSE para obtener más detalles.