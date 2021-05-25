# JSON Web Toke Demo Using Gin And jwt-go

A demo jwt project using:
* github.com/gin-gonic/gin v1.7.2
* github.com/dgrijalva/jwt-go v3.2.0+incompatible

## Usage

### Install dependency
```shell
go mod tidy
```

### Setup jwt key
Set up jwt key in code file `controller/login`
```go
var JwtKey = []byte("gin-jwt")  // use your own key to replace "gin-jwt"
```

### Run project
Use `go run .` or debug in Vscode

### Structure Description

#### Function for signing and verifying jwt

utils/auth.go
```go
\\ Use to sign jwt
func SignJwt(key []byte, user service.User) string


\\ Use to validate jwt
func ValidJwt(tokenString string, key []byte) (user UserInfo, err error)

```
#### Middleware for validate jwt

middleware/auth.go
```go
func Auth() gin.HandlerFunc
```