# TP - TDL

En este repo se encuentra el proyecto integrador de GO para la materia Teoría del Lenguaje. Acá van
a poder encontrar dos versiones del proyecto, el cual es una idea del ageofempires implementado en
Go, además de un par de ejemplos y el frontend.

## AGE OF EMPIRES

Esta carpeta contiene la versión principal de la implementación del proyecto integrador.

Para levantar este proyecto ya en la raíz de la carpeta `age-of-empires` se debe correr:

```bash
$ go install .
$ ageofempires
```

Y levantamos el front parados en la carpeta `front` con:

```
$ npm start
```

Por último abrimos nuestro navegador en `localhost:3000`.

## AGE OF EMPIRES CHANNELS

Esta es la otra versión del proyecto integrador. La diferencia con la anterior es el uso de channels
y statefuls goroutines para la implementación del Warehouse, cuando en la versión principal usamos
solo un mutex para el mismo.

Para levantarlo es lo mismo que en el otro caso. Ya en la carpeta `age-of-empires-channels` corremos:

```bash
$ go install .
$ ageofempires
```

Y levantamos el front parados en la carpeta `front` con:

```
$ npm start
```

Por último abrimos nuestro navegador en `localhost:3000`.

## PRODUCER CONSUMER

Es esta carpeta hay varios ejemplos de productor y consumidor para mostrar diferentes formas en la
que se puede hacer.

Los ejemplos que se encuentran sueltos se pueden correr con:

```bash
$ go run archivo_ejemplo.go
```

Y el de la carpeta `only-channels`, ya dentro de esta carpeta, se puede correr con:

```bash
$ go run *.go
```

## WEBSOCKETS

Acá se encuentra un ejemplo corto sobre cómo usar websockets para conectarse a un front.

## PRESENTACIÓN

Carpeta con los archivos necesarios para la presentación que se hizo sobre el lenguaje.
En esta se puede encontrar un README.md con más detalles sobre cómo levantar este.

## HELLO API

Un mini ejemplo de cómo usar http desde GO.
Dentro de esta carpeta hay un README.md con detalles sobre cómo levantar el mismo.
