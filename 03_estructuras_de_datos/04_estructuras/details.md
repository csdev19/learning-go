# Estructura

- Son una secuencia de elementos nombrados, que se llaman campos. Cada uno de los cuales tienen un nombre y un tipo.
- Una **Estructura** es el "equivalente" a las clases de otros lenguajes de programacion.
- Para crearlas usamos la palabra reservada **type** 
    ```go
    package main

    import "fmt"

    type Point struct {
        pointX int
        pointY int
    }

    func main (){
        point = Point{3,2}
        fmt.Println(point)
    }
    ```
- Un modelo seria:
    ```go
    package main

    import "fmt"

    type <nombre-de-la-estructura> struct {
        <nombre-atributo> <tipo-de-dato>(string)
        <nombre-atributo> <tipo-de-dato>(int)
    }
    ```
- Y para usar una estructura
    ```go
    package main

    import "fmt"

    type Example struct {
        hola string
        hola2 string
    }

    func main(){
        var <nombre-var> <nombre-de-la-estructura>(Example)
        // y para acceder a un atributo de esta
        <nombre-var>.hola = <valor-string-este-caso>
        <nombre-var>.hola2 = <otro-valor>
    }
    ```
- De forma abreviada:
    ```go
    func main(){
        <nombre-var> := <nombre-estructura>{<valor1>,<valor2>}
    }
    ````
- Para las estructuras debemos tener cuidado que valores declaramos inicialmente porque go es muy especial para estos asuntos.
- Podemos asignar los valores por nombre y por posicion
    ```go
    type Point struct {
        x int
        y int
        z int
    }

    func main(){
        point1 := Point{
            1,
            2,
            3,
        }
        // is the same as
        point2 := Point{
            x: 1,
            y: 2,
            z: 3,
        }
    }
    ```
    - Y cuando usemos **"las llaves"** para asignar valores podemos usar cualquier orden.
- Al momento de crear la estructura el orden en el que declaras los valeos sera el orden en el que aparezcan cuando los imprimas con **Println**.
- En las estructuras puedes agrupar varios tipos de elementos.
