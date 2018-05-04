# Cosas importantes al momento de crear variables

- Declaracion tradicional de una variable 'cuerpo'
    ```go
    var <nombre-variable> <tipo-de-variable>
    ```

- Variables inicializadas
    - Al crear variables podemos asignarles valores iniciales
    ```go
    var a, b, c string = 'hola', 'como', 'estas'
    ```
    -  Cada una toma el valor por orden
           a = 'hola'
           b = 'como'
           c = 'estas'
    - Y tambien se pueden autoasignar
        ```go
        var uno, dos, tres = true, false, 21
        ```
    - Las variables automaticamente detectaran que valor se deposita en ellas

- Declaracion implicita de variables (permite el tipado dinamico -> ':=')
    - Solo se podra usar dentro de una funcion de lo contrario se devera usar la palabra var
    - La declaracion implicita nos permite obviar el usar la palabra **var** y detallar el tipo de dato, este se autoasignara  ejem:
        ```go
        func main () {
            a := 1
            b := 'hola'
            c := true
        }
        ```

- Declaracion de variables consecutivas
   - Al momento de crear variables variables podemos optar por:
    ```go
     var <var1> <tipo-var>(string), <var2> <tipo-var>(string)
    ```
   - Si las variables consecutivas son del mismo tipo podemos...
    ```go
     var <var1>, <var2> <tipo-var>(string)
    ```

- Definiendo multiples variables
    ```go
    func main() {
        var (
            a = 14
            b = 15
            c = 51
            d = 19
        )
    }
    ```

- **NOTA:** Para no usar una variable usaremos el **"_"**. El ejemplo esta en la seccion del ciclo **for**
