# Maps

- Son como diccionarios de otros lenguajes
- La forma de crearlos es:
    ```go
    package main

    import "fmt"

    func main(){
        var dicc = make(map(string)string)
        fmt.Println(dicc)
    }
    ````
    - Donde se tiene que especificar los tipos de datos que se van a usar tanto en las llaves como en los valores
        ```go
        var <nombre-de-var> = make(map(<tipo-de-dato-de-llave>)<tipo-de-dato-del-valor>)
        ```

- Una forma de declarar un map con valores iniciales es:
    ```go
    package main

    import "fmt"

    func main(){
        var dicc = map(string)string{"hola1":"cristian", "hola2": "diego"}
        fmt.Println(dicc)
    }
    ````
- Para llamar a los valores de un map se usan **[]** y dentro de estos iria el valor de la llave. Como se vio en ejemplos anteriores para guardar cualquier valor primero se dice el valor de la llave.
- Una cosa interesante acerca de los maps es:
    ```go
    package main

    import "fmt"

    func main(){
        var dias1 = make(map([string]string))
        dias1["lunes"] = "Lunes"
        dias1[martes] = "Martes"
        dias1["miercoles"] = "Miercoles"

        //
        var dias2 = map([string]string){"lunes":"Lunes","martes":"Martes","miercoles":"Miercoles"}

        // pueden resultar distitos en cuanto a ordenamiento
        fmt.Println(dias1)
        fmt.Println(dias2)
    }
    ```
    - Y porque pasa esto?. Porque en go los diccionarios se iteran aleotiramente asi que nos obliga a especificar que llave queremos que busque.

- Si queremos borrar un elemento del mapa usamos DELETE()
    ```go
    package main

    import "fmt"

    func main(){
        dias := map[string]string{
            "lun" : "Lunes",
            "mar" : "Martes",
            "mier": "Miercoles",
        }
        fmt.Println(dias)

        delete(dias, "mar")

        fmt.Pritln(dias)
    }
    ```
- Importante!! Cuando se llama a un valor del mapa por su llave esta devuelve **su VALOR** y **un BOOLEANO** (para saber si existe o no)
- Para validar si una llave existe podemos hacer:
    ```go
    func main(){
        dias = map([string]string){"lunes":"Lunes", "martes": "Martes"}
        if dia, ok := dias["miercoles"]{
            // y esto pasa porque devuelve dos variables
            fmt.Println(dia, "existe", "y es", ok)
        }
    }
    ```
