# Slices en GO

- Los slices representan una secuencia de tamaño de variable de elementos del mismo tipo. Por decirlo asi un array dinamico
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

- Dato interesante:
    - las variables se alamcenan en memoria ram y por esto se debe tener un control de que tantas se usan
