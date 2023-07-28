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

---- Locking
Two other useful atomic functions are LoadInt64 and StoreInt64. These functions provide a safe way to read and write to an integer value. Here’s an example using LoadInt64 and StoreInt64 to create a synchronous flag that can alert multiple goroutines of a special condition in a program.

--Mutex

Another way to synchronize access to a shared resource is by using a mutex. A mutex is named after the concept of mutual exclusion. A mutex is used to create a critical section around code that ensures only one goroutine at a time can execute that code section. 


A *mutex* is a synchronization object. You acquire a lock on a mutex at the beginning of a section of code, and release it at the end, in order to ensure that no other thread is accessing the same data at the same time. A mutex typically has a lifetime equal to that of the data it is protecting, and that one mutex is accessed by multiple threads.

A *lock* object is an object that encapsulates that lock. When the object is constructed it acquires the lock on the mutex. When it is destructed the lock is released. You typically create a new lock object for every access to the shared data.

5. Channels

channels that synchronize goroutines as they send and receive the resources they need to share between each other.

When declaring a channel, the type of data that will be shared needs to be specified. Values and pointers of built-in, named, struct, and reference types can be shared through a channel.

Each channel has a type associated with it. This type is the type of data that the channel is allowed to transport. No other type is allowed to be transported using the channel.

data := <- a // read from channel a  
a <- data // write to channel a  

https://golangbot.com/channels/

Sends and receives are blocking by default
Sends and receives to a channel are blocking by default. What does this mean? When data is sent to a channel, the control is blocked in the send statement until some other Goroutine reads from that channel. Similarly, when data is read from a channel, the read is blocked until some Goroutine writes data to that channel.


6.  Buffer

7. Select

The select statement lets a goroutine wait on multiple communication operations.

A select blocks until one of its cases can run, then it executes that case. It chooses one at random if multiple are ready.

8. Mutex
https://www.sohamkamani.com/golang/mutex/
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










