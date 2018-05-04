# Arrays en GO

- Los arrays se caracterizan por tener un numero limitado de espacio el cual es definido al momento de ser creado.
- var <nombre> [<valor-nummerico>]<tipo-de-dato>
- Todo los valores deben ser del mismo tipo
- El tamaño es fijo y se declara junto con la variable
- El indice de los arrays comienza en 0
- Para asignar un indice se usa:
    ```
    go[<valor-del-indice>]
    ```
- Tambien podemos crear un array de manera dinamica
    De la sgte manera:
    ```go
    <var> := [<valor-num>]<tipo-de-var>{valores,del,array}
    ```
- Para ver el tamaño del array se usa la funcion LEN() y esta nos devuelve la capacidad total del array
- El valor del array cuando esta vacio es el valor vacio del tipo de dato con el que se creo
