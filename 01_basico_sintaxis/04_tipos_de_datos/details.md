# Tipos de datos

## [Referencia] (https://golang.org/ref/spec#Types)

- Booleanos
    ```go
    fun main () {
        // valores unicos -> true/false

        // Operadores  &&(and) ||(or) !(not)
        fmt.Println(true&&true)
        fmt.Println(true&&false)
        fmt.Println(true||true)
        fmt.Println(!true)
    }
    ```
- Numerico
    - uint(8|16|32|64): numero sin signo (aprox 0 to 2x9x10^20)
    - int(8|16|32|64): enteros (aprox -9x10^20 to +9x10^20)
    - float(32|64): numero de punto flotante (decimales)
    - complex(64|128): numeros imaginarios
    - byte: alias para uint8
    - rune: alias para  int32


- En go debes ser exacto al momento de sumar variables
    - puedes sumar: int + int
    - pero no: int8 + int || int8 + int16 || ... etc ...

    - asi que en todo caso debes castear(cambio TEMPORAL de una variablen)
     ```go
     <tipo-de-dato-a-cambiar>(<variable>)
     ```
     OJO: tomar en cuenta las restricciones del tipo de dato como int8|16|32|64 tienen diferentes limites
    - prioridad aritmetica (  * / + - )


- Para formear un texto se usa:
    ```go
     Printf("hola $uno $dos", uno, dos)
    ```
- Y para ver el tipo de una variable
    ```go
    fmt.Printf("var es de tipo: %T", var)
    ```

## Referencia 2 [Introducing Go]()

- Strings
    ```go
    func main () {
        var texto string
        // ojo debemos usar las dobles comillas "" y no '' porque las comillas '' solo sirven para caracteres unicos
        texto = "hola"

        num := len(texto) // mira el tamaño de la cadena de texto
        letra := texto[2] // nos devuelve un caracter en particular de la cadena
        texto_largo := texto + "mundo" // y se puede concatenar

        fmt.Println(texto_largo)

    }
    ```
