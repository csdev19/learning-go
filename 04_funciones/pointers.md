# Punteros

- Un puntero es la direccion hacia un espacio de memoria 
- Cuando nosotros llamamos a una funcion, esta requiere argumentos y estos argumentos son copiados hacia la funcion:
    ```go
    func zero(x int) {
        x = 0
    }

    func main() {
        x := 5
        zero(x)
        fmt.Println(x) // sigue siendo 5
    }
    ```
- En este programa, la funcion **zero** no modifica el **x** que esta dentro de la **funcion principal (func main)**. Pero como le hariamos si queremos que si se cambie? una manera de hacer esto es usando un tipo de dato especial llamado **pointer**:
    ```go
    func zero(xPtr *int) {
        *xPtr = 0
    }
    func main() {
        x := 5
        zero(&x)
        fmt.Println(x) // ahora si es 0
    }
    ```
- Los punteros referencian hacia una posicion de memoria donde se guardan valores en lugar del valor en si mismo. Usando un puntero **(*int)**, la funcion es capas de modificar el valor original.

## The * and & operators

- En **GO**, un puntero es representado usando un asterisco  **(*)** seguido del tipo del valor guardado. En la funcion zero, xPtr es un puntero hacia un **int**.
- Un asterisco tambien es usado para eliminar una referencias de variables hacia un puntero. Desreferenciar un puntero nos da acceso al valor que se esta apuntando. Cuando escribirmos **(*xPtr = 0)** , nosotros estamos diciendo: "guarda el int 0 dentro de la locacion de memoria a la cual xPtr es referida". Si nosotros tratamos de ejecutar xPtr=0 nosotors recibiremos un **compile-time error** porque xẗr no es un entero **(int)**, es un **(*int)**, el cual solo puede recibir otro **(*int)**.
- Finalmente, usamos el operador **&** para encontrar la direccion de la variable. **&numero** retorna un **(*int)**(puntero hacia un int) porque numero es un numero entero **int**. Esto nos permite modificar la variable original. **&x** en **func main** y **xPtr** en **zero** son referidos hacia la mismo locacion de memoria.

### New
- Otro camino para obtener un puntero es usando una funcion prefabricada llamada **new**:
    ```go
    func one(xPtr, *int) {
        *xPtr = 1
    }

    func main() {
        xPtr := new(int)
        one(xPtr)
        fmt.Println(*xPtr) // x es 1
    }
    ```
- **new** toma un tipo de dato como argumento, asigna suficiente memoria para ajustarse a un valor de ese tipo y devuelve un puntero a el.
- En otros lenguajes de programacion, hay una diferencias muy significante entre usar **new** y **&**, con el cuidado necesario para eventualmente eliminar cualquier cosa creada con **new**. No nos debemos preocupar acerca de si con **GO** (y su **garbage-collected programming languaje**), lo que significa que la memoria es limpiada automaticamente cuando nada es referencia a ella nunca mas.
- Punteros son raramente usados en **GO** con tipos de datosprefabricados, pero son extremadamente utilas cuando son usados junto con las estructuras **struct**.
