# Errores

- **GO** tiene un tipo de dato preconstruido para los **errores**. Y nosotros podemos crear nuestros propios **errores** usando la funcion **New** en el paquete de errores(errors package).
```go
package main

import (
    "fmt"
    "errors"
)

func main(){
    err := errors.New("Error message")
    fmt.Printf("%T\n", err)
}
```

- El valor zero del tipo error es **nil**.
- Ojo solo pasar **string** como parametro al metodo **New** del caso contrario saldria un error como el que se vera a continuacion:
```go
package main

import (
	"errors"
	"fmt"
)

func main() {
	err1 := errors.New("error message")
	err2 := errors.New(12)
	err3 := errors.New(12.45)
	err4 := errors.New(true)
	fmt.Printf("%T\n", err1)
	fmt.Printf("%T\n", err2)
	fmt.Printf("%T\n", err3)
	fmt.Printf("%T\n", err4)
}
``

```console
# command-line-arguments
./main.go:10:21: cannot use 12 (type int) as type string in argument to errors.New
./main.go:11:21: cannot use 12.45 (type float64) as type string in argument to errors.New
./main.go:12:20: cannot use true (type bool) as type string in argument to errors.New
```

- Ahora veamos un ejemplo mas real
```go
package main

import (
    "fmt"
    "errors"
)

func main(){
    resultado1, err1 := division(15, 2)

    if err1 != nil {
        fmt.Println("oh oh hubo un error")
    }

    fmt.Println(resultado1, err)

}

func division(dividendo, divisor float64) (resultado float64, err error){
    if divisor == 0 {
        err = errors.New("no se puede dividir entre 0")
    }
    resultado = dividendo / divisor
    // los dos son lo mismo
    //return (resultado, err)
}
```



