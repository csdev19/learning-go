# Metodos en GO

- En **GO** podemos agregarle metodos a cualquier tipo de dato.
- En **GO** un metodo es una funcion que esta asociada a un tipo de dato.
- Ejemplo:
    ```go
    package main

    import "fmt"

    type Persona struct {
        nombre string
        edad int8
    }

    func (p Persona) saludar() {
        fmt.Println("hola soy una persona")
    }

    func main(){
        var persona Persona

        persona.saludar()
    }

    ```
- Se puede agregar metodos a tipos de datos que no son estructuras.
    ```go
    package main

    import "fmt"

    type Numero int

    func (n Numero) presentacion() {
        fmt.Println("hola soy un numero")
    }

    func main(){
        var numero Numero
        numero.presentacion() // hola soy un numero
    }
    ```
- Como se puede apreciar en los ejemplo anterioir cuando a un tipo de dato se le asigna una funcion es de la siguiente manera: **func (variable Nombre-tipo-de-dato) nombre-funcion()** pero **OJO** lo que esta entre **parentesis** se llama **receiver** y este es un parametro que nos permite enlazar una funcion a un tipo de dato.
- Y dentro de este **receiver** estara una variable del **tipo de dato** que queremos crear el metodo. Pero una particularidad es que si queremos cambiar el contenido de un campo de una estructura en particular debemos usar punteros.
    ```go
    package main

    import "fmt"

    type Punto struct {
        x int
        y int
    }

    func (p Punto) set_ubicacion1(x, y int) {
        p.x = x
        p.y = y
    }

    func (p Punto) set_ubicacion2(x, y int) {
        p.x = x
        p.y = y
    }

    func main(){
        var punto1 Punto
        var punto2 Punto

        fmt.Println(punto1)
        punto1.set_ubicacion1(5,5)
        fmt.Println(punto1)

        fmt.Println("-------------------")

        fmt.Pritnln(punto2)
        punto2.set_ubicacion2(5,5)
        fmt.Println(punto2)
    }
    ```
- Del ejemplo anterior cree dos metodos para demostrar que con punteros se podria modificar los valores de los campos de la estructura y sin punteros no se podria.
