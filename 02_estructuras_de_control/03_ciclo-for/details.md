# For

## For clasico

- Es este:
    ```go
    package main

    import "fmt"

    func main(){
        for i:=0; i<10; i++ {
            fmt.Println(i)
        }
    }
    ```
- En **GO** el for no usa parentesis **()** ni siquiera son opcionales **no los pongas**.
- En go tambien puedes dejar **la inicializacion** y **el incremento** vacios.
    ```go
    package main

    import "fmt"

    func main(){
        sum := 1
        for ;sum < 1000; {
            sum += sum
        }
        fmt.Println(sum)
    }
    ```
- For es el **"while"** de Go
    ```go
    func main() {
        sum := 1
        for sum < 100 {
            sum += sum
        }
        fmt.Println(sum)
    }
    ```
- Para romper el ciclo podemos usar la palabra reservada **break**

## For continuo

- Es como un while en otros lenguajes
    ```go
    func main() {
        a := 0
        for a < 100 {
            a++
            fmt.Println(a)
        }
    }
    ```

## For eterno

- Ocurre cuando no se pone una condicion para terminar el for
    ```go
    func main() {
        for {
            fmt.Println("hola")
        }
    }
    ```

## For range

- Es un **ForEach** de otros lenguajes
- Un ejemplo
    ```go
    package main

    import "fmt"

    var pow = []int{1, 2, 4, 8, 16, 32, 64, 128}

    func main() {
        for i, v := range pow {
            fmt.Printf("2**%d = %d\n", i, v)
        }
    }
    ```
- Range nos devuelve dos elemento **la llave** y **el valor**
- **NOTA:** Y en **GO** cuando no queremos usar un valor debemos usar el **underscore** osea **"_"**. La cual sirve para que al no usar la variable no nos salte un error.
    ```go
    package main

    import "fmt"

    var pow = []int{1, 2, 4, 8, 16, 32, 64, 128}

    func main() {
        // usamos el _ para denotar que no lo usaremos y que no salte un error
        for _, v := range pow {
            fmt.Println(v)
        }
        // porque si no usariamos el _ nos saltaria un error
        for i,v := range pow {
            fmt.Println(v)
        }
    }
    ```
- En el for con range podemos omitir el valor y seguira funcionando.
    ```go
    for i := range <any-iter-item>{}
    ```
