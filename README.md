# Saludos en go

Este paquete proporciona una forma simple de obtener saludo personalizados en Go.

## Instalacion
Ejecuta el siguiente comado para instalar paquete:
```bash
go get -u github.com/cc-santiago-alvarez/greetings
```

## Uso
Aqui tienes un ejempl de como utilizar el paquete en tu código:

```go
package main

import (
	"fmt"

	"github.com/cc-santiago-alvarez/greetings"
)

func main() {

	message, err := greetings.Hello("Juan")

	if err != nil {
		fmt.Println("Ocurrio un error", err)
		return
	}

	fmt.Println(message)
}
```
