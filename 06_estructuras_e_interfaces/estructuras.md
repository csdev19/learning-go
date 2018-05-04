# Estructuras (struct)

- Aunque puede ser posible para nosotros escribir programas solo usando los **tipos de datos prefabricados de Go**, en cierto punto llegara a ser un poco tedioso. Considerando un programa que interactura con formas.
    ```go
    package main

    import (
        "fmt"
        "math"
    )

    func distance(x1, y1, x2, y2 float64) float64 {
        a := x2 – x1
        b := y2 – y1
        return math.Sqrt(a*a + b*b)
    }

    func rectangleArea(x1, y1, x2, y2 float64) float64 {
        l := distance(x1, y1, x1, y2)
        w := distance(x1, y1, x2, y1)
        return l * w
    }

    func circleArea(x, y, r float64) float64 {
        return math.Pi * r*r
    }

    func main() {
        var rx1, ry1 float64 = 0, 0
        var rx2, ry2 float64 = 10, 10
        var cx, cy, cr float64 = 0, 0, 5
        fmt.Println(rectangleArea(rx1, ry1, rx2, ry2))
        fmt.Println(circleArea(cx, cy, cr))
    }
    ```
- Este programa encuentra el area de un rectangulo y un circulo. Hacer un seguimiento de todas las coordenadas hace dificil de ver lo uqe el programa esta haciendo y probablemente conduzca a errores.

## Structs
- Una manera facil de hacer este programa mejor es usando estructuras. Una estructura es un tipo de dato que contiene campos con sus respectivos nombres. Por ejemplo, podemos representar a un circulo de esta manera:
    ```go
    type Circle struct {
        x float64
        y float64
        z float64
    }
    ```
- La palabra clave **type** introduce un nuevo tipo. Y es seguido por el nombre del tipo (Circle), y la palabra clave **struct** que indica que estamos definiendo una estructura, y una lista de campos dentro de los **{}**
- Los campos son como un conjunto de variables agrupadas,dichos campos tienen un nombre, un tipo y es almacenado adyancente a los otros campos en la estructura. Al igual que con las funciones, podemos colapsar campos que tienen el mismo tipo:
    ```go
    type Circle struct {
        x ,y, r float64
    }
    ```

### Inicializacion
- Nosotros podemos crear una instancia de nuestro tipo **new Circle** en una serie de diferentes maneras:
    ```go
    var instancia Circle
    ```
- Como con otros tipos de datos, esto creara una variable local **Circle** y es por defecto igual a **zero**. Para una estructura, **zero** significa que cada uno de sus campos estara seteado con su respectivo valor **zero** (0 para int, 0.0 para floats , "" para strings, nil para punteros, etc.). Podemos tambien usar la funcion:
    ```go
    c := new(Circle)
    ```
- Esto asigno memoria para todos los campos, setea(establece) cada valor, y retorna un puntero hacia la estructura **(*Circle)**. Los punteros suelen ser usados ocn estructuras que con funciones puedan modificar sus contenidos.
- Usando **new** de esta forma es muy poco comun. Generalmente, cuando queremos darle valores iniciales a cada uno de los campos. Tenemos dos caminos:
    ```go
    // de primera opcion con sus respectivas llaves
    c := Circle{x:0, y:0, z:0}
    // Y de esta manera creamos el mismo circulo que en la linea anterior. y si quieres un puntero en la estructura usa &
    c := &Circle{0,0,0}
    ```

## Fields

- Podemos accesar a los campos usando el operador punto **(.)**:
    ```go
    type Point struct {
        x int
        y int
    }
    func main() {
        punto := Point{0,0}
        fmt.Println(punto.x, punto.y) // 0 0
        punto.x = 12
        punto.y = 10
        fmt.Println(punto.x, punto.y) // 12 10
    }
    ```
- Vamos a modificar la funcion distanciaPuntos que usa las estructuras Punto que ya creamos anteriormente
    ```go
    package main

    import (
	    "fmt"
	    "math"
    )

    type Punto struct {
    	x float64
    	y float64
    }

    func main() {
        // aqui puede ingresar los valores de los puntos para hallar la distancia
	    punto1 := Punto{0, 0}
	    punto2 := Punto{10, 0}

	    dist := distancia(punto1, punto2)
	    fmt.Println(dist)
    }

    func distancia(p1 Punto, p2 Punto) float64 {
	    x := p1.x - p2.x
	    y := p1.y - p2.y
	    result := math.Sqrt(x*x - y*y)
	    return result
    }
    ```
- Una cosa que tener en cuenta es que los argumentos pasados a la funcion son copiados en **GO**. Y si queremos modificarlos el valor original de la variable debemos usar punteros .
- Y tambien tener cuidado con los tipos de datos porque **Go** es muy exigente con eso.

## Methods

- 
-
