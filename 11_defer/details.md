# Defer

- **GO** tiene una funcion que nos permite dejar en espera una accion.
- Este aspecto de las funcion ya fue cubierto en [funciones](https://github.com/pystudent1913/learning-go/blob/master/04_funciones/details.md) de manera basica y aqui tratare de profundizar en lo que pueda.
- Y en caso de usar **varios defer** usaremos la regla **LIFO(last in first out)** y esto que significa ?. Que el **ultimo defer** sera el primero en ejecutarse despues de que **la funcion main** se termine de ejecutar.
- **OJO**: En **GO** una funcion generalmente se termina cuando aparece la palabra reservada **return**.
