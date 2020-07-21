# Present tool

## Estructura de carpetas

El archivo fuente de la presentación es `presentacion.slide`. Es un archivo de texto plano, el formato es bastante sencillo, parecido a markdown. En la presentación hay varios ejemplos de lo que se puede hacer en la presentación.

[Acá](https://pkg.go.dev/golang.org/x/tools/present) está todo lo que se puede hacer en la presentación, hay varios comandos con varias opciones cada uno.

La carpeta `examples` contiene los ejemplos de código que se linkean desde la presentación. Si los queremos correr, deben ser un ejemplo completo (con main, imports, etc.). Después se puede elegir mostrar una parte del archivo con START OMIT y END OMIT (ver ejemplos en el archivo de la presentación).

La carpeta `images` contiene las imágenes para linkear desde la presentación. Ahí es donde van los memes.

## Setup

Para correr la presentación se usa el paquete `present`, se instala con el siguiente comando:

```
go get golang.org/x/tools/cmd/present
```

## Correr la presentación

(Para Ubuntu) Primero, agregamos la carpeta de los paqutes de Go al PATH. Por ejemplo, podemos agregar esta línea en el archivo .bashrc:

```
export PATH=~/go/bin:$PATH
```

Pararse en la carpeta en donde está el archivo `presentacion.slide` y correr el comando

```
present
```

Luego, abrir un navegador, ir a: `http://127.0.0.1:3999` y seleccionar el archivo de la presentación.  
**Importante**: Es importante entrar por `127.0.0.1` y **no** por `localhost`, ya que si se entra por localhost no se puede ejecutar el código en las diapositivas.