# Slices en GO

- Los slices si pueden variar en cuanto a tamaño.
- Un slice es un segmento de un array. Y como los arrays los slices son indexables y tienen un tamaño. En cambio estos pueden variar.
- [Article of Go slides] (https://blog.golang.org/go-slices-usage-and-internals)
- Los slices representan una secuencia de tamaño de variable de elementos del mismo tipo. Por decirlo asi un array dinamico
    ```go
    func main (){
        // declaracion normal
        var x []float32
        // declaracion dinamica con la funcion make
        y := make([]float32)
        // declaracion dinamica
        z := []float32{<values>}
    }
    ```
- Tienen un apuntador a un array
- Los slices tienen 3 elementos especificos:
    1. Apuntador a un array
    2. Tamaño:  No es fijo(NO SE DECLARA)
    3. Capacidad
- Cuando tu creas un slice este tiene la capacidad de almacenar el doble de su tamaño original
- Go para casos asi administra muy bien la memoria, en caso del slice la duplica cada vez que se ve excedido
- Los slices tienen un metodo distinto para agregar info debemos usar la funcion APPEND(). De la sgte manera:
    ```go
    <nombre-var> = append(<nombre-var>, <elemento-a-agregar>)
    ```
- Para saber la capacidad de un slice se usa la funcion CAP(). Y como podriamos modificar un elemento de los slices
    ```go
    <nombre-var>[<indice>] = <valor>
    ```
- Forma comun con la que se crean los slices con la funcion MAKE():
    ```go
    <nombre-var> := make([]<tipo-de-dato>, <tamaño-inic>, <capacidad-inicial>)
    ```
    - A la funcion le pasamos 3 valores: 2 primeros obligatorios y ultimo opcional porque se puede decir que es implicito

- Slices **nil**:
    - El valor por defecto de un slice es nil.
    - Un slice nil tiene un tamaño y una longitud de 0.
    ```go
    package main

    import "fmt"

    func main() {
        var z []int
        fmt.Println(z, len(z), cap(z))
        if z == nil {
            fmt.Println("¡nulo!")
        }
    }
    ```


- Dato interesante:
    - las variables se alamcenan en memoria ram y por esto se debe tener un control de que tantas se usan

