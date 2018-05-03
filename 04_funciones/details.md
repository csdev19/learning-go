# Funciones

- Las funciones de declaran con la palabra reservada **func** seguido del nombre de la funcion **()** y las llaves que encierran el dominio de la funcion.
    ```go
    package main

    import "fmt"

    func <nombre-de-funcion> () {}
    func main() {}
    ```
- La funcion puede estar abajo o arriba de la funcion principal **(func main())**.
- En las funciones toda variable creada o mandada hacia esta se queda dentro de ese **scope**.
- Asi como en otros lenguajes las funciones pueden ser pasadas como valor a una variable
    ```go
    package main

    import (
        "fmt"
        "math"
    )

    func main() {
        hypot := func(x, y float64) float64 {
            return math.Sqrt(x*x + y*y)
        }

        fmt.Println(hypot(3, 4))
    }
    ```

## Funciones con parametros
- A las funciones les podemos pasar parametros, los cuales deben ser establecidos junto con la funcion.
    ```go
    package main

    import "fmt"

    func saludo(nombre string) {
        fmt.Printf("Hola %s", string)
    }

    func main() {
        saludo("Cristian")
    }
    ```

## Funciones que retornan valores
- Para retornar debemos usar la palabra reservada **return**
- Debemos decir de dato queremos devolver como se vera a continuacion
    ```go
    package main

    import "fmt"

    func main() {
        sum := suma(3,5)
        fmt.Println(sum)
    }
    func suma (x int, y int) int {
        return x + y
    }

    ```
- Tambien podemos recibir un tipo y devolver otro completamente diferente
    ```go
    func esMenor(edad uint8) bool {
        if edad < 20 {
            return true
        }else {
            return false
        }
    }
    ```


## Funciones que retornan multiples valores
- Por lo cual devuelve un slice []**tipo-de-dato**
    ```go
    package main

    import "fmt"

    func main() {
        num, ok = isPositive(12)
    }

    func isPositive(integer int)  (int, bool) {
        if integer > 0 {
            return integer,true
        }
    }
    ```
- Como en otros lenguajes aceptan listas y diccionarios. Go acepta pasar slices

