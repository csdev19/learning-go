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

- Switch sin condicion: Un switch sin condición es lo mismo que switch true. Esta construcción puede ser una manera clara de escribir cadenas largas if-then-else.
    ```go
    package main

    import (
        "fmt"
        "time"
    )

    func main() {
        t := time.Now()
        switch {
        case t.Hour() < 12:
            fmt.Println("¡Buenos días!")
        case t.Hour() < 17:
            fmt.Println("Buenas tardes.")
        default:
            fmt.Println("Buenas noches.")
        }
    }
    ```
