# Las estructuras condicionales

- IF

    - Solo acepta el valor booleano TRUE

    - Usa los operadores AND/OR para determinar el valor

    - Se puede crear una variable en medio del enunciado

    - Y cuando declaramos la variable esta se puede acceder desde los (else-if) y (else)
    ```go
    if v := 0; v < 10 {
        return v
    }
    ```

- ELSE

    - ALERTA!! el else tiene que ir exactamente despues del if sino no funciona. ENSERIO!
    ```go
    if {
   } else {}
    ```
    - El else se ejecuta si el if es false

