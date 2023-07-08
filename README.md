1. go mod init


#Basic 
1. Conditionals

cd .\Basic\1-Conditionals\
go run .

Site:

2. Functions

cd .\Basic\2-Functions\
go run main.go

Site:

3. Packages

- Crear carpeta 3-Packages
cd 3-Packages
- Crear carpeta Math
- crear funcion math.go
- ejecutar comando dentro del folder math

referenciar desde el otro archivo

- para documentar un paquete (Package)
¿Cómo instalar godoc?
$go get -u golang.org/x/tools/cmd/godoc
¿Cómo ver la documentación de un archivo en Go?
$ godoc archivo.go


4. Type Inference

5. Type Conversion
Type Casting or Type Conversion in Golang
The expression T(v) converts the value v to the type T.

Some numeric conversions:

var i int = 42
var f float64 = float64(i)
var u uint = uint(f)
Or, put more simply:

i := 42
f := float64(i)
u := uint(f)

6. Slices

7. Maps

8. Make

9. Structus

10. For Loop

- Golang - traditional for Statement
- for loop with condition
- for loop with range
- for loop with range and ignore index
- Infinite loop


11. Range

- Basic for-each loop (slice or array)
- String iteration: runes or bytes
- loop over individual bytes
- Map iteration: keys and values
- Channel iteration
//******Gotchas********
- Unexpected values in range loop
- Can’t change entries in range loop

12. Errors
- Constructing Errors
- Type Errors
- Defining Expected Errors
- Wrapping Errors

13. Panic and Recover


# Going Deeper

1. Types and type assertions

proporciona acceso al valor concreto subyacente de un valor de interfaz.
Básicamente, se utiliza para eliminar la ambigüedad de las variables de la interfaz.

2. Interfaces



3. Context

 Un contexto incluye plazos, señales de cancelación y otros valores del ámbito de la solicitud a través de los límites y las rutinas de la API.

 cd 

 go run main.go

 abrir el navegador


4. Goroutines

syntax:

// Anonymous function call
go func (parameter_list){
// statement
}(arguments)


 wg -> WaitGroup
 wg is used to wait for the program to finish.

5. Channels

6.  Buffer

7. Select

8. Mutex

9. Working with json

10. Go Modules

# ORMs

1. GORM

# Loggin

1. Zerolog

2. Zap

# Realtime Communication

1. Melody

2. Centrifugo

# Rest Clients

1. Heimdall

2. GRequest

# GrapghQL Clients
1. Graphql-go

# Microservices










