# Verificador de Palíndromos

![technology Go](https://img.shields.io/badge/technology-go-blue.svg)

Este programa en Go verifica si un número entero ingresado es un palíndromo.

## ¿Qué es un palíndromo?

Un palíndromo es una secuencia de caracteres que se lee igual de izquierda a derecha y de derecha a izquierda. En el caso de este programa, se trata de verificar si un número entero dado es un palíndromo cuando se lee en ambas direcciones.

## Cómo funciona el programa

1. El programa muestra un menú en la consola con las siguientes opciones:
   - Opción 1: Calcular palíndromo de un número entero.
   - Opción 2: Salir del programa.

2. Si el usuario elige la opción 1, se le solicita ingresar un número entero.

3. El programa verifica si el número ingresado es un palíndromo utilizando el siguiente enfoque:
   - Convierte el número entero en una cadena de texto.
   - Invierte la cadena de texto.
   - Compara la cadena original con la cadena invertida.
   - Si las cadenas son iguales, el número es un palíndromo. De lo contrario, no lo es.

4. El programa muestra un mensaje indicando si el número ingresado es un palíndromo o no.

5. El usuario puede repetir el proceso seleccionando una opción válida del menú o puede elegir la opción 2 para salir del programa.

## Requisitos
- Go (es compatible con cualquier versión de Go que sea igual o posterior a la versión 1.0).
*No se utilizan características específicas de una versión en particular, por lo que debería funcionar sin problemas en las versiones más recientes de Go.*

## Instalación
1. Clona este repositorio o descarga los archivos del programa.
2. Navega hasta el directorio del programa.

## Uso
1. Abre una terminal y navega hasta el directorio del programa.
2. Ejecuta el programa con el siguiente comando:
```go
go run main.go
```
3. Sigue las instrucciones en pantalla para interactuar con el programa. 
4. Elige la opción "1" para calcular si un número entero es un palíndromo. 
5. Ingresa un número entero cuando se te solicite. 
6. El programa te mostrará si el número ingresado es un palíndromo o no. 
7. Para salir del programa, elige la opción "2". 

***Si ingresas una opción inválida o no ingresas un número entero válido, el programa te mostrará un mensaje de error y te solicitará ingresar nuevamente el número o la opción correspondiente.***

## Funcionalidades
* Opción 1: Permite al usuario ingresar un número entero y verifica si es un palíndromo.
* Opción 2: Permite al usuario salir del programa.

## Contribuciones
Las contribuciones son bienvenidas. Si deseas mejorar o agregar nuevas funcionalidades al programa, por favor sigue los siguientes pasos:

1. Crea un fork del repositorio. 
2. Crea una nueva rama para tu funcionalidad (git checkout -b nueva-funcionalidad). 
3. Realiza los cambios necesarios y haz commits. 
4. Envía un pull request con tus cambios.

## Autor
- Nombre: Oscar Iván Rada Osorio
- Correo electrónico: [oscar.rada@mercadolibre.com.co](mailto:oscar.rada@mercadolibre.com.co)
