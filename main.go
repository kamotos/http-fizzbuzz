package main

import (
	"fmt"
	"net/http"

	"github.com/go-ozzo/ozzo-validation/v4"
	"github.com/labstack/echo"
)

// FizzBuzz serves as a request validation for the user's input
type FizzBuzz struct {
	Int1  int    `query:"int1" json:"int1"`
	Int2  int    `query:"int2" json:"int2"`
	Str1  string `query:"str1" json:"str1"`
	Str2  string `query:"str2" json:"str2"`
	Limit int    `query:"limit" json:"limit"`
}


// Validate user input
func (f FizzBuzz) Validate() error {
	return validation.ValidateStruct(&f,
		validation.Field(&f.Int1, validation.Required, validation.Min(1)),
		validation.Field(&f.Int2, validation.Required, validation.Min(1)),
		validation.Field(&f.Str1, validation.Required),
		validation.Field(&f.Str2, validation.Required),
		validation.Field(&f.Limit, validation.Required, validation.Min(1)),
	)
}

func main() {
	e := echo.New()
	e.GET("/", FizzBuzzHandler)
	e.Logger.Fatal(e.Start(":1323"))
}

func FizzBuzzHandler(c echo.Context) (err error) {
	f := new(FizzBuzz)
	// Error while binding request
	if err = c.Bind(f); err != nil {
		return
	}
	// Invalid request
	if err := f.Validate(); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	var result []string

	for i := 1; i < f.Limit+1; i++ {
		if i%f.Int1 == 0 && i%f.Int2 == 0 {
			result = append(result, f.Str1+f.Str2)
			continue
		}

		if i%f.Int1 == 0 {
			result = append(result, f.Str1)
			continue
		}

		if i%f.Int2 == 0 {
			result = append(result, f.Str2)
			continue
		}

		result = append(result, fmt.Sprintf("%d", i))
	}

	return c.JSON(http.StatusOK, result)
}
