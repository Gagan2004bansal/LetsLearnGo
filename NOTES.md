# 🐹 Let's Learn Go

> A personal Go learning handbook for backend development and interview preparation

```
╔══════════════════════════════════════════════════════════════════════════════╗
║                                                                              ║
║    🐹 Welcome! This handbook documents my journey learning Go                ║
║      for backend development and production-oriented systems.                ║
║                                                                              ║
║   Not generic tutorials—real concepts I've studied deeply                    ║
║   with practical examples, backend connections, and interview Q&As.          ║
║                                                                              ║
╚══════════════════════════════════════════════════════════════════════════════╝
```

---

## 📚 About This Handbook

**What I'm Learning:** Go (Golang) — a systems programming language designed for clarity, concurrency, and backend development.

**Why Go for Backend?**
- ✅ Fast compilation → fast deployments
- ✅ Built-in concurrency → handle thousands of goroutines simultaneously
- ✅ Simple syntax → easier to write reliable code
- ✅ Static typing → catch bugs at compile time
- ✅ Standard library → rich tools for HTTP, testing, logging

**My Learning Goals:**
1. Master Go fundamentals deeply
2. Build production-oriented backend systems (HTTP servers, APIs, databases)
3. Understand concurrency patterns for real workloads
4. Ace Go technical interviews
5. Write clean, idiomatic Go code

**Backend Focus:** Everything connects to building real services—HTTP handlers, service layers, database integration, concurrency, and deployment.

---

## 📖 Table of Contents

