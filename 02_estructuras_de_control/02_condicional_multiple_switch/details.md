# Switch

- En go NO SE PONE UN BREAK porque es implicito

- Tambien podemos declarar variables dentro del switch asi swtich a:= <valor>; {} **PERO** al no poder poner la variable a evaluar debemos usar condicionales en los case pues esta si tiene acceso al scope
    - Resumen: Si declaras una variables usa condicionales sino no reconocera la variable

- Un ejemplo de switch (de go-tour)
    ```go
    package main

import (
    "fmt"
    "runtime"
)

func main() {
    fmt.Print("Go runs on ")
    switch os := runtime.GOOS; os {
    case "darwin":
        fmt.Println("OS X.")
    case "linux":
        fmt.Println("Linux.")
    default:
        // freebsd, openbsd,
        // plan9, windows...
        fmt.Printf("%s.", os)
    }
}
    ```
