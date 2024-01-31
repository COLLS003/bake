package database

import (
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

// Open the database and establish the connection
func Init() *gorm.DB {
	// Specify connection properties.
	dsn := "user=jwsqrlln password=7n6ooKGLIH3SDEdrdSrqehDP84u6ZSgf dbname=jwsqrlln host=hattie.db.elephantsql.com port=5432 sslmode=require"

	// Open the database connection.
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}

	DB = db
	return DB
}

// Function to get the database connection
func GetConnection() *gorm.DB {
	return DB
}

// Binder
func Bind(c *gin.Context, object interface{}) error {
	binder := binding.Default(c.Request.Method, c.ContentType())
	return c.ShouldBindWith(object, binder)
}

// Return customized error info
// Helps return customized error info
type CommonError struct {
	Errors map[string]interface{} `json:"errors"`
}

// Validators
func NewValidatorError(err error) CommonError {
	res := CommonError{}
	res.Errors = make(map[string]interface{})
	errs := err.(validator.ValidationErrors)
	for _, v := range errs {
		res1 := fmt.Sprintf("{%v: %v}", v.Tag(), v.Param())
		fmt.Println(res1)
		if v.Param() != "" {
			res.Errors[v.Field()] = fmt.Sprintf("{%v: %v}", v.Tag(), v.Param())
		} else {
			res.Errors[v.Field()] = fmt.Sprintf("{key: %v}", v.Tag())
		}
	}
	return res
}

// Wrap the error info in an object
func NewError(key string, err error) CommonError {
	res := CommonError{}
	res.Errors = make(map[string]interface{})
	res.Errors[key] = err.Error()
	return res
}