### Fundamentals (🟢 Complete)
- [Go Installation & Project Setup](#go-installation--project-setup)
- [Variables, Constants & Types](#variables-constants--types)
- [Control Flow](#control-flow)
- [Functions](#functions)
- [Data Structures: Arrays, Slices & Maps](#data-structures-arrays-slices--maps)
- [Structs](#structs)
- [Pointers](#pointers)

### Core Concepts (🟢 Complete)
- [Methods & Receivers](#methods--receivers)
- [Interfaces](#interfaces)
- [Error Handling](#error-handling)
- [Packages & Project Organization](#packages--project-organization)

### Backend Essentials (🟡 Learning)
- [HTTP Servers & REST APIs](#http-servers--rest-apis)
- [Handlers & Middleware](#handlers--middleware)
- [Layered Architecture](#layered-architecture)
- [Configuration Management](#configuration-management)

### Concurrency & Scalability (🟡 Learning)
- [Goroutines](#goroutines)
- [Channels](#channels)
- [Synchronization](#synchronization)
- [Context](#context)

### Testing & Quality (🟡 Learning)
- [Testing Fundamentals](#testing-fundamentals)

### Next Steps (🔴 Coming Soon)
- [What To Learn Next](#-what-to-learn-next)

### Roadmap & Practice
- [Go Mental Model](#-go-mental-model)
- [Go Backend Roadmap](#-go-backend-roadmap)
- [Practice Projects](#-practice-projects)

---

<div align="center">

## 🟢 FUNDAMENTALS

*The foundation everything else builds on.*

</div>

---

## Go Installation & Project Setup

### What is it?

Setting up Go environment, managing code with modules, and understanding the build system (`go run`, `go build`).

### Why does Go have it?

Early versions of other languages used complex dependency systems (npm, pip, maven). Go wanted to be **fast to compile** and **simple to manage dependencies**. Go modules solve the "dependency hell" problem with a single `go.mod` file that locks versions globally.

### Syntax / Example

```go
// Initialize a new Go module project
// $ go mod init github.com/username/myapp

// go.mod (version-locked dependency file)
module github.com/username/myapp

go 1.22

require (
    github.com/joho/godotenv v1.5.1
)

// Run code directly (compiles + executes in memory)
// $ go run main.go

// Build to binary
// $ go build -o myapp
// $ ./myapp

// Add a dependency
// $ go get github.com/user/package@v1.0.0
```

### How it works

1. **`go.mod`** = manifest file listing all dependencies + Go version
2. **`go.sum`** = checksum file ensuring dependency integrity (don't edit manually)
3. **`go run`** = compile + execute in one step (development mode)
4. **`go build`** = create a standalone binary (production deployment)
5. **`go get`** = download + add dependencies automatically

### Common mistakes

- ⚠️ **Committing `go.sum`** → Don't ignore it! Commit it. It ensures reproducible builds.
- ⚠️ **Importing packages before `go get`** → Will get "package not found" error.
- ⚠️ **Wrong module path** → Use your GitHub path: `github.com/yourname/project`

### Backend relevance

In production:
- You deploy with `go build` → creates a single binary (no runtime needed, unlike Node/Python)
- Dependencies are versioned in `go.mod` → team consistency, no version conflicts
- Docker build: `RUN go build -o app` → super lightweight container

### Interview questions

**Q: What's the difference between `go run` and `go build`?**
A: `go run` compiles to memory and executes immediately (dev only). `go build` creates a persistent binary file (production deployment).

**Q: Why is `go.sum` important?**
A: It locks exact dependency versions. Without it, the same dependency version could download different code later (security risk).

### Quick revision

-  `go mod init` initializes a project
-  `go.mod` locks dependency versions across your team
-  `go run` for development, `go build` for deployment
-  Always commit `go.sum` to version control
-  `go get package@version` adds dependencies

**[↑ Back to Top](#-lets-learn-go)**

---

## Variables, Constants & Types

### What is it?

Go has explicit type declarations. Every variable has a type (`int`, `string`, `bool`, etc.), and you must know the type at compile time.

### Why does Go have it?

🤔 Trade-off: Explicit types seem verbose but catch **type errors at compile time**, not at runtime. A Python script that mixes strings and numbers might crash after deployment. Go catches it before.

### Syntax / Example

```go
package main

import "fmt"

func main() {
    // Explicit type declaration
    var name string = "Gagan"
    var age int = 21
    var height float64 = 5.9
    var isProgrammer bool = true

    // Short declaration (type inferred)
    city := "Chandigarh"
    gpa := 9.41

    // Constants (immutable)
    const PI = 3.14159
    const MaxConnections = 1000

    // Zero values (default when not initialized)
    var unsetString string        // ""
    var unsetInt int              // 0
    var unsetFloat float64        // 0.0
    var unsetBool bool            // false

    // Type conversion
    var intAge int = 21
    var floatAge float64 = float64(intAge)

    fmt.Println(name, age, city)
}
```

### How it works

1. **Variables** = mutable storage with explicit type
2. **Constants** = immutable values defined at compile time
3. **Type inference** = `:=` operator lets Go figure out the type
4. **Zero values** = sensible defaults when you don't initialize
5. **Type conversion** = explicit casting required (no automatic upcasting)

### Common mistakes

- ⚠️ **Reassigning with `var` inside a scope** → Use `:=` for new variables, `=` for existing ones
- ⚠️ **Automatic type conversion** → Go doesn't auto-convert. `int + float64` is a compile error.
- ⚠️ **Forgetting type in function return** → Must specify return type explicitly

### Backend relevance

Type safety prevents bugs in production:
- Configuration parsing always knows the type (no runtime surprises)
- Database column types match Go struct types (caught at compile time)
- API responses are type-safe (no `undefined` crashes)

### Interview questions

**Q: What's the difference between `var` and `:=`?**
A: `var` explicitly declares a variable (can be used at package scope). `:=` is short declaration (only inside functions, infers type).

**Q: Can Go automatically convert types?**
A: No. `int` and `float64` don't mix automatically. You must explicitly cast: `float64(myInt)`.

### Quick revision

-  Go is **statically typed** — type checked at compile time
-  Use `:=` for new variables inside functions
-  Constants are immutable and compile-time safe
-  Zero values provide safe defaults
-  Type conversion is explicit

**[↑ Back to Top](#-lets-learn-go)**

---

## Control Flow

### What is it?

`if/else`, `switch`, and `for` loops — how Go makes decisions and repeats code.

### Why does Go have it?

Every programming language needs control flow. Go's twist: **no `while` loop** (only `for`), and **`switch` without `fallthrough` by default**.

### Syntax / Example

```go
package main

import "fmt"

func main() {
    // if/else
    age := 21
    if age >= 18 {
        fmt.Println("Adult")
    } else {
        fmt.Println("Minor")
    }

    // switch (no fallthrough by default)
    role := "admin"
    switch role {
    case "admin":
        fmt.Println("Full access")
    case "user":
        fmt.Println("Read access")
    default:
        fmt.Println("No access")
    }

    // for loop (traditional)
    for i := 0; i < 5; i++ {
        fmt.Println(i)
    }

    // for loop as while
    count := 0
    for count < 5 {
        fmt.Println(count)
        count++
    }

    // infinite loop (with break)
    for {
        fmt.Println("Running...")
        break
    }

    // range over slice
    names := []string{"Alice", "Bob", "Charlie"}
    for index, name := range names {
        fmt.Println(index, name)
    }
}
```

### How it works

1. **`if/else`** = conditional branching (no parentheses required)
2. **`switch`** = multi-way branching (no automatic fallthrough)
3. **`for`** = loop with init/condition/increment OR while-style
4. **`range`** = iterate over collections (returns index + value for slices/maps)
5. **`break`/`continue`** = exit or skip loop iteration

### Common mistakes

- ⚠️ **Expecting fallthrough in `switch`** → Go doesn't auto-fallthrough. Use explicit `fallthrough` keyword if needed.
- ⚠️ **Parentheses in `if` conditions** → Not needed! `if age > 18` not `if (age > 18)`
- ⚠️ **Ignoring one value from `range`** → Use `_` blank identifier: `for _, name := range names`

### Backend relevance

In backend servers:
- Request routing: `switch req.Method { case "GET": ... case "POST": ... }`
- Conditional error handling: `if err != nil { return err }`
- Iterating over database rows: `for rows.Next() { ... }`

### Interview questions

**Q: Why doesn't Go have a `while` loop?**
A: Design philosophy: "one way to do things." `for condition { }` replaces `while`. Reduces complexity, easier to learn.

**Q: What's the difference between `range` on a slice vs a map?**
A: Slice: `for index, value := range slice`. Map: `for key, value := range map`.

### Quick revision

-  No parentheses needed in `if/else` conditions
-  `switch` doesn't auto-fallthrough (good for avoiding bugs)
-  `for` is the only loop construct (used for all looping)
-  `range` iterates with index + value (use `_` to ignore)
-  `break` and `continue` work inside loops

**[↑ Back to Top](#-lets-learn-go)**

---

## Functions

### What is it?

Reusable blocks of code. Functions are first-class citizens in Go—they can be passed around, returned, and stored in variables.

### Why does Go have it?

Like every language, functions organize code and enable reuse. Go's special feature: **multiple return values** (e.g., `value, error`) without tuples or exception handling.

### Syntax / Example

```go
package main

import "fmt"

// Basic function
func greet(name string) {
    fmt.Println("Hello, " + name)
}

// Function with return value
func add(a int, b int) int {
    return a + b
}

// Multiple return values (Go pattern!)
func divide(a float64, b float64) (float64, error) {
    if b == 0 {
        return 0, errors.New("division by zero")
    }
    return a / b, nil
}

// Named return values
func getCoordinates() (x int, y int) {
    x = 10
    y = 20
    return  // returns x, y automatically
}

// Variadic function (variable arguments)
func sum(numbers ...int) int {
    total := 0
    for _, num := range numbers {
        total += num
    }
    return total
}

// Anonymous function / closure
func main() {
    greet("Gagan")
    
    result := add(5, 3)
    fmt.Println(result)

    ans, err := divide(10, 2)
    if err != nil {
        fmt.Println("Error:", err)
    } else {
        fmt.Println("Result:", ans)
    }

    total := sum(1, 2, 3, 4, 5)  // Pass multiple args
    fmt.Println(total)

    // Anonymous function
    multiply := func(a, b int) int {
        return a * b
    }
    fmt.Println(multiply(3, 4))

    // Closure (inner function accessing outer scope)
    counter := 0
    increment := func() {
        counter++
    }
    increment()
    fmt.Println(counter)  // 1

    // defer (runs at end of function)
    defer fmt.Println("Done!")
    fmt.Println("Starting...")
    // Output: Starting... Done!
}
```

### How it works

1. **Function declaration** = `func name(params) returnType { body }`
2. **Multiple returns** = `(type1, type2)` — Go standard for error handling
3. **Named returns** = return values get names, auto-returned at end
4. **Variadic** = `...type` accepts any number of arguments
5. **Anonymous functions** = `func() { ... }` defined inline
6. **Closures** = inner functions can access outer scope
7. **`defer`** = queues function call to run when enclosing function exits (useful for cleanup)

### Common mistakes

- ⚠️ **Ignoring error returns** → `val, err := someFunc(); val + 1` will panic if `err != nil`. Always check!
- ⚠️ **Using named returns everywhere** → Keep it simple; name returns only when it clarifies intent.
- ⚠️ **Forgetting to call `defer` functions** → They run automatically; don't add `()` at definition.

### Backend relevance

Backend patterns:
```go
// Handler function signature
func handleUser(w http.ResponseWriter, r *http.Request) error {
    user, err := fetchUser(r.Context(), userID)
    if err != nil {
        return err  // Multiple return for error handling
    }
    return json.NewEncoder(w).Encode(user)
}

// Service layer returns (value, error)
func (s *UserService) GetUser(ctx context.Context, id string) (*User, error) {
    user, err := s.repo.FindByID(ctx, id)
    if err != nil {
        return nil, fmt.Errorf("failed to fetch user: %w", err)
    }
    return user, nil
}

// defer for cleanup
func (s *DB) Query(sql string) (*Rows, error) {
    conn, err := s.getConnection()
    if err != nil {
        return nil, err
    }
    defer conn.Close()  // Ensures connection closes
    
    return conn.Query(sql)
}
```

### Interview questions

**Q: Why does Go use multiple return values instead of exceptions?**
A: Explicit error handling → errors are part of the function contract. Caller must handle them. No surprise exceptions that crash code.

**Q: What does `defer` do?**
A: Schedules a function to run when its enclosing function returns. Used for cleanup (closing files, connections, releasing locks).

**Q: What's the difference between named and unnamed return values?**
A: Named returns let you assign values in the function body and `return` without arguments. Unnamed requires explicit values: `return value, error`.

### Quick revision

-  Functions are first-class values (can pass them around)
-  **Multiple return values** = Go's answer to error handling
-  `defer` ensures cleanup code runs (even if function panics)
-  Variadic `...type` accepts unlimited arguments
-  Closures can access and modify outer scope

**[↑ Back to Top](#-lets-learn-go)**

---

## Data Structures: Arrays, Slices & Maps

### What is it?

- **Arrays** = fixed-size collections (size part of type)
- **Slices** = dynamic arrays (grow/shrink at runtime) ✨ *Most common*
- **Maps** = key-value stores (like dictionaries/hash tables)

### Why does Go have it?

Arrays are too rigid (fixed size at compile time). Slices are flexible (dynamic). Maps enable fast lookups by key. Together they cover 90% of Go data structure needs.

### Syntax / Example

```go
package main

import "fmt"

func main() {
    // ARRAYS (fixed size, must know size at compile time)
    var scores [3]int = [3]int{85, 90, 78}
    var names [2]string = [2]string{"Alice", "Bob"}
    
    scores[0] = 95  // Modify
    fmt.Println(scores)  // [95 90 78]

    // SLICES (dynamic, flexible)
    var numbers []int = []int{1, 2, 3}
    var words []string = []string{"hello", "world"}

    // Slice declaration (no size!)
    emptySlice := []int{}
    
    // make() to create slice with capacity
    slice := make([]int, 5)          // len=5, cap=5
    slice := make([]int, 5, 10)      // len=5, cap=10

    // append (grows slice)
    numbers = append(numbers, 4, 5)
    fmt.Println(numbers)  // [1 2 3 4 5]

    // Slice from slice (start:end)
    subset := numbers[1:3]  // [2 3]

    // len() and cap()
    fmt.Println(len(numbers))  // 5
    fmt.Println(cap(numbers))  // 10

    // MAPS (key-value)
    var users map[string]int = map[string]int{
        "alice": 25,
        "bob":   30,
    }

    // or with make()
    config := make(map[string]string)

    // Set and Get
    users["charlie"] = 28
    age := users["alice"]  // 25
    
    // Check if key exists
    age, exists := users["david"]  // age=0, exists=false
    if exists {
        fmt.Println("David's age:", age)
    }

    // Delete from map
    delete(users, "bob")

    // Iterate over map
    for name, age := range users {
        fmt.Println(name, age)
    }

    // Iterate over slice
    for i, name := range words {
        fmt.Println(i, name)
    }
}
```

### How it works

**Slices** are dynamic arrays backed by an underlying array:
- **len** = number of elements
- **cap** = capacity of underlying array
- When you `append` and exceed capacity, Go allocates new (larger) array

**Maps** are hash tables:
- O(1) average lookup by key
- Unordered (range order not guaranteed)
- Can safely check key existence: `val, ok := map[key]`

### Common mistakes

- ⚠️ **Confusing array and slice** → Arrays have `[n]Type`, slices have `[]Type`
- ⚠️ **Map iteration order** → Maps are intentionally randomized. Don't rely on order.
- ⚠️ **Nil vs empty slice** → `var s []int` is nil (no allocation), `[]int{}` is empty but allocated
- ⚠️ **Out-of-bounds panic** → Accessing `slice[10]` when len=5 will crash. Use bounds checking.

### Backend relevance

Database/API patterns:
```go
// Fetch multiple users (slice)
users, err := userService.GetAllUsers(ctx)
for _, user := range users {
    // Process each user
}

// Configuration cache (map)
configCache := make(map[string]interface{})
configCache["db_host"] = "localhost"
configCache["db_port"] = 5432

// Parse JSON array into slice
var orders []Order
json.Unmarshal(data, &orders)

// Group by ID (map)
usersById := make(map[string]*User)
for _, user := range users {
    usersById[user.ID] = &user
}
```

### Interview questions

**Q: What's the difference between a slice and an array?**
A: Array has fixed size `[5]int`, slice is dynamic `[]int`. Slices are more common in Go because they're flexible.

**Q: What happens when you `append` to a slice at capacity?**
A: Go allocates a new underlying array (usually ~2x capacity), copies data, returns new slice.

**Q: Why is map iteration order randomized?**
A: Intentional design. Prevents developers from relying on order (which isn't guaranteed in hash tables).

### Quick revision

-  **Arrays** = fixed size `[n]Type`, rarely used directly
-  **Slices** = dynamic `[]Type`, backed by arrays, most common
-  **Maps** = key-value `map[KeyType]ValueType`, fast lookups
-  `append()` grows slices; may reallocate underlying array
-  Check map key existence: `val, ok := myMap[key]`

**[↑ Back to Top](#-lets-learn-go)**

---

## Structs

### What is it?

Structs are custom data types that group related fields together. Think "templates" for objects.

### Why does Go have it?

Go is not object-oriented (no inheritance, no classes). Structs + methods replace objects. They let you model real-world concepts (User, Order, Config) as data.

### Syntax / Example

```go
package main

import "fmt"

// Define a struct
type User struct {
    ID       int
    Name     string
    Email    string
    Age      int
    IsActive bool
}

// Nested struct
type Address struct {
    Street string
    City   string
    Zip    string
}

type Person struct {
    Name    string
    Address Address  // nested
}

// Struct tags (metadata for JSON, database)
type Product struct {
    ID    int     `json:"id"`
    Name  string  `json:"name"`
    Price float64 `json:"price"`
}

func main() {
    // Create struct with all fields
    user1 := User{
        ID:       1,
        Name:     "Gagan",
        Email:    "gagan@example.com",
        Age:      21,
        IsActive: true,
    }

    // Positional (not recommended, less readable)
    user2 := User{1, "Alice", "alice@ex.com", 25, false}

    // Zero values
    var user3 User  // ID=0, Name="", Age=0, IsActive=false

    // Access fields
    fmt.Println(user1.Name)
    user1.Age = 22

    // Nested struct
    person := Person{
        Name: "Bob",
        Address: Address{
            Street: "123 Main St",
            City:   "Chandigarh",
            Zip:    "160001",
        },
    }
    fmt.Println(person.Address.City)

    // JSON encoding (uses tags)
    data, _ := json.Marshal(product)
    fmt.Println(string(data))
    // Output: {"id":1,"name":"Gopher","price":9.99}
}
```

### How it works

1. **Definition** = `type StructName struct { fields }`
2. **Fields** = named and typed
3. **Instantiation** = `StructName{ field1: val1, field2: val2 }`
4. **Zero values** = empty struct has all fields at zero
5. **Tags** = metadata (backticks) for JSON marshaling, database mapping, etc.

### Common mistakes

- ⚠️ **Exported vs unexported** → `ID` (capital) is exported (public), `id` (lowercase) is unexported (private to package)
- ⚠️ **Tags ignored on lowercase fields** → Struct tags only work on exported fields
- ⚠️ **Modifying struct in function** → Structs are copied by value. Use pointers to modify!

### Backend relevance

API request/response modeling:
```go
// Request model
type CreateUserRequest struct {
    Name  string `json:"name"`
    Email string `json:"email"`
}

// Response model
type UserResponse struct {
    ID        int       `json:"id"`
    Name      string    `json:"name"`
    Email     string    `json:"email"`
    CreatedAt time.Time `json:"created_at"`
}

// Database model
type User struct {
    ID        int       `db:"id"`
    Name      string    `db:"name"`
    Email     string    `db:"email"`
    CreatedAt time.Time `db:"created_at"`
}

// Handler receives struct, returns struct
func (h *Handler) CreateUser(w http.ResponseWriter, r *http.Request) {
    var req CreateUserRequest
    json.NewDecoder(r.Body).Decode(&req)
    
    user := h.service.CreateUser(req.Name, req.Email)
    
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(UserResponse{...})
}
```

### Interview questions

**Q: What's the difference between exported and unexported struct fields?**
A: Exported (capital letter) visible outside package. Unexported (lowercase) private to package only.

**Q: What are struct tags used for?**
A: Metadata for serialization. `json:"field_name"` tells JSON marshaler how to encode/decode. Also used for database mapping.

### Quick revision

-  Structs group related fields (like lightweight objects)
-  Exported fields: capital letter `ID`, unexported: lowercase `id`
-  Tags provide metadata for JSON/database mapping
-  Structs copied by value (pass pointers to modify)
-  Nested structs enable composition

**[↑ Back to Top](#-lets-learn-go)**

---

## Pointers

### What is it?

A pointer is a variable that holds the **memory address** of another variable. When you need to modify data in a function, you pass a pointer, not a copy.

### Why does Go have it?

Performance and intent. Passing large structs by value copies all data. Pointers pass 8 bytes (the address). Also enables actual modification of original data.

### Syntax / Example

```go
package main

import "fmt"

func main() {
    // Declare a variable
    name := "Gagan"
    age := 21

    // Get address with &
    var namePtr *string = &name      // pointer to string
    var agePtr *int = &age           // pointer to int

    // Dereference with *
    fmt.Println(*namePtr)  // "Gagan" (dereference to get value)
    fmt.Println(*agePtr)   // 21

    // Modify through pointer
    *agePtr = 22
    fmt.Println(age)  // 22 (original changed!)

    // nil pointer (zero value)
    var emptyPtr *string  // nil (not pointing to anything)
    if emptyPtr == nil {
        fmt.Println("Pointer is nil")
    }

    // Create pointer to struct
    user := User{ID: 1, Name: "Alice"}
    userPtr := &user

    // Access struct fields through pointer (Go auto-dereferences)
    fmt.Println(userPtr.Name)  // "Alice" (not (*userPtr).Name)
    userPtr.Age = 25

    // Function that modifies
    modifyUser(userPtr)
    fmt.Println(user.Age)  // Changed!
}

// Function receives pointer
func modifyUser(u *User) {
    u.Age = 30
}

// vs. value receiver (no modification)
func printUser(u User) {
    fmt.Println(u.Name)  // Just read, can't modify original
}
```

### How it works

1. **`&` operator** = get address of a variable
2. **`*Type`** = pointer type (points to a value of `Type`)
3. **`*` operator** = dereference (get value at address)
4. **nil pointer** = points to nothing (zero value)
5. **Struct fields through pointers** = Go auto-dereferences (`ptr.Field` not `(*ptr).Field`)

### Common mistakes

- ⚠️ **Dereferencing nil pointers** → `var ptr *int; *ptr = 5` will panic. Check `if ptr == nil` first.
- ⚠️ **Confusing `&` and `*`** → `&x` gives address, `*ptr` gives value at address
- ⚠️ **Copying vs modifying** → Pass pointer to modify, pass value to just read
- ⚠️ **Pointer to loop variable** → Captures address, not value! Use caution in loops.

### Backend relevance

Critical for backend patterns:
```go
// Service layer passes pointers to avoid copying large structs
func (s *UserService) GetUser(ctx context.Context, id string) (*User, error) {
    user := &User{}  // Allocate pointer
    // ... fetch from database ...
    return user, nil  // Return pointer
}

// Handler receives pointer
func (h *Handler) GetUserHandler(w http.ResponseWriter, r *http.Request) {
    user, err := h.service.GetUser(r.Context(), userID)
    if err != nil {
        // handle
    }
    // Modify without copying
    user.LastAccessedAt = time.Now()
    h.service.Save(r.Context(), user)
}

// Middleware chain often uses pointers
type Middleware func(http.Handler) http.Handler

func LoggingMiddleware(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        // r is a *Request (pointer)
        next.ServeHTTP(w, r)
    })
}
```

### Interview questions

**Q: When should you use a pointer vs a value?**
A: Use pointer if function needs to modify data OR if struct is large (avoid copying). Use value for small immutable data.

**Q: What happens if you dereference a nil pointer?**
A: Runtime panic! Always check `if ptr == nil` before dereferencing.

**Q: Can you take a pointer to a pointer?**
A: Yes, `**int` is a pointer to a pointer. Rarely used in Go (complex, confusing).

### Quick revision

-  `&variable` = get address (create pointer)
-  `*pointerType` = dereference (get value)
-  Pass pointers to modify data or avoid large copies
-  nil pointer = zero value, dereferencing it crashes
-  Go auto-dereferences struct fields: `ptr.Field` ≈ `(*ptr).Field`

**[↑ Back to Top](#-lets-learn-go)**

---

<div align="center">

## 🟢 CORE CONCEPTS

*Beyond basics—how Go's type system works.*

</div>

---

## Methods & Receivers

### What is it?

Methods are functions attached to a type (struct). The "receiver" is the type that the method belongs to. Go's way of adding behavior to data without inheritance.

### Why does Go have it?

Go isn't OOP but still needs to attach functions to types. Methods + interfaces = a simpler, more explicit type system than traditional OOP.

### Syntax / Example

```go
package main

import "fmt"

type User struct {
    ID   int
    Name string
    Age  int
}

// VALUE RECEIVER (receives copy of User)
func (u User) Greet() {
    fmt.Printf("Hello, I'm %s and I'm %d years old\n", u.Name, u.Age)
}

func (u User) IsAdult() bool {
    return u.Age >= 18
}

// POINTER RECEIVER (receives address of User)
func (u *User) BirthdayUp() {
    u.Age++  // Modifies original
}

func (u *User) UpdateEmail(email string) {
    // This is a setter (modifies state)
    u.Name = email
}

// Receiver can be any type (not just struct!)
type Counter int

func (c *Counter) Increment() {
    *c++  // Modify the underlying int
}

func (c Counter) Value() int {
    return int(c)
}

func main() {
    user := User{ID: 1, Name: "Alice", Age: 25}

    // Call value receiver method
    user.Greet()           // Works
    fmt.Println(user.IsAdult())  // true

    // Call pointer receiver method (Go auto-takes address)
    user.BirthdayUp()      // Same as (&user).BirthdayUp()
    fmt.Println(user.Age)  // 26

    // Call through pointer explicitly
    userPtr := &user
    userPtr.BirthdayUp()
    fmt.Println(user.Age)  // 27

    // Counter example
    var count Counter = 0
    count.Increment()
    fmt.Println(count.Value())  // 1
}
```

### How it works

1. **Value receiver** = `(u User)` receives a copy. Modifications don't affect original.
2. **Pointer receiver** = `(u *User)` receives address. Modifications affect original.
3. **Method call** = Go auto-dereferences/addresses as needed.
4. **By convention** = use pointer receivers unless receiver is small and immutable

### Common mistakes

- ⚠️ **Mixing receiver types** → If you use pointer receiver for modification, always use pointer receiver consistently
- ⚠️ **Value receiver forgetting it's a copy** → Modifying a value receiver doesn't affect original
- ⚠️ **Overcomplicating with too many setters** → Go prefers simple assignment, not getter/setter pattern

### Backend relevance

Service layer pattern:
```go
// Service with pointer receiver (stateful)
type UserService struct {
    db *Database
    cache *Cache
}

func (s *UserService) GetUser(ctx context.Context, id string) (*User, error) {
    // Can access s.db, s.cache
    user, _ := s.cache.Get(id)
    if user == nil {
        user, _ = s.db.FindUser(id)
    }
    return user, nil
}

// Handler layer
type UserHandler struct {
    service *UserService  // Injected dependency
}

func (h *UserHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
    user, err := h.service.GetUser(r.Context(), userID)
    // ...
}
```

### Interview questions

**Q: What's the difference between value and pointer receivers?**
A: Value receiver = copy (modifications lost). Pointer receiver = address (modifications affect original).

**Q: When should you use pointer receivers?**
A: When the method needs to modify the receiver OR the receiver is large (avoid copying).

### Quick revision

-  Methods attach functions to types: `func (receiver Type) MethodName()`
-  Value receiver = copy, pointer receiver = address
-  Pointer receivers modify original, value receivers modify copy
-  Use pointer receivers for consistency and clarity
-  Go auto-dereferences/takes address as needed

**[↑ Back to Top](#-lets-learn-go)**

---

## Interfaces

### What is it?

An interface defines a **contract** of methods. Any type that implements all methods of an interface automatically satisfies the interface (implicit implementation).

### Why does Go have it?

Decoupling! Instead of knowing concrete types, code can work with abstractions. No inheritance needed. It's "if it walks like a duck, quacks like a duck, it's a duck."

### Syntax / Example

```go
package main

import (
    "fmt"
    "io"
)

// Define interface
type Reader interface {
    Read(p []byte) (n int, err error)
}

type Writer interface {
    Write(p []byte) (n int, err error)
}

// Any type with Read() method satisfies Reader
type FileReader struct {
    filename string
}

func (fr *FileReader) Read(p []byte) (int, error) {
    // Implementation
    return len(p), nil
}

type NetworkReader struct {
    addr string
}

func (nr *NetworkReader) Read(p []byte) (int, error) {
    // Implementation
    return len(p), nil
}

// Function that works with ANY Reader
func CopyData(reader Reader, writer Writer) error {
    buffer := make([]byte, 1024)
    n, err := reader.Read(buffer)
    if err != nil {
        return err
    }
    _, err = writer.Write(buffer[:n])
    return err
}

// Empty interface (accepts anything)
type Any interface{}  // or just "any" keyword in Go 1.18+

func Print(v any) {
    fmt.Println(v)
}

// Type assertion
func ProcessValue(v any) {
    // Assert it's an int
    if intVal, ok := v.(int); ok {
        fmt.Println("It's an int:", intVal)
    }
    
    // Assert it's a string
    if strVal, ok := v.(string); ok {
        fmt.Println("It's a string:", strVal)
    }
}

// Type switch
func Handle(v any) {
    switch val := v.(type) {
    case int:
        fmt.Println("Integer:", val)
    case string:
        fmt.Println("String:", val)
    case error:
        fmt.Println("Error:", val)
    default:
        fmt.Println("Unknown type")
    }
}

func main() {
    file := &FileReader{filename: "data.txt"}
    network := &NetworkReader{addr: "127.0.0.1"}

    // Both satisfy Reader interface (implicit!)
    CopyData(file, network)

    Print(42)
    Print("hello")
    ProcessValue(100)
    Handle("test")
}
```

### How it works

1. **Interface definition** = set of methods
2. **Implicit implementation** = any type with all methods automatically satisfies interface
3. **Empty interface** = `interface{}` or `any` = accepts anything
4. **Type assertion** = `value.(Type)` extracts concrete type
5. **Type switch** = `switch v.(type)` pattern matches on type

### Common mistakes

- ⚠️ **Expecting explicit implementation** → Go interfaces are implicit. Easy to accidentally satisfy an interface!
- ⚠️ **Interface with too many methods** → Violates "Interface Segregation Principle." Small, focused interfaces are better.
- ⚠️ **Panicking on type assertion failure** → Always use comma-ok: `val, ok := v.(Type)`, not `val := v.(Type)`

### Backend relevance

Dependency injection pattern:
```go
// Define interface (behavior contract)
type Database interface {
    FindUser(ctx context.Context, id string) (*User, error)
    SaveUser(ctx context.Context, user *User) error
}

type Cache interface {
    Get(key string) (string, bool)
    Set(key string, value string) error
}

// Service works with interfaces (decoupled)
type UserService struct {
    db    Database
    cache Cache
}

func (s *UserService) GetUser(ctx context.Context, id string) (*User, error) {
    // Works with ANY database implementation
    return s.db.FindUser(ctx, id)
}

// Concrete implementations
type PostgresDB struct { /* ... */ }
func (p *PostgresDB) FindUser(ctx context.Context, id string) (*User, error) { /* ... */ }

type MockDB struct { /* ... */ }
func (m *MockDB) FindUser(ctx context.Context, id string) (*User, error) { /* ... */ }

// Easy to swap for testing
func TestGetUser(t *testing.T) {
    mockDB := &MockDB{ /* ... */ }
    service := &UserService{db: mockDB}
    
    user, _ := service.GetUser(context.Background(), "1")
    // Test without touching real database!
}
```

### Interview questions

**Q: What does "implicit interface implementation" mean?**
A: You don't declare `implements UserRepository`. If your type has all the methods, it's automatically an interface. No explicit keyword needed.

**Q: Why use interfaces in backend development?**
A: Decouple code from concrete types. Swap database, cache, logger implementations without changing service layer.

**Q: What's the difference between `interface{}` and a specific interface?**
A: `interface{}` accepts anything (loses type info). Specific interface enforces contract (type-safe).

### Quick revision

-  Interfaces define method contracts (no inheritance)
-  Implementation is implicit (no `implements` keyword)
-  Small, focused interfaces >> large interfaces
-  Type assertion: `val, ok := v.(ConcreteType)`
-  Perfect for dependency injection and testing

**[↑ Back to Top](#-lets-learn-go)**

---

## Error Handling

### What is it?

Go treats errors as values (not exceptions). Functions return `(value, error)`. Caller must explicitly handle the error.

### Why does Go have it?

No try/catch (by design). Explicit error handling → forces you to think about what could go wrong. No surprise exceptions crashing your code.

### Syntax / Example

```go
package main

import (
    "errors"
    "fmt"
    "log"
)

// Create errors
func divide(a, b float64) (float64, error) {
    if b == 0 {
        return 0, errors.New("division by zero")
    }
    return a / b, nil
}

// fmt.Errorf creates error with formatting
func fetchUser(id string) (*User, error) {
    if id == "" {
        return nil, fmt.Errorf("invalid user id: %s", id)
    }
    return &User{ID: id}, nil
}

// Error wrapping (preserves context)
func processData() error {
    result, err := divide(10, 0)
    if err != nil {
        return fmt.Errorf("division failed: %w", err)  // %w wraps
    }
    return nil
}

// Check specific error
func handleRequest() {
    err := processData()
    
    // errors.Is (v1.13+)
    if errors.Is(err, io.EOF) {
        fmt.Println("End of file")
    }

    // errors.As (extract wrapped error)
    var pathErr *os.PathError
    if errors.As(err, &pathErr) {
        fmt.Println("Path error:", pathErr.Path)
    }
}

// Custom error type
type ValidationError struct {
    Field string
    Value string
    Err   string
}

func (ve *ValidationError) Error() string {
    return fmt.Sprintf("validation error: %s=%s: %s", ve.Field, ve.Value, ve.Err)
}

// Any type with Error() string method is an error
func validateEmail(email string) error {
    if !strings.Contains(email, "@") {
        return &ValidationError{
            Field: "email",
            Value: email,
            Err:   "missing @",
        }
    }
    return nil
}

// idiomatic: handle error immediately
func main() {
    result, err := divide(10, 2)
    if err != nil {
        log.Fatal(err)  // Log and exit
    }
    fmt.Println(result)

    user, err := fetchUser("123")
    if err != nil {
        fmt.Println("Error:", err)
        return
    }
    fmt.Println(user)

    err = validateEmail("invalid")
    if err != nil {
        fmt.Println(err)
    }
}
```

### How it works

1. **Multiple returns** = `(value, error)` pattern
2. **Check immediately** = `if err != nil { return err }`
3. **errors.New()** = create simple error
4. **fmt.Errorf()** = error with message formatting
5. **Error wrapping** = `%w` preserves context chain
6. **errors.Is()** = check specific error type
7. **errors.As()** = extract wrapped error
8. **Custom errors** = implement `Error() string` method

### Common mistakes

- ⚠️ **Ignoring errors** → `result, _ := somFunc()` loses info. Always check!
- ⚠️ **Panicking instead of returning** → Return error instead of `panic()` (except unrecoverable)
- ⚠️ **Not wrapping context** → Bare `return err` loses context. Use `fmt.Errorf("operation: %w", err)`
- ⚠️ **Shadowing error in recovery** → `err := recover()` may not be `error` type. Check it.

### Backend relevance

HTTP handler error pattern:
```go
func (h *Handler) GetUser(w http.ResponseWriter, r *http.Request) {
    id := r.URL.Query().Get("id")
    if id == "" {
        http.Error(w, "missing id", http.StatusBadRequest)
        return
    }

    user, err := h.service.GetUser(r.Context(), id)
    if err != nil {
        // Log error for debugging
        h.logger.Error("get user failed", "error", err, "id", id)
        
        // Return 500 to client
        http.Error(w, "internal server error", http.StatusInternalServerError)
        return
    }

    // Success
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(user)
}

// Database layer with error wrapping
func (r *UserRepository) FindByID(ctx context.Context, id string) (*User, error) {
    var user User
    err := r.db.QueryRowContext(ctx, "SELECT * FROM users WHERE id=$1", id).
        Scan(&user.ID, &user.Name)
    if err != nil {
        if err == sql.ErrNoRows {
            return nil, fmt.Errorf("user not found: %w", err)
        }
        return nil, fmt.Errorf("database error: %w", err)
    }
    return &user, nil
}
```

### Interview questions

**Q: Why does Go return errors instead of throwing exceptions?**
A: Explicit error handling forces you to consider failure cases. No surprise exceptions. Errors are part of the function contract.

**Q: What's the difference between `errors.Is()` and `errors.As()`?**
A: `errors.Is()` checks if error equals a specific value. `errors.As()` extracts a wrapped error of a specific type.

**Q: How should you wrap errors?**
A: Use `%w` format verb: `fmt.Errorf("operation failed: %w", err)`. Preserves error chain for debugging.

### Quick revision

-  Errors are values: `(value, error)` return pattern
-  Check `if err != nil` immediately
-  Wrap errors with context: `fmt.Errorf("msg: %w", err)`
-  Custom errors implement `Error() string` method
-  `errors.Is()` / `errors.As()` for error inspection

**[↑ Back to Top](#-lets-learn-go)**

---

## Packages & Project Organization

### What is it?

Go organizes code into packages. A package is a directory with `.go` files. Import packages to use them. The `internal` package prevents external imports.

### Why does Go have it?

Large projects need organization. Packages provide namespacing and control over visibility. The `internal` package is unique—Go prevents imports from outside your module.

### Syntax / Example

```go
// File: main.go (in github.com/username/myapp)
package main

import (
    "fmt"
    
    // Import from go.mod module
    "github.com/username/myapp/internal/service"
    "github.com/username/myapp/internal/repository"
)

func main() {
    repo := repository.NewUserRepository()
    svc := service.NewUserService(repo)
    user, _ := svc.GetUser("1")
    fmt.Println(user)
}

// File: internal/service/user.go
package service  // Same package as other files in same dir

import (
    "github.com/username/myapp/internal/repository"
)

type UserService struct {
    repo repository.UserRepository
}

func NewUserService(repo repository.UserRepository) *UserService {
    return &UserService{repo: repo}
}

// Exported function (capital letter)
func (s *UserService) GetUser(id string) (*User, error) {
    return s.repo.FindByID(id)
}

// Private helper (lowercase)
func (s *UserService) validateID(id string) bool {
    return id != ""
}

// File: internal/repository/user.go
package repository

type UserRepository interface {
    FindByID(id string) (*User, error)
}

type User struct {
    ID   string
    Name string
}

// Concrete implementation
type PostgresUserRepository struct {
    // connection details
}

func (r *PostgresUserRepository) FindByID(id string) (*User, error) {
    // Query database
    return &User{ID: id}, nil
}

// File: internal/config/config.go
package config

type Config struct {
    DBHost     string
    DBPort     int
    HTTPPort   int
}

func Load() *Config {
    return &Config{
        DBHost:   "localhost",
        DBPort:   5432,
        HTTPPort: 8080,
    }
}
```

**Project Structure:**
```
myapp/
├── go.mod
├── go.sum
├── main.go                    # Entry point
├── cmd/
│   └── server/
│       └── main.go            # Alt entry point
├── internal/                  # Private to module (cannot import from outside)
│   ├── service/
│   │   └── user.go
│   ├── repository/
│   │   └── user.go
│   ├── config/
│   │   └── config.go
│   └── model/
│       └── user.go
├── pkg/                       # Public (can import from outside)
│   └── utils/
│       └── helpers.go
└── migrations/
    └── 001_users.sql
```

### How it works

1. **Package = directory** with same package name in all `.go` files
2. **Exported = capital letter** `GetUser()` visible outside package
3. **Unexported = lowercase** `getUser()` private to package
4. **`internal/` package** = Go enforces import restrictions (can't import from outside module)
5. **`cmd/` package** = separate executable entry points
6. **`pkg/` package** = public libraries within module

### Common mistakes

- ⚠️ **Mixing multiple packages in one directory** → All files in directory must have same `package name`
- ⚠️ **Trying to import `internal` from outside** → Go compiler prevents it. Use `pkg/` instead.
- ⚠️ **Circular imports** → Package A imports B, B imports A. Refactor to break cycle.
- ⚠️ **Too many packages** → Don't create a package for one file. Group logically.

### Backend relevance

Standard backend structure:
```go
// cmd/api/main.go - HTTP API server
func main() {
    cfg := config.Load()
    db := database.Connect(cfg)
    
    userRepo := repository.NewUserRepository(db)
    userService := service.NewUserService(userRepo)
    
    router := http.NewServeMux()
    handler := NewHandler(userService)
    handler.RegisterRoutes(router)
    
    server := &http.Server{Addr: ":" + cfg.HTTPPort}
    server.ListenAndServe()
}

// cmd/worker/main.go - Background job processor
func main() {
    cfg := config.Load()
    cache := cache.NewRedis(cfg)
    
    worker := NewWorker(cache)
    worker.Start()
}

// internal/service - Business logic
type UserService struct { /* ... */ }
func (s *UserService) CreateUser(ctx context.Context, req *CreateUserRequest) (*User, error) { /* ... */ }

// internal/repository - Data access
type UserRepository interface { /* ... */ }
```

### Interview questions

**Q: What's the difference between `internal` and `pkg` packages?**
A: `internal/` is private to the module (Go prevents imports from outside). `pkg/` is public library code.

**Q: Can you have multiple packages in one directory?**
A: No. All `.go` files in a directory must have the same `package` name.

### Quick revision

-  Package = directory with same `package name` in all files
-  Exported (capital) vs unexported (lowercase) controls visibility
-  `internal/` = private package (Go enforces, can't import externally)
-  `cmd/` = executable entry points
-  Group related code → fewer, larger packages better than many tiny ones

**[↑ Back to Top](#-lets-learn-go)**

---

<div align="center">

## 🟡 BACKEND ESSENTIALS

*Real-world patterns for building APIs and servers.*

</div>

---

## HTTP Servers & REST APIs

### What is it?

Go's `net/http` package provides HTTP server and client functionality. Build REST APIs by handling HTTP requests and returning JSON responses.

### Why does Go have it?

HTTP is the language of the web. Go's standard library includes everything needed to build HTTP services without external dependencies (though frameworks exist).

### Syntax / Example

```go
package main

import (
    "encoding/json"
    "fmt"
    "log"
    "net/http"
    "strconv"
)

// Model
type User struct {
    ID    int    `json:"id"`
    Name  string `json:"name"`
    Email string `json:"email"`
}

// In-memory "database"
var users = map[int]*User{
    1: {ID: 1, Name: "Alice", Email: "alice@ex.com"},
    2: {ID: 2, Name: "Bob", Email: "bob@ex.com"},
}

// GET /users
func GetUsers(w http.ResponseWriter, r *http.Request) {
    if r.Method != http.MethodGet {
        http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
        return
    }

    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(users)
}

// GET /users/:id
func GetUser(w http.ResponseWriter, r *http.Request) {
    if r.Method != http.MethodGet {
        http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
        return
    }

    // Parse ID from query params or path
    id, _ := strconv.Atoi(r.URL.Query().Get("id"))

    user, exists := users[id]
    if !exists {
        http.Error(w, "Not found", http.StatusNotFound)
        return
    }

    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(user)
}

// POST /users
func CreateUser(w http.ResponseWriter, r *http.Request) {
    if r.Method != http.MethodPost {
        http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
        return
    }

    var req struct {
        Name  string `json:"name"`
        Email string `json:"email"`
    }

    if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
        http.Error(w, "Invalid JSON", http.StatusBadRequest)
        return
    }

    // Validate
    if req.Name == "" || req.Email == "" {
        http.Error(w, "Missing fields", http.StatusBadRequest)
        return
    }

    // Create
    newID := len(users) + 1
    user := &User{ID: newID, Name: req.Name, Email: req.Email}
    users[newID] = user

    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusCreated)
    json.NewEncoder(w).Encode(user)
}

func main() {
    // Register handlers
    http.HandleFunc("/users", GetUsers)
    http.HandleFunc("/user", GetUser)
    http.HandleFunc("/users/create", CreateUser)

    // Start server
    log.Println("Server starting on :8080")
    if err := http.ListenAndServe(":8080", nil); err != nil {
        log.Fatal(err)
    }
}

// Test:
// curl http://localhost:8080/users
// curl http://localhost:8080/user?id=1
// curl -X POST http://localhost:8080/users/create \
//   -H "Content-Type: application/json" \
//   -d '{"name":"Charlie","email":"charlie@ex.com"}'
```

### How it works

1. **`http.HandleFunc(pattern, handler)`** = register handler for route
2. **Handler signature** = `func(w http.ResponseWriter, r *http.Request)`
3. **Request** = `*http.Request` contains method, headers, body, query params
4. **Response** = `http.ResponseWriter` to write headers, status code, body
5. **`json.NewDecoder().Decode()`** = parse JSON from request body
6. **`json.NewEncoder().Encode()`** = serialize struct to JSON response

### Common mistakes

- ⚠️ **Not checking request method** → Always check `r.Method` to handle GET/POST/PUT/DELETE correctly
- ⚠️ **Writing response after setting status** → Set status code before `w.Write()`
- ⚠️ **Forgetting to set Content-Type** → Browsers expect `application/json`
- ⚠️ **Not closing request body** → Use `defer r.Body.Close()` if reading multiple times

### Backend relevance

Real production API:
```go
// handler.go
type UserHandler struct {
    service *UserService
}

func (h *UserHandler) CreateUser(w http.ResponseWriter, r *http.Request) {
    // Parse request
    var req CreateUserRequest
    if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
        http.Error(w, "Invalid JSON", http.StatusBadRequest)
        return
    }

    // Call service
    user, err := h.service.CreateUser(r.Context(), req)
    if err != nil {
        h.handleError(w, err)
        return
    }

    // Return response
    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusCreated)
    json.NewEncoder(w).Encode(user)
}

// main.go
func main() {
    db := database.Connect()
    userService := service.NewUserService(db)
    handler := &UserHandler{service: userService}

    mux := http.NewServeMux()
    mux.HandleFunc("/users", handler.CreateUser)

    server := &http.Server{
        Addr:    ":8080",
        Handler: mux,
    }
    
    log.Fatal(server.ListenAndServe())
}
```

### Interview questions

**Q: What's the signature of an HTTP handler in Go?**
A: `func(w http.ResponseWriter, r *http.Request)`. `w` writes response, `r` contains request data.

**Q: How do you read JSON from a request?**
A: `json.NewDecoder(r.Body).Decode(&struct)`. Always check error.

### Quick revision

-  `http.HandleFunc()` registers handlers for routes
-  Handler signature: `func(w http.ResponseWriter, r *http.Request)`
-  Check `r.Method` for HTTP verb (GET, POST, etc.)
-  `json.NewDecoder()` parses request body
-  Set headers before writing response body

**[↑ Back to Top](#-lets-learn-go)**

---

## Handlers & Middleware

### What is it?

**Handlers** = functions that process HTTP requests. **Middleware** = wraps handlers to add cross-cutting concerns (logging, auth, timeout) without repeating code.

### Why does Go have it?

Every request needs logging, auth checks, CORS headers, etc. Middleware keeps handler logic clean and reusable.

### Syntax / Example

```go
package main

import (
    "fmt"
    "log"
    "net/http"
    "time"
)

// Simple handler
func HelloHandler(w http.ResponseWriter, r *http.Request) {
    w.WriteHeader(http.StatusOK)
    fmt.Fprintf(w, "Hello, %s!\n", r.URL.Query().Get("name"))
}

// Middleware as higher-order function
// Returns a handler that wraps another handler
func LoggingMiddleware(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        start := time.Now()
        
        log.Printf("-> %s %s", r.Method, r.RequestURI)
        
        next.ServeHTTP(w, r)  // Call wrapped handler
        
        duration := time.Since(start)
        log.Printf("<- %s (took %v)", r.Method, duration)
    })
}

// Auth middleware
func AuthMiddleware(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        token := r.Header.Get("Authorization")
        if token == "" {
            http.Error(w, "Unauthorized", http.StatusUnauthorized)
            return
        }

        // In real app, validate token here
        log.Printf("User authorized with token: %s", token)

        next.ServeHTTP(w, r)
    })
}

// Timeout middleware
func TimeoutMiddleware(timeout time.Duration) func(http.Handler) http.Handler {
    return func(next http.Handler) http.Handler {
        return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
            ctx, cancel := context.WithTimeout(r.Context(), timeout)
            defer cancel()

            r = r.WithContext(ctx)
            next.ServeHTTP(w, r)
        })
    }
}

// Chaining middleware (bottom → top execution)
func Chain(h http.Handler, middleware ...func(http.Handler) http.Handler) http.Handler {
    for i := len(middleware) - 1; i >= 0; i-- {
        h = middleware[i](h)
    }
    return h
}

func main() {
    mux := http.NewServeMux()

    // Register handler
    baseHandler := http.HandlerFunc(HelloHandler)

    // Apply middleware chain
    handler := Chain(
        baseHandler,
        LoggingMiddleware,
        AuthMiddleware,
    )

    mux.Handle("/hello", handler)

    log.Println("Server on :8080")
    http.ListenAndServe(":8080", mux)
}

// Test:
// curl -H "Authorization: token123" http://localhost:8080/hello?name=Gagan
// Output: User authorized with token: token123
//         -> GET /hello?name=Gagan
//         <- GET (took 1.234ms)
```

### How it works

1. **Handler** = `func(http.ResponseWriter, *http.Request)` processes request
2. **Middleware** = function taking handler, returning new handler
3. **Chain** = apply multiple middleware in order
4. **Context** = pass data through request (e.g., user info, request ID)
5. **Order matters** = middleware executed in reverse order of registration

```
Request → Auth → Logging → Handler → Response
```

### Common mistakes

- ⚠️ **Forgetting to call `next.ServeHTTP()`** → Middleware must call next handler
- ⚠️ **Modifying request after calling next** → Changes ignored
- ⚠️ **Not using context** → Data between middleware lost
- ⚠️ **Middleware order** → Auth should run before logging (for user context)

### Backend relevance

Real production middleware:
```go
// Logging middleware
func LoggingMiddleware(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        requestID := uuid.New().String()
        
        // Inject request ID into context
        ctx := context.WithValue(r.Context(), "requestID", requestID)
        r = r.WithContext(ctx)

        log.WithFields(log.Fields{
            "request_id": requestID,
            "method":     r.Method,
            "path":       r.URL.Path,
        }).Info("request started")

        next.ServeHTTP(w, r)
    })
}

// Auth middleware (JWT validation)
func JWTMiddleware(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        token := r.Header.Get("Authorization")
        claims, err := validateToken(token)
        if err != nil {
            http.Error(w, "Unauthorized", http.StatusUnauthorized)
            return
        }

        // Pass user to handler via context
        ctx := context.WithValue(r.Context(), "user", claims.UserID)
        r = r.WithContext(ctx)

        next.ServeHTTP(w, r)
    })
}

// Use in handler
func GetUserHandler(w http.ResponseWriter, r *http.Request) {
    userID := r.Context().Value("user").(string)
    requestID := r.Context().Value("requestID").(string)
    
    log.Infof("[%s] Fetching user %s", requestID, userID)
    // ...
}
```

### Interview questions

**Q: How do you pass data between middleware?**
A: Use `context.WithValue()` to add values to request context, retrieve with `r.Context().Value()`.

**Q: In what order do middleware execute?**
A: In reverse order. Last middleware registered executes first (closest to handler).

### Quick revision

-  Middleware = handler wrapper for cross-cutting concerns
-  Signature: `func(http.Handler) http.Handler`
-  Chain middleware to apply multiple in order
-  Use `context.WithValue()` to pass data
-  Call `next.ServeHTTP()` to invoke wrapped handler

**[↑ Back to Top](#-lets-learn-go)**

---

## Layered Architecture

### What is it?

Organizing backend code into layers:
1. **Handler** (HTTP) → Request/Response
2. **Service** (Business Logic) → Rules, calculations, validation
3. **Repository** (Data Access) → Database queries

Each layer talks to the one below. Loose coupling, easy testing.

### Why does Go have it?

Separates concerns. Handler doesn't know SQL. Service doesn't know HTTP. Easy to test (mock repository), swap implementations (SQL → NoSQL), maintain code.

### Syntax / Example

```go
// model.go
package internal

type User struct {
    ID    string
    Name  string
    Email string
}

// repository.go (data access layer)
package repository

import "context"

type UserRepository interface {
    FindByID(ctx context.Context, id string) (*User, error)
    Save(ctx context.Context, user *User) error
    Delete(ctx context.Context, id string) error
}

// Concrete implementation (could swap with another)
type PostgresUserRepository struct {
    db *sql.DB
}

func (r *PostgresUserRepository) FindByID(ctx context.Context, id string) (*User, error) {
    var user User
    err := r.db.QueryRowContext(
        ctx,
        "SELECT id, name, email FROM users WHERE id = $1",
        id,
    ).Scan(&user.ID, &user.Name, &user.Email)
    
    if err == sql.ErrNoRows {
        return nil, fmt.Errorf("user not found")
    }
    return &user, err
}

func (r *PostgresUserRepository) Save(ctx context.Context, user *User) error {
    _, err := r.db.ExecContext(
        ctx,
        "INSERT INTO users (id, name, email) VALUES ($1, $2, $3)",
        user.ID, user.Name, user.Email,
    )
    return err
}

// service.go (business logic layer)
package service

import "context"

type UserService struct {
    repo repository.UserRepository
}

func NewUserService(repo repository.UserRepository) *UserService {
    return &UserService{repo: repo}
}

func (s *UserService) GetUser(ctx context.Context, id string) (*User, error) {
    // Business logic: validation, caching, etc.
    if id == "" {
        return nil, fmt.Errorf("empty user id")
    }

    user, err := s.repo.FindByID(ctx, id)
    if err != nil {
        return nil, fmt.Errorf("failed to fetch user: %w", err)
    }

    return user, nil
}

func (s *UserService) CreateUser(ctx context.Context, name, email string) (*User, error) {
    // Validation
    if name == "" || email == "" {
        return nil, fmt.Errorf("invalid input")
    }

    // Create user
    user := &User{
        ID:    uuid.New().String(),
        Name:  name,
        Email: email,
    }

    // Save to database
    if err := s.repo.Save(ctx, user); err != nil {
        return nil, fmt.Errorf("failed to save user: %w", err)
    }

    return user, nil
}

// handler.go (HTTP layer)
package handler

import (
    "encoding/json"
    "net/http"
)

type UserHandler struct {
    service *service.UserService
}

func NewUserHandler(service *service.UserService) *UserHandler {
    return &UserHandler{service: service}
}

func (h *UserHandler) GetUser(w http.ResponseWriter, r *http.Request) {
    id := r.URL.Query().Get("id")

    user, err := h.service.GetUser(r.Context(), id)
    if err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }

    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(user)
}

func (h *UserHandler) CreateUser(w http.ResponseWriter, r *http.Request) {
    var req struct {
        Name  string `json:"name"`
        Email string `json:"email"`
    }

    if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
        http.Error(w, "Invalid JSON", http.StatusBadRequest)
        return
    }

    user, err := h.service.CreateUser(r.Context(), req.Name, req.Email)
    if err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }

    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusCreated)
    json.NewEncoder(w).Encode(user)
}

// main.go (dependency wiring)
package main

func main() {
    db := sql.Open("postgres", "...")
    
    // Wire layers
    repo := repository.NewPostgresUserRepository(db)
    svc := service.NewUserService(repo)
    handler := handler.NewUserHandler(svc)

    mux := http.NewServeMux()
    mux.HandleFunc("/users", handler.GetUser)
    mux.HandleFunc("/users/create", handler.CreateUser)

    http.ListenAndServe(":8080", mux)
}

// test.go (easy to mock)
package service

import "testing"

type MockRepository struct{}

func (m *MockRepository) FindByID(ctx context.Context, id string) (*User, error) {
    return &User{ID: "1", Name: "Mock"}, nil
}

func TestGetUser(t *testing.T) {
    mockRepo := &MockRepository{}
    service := NewUserService(mockRepo)

    user, _ := service.GetUser(context.Background(), "1")
    if user.Name != "Mock" {
        t.Error("Expected mock user")
    }
}
```

### How it works

```
Request
  ↓
Handler (parse request)
  ↓
Service (business logic)
  ↓
Repository (database)
  ↓
Response
```

Each layer is independent. Swap repository → easy to change database.

### Common mistakes

- ⚠️ **Handler logic in service** → Service should be database-agnostic
- ⚠️ **Repository exposing SQL details** → Repository should hide database implementation
- ⚠️ **Tight coupling** → Use interfaces, not concrete types
- ⚠️ **Forgetting dependency injection** → Pass dependencies via constructors

### Backend relevance

Standard Go backend layout:
```
myapp/
├── cmd/server/main.go
├── internal/
│   ├── handler/
│   │   ├── user.go
│   │   └── order.go
│   ├── service/
│   │   ├── user.go
│   │   └── order.go
│   ├── repository/
│   │   ├── user.go
│   │   └── order.go
│   └── model/
│       └── user.go
```

### Interview questions

**Q: Why separate handler, service, and repository?**
A: Separation of concerns. Handler knows HTTP, service knows business rules, repository knows database. Easy to test, modify, and swap.

**Q: How do you test without a real database?**
A: Create mock repository implementing the interface. Service uses mock instead of real database.

### Quick revision

-  **Handler** = HTTP parsing + response formatting
-  **Service** = business logic, validation, orchestration
-  **Repository** = database queries (abstracted via interface)
-  Loose coupling → swap implementations easily
-  Test service with mock repository (no database needed)

**[↑ Back to Top](#-lets-learn-go)**

---

## Configuration Management

### What is it?

Loading application settings (database URL, port, secrets) from environment variables, config files, or flags.

### Why does Go have it?

Different environments (dev, prod) need different configs. Hardcoding secrets is dangerous. Environment variables keep secrets out of source code.

### Syntax / Example

```go
package main

import (
    "fmt"
    "log"
    "os"
    "strconv"

    "github.com/joho/godotenv"
)

// Config struct
type Config struct {
    DBHost       string
    DBPort       int
    DBUser       string
    DBPassword   string
    HTTPPort     int
    LogLevel     string
    JWTSecret    string
    RedisURL     string
}

// Load from environment
func LoadConfig() *Config {
    // Load .env file (dev only)
    godotenv.Load()  // Ignore error if no .env

    return &Config{
        DBHost:     getEnv("DB_HOST", "localhost"),
        DBPort:     getEnvInt("DB_PORT", 5432),
        DBUser:     getEnv("DB_USER", "postgres"),
        DBPassword: getEnv("DB_PASSWORD", ""),  // Required in prod
        HTTPPort:   getEnvInt("HTTP_PORT", 8080),
        LogLevel:   getEnv("LOG_LEVEL", "info"),
        JWTSecret:  getEnv("JWT_SECRET", ""),  // Must exist
        RedisURL:   getEnv("REDIS_URL", "redis://localhost:6379"),
    }
}

// Helper functions
func getEnv(key, defaultVal string) string {
    if value, exists := os.LookupEnv(key); exists {
        return value
    }
    return defaultVal
}

func getEnvInt(key string, defaultVal int) int {
    valStr := getEnv(key, "")
    if val, err := strconv.Atoi(valStr); err == nil {
        return val
    }
    return defaultVal
}

// Validate config
func (c *Config) Validate() error {
    if c.DBPassword == "" {
        return fmt.Errorf("DB_PASSWORD is required")
    }
    if c.JWTSecret == "" {
        return fmt.Errorf("JWT_SECRET is required")
    }
    if c.HTTPPort <= 0 || c.HTTPPort > 65535 {
        return fmt.Errorf("invalid HTTP_PORT")
    }
    return nil
}

func main() {
    cfg := LoadConfig()
    
    if err := cfg.Validate(); err != nil {
        log.Fatalf("Config error: %v", err)
    }

    fmt.Printf("Database: %s:%d\n", cfg.DBHost, cfg.DBPort)
    fmt.Printf("HTTP Port: %d\n", cfg.HTTPPort)
    fmt.Printf("Log Level: %s\n", cfg.LogLevel)
}
```

**.env file (development):**
```
DB_HOST=localhost
DB_PORT=5432
DB_USER=postgres
DB_PASSWORD=devpassword
HTTP_PORT=8080
LOG_LEVEL=debug
JWT_SECRET=dev-secret-not-for-production
REDIS_URL=redis://localhost:6379
```

**Production deployment (environment variables):**
```bash
export DB_HOST=prod-db.example.com
export DB_PORT=5432
export DB_USER=prod_user
export DB_PASSWORD=<secure-password>
export HTTP_PORT=8080
export LOG_LEVEL=info
export JWT_SECRET=<prod-secret>
export REDIS_URL=redis://prod-redis:6379

./app  # Reads from environment
```

### How it works

1. **`os.LookupEnv()`** = read environment variable
2. **`godotenv.Load()`** = load from `.env` file (dev convenience)
3. **Defaults** = sensible values if not set
4. **Validation** = check required configs at startup
5. **Never hardcode secrets** = use environment variables

### Common mistakes

- ⚠️ **Committing .env files** → Add to `.gitignore`! Secrets leak otherwise.
- ⚠️ **Missing validation** → Check required configs at startup (not when used)
- ⚠️ **Type conversion errors** → Use helper functions for int/bool parsing
- ⚠️ **No defaults** → Provide sensible defaults for dev convenience

### Backend relevance

Production-ready config:
```go
// config.go
type Config struct {
    Database DatabaseConfig
    Redis    RedisConfig
    HTTP     HTTPConfig
    Auth     AuthConfig
}

type DatabaseConfig struct {
    Host     string
    Port     int
    User     string
    Password string
    Name     string
}

func LoadConfig() *Config {
    return &Config{
        Database: DatabaseConfig{
            Host:     getEnv("DB_HOST", "localhost"),
            Port:     getEnvInt("DB_PORT", 5432),
            User:     getEnv("DB_USER", "postgres"),
            Password: getEnv("DB_PASSWORD", ""),
            Name:     getEnv("DB_NAME", "myapp"),
        },
        Redis: RedisConfig{
            URL: getEnv("REDIS_URL", "redis://localhost:6379"),
        },
        // ...
    }
}

// main.go
func main() {
    cfg := LoadConfig()
    if err := cfg.Validate(); err != nil {
        log.Fatalf("Invalid config: %v", err)
    }

    db := database.Connect(cfg.Database)
    cache := redis.Connect(cfg.Redis)
    
    // Use config throughout app
}
```

### Interview questions

**Q: Should you commit .env files?**
A: No! Add to .gitignore. Use environment variables in production. .env is dev convenience only.

**Q: How do you provide secrets in production?**
A: Environment variables (cloud provider sets them). Or secret management (Vault, AWS Secrets Manager).

### Quick revision

-  Use environment variables for configs
-  `.env` file for development convenience
-  Never commit `.env` (add to `.gitignore`)
-  Validate required configs at startup
-  Provide sensible defaults for optional configs

**[↑ Back to Top](#-lets-learn-go)**

---

<div align="center">

## 🟡 CONCURRENCY & SCALABILITY

*Making Go fast—handling thousands of requests.*

</div>

---

## Goroutines

### What is it?

Goroutines are lightweight "threads" managed by the Go runtime. Thousands run concurrently on a few OS threads (unlike OS threads, which are expensive).

### Why does Go have it?

Handling 10,000 simultaneous connections needs concurrency, but OS threads are heavy. Goroutines are cheap (~2KB each vs ~2MB for OS thread). Go scheduler multiplexes goroutines onto OS threads automatically.

### Syntax / Example

```go
package main

import (
    "fmt"
    "sync"
    "time"
)

func printNumbers(name string, count int) {
    for i := 1; i <= count; i++ {
        fmt.Printf("[%s] %d\n", name, i)
        time.Sleep(100 * time.Millisecond)
    }
}

func main() {
    // Single goroutine (blocks)
    printNumbers("main", 3)

    fmt.Println("---")

    // Concurrent goroutines (non-blocking)
    go printNumbers("goroutine1", 3)
    go printNumbers("goroutine2", 3)

    // Give goroutines time to finish
    time.Sleep(1 * time.Second)

    fmt.Println("Main done!")
}

// Output:
// [goroutine1] 1
// [goroutine2] 1
// [goroutine1] 2
// [goroutine2] 2
// ... (interleaved)

// Better: wait for goroutines with WaitGroup
func main() {
    var wg sync.WaitGroup

    // Add 2 goroutines to wait for
    wg.Add(2)

    go func() {
        defer wg.Done()  // Mark as done
        printNumbers("g1", 3)
    }()

    go func() {
        defer wg.Done()
        printNumbers("g2", 3)
    }()

    wg.Wait()  // Block until all Done()
    fmt.Println("All goroutines finished!")
}
```

### How it works

1. **`go function()`** = start goroutine (non-blocking)
2. **Goroutine = lightweight** → thousands can run concurrently
3. **Go scheduler** = maps goroutines to OS threads automatically
4. **`sync.WaitGroup`** = wait for goroutines to complete
5. **Main function exits** → goroutines are cancelled (no background tasks!)

### Common mistakes

- ⚠️ **Main exits before goroutines finish** → Use `sync.WaitGroup` or channels
- ⚠️ **Race conditions** → Multiple goroutines access same variable without synchronization
- ⚠️ **Goroutine leaks** → Goroutine blocked forever (never completes, wastes memory)
- ⚠️ **Too many goroutines** → Creating 1M goroutines may still run but waste resources

### Backend relevance

Handling concurrent requests:
```go
// HTTP handler (runs in goroutine automatically)
func (h *Handler) GetUser(w http.ResponseWriter, r *http.Request) {
    // This runs in a separate goroutine per request!
    user, _ := h.service.GetUser(r.Context(), userID)
    json.NewEncoder(w).Encode(user)
}

// ListenAndServe creates goroutine per request
server := &http.Server{Addr: ":8080"}
server.ListenAndServe()  // Handles 10,000+ concurrent requests

// Parallel database queries
var wg sync.WaitGroup
var users []User
var orders []Order

wg.Add(2)
go func() {
    defer wg.Done()
    users, _ = userService.GetAllUsers(ctx)
}()
go func() {
    defer wg.Done()
    orders, _ = orderService.GetAllOrders(ctx)
}()
wg.Wait()  // Wait for both to complete

return users, orders
```

### Interview questions

**Q: What's the difference between goroutines and OS threads?**
A: Goroutines are managed by Go runtime (lightweight, ~2KB). OS threads are created by OS (heavy, ~2MB). Go can run thousands of goroutines on few OS threads.

**Q: What happens if the main function exits before goroutines finish?**
A: Goroutines are cancelled immediately. Use `sync.WaitGroup` to wait.

### Quick revision

-  `go function()` starts concurrent goroutine
-  Goroutines are lightweight (thousands can run)
-  `sync.WaitGroup` waits for goroutines to finish
-  Main exits → all goroutines cancelled
-  Perfect for handling concurrent HTTP requests

**[↑ Back to Top](#-lets-learn-go)**

---

## Channels

### What is it?

Channels are typed "pipes" that goroutines use to communicate safely. Send data on one end, receive on the other.

### Why does Go have it?

Goroutines need coordination. Channels prevent race conditions by ensuring only one goroutine accesses data at a time (safe synchronization).

### Syntax / Example

```go
package main

import (
    "fmt"
)

func main() {
    // Create channel
    messages := make(chan string)

    // Send in goroutine
    go func() {
        messages <- "Hello from goroutine!"
    }()

    // Receive (blocks until data available)
    msg := <-messages
    fmt.Println(msg)

    // Buffered channel (can hold N values)
    buffer := make(chan int, 2)
    buffer <- 1
    buffer <- 2
    fmt.Println(<-buffer)  // 1
    fmt.Println(<-buffer)  // 2

    // Iterate over channel
    results := make(chan int)
    go func() {
        results <- 1
        results <- 2
        results <- 3
        close(results)  // Signal no more data
    }()

    for val := range results {
        fmt.Println(val)
    }

    // Channel as function parameter (direction)
    go produce(messages)  // Send-only
    consume(messages)     // Receive-only
}

// Send-only channel (can't receive)
func produce(ch chan<- string) {
    ch <- "data"
}

// Receive-only channel (can't send)
func consume(ch <-chan string) {
    msg := <-ch
    fmt.Println(msg)
}

// Example: fetching data in parallel
func fetchUserData(userID string) (string, error) {
    userCh := make(chan string, 1)
    ordersCh := make(chan string, 1)

    go func() {
        user, _ := fetchUser(userID)
        userCh <- user
    }()

    go func() {
        orders, _ := fetchOrders(userID)
        ordersCh <- orders
    }()

    user := <-userCh
    orders := <-ordersCh
    return user + orders, nil
}
```

### How it works

1. **`make(chan Type)`** = create unbuffered channel
2. **`make(chan Type, N)`** = buffered channel (holds N values)
3. **`ch <- value`** = send value (blocks if buffer full)
4. **`value := <-ch`** = receive value (blocks if empty)
5. **`close(ch)`** = signal no more data
6. **`<-chan Type`** = receive-only, `chan<- Type` = send-only

### Common mistakes

- ⚠️ **Sending on closed channel** → Panic! Only sender should close.
- ⚠️ **Deadlock** → Both goroutines waiting on channel. Use buffered channels or goroutines.
- ⚠️ **Leaking goroutines** → Goroutine blocked on channel forever (memory leak)
- ⚠️ **Not closing channel** → Receivers can't know when to stop (if using range)

### Backend relevance

Real patterns:
```go
// Worker pool (handle jobs concurrently)
type JobQueue struct {
    jobs    chan Job
    results chan Result
}

func (q *JobQueue) ProcessJobs(workerCount int) {
    for i := 0; i < workerCount; i++ {
        go q.worker()  // Multiple workers
    }
}

func (q *JobQueue) worker() {
    for job := range q.jobs {  // Wait for jobs
        result := processJob(job)
        q.results <- result  // Send result back
    }
}

// Timeout handling
func FetchWithTimeout(userID string, timeout time.Duration) (*User, error) {
    userCh := make(chan *User, 1)
    errCh := make(chan error, 1)

    go func() {
        user, err := fetchUser(userID)
        if err != nil {
            errCh <- err
        } else {
            userCh <- user
        }
    }()

    select {
    case user := <-userCh:
        return user, nil
    case err := <-errCh:
        return nil, err
    case <-time.After(timeout):
        return nil, fmt.Errorf("timeout")
    }
}
```

### Interview questions

**Q: What's the difference between buffered and unbuffered channels?**
A: Unbuffered blocks until sender/receiver ready. Buffered holds N values (doesn't block until full).

**Q: What happens if you send on a closed channel?**
A: Panic! Only sender should close. Receivers check `val, ok := <-ch`.

### Quick revision

-  Channels = typed pipes for goroutine communication
-  Unbuffered = blocks until both sides ready
-  Buffered = holds N values (async communication)
-  Close signals no more data (receivers can detect)
-  `select` waits on multiple channels

**[↑ Back to Top](#-lets-learn-go)**

---

## Synchronization

### What is it?

Mutexes (mutual exclusion locks) ensure only one goroutine accesses shared data at a time. Prevents race conditions.

### Why does Go have it?

Without locks, concurrent access to shared variables causes corruption. Mutex ensures safe, ordered access.

### Syntax / Example

```go
package main

import (
    "fmt"
    "sync"
)

// Counter protected by mutex
type SafeCounter struct {
    mu    sync.Mutex  // Protects value
    value int
}

// Increment safely
func (c *SafeCounter) Increment() {
    c.mu.Lock()
    defer c.mu.Unlock()
    c.value++
}

// Get value safely
func (c *SafeCounter) Get() int {
    c.mu.Lock()
    defer c.mu.Unlock()
    return c.value
}

// RWMutex for read-heavy workloads
type Cache struct {
    mu    sync.RWMutex
    data  map[string]interface{}
}

// Multiple readers (no lock)
func (c *Cache) Get(key string) interface{} {
    c.mu.RLock()  // Read lock (multiple allowed)
    defer c.mu.RUnlock()
    return c.data[key]
}

// Single writer (exclusive lock)
func (c *Cache) Set(key string, value interface{}) {
    c.mu.Lock()  // Write lock (exclusive)
    defer c.mu.Unlock()
    c.data[key] = value
}

func main() {
    counter := &SafeCounter{}

    var wg sync.WaitGroup
    wg.Add(100)

    // 100 goroutines increment counter
    for i := 0; i < 100; i++ {
        go func() {
            defer wg.Done()
            counter.Increment()
        }()
    }

    wg.Wait()
    fmt.Println("Final:", counter.Get())  // 100 (safe!)
}

// Example: Cache with RWMutex
cache := &Cache{data: make(map[string]interface{})}

// Many readers (parallel)
for i := 0; i < 100; i++ {
    go func(id int) {
        val := cache.Get("key")  // RLock, no serialization
        fmt.Printf("Reader %d: %v\n", id, val)
    }(i)
}

// Few writers (exclusive)
go func() {
    cache.Set("key", "value")  // Lock, exclusive access
}()
```

### How it works

1. **`sync.Mutex`** = exclusive lock (one holder at a time)
2. **`Lock()` / `Unlock()`** = acquire/release lock
3. **`defer Unlock()`** = ensure lock released even if panic
4. **`sync.RWMutex`** = read-write lock (multiple readers OR one writer)
5. **`RLock() / RUnlock()`** = read lock (shared)
6. **`Lock() / Unlock()`** = write lock (exclusive)

### Common mistakes

- ⚠️ **Forgetting to unlock** → Deadlock! Use `defer Unlock()`
- ⚠️ **Holding lock too long** → Blocks other goroutines. Keep critical section small.
- ⚠️ **Using Mutex when not needed** → Channels often simpler for communication
- ⚠️ **Deadlock with multiple locks** → Lock order must be consistent across goroutines

### Backend relevance

Cache layer with RWMutex:
```go
type UserCache struct {
    mu    sync.RWMutex
    cache map[string]*User
}

func (c *UserCache) Get(id string) (*User, bool) {
    c.mu.RLock()  // Many concurrent readers
    defer c.mu.RUnlock()
    user, exists := c.cache[id]
    return user, exists
}

func (c *UserCache) Set(id string, user *User) {
    c.mu.Lock()  // Exclusive write
    defer c.mu.Unlock()
    c.cache[id] = user
}

// Typical flow: many reads, few writes
for i := 0; i < 1000; i++ {
    go func() {
        user, _ := cache.Get("user1")  // No contention
    }()
}

cache.Set("user1", user)  // Blocks reads briefly
```

### Interview questions

**Q: When should you use Mutex vs channels?**
A: Mutex for protecting shared state (caches, counters). Channels for communication between goroutines.

**Q: What's the difference between Mutex and RWMutex?**
A: Mutex = exclusive lock. RWMutex = multiple readers OR one writer (better for read-heavy workloads).

### Quick revision

-  `sync.Mutex` = exclusive lock (one holder)
-  `sync.RWMutex` = multiple readers OR one writer
-  Always `defer Unlock()` to prevent deadlocks
-  Keep critical section small
-  Channels often better than mutexes for communication

**[↑ Back to Top](#-lets-learn-go)**

---

## Context

### What is it?

`context.Context` is a way to pass request-scoped values (user info, timeouts, cancellation) through your application layers.

### Why does Go have it?

Requests need timeouts (kill slow queries), cancellation (user closes browser), and metadata (request ID, user). Context flows through handler → service → database cleanly.

### Syntax / Example

```go
package main

import (
    "context"
    "fmt"
    "time"
)

// Timeout context
func FetchUserWithTimeout() {
    ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
    defer cancel()  // Clean up

    // If takes > 5s, ctx.Done() becomes readable
    user, err := fetchUserFromDB(ctx, "123")
    if err == context.DeadlineExceeded {
        fmt.Println("Query timeout!")
    }
    fmt.Println(user)
}

// Fetch from DB with timeout
func fetchUserFromDB(ctx context.Context, id string) (*User, error) {
    // Simulate query
    select {
    case <-time.After(10 * time.Second):
        return &User{ID: id}, nil
    case <-ctx.Done():
        return nil, ctx.Err()  // Timeout/cancellation
    }
}

// Cancel context
func FetchWithCancellation() {
    ctx, cancel := context.WithCancel(context.Background())

    go func() {
        time.Sleep(2 * time.Second)
        cancel()  // Signal cancellation
    }()

    result := doLongWork(ctx)
    fmt.Println(result)
}

func doLongWork(ctx context.Context) string {
    for i := 0; i < 100; i++ {
        select {
        case <-time.After(100 * time.Millisecond):
            fmt.Printf("Doing work... %d\n", i)
        case <-ctx.Done():
            return "Cancelled!"
        }
    }
    return "Done!"
}

// Pass values through context
func HandleRequest(w http.ResponseWriter, r *http.Request) {
    ctx := r.Context()
    
    // Add request ID
    requestID := generateID()
    ctx = context.WithValue(ctx, "requestID", requestID)

    // Pass to service
    user, _ := userService.GetUser(ctx, userID)
    fmt.Fprintf(w, "User: %s\n", user.Name)
}

func (s *UserService) GetUser(ctx context.Context, id string) (*User, error) {
    // Retrieve request ID from context
    requestID := ctx.Value("requestID").(string)
    log.Printf("[%s] Fetching user %s", requestID, id)

    // Pass context to database layer
    return s.repo.FindByID(ctx, id)
}

func (r *UserRepository) FindByID(ctx context.Context, id string) (*User, error) {
    // Database layer respects timeout
    var user User
    err := r.db.QueryRowContext(
        ctx,  // Pass context for timeout/cancellation
        "SELECT * FROM users WHERE id = $1",
        id,
    ).Scan(&user.ID, &user.Name)
    return &user, err
}
```

### How it works

1. **`context.Background()`** = root context (never times out)
2. **`context.WithTimeout()`** = timeout context (cancels after duration)
3. **`context.WithDeadline()`** = deadline context (cancels at specific time)
4. **`context.WithCancel()`** = manual cancellation
5. **`context.WithValue()`** = pass data (request ID, user info)
6. **`<-ctx.Done()`** = channel closes when context cancelled
7. **`ctx.Err()`** = reason for cancellation (timeout, manual, etc.)

### Common mistakes

- ⚠️ **Using background context for requests** → HTTP handlers pass request context
- ⚠️ **Not passing context down** → Service and repository must receive context
- ⚠️ **Ignoring timeout** → Don't check `ctx.Done()` in loops
- ⚠️ **Storing large values** → Context values should be small (request ID, not whole user struct)

### Backend relevance

HTTP middleware passes context:
```go
func TimeoutMiddleware(timeout time.Duration) func(http.Handler) http.Handler {
    return func(next http.Handler) http.Handler {
        return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
            ctx, cancel := context.WithTimeout(r.Context(), timeout)
            defer cancel()

            r = r.WithContext(ctx)  // Replace context
            next.ServeHTTP(w, r)
        })
    }
}

// Handler uses context
func (h *Handler) GetUser(w http.ResponseWriter, r *http.Request) {
    user, err := h.service.GetUser(r.Context(), userID)
    if err == context.DeadlineExceeded {
        http.Error(w, "Request timeout", http.StatusGatewayTimeout)
        return
    }
    // ...
}

// Service respects context
func (s *UserService) GetUser(ctx context.Context, id string) (*User, error) {
    user, err := s.repo.FindByID(ctx, id)
    // If ctx times out, database will cancel query
    return user, err
}
```

### Interview questions

**Q: What's the difference between timeout and deadline context?**
A: Timeout = relative duration (5 seconds from now). Deadline = absolute time (3:00 PM exactly).

**Q: Should you store large objects in context?**
A: No. Context values should be small (request ID, user ID). Don't use context as data store.

### Quick revision

-  `context.Background()` = root context
-  `WithTimeout()` / `WithDeadline()` for time limits
-  `WithCancel()` for manual cancellation
-  Pass context through handler → service → database
-  Check `<-ctx.Done()` to respect cancellation

**[↑ Back to Top](#-lets-learn-go)**

---

## Testing Fundamentals

### What is it?

Writing tests to verify your code works correctly. Go's standard library `testing` package provides utilities.

### Why does Go have it?

Prevents bugs from reaching production. Ensures refactoring doesn't break functionality. Confidence when deploying.

### Syntax / Example

```go
package user

import (
    "context"
    "testing"
)

// Simple test
func TestGetUser(t *testing.T) {
    mockRepo := &MockUserRepository{
        users: map[string]*User{
            "1": {ID: "1", Name: "Alice"},
        },
    }
    service := NewUserService(mockRepo)

    user, err := service.GetUser(context.Background(), "1")
    if err != nil {
        t.Fatalf("Unexpected error: %v", err)
    }

    if user.Name != "Alice" {
        t.Errorf("Expected 'Alice', got '%s'", user.Name)
    }
}

// Test not found
func TestGetUserNotFound(t *testing.T) {
    mockRepo := &MockUserRepository{users: map[string]*User{}}
    service := NewUserService(mockRepo)

    _, err := service.GetUser(context.Background(), "999")
    if err == nil {
        t.Error("Expected error, got nil")
    }
}

// Table-driven tests (multiple cases)
func TestCreateUserValidation(t *testing.T) {
    tests := []struct {
        name      string
        input     CreateUserRequest
        wantError bool
    }{
        {
            name:      "valid request",
            input:     CreateUserRequest{Name: "Alice", Email: "alice@ex.com"},
            wantError: false,
        },
        {
            name:      "missing name",
            input:     CreateUserRequest{Name: "", Email: "alice@ex.com"},
            wantError: true,
        },
        {
            name:      "missing email",
            input:     CreateUserRequest{Name: "Alice", Email: ""},
            wantError: true,
        },
    }

    for _, tc := range tests {
        t.Run(tc.name, func(t *testing.T) {
            service := NewUserService(&MockUserRepository{})
            _, err := service.CreateUser(context.Background(), tc.input)

            if (err != nil) != tc.wantError {
                t.Errorf("Got error %v, want error %v", err, tc.wantError)
            }
        })
    }
}

// Mock implementation (for testing without database)
type MockUserRepository struct {
    users map[string]*User
}

func (m *MockUserRepository) FindByID(ctx context.Context, id string) (*User, error) {
    user, exists := m.users[id]
    if !exists {
        return nil, fmt.Errorf("user not found")
    }
    return user, nil
}

func (m *MockUserRepository) Save(ctx context.Context, user *User) error {
    m.users[user.ID] = user
    return nil
}

// Run tests:
// $ go test ./...
// $ go test -v  (verbose)
// $ go test -run TestGetUser  (specific test)
```

### How it works

1. **Test file** = `*_test.go` (e.g., `user_test.go`)
2. **Test function** = `func TestXXX(t *testing.T)`
3. **`t.Error()`** = fail test with message
4. **`t.Fatalf()`** = fail immediately
5. **Mock implementations** = fake repository for testing without database
6. **Table-driven tests** = multiple test cases in one test function

### Common mistakes

- ⚠️ **Not testing error cases** → Test both success and failure
- ⚠️ **Hardcoding test data** → Use table-driven tests for multiple scenarios
- ⚠️ **Testing implementation, not behavior** → Test "what it does", not "how it works"
- ⚠️ **Real database in tests** → Use mocks, not real database (slow, unreliable)

### Backend relevance

HTTP handler testing:
```go
func TestCreateUserHandler(t *testing.T) {
    mockService := &MockUserService{
        createUserFn: func(ctx context.Context, req *CreateUserRequest) (*User, error) {
            return &User{ID: "1", Name: req.Name}, nil
        },
    }
    handler := NewUserHandler(mockService)

    // Create request
    body := `{"name":"Alice","email":"alice@ex.com"}`
    req := httptest.NewRequest("POST", "/users", strings.NewReader(body))
    w := httptest.NewRecorder()

    // Call handler
    handler.CreateUser(w, req)

    // Assert response
    if w.Code != http.StatusCreated {
        t.Errorf("Expected 201, got %d", w.Code)
    }

    var user User
    json.NewDecoder(w.Body).Decode(&user)
    if user.Name != "Alice" {
        t.Errorf("Expected Alice, got %s", user.Name)
    }
}
```

### Interview questions

**Q: What are mocks used for?**
A: Replace real dependencies (database, external API) to test in isolation without side effects.

**Q: Why table-driven tests?**
A: Multiple test cases in one test function. DRY principle. Easy to add new cases.

### Quick revision

-  Test files: `*_test.go`, functions: `TestXXX(t *testing.T)`
-  Mock repositories for testing without database
-  Table-driven tests for multiple scenarios
-  Test error cases, not just happy path
-  Run with `go test ./...`

**[↑ Back to Top](#-lets-learn-go)**

---

<div align="center">

## 🧠 Go Mental Model

*How it all connects.*

</div>

When you build a Go backend system, everything flows together:

```
┌─────────────────────────────────────────────────────────────┐
│                                                             │
│  Incoming HTTP Request                                     │
│           ↓                                                │
│  Middleware Chain (Logging, Auth, Timeout)                │
│           ↓                                                │
│  Handler Layer (Parse Request, Validate)                  │
│           ↓                                                │
│  Service Layer (Business Logic, Validation)               │
│           ↓                                                │
│  Repository Layer (Database Queries)                      │
│           ↓                                                │
│  Database / Cache                                          │
│           ↓                                                │
│  Response (JSON) ← Back up through layers                 │
│                                                             │
└─────────────────────────────────────────────────────────────┘
```

### Each Layer's Job

**Handler** (HTTP)
- Parse request body (JSON)
- Extract URL parameters/query strings
- Validate inputs
- Call service layer
- Format response (JSON)
- Return HTTP status code

```go
func (h *UserHandler) CreateUser(w http.ResponseWriter, r *http.Request) {
    var req CreateUserRequest
    json.NewDecoder(r.Body).Decode(&req)  // Parse
    
    user, err := h.service.CreateUser(r.Context(), req)  // Call service
    if err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }
    
    w.WriteHeader(http.StatusCreated)
    json.NewEncoder(w).Encode(user)  // Format response
}
```

**Service** (Business Logic)
- Validate business rules
- Orchestrate database calls
- Apply transformations
- Handle errors with context

```go
func (s *UserService) CreateUser(ctx context.Context, req CreateUserRequest) (*User, error) {
    if req.Email == "" {
        return nil, fmt.Errorf("email required")
    }
    
    user := &User{ID: uuid.New().String(), Name: req.Name}
    if err := s.repo.Save(ctx, user); err != nil {
        return nil, fmt.Errorf("failed to save: %w", err)
    }
    
    return user, nil
}
```

**Repository** (Data Access)
- Write SQL queries
- Map rows to structs
- Handle database-specific errors
- No business logic here

```go
func (r *UserRepository) Save(ctx context.Context, user *User) error {
    _, err := r.db.ExecContext(ctx,
        "INSERT INTO users (id, name, email) VALUES ($1, $2, $3)",
        user.ID, user.Name, user.Email)
    return err
}
```

### Concurrency Throughout

**HTTP Server** automatically launches goroutine per request:
- 1000 concurrent users = 1000 goroutines (no problem!)

**Goroutines + Channels** for parallelism within handler:
- Fetch user + orders simultaneously
- Both complete → respond to client

**Context** flows through all layers:
- Timeout set by middleware
- Respected by database layer
- Slow query gets cancelled

**Mutexes/Channels** protect shared state:
- Cache hits from 1000s of goroutines safely
- Work queue distributes jobs to workers

### Error Propagation

Errors bubble up with context:

```
Database Error (connection failed)
        ↓
Repository wraps: "failed to fetch user: connection failed"
        ↓
Service wraps: "failed to load user profile: failed to fetch user: ..."
        ↓
Handler logs + responds HTTP 500
```

### Request Lifecycle

1. **HTTP Request arrives** → Go automatically creates goroutine
2. **Middleware chain** → logging, auth, timeout (context added)
3. **Handler** → parses JSON, validates, calls service
4. **Service** → applies business rules, calls repository
5. **Repository** → queries database using context (respects timeout)
6. **Response flows back** → service → handler → JSON → HTTP 200
7. **Goroutine exits** → memory freed

This is the Go way. Clean, simple, fast.

---

<div align="center">

## Go Backend Roadmap

*Your learning path to production-ready systems.*

</div>

### ✅ Beginner (Fundamentals Complete)

- [x] Variables, constants, types
- [x] Functions, multiple returns
- [x] Structs, methods, receivers
- [x] Pointers, dereferencing
- [x] Control flow (if, switch, for)
- [x] Slices, maps, arrays
- [x] Error handling
- [x] Packages, imports

**Current Level:** Comfortable with Go syntax. Can write simple programs.

---

### 🟡 Intermediate (Core Concepts)

- [x] Interfaces, implicit implementation
- [x] Goroutines, concurrency model
- [x] Channels, synchronization
- [x] HTTP servers, REST APIs
- [x] Handlers, middleware
- [x] Configuration management
- [x] Context, timeouts
- [x] Testing basics

**Current Level:** Can build HTTP servers. Understand concurrency. Write tests.

---

### 🟠 Backend (Backend Patterns)

- [x] Layered architecture (handler → service → repository)
- [x] Dependency injection
- [x] Request/response models
- [ ] **PostgreSQL integration** (use `database/sql`)
- [ ] **Connection pooling** (max open connections)
- [ ] **Transaction handling** (BEGIN, COMMIT, ROLLBACK)
- [ ] **Redis caching** (cache-aside pattern)
- [ ] **Middleware chains** (chaining multiple middlewares)
- [ ] **Logging** (structured logging with slog)
- [ ] **Error handling** (HTTP error responses)

**Next Priority:** Connect to real database. Add caching. Logging.

---

### 🔴 Production (Deployment)

- [ ] **Graceful shutdown** (drain connections, finish requests)
- [ ] **Docker** (containerize Go app)
- [ ] **Health checks** (`/health` endpoint)
- [ ] **Metrics / observability** (Prometheus)
- [ ] **Rate limiting** (limit requests per user)
- [ ] **Authentication** (JWT tokens)
- [ ] **CORS middleware** (cross-origin requests)
- [ ] **Request ID tracking** (end-to-end tracing)
- [ ] **Database migrations** (version schema)
- [ ] **Environment-specific config** (dev, staging, prod)

**Ready for:** Small production deployments. Monitoring. Multi-user systems.

---

### 💎 Advanced (Optimization)

- [ ] **Advanced concurrency** (worker pools, semaphores)
- [ ] **gRPC** (protobuf, binary protocol)
- [ ] **Message queues** (RabbitMQ, Kafka)
- [ ] **Event sourcing** (immutable event store)
- [ ] **Distributed tracing** (Jaeger, OpenTelemetry)
- [ ] **Load testing** (Apache Bench, k6)
- [ ] **Code generation** (wire, stringer)
- [ ] **Profiling** (pprof, flame graphs)

---

---

<div align="center">

## 🛠️ Practice Projects

*Build real skills through hands-on projects.*

</div>

### 1️⃣ **Beginner: URL Shortener** ✅ (You did this!)

**What you learn:**
- HTTP handlers
- JSON encoding/decoding
- In-memory storage (map)
- URL validation
- Basic error handling

**Goal:** Build a service that converts long URLs to short codes (like bit.ly).

```go
// POST /shorten
// {"url":"https://example.com/very/long/path"}
// Response: {"shortCode":"abc123"}

// GET /redirect/:code
// Redirect to original URL
```

**Challenges:**
- Handle duplicate URLs (return same code)
- Validate URL format
- Track creation time
- Implement expiration (old links expire)

---

### 2️⃣ **Beginner-Intermediate: Expense Tracker** ✅ (You did this!)

**What you learn:**
- Layered architecture (handler → service → repository)
- Struct models + JSON tags
- Request validation
- In-memory database
- Multiple handlers

**Goal:** Track daily expenses with categories.

```go
// POST /expenses
// {"amount":50,"category":"food","description":"lunch"}

// GET /expenses
// Returns all expenses with totals by category

// GET /expenses/summary
// {"food":150,"transport":200}
```

**Challenges:**
- Validate amounts (positive numbers)
- Group by category
- Filter by date range
- Delete/update expenses

---

### 3️⃣ **Intermediate: TODO API with Persistence**

**What you learn:**
- Connect to PostgreSQL database
- Database migrations
- CRUD operations (Create, Read, Update, Delete)
- SQL queries with proper escaping
- Connection pooling

**Goal:** Persistent TODO list API.

```go
// POST /todos
// {"title":"Buy groceries","completed":false}

// GET /todos
// Returns all todos

// PUT /todos/:id
// Update a todo

// DELETE /todos/:id
// Delete a todo

// GET /todos?completed=false
// Filter todos
```

**Schema:**
```sql
CREATE TABLE todos (
    id SERIAL PRIMARY KEY,
    title VARCHAR(255) NOT NULL,
    completed BOOLEAN DEFAULT false,
    created_at TIMESTAMP DEFAULT NOW()
);
```

**Challenges:**
- Pagination (limit, offset)
- Sorting (by created_at, title)
- Transaction handling
- Error handling for database errors

---

### 4️⃣ **Intermediate-Advanced: User Service with Auth**

**What you learn:**
- Middleware (auth, logging, timeout)
- JWT token generation/validation
- Password hashing (bcrypt)
- Middleware chains
- Request/response models
- Structured logging

**Goal:** User management service with authentication.

```go
// POST /register
// {"email":"user@ex.com","password":"secret"}

// POST /login
// Returns JWT token

// GET /me (requires auth)
// Returns current user

// GET /users (admin only)
// List all users

// PUT /users/:id (own account or admin)
// Update user
```

**Challenges:**
- Password validation (strong passwords)
- Token expiration
- Refresh tokens
- Admin roles
- Rate limiting on login attempts

---

### 5️⃣ **Advanced: E-Commerce Backend**

**What you learn:**
- Complex database schema (users, products, orders, order_items)
- Transactions (consistent order creation)
- Caching (Redis) for frequently accessed products
- Async processing (background jobs)
- Webhook handling (payment notifications)

**Goal:** Full e-commerce API.

```
Users → Products → Shopping Cart → Orders
                 ↓
            Payments (Stripe API)
                 ↓
           Notifications (Email)
```

**Key features:**
- User authentication + profiles
- Product catalog (search, filter, sort)
- Shopping cart (persistent)
- Order placement (transaction)
- Payment processing (integrate Stripe/Razorpay)
- Email notifications
- Admin dashboard

**Database:**
```sql
users
products
categories
shopping_carts
orders
order_items
payments
```

**Challenges:**
- Transaction safety (order + payment atomic)
- Inventory management (prevent overselling)
- Concurrent order processing
- Payment webhook validation
- Email queue system

---

### 6️⃣ **Advanced: Distributed System - Chat Server**

**What you learn:**
- WebSocket connections (real-time)
- Message broadcasting
- Concurrency patterns (message channels)
- Database persistence
- Session management
- Horizontal scaling (Redis Pub/Sub)

**Goal:** Real-time chat application.

**Features:**
- User authentication
- Create/join chat rooms
- Real-time message broadcasting
- Message persistence
- User presence (online/offline)
- Typing indicators

**Architecture:**
```
Client (WebSocket)
   ↓
API Server (goroutine per connection)
   ↓
Message Hub (channels for broadcast)
   ↓
Database (message storage)
   ↓
Redis Pub/Sub (multi-server sync)
```

---

### Progression Path

```
URL Shortener (simple)
       ↓
Expense Tracker (structured)
       ↓
TODO API (database)
       ↓
User Service (auth, middleware)
       ↓
E-Commerce (complex)
       ↓
Chat Server (distributed)
```

Each project builds on previous skills. Start with #3 when you finish current projects.

---

<div align="center">

## 📚 What To Learn Next

*The missing pieces for production-ready backend.*

</div>

You've learned the fundamentals. Now focus on these backend essentials:

### 🟴 **Database Integration (Do This First)**

#### PostgreSQL with Go

```go
// database/sql package (standard library)
import "database/sql"
import _ "github.com/lib/pq"  // PostgreSQL driver

// Connection pooling
db, err := sql.Open("postgres", "postgres://user:pass@localhost/dbname")
db.SetMaxOpenConns(25)
db.SetMaxIdleConns(5)

// Query with context (respects timeout)
var user User
err := db.QueryRowContext(ctx,
    "SELECT id, name FROM users WHERE id = $1",
    userID,
).Scan(&user.ID, &user.Name)

// Transactions
tx, _ := db.BeginTx(ctx, nil)
// INSERT, UPDATE
tx.Commit()
```

**Must Learn:**
- Connection pooling (prevent resource exhaustion)
- Query escaping (prevent SQL injection)
- Transactions (atomic operations)
- Migrations (version your schema)

---

### 🟠 **Caching Layer**

#### Redis with Go

```go
// go-redis package
import "github.com/redis/go-redis/v9"

client := redis.NewClient(&redis.Options{
    Addr: "localhost:6379",
})

// Cache-aside pattern
func GetUser(ctx context.Context, id string) (*User, error) {
    // Try cache first
    cached, err := client.Get(ctx, "user:"+id).Result()
    if err == nil {
        json.Unmarshal([]byte(cached), &user)
        return &user, nil
    }

    // Cache miss → fetch from database
    user, err := db.GetUser(ctx, id)
    if err != nil {
        return nil, err
    }

    // Update cache
    data, _ := json.Marshal(user)
    client.Set(ctx, "user:"+id, data, 1*time.Hour)

    return user, nil
}
```

**Must Learn:**
- Cache invalidation (when to clear)
- TTL (time-to-live)
- Hot keys (protect from thundering herd)

---

### 🟠 **Authentication**

#### JWT Tokens

```go
// github.com/golang-jwt/jwt/v5
import "github.com/golang-jwt/jwt/v5"

// Create token
claims := jwt.MapClaims{
    "userID": user.ID,
    "email":  user.Email,
    "exp":    time.Now().Add(24 * time.Hour).Unix(),
}

token, err := jwt.NewWithClaims(jwt.SigningMethodHS256, claims).
    SignedString([]byte(secretKey))

// Verify token (in middleware)
token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
    return []byte(secretKey), nil
})

if claims, ok := token.Claims.(jwt.MapClaims); ok {
    userID := claims["userID"].(string)
}
```

**Must Learn:**
- Token generation
- Token validation
- Token refresh
- Secret key management

---

### 🟡 **Structured Logging**

#### log/slog (Go 1.21+)

```go
import "log/slog"

logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))

logger.Info("user created",
    "userID", user.ID,
    "email", user.Email,
    "duration_ms", elapsed.Milliseconds(),
)

// Output: {"time":"...","level":"INFO","msg":"user created","userID":"123",...}
```

**Must Learn:**
- Structured logging (JSON format)
- Log levels (DEBUG, INFO, WARN, ERROR)
- Request ID correlation
- Performance impact

---

### 🟡 **Error Handling at Scale**

```go
// Wrap errors with context
if err != nil {
    return fmt.Errorf("failed to fetch user %s: %w", userID, err)
}

// Custom errors
type ValidationError struct {
    Field string
    Value string
}

func (e *ValidationError) Error() string {
    return fmt.Sprintf("invalid %s: %s", e.Field, e.Value)
}

// HTTP error responses
func handleError(w http.ResponseWriter, err error) {
    if err == sql.ErrNoRows {
        http.Error(w, "not found", http.StatusNotFound)
    } else if _, ok := err.(*ValidationError); ok {
        http.Error(w, err.Error(), http.StatusBadRequest)
    } else {
        http.Error(w, "internal server error", http.StatusInternalServerError)
    }
}
```

---

### 🟡 **Graceful Shutdown**

```go
server := &http.Server{Addr: ":8080"}

// Handle shutdown signal
sigChan := make(chan os.Signal, 1)
signal.Notify(sigChan, os.Interrupt, syscall.SIGTERM)

go func() {
    <-sigChan
    log.Println("Shutdown signal received")
    
    ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
    defer cancel()
    
    server.Shutdown(ctx)  // Wait for ongoing requests
}()

server.ListenAndServe()  // Will exit cleanly
```

**Why:** Allow in-flight requests to complete before stopping.

---

### 🔴 **Rate Limiting**

```go
import "golang.org/x/time/rate"

limiter := rate.NewLimiter(rate.Every(time.Second), 10)  // 10 requests/sec

func RateLimitMiddleware(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        if !limiter.Allow() {
            http.Error(w, "too many requests", http.StatusTooManyRequests)
            return
        }
        next.ServeHTTP(w, r)
    })
}
```

---

### 🔴 **Observability**

#### Metrics (Prometheus)

```go
import "github.com/prometheus/client_golang/prometheus"

requestCount := prometheus.NewCounterVec(
    prometheus.CounterOpts{Name: "http_requests_total"},
    []string{"method", "status"},
)

// In handler
requestCount.WithLabelValues("GET", "200").Inc()
```

#### Tracing (Request ID)

```go
requestID := uuid.New().String()
ctx := context.WithValue(r.Context(), "requestID", requestID)

// Log throughout the request
logger.Info("processing request", "requestID", requestID)
// ... service, database
logger.Info("request completed", "requestID", requestID, "duration_ms", elapsed)
```

---

### Learning Path

**Start:**
1. PostgreSQL integration (database/sql)
2. Basic error handling patterns
3. Structured logging (log/slog)

**Then:**
4. Redis caching
5. JWT authentication
6. Middleware chains

**Finally:**
7. Graceful shutdown
8. Rate limiting
9. Metrics & observability

---

<div align="center">

## 🎯 Checklist - Your Progress

Use this to track what you've completed:

</div>

### Fundamentals
- [x] Installation & project setup
- [x] Variables, constants, types
- [x] Functions & multiple returns
- [x] Control flow (if, switch, for)
- [x] Arrays, slices, maps
- [x] Structs & pointers
- [x] Methods & receivers

### Core Concepts
- [x] Interfaces
- [x] Error handling
- [x] Packages & organization

### Backend Essentials
- [x] HTTP servers & REST APIs
- [x] Handlers & middleware (basic)
- [x] Layered architecture
- [x] Configuration management

### Concurrency
- [x] Goroutines
- [x] Channels
- [x] Synchronization (mutex)
- [x] Context

### Quality
- [x] Testing basics

### Next Phase
- [ ] PostgreSQL integration
- [ ] Connection pooling
- [ ] Transactions
- [ ] Redis caching
- [ ] JWT authentication
- [ ] Structured logging
- [ ] Graceful shutdown
- [ ] Rate limiting
- [ ] Docker deployment
- [ ] Health checks

---

---

<div align="center">

## 📖 Quick Reference

*Bookmark this for common patterns.*

</div>

### Project Setup

```bash
# Initialize module
go mod init github.com/username/myapp

# Add dependency
go get github.com/lib/pq

# Run
go run main.go

# Build
go build -o myapp

# Test
go test ./...
```

### HTTP Handler Template

```go
func (h *Handler) HandleRequest(w http.ResponseWriter, r *http.Request) {
    // 1. Parse request
    var req RequestStruct
    if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
        http.Error(w, "invalid JSON", http.StatusBadRequest)
        return
    }

    // 2. Validate
    if req.Field == "" {
        http.Error(w, "missing field", http.StatusBadRequest)
        return
    }

    // 3. Call service
    result, err := h.service.DoSomething(r.Context(), req)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    // 4. Return response
    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusOK)
    json.NewEncoder(w).Encode(result)
}
```

### Test Template

```go
func TestSomething(t *testing.T) {
    // 1. Setup
    mockRepo := &MockRepository{}
    service := NewService(mockRepo)

    // 2. Execute
    result, err := service.DoSomething()

    // 3. Assert
    if err != nil {
        t.Fatalf("Unexpected error: %v", err)
    }
    if result != expected {
        t.Errorf("Got %v, want %v", result, expected)
    }
}
```

### Middleware Template

```go
func MyMiddleware(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        // Before
        ctx := context.WithValue(r.Context(), "key", "value")
        r = r.WithContext(ctx)

        // Call next
        next.ServeHTTP(w, r)

        // After
    })
}
```

---

---

## My Go Journey

This handbook documents **is my real journey** learning Go for backend development.

**Remember:**
- Go values **simplicity** over complexity
- **Explicit error handling** catches bugs early
- **Interfaces** enable clean, testable code
- **Goroutines** make concurrency accessible
- **Context** ties it all together


---

**Last updated:** my learning is ongoing 