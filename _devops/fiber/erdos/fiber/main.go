package main

import (
	"crypto/sha256"
	"crypto/subtle"
	"erdos/mathlogic"
	"erdos/apiresponse"
	"fmt"
	"os"
	"math"
// 	"strings"
    "strconv"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/keyauth/v2"
	log "github.com/sirupsen/logrus"
	// "github.com/gofiber/fiber/v2/middleware/basicauth"

    swagger "github.com/arsmn/fiber-swagger/v2"
//     _ "github.com/minusnaplus/api/docs"
)

var (
	apiKey = "public-key-123"
	// curl --cookie "token=public-key-123" http://erdos.localhos/v1/api/hello
    // curl --cookie "token=public-key-123" http://erdos.localhost/v1/api/add/?x=100&y=999
)

var (
	WarningLogger *log.Logger
	InfoLogger    *log.Logger
	ErrorLogger   *log.Logger
)

func addSecuredKeys() string {
	key := os.Getenv("SECRET_API_KEY")
// 	return strings.Split(keys, ":")
    return key
}

func parseInt64(str string) (int64, error) {
	val, err := strconv.ParseInt(str, 10, 64)
	if err != nil {
		return 0, fmt.Errorf("conversion error for %s: %v", str, err)
	}
	return val, nil
}

func handleQueryInt64(c *fiber.Ctx, queryParam string) (int64, error) {
	valStr := c.Query(queryParam)
	val, err := parseInt64(valStr)
	if err != nil {
		return 0, err
	}
	return val, nil
}

func validateAPIKey(c *fiber.Ctx, key string) (bool, error) {
	hashedAPIKey := sha256.Sum256([]byte(apiKey))
	hashedKey := sha256.Sum256([]byte(key))

	if subtle.ConstantTimeCompare(hashedAPIKey[:], hashedKey[:]) == 1 {
		return true, nil
	}
	log.WithFields(log.Fields{
		"key": key,
	}).Info("A wrong key")
	return false, keyauth.ErrMissingOrMalformedAPIKey
}
// @title Minus Plus API
// @version 1.0
// @description This is a sample swagger for Fiber
// @termsOfService http://swagger.io/terms/
// @contact.name API Support
// @contact.email fiber@swagger.io
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
// @host localhost:3001
// @BasePath /
func main() {
	app := fiber.New(fiber.Config{
		Prefork:       true,
		CaseSensitive: true,
		StrictRouting: true,
		ServerHeader:  "NestedGin",
		AppName:       "plusy-ujemne",
	})

	keys := addSecuredKeys()

	app.Get("/v1/api/hello", func(c *fiber.Ctx) error {
		return c.SendString("Hello fiber without api-key!" + keys)
	})

    app.Get("/swagger/*", swagger.HandlerDefault) // default

    app.Get("/swagger/*", swagger.New(swagger.Config{ // custom
        URL:         "http://example.com/doc.json",
        DeepLinking: false,
        // Expand ("list") or Collapse ("none") tag groups by default
        DocExpansion: "none",
        // Prefill OAuth ClientId on Authorize popup
        OAuth: &swagger.OAuthConfig{
            AppName:  "OAuth Provider",
            ClientId: "21bb4edc-05a7-4afc-86f1-2e151e4ba6e2",
        },
        // Ability to change OAuth2 redirect uri location
        OAuth2RedirectUrl: "http://localhost:3001/swagger/oauth2-redirect.html",
    }))

	app.Use(keyauth.New(keyauth.Config{
		KeyLookup: "cookie:token",
		Validator: validateAPIKey,
	}))

	app.Get("/v1/api/healthy", func(c *fiber.Ctx) error {
		resp := apiresponse.CreateApiResponse(true, "healthy", "fiber node")
		return c.JSON(resp)
	})

    // Add godoc
    // @Summary Add two numbers
    // @Description Add two numbers passed as query parameters
    // @ID add
    // @Tags math
    // @Accept json
    // @Produce json
    // @Param x query integer true "The first number"
    // @Param y query integer true "The second number"
    // @Success 200 {object} apiresponse.ApiResponse
    // @Failure 400 {object} apiresponse.ApiResponse
    // @Router /v1/api/add/ [get]

	app.Get("/v1/api/add/", func(c *fiber.Ctx) error {
		c.Accepts("application/json")
		valX, errX := apiresponse.ParseInt64(c.Query("x"))
		if errX != nil {
			resp := apiresponse.CreateApiResponse(false, "Conversion error for X: "+errX.Error(), nil)
			return c.Status(fiber.StatusBadRequest).JSON(resp)
		}
		valY, errY := apiresponse.ParseInt64(c.Query("y"))
		if errY != nil {
			resp := apiresponse.CreateApiResponse(false, "Conversion error for Y: "+errY.Error(), nil)
			return c.Status(fiber.StatusBadRequest).JSON(resp)
		}
		val := mathlogic.Add(valX, valY)
		resp := apiresponse.CreateApiResponse(true, "Success", val)
		return c.JSON(resp)
	})
	// Multiply godoc
    // @Summary Multiply two numbers
    // @Description Multiply two numbers passed as query parameters
    // @ID multiply
    // @Tags math
    // @Accept json
    // @Produce json
    // @Param x query string true "The first number to multiply"
    // @Param y query string true "The second number to multiply"
    // @Success 200 {object} apiresponse.ApiResponse
    // @Failure 400 {object} apiresponse.ApiResponse
    // @Router /v1/api/mul/ [get]

	app.Get("/v1/api/mul/", func(c *fiber.Ctx) error {
        c.Accepts("application/json")
        valX, errX := apiresponse.ParseBigInt(c.Query("x"))
        if errX != nil {
            resp := apiresponse.CreateApiResponse(false, "Conversion error for X: "+errX.Error(), nil)
            return c.Status(fiber.StatusBadRequest).JSON(resp)
        }
        valY, errY := apiresponse.ParseBigInt(c.Query("y"))
        if errY != nil {
            resp := apiresponse.CreateApiResponse(false, "Conversion error for Y: "+errY.Error(), nil)
            return c.Status(fiber.StatusBadRequest).JSON(resp)
        }
        val := mathlogic.Mul(valX, valY)
        resp := apiresponse.CreateApiResponse(true, "Success", val)
        return c.JSON(resp)
    })
    // Divide godoc
    // @Summary Divide two numbers
    // @Description Divide two numbers passed as query parameters
    // @ID divide
    // @Tags math
    // @Accept json
    // @Produce json
    // @Param x query string true "The numerator"
    // @Param y query string true "The denominator"
    // @Success 200 {object} apiresponse.ApiResponse
    // @Failure 400 {object} apiresponse.ApiResponse
    // @Failure 500 {object} apiresponse.ApiResponse
    // @Router /v1/api/div/ [get]


    app.Get("/v1/api/div/", func(c *fiber.Ctx) error {
        c.Accepts("application/json")
        valX, errX := strconv.ParseInt(c.Query("x"), 10, 64)
        if errX != nil {
            resp := apiresponse.CreateApiResponse(false, "Conversion error for X: "+errX.Error(), nil)
            return c.Status(fiber.StatusBadRequest).JSON(resp)
        }
        valY, errY := strconv.ParseInt(c.Query("y"), 10, 64)
        if errY != nil {
            resp := apiresponse.CreateApiResponse(false, "Conversion error for Y: "+errY.Error(), nil)
            return c.Status(fiber.StatusBadRequest).JSON(resp)
        }
        if valY == 0 {
            resp := apiresponse.CreateApiResponse(false, "Division by zero", nil)
            return c.Status(fiber.StatusBadRequest).JSON(resp)
        }
        val := mathlogic.Divide(valX, valY)
        resp := apiresponse.CreateApiResponse(true, "Success", val)
        return c.JSON(resp)
    })
    // Subtract godoc
    // @Summary Subtract two numbers
    // @Description Subtract the value of Y from the value of X
    // @ID subtract
    // @Tags arithmetic
    // @Accept json
    // @Produce json
    // @Param x query int true "The first number to subtract from (X)"
    // @Param y query int true "The second number to subtract (Y)"
    // @Success 200 {object} ApiResponse
    // @Failure 400 {object} ApiResponse
    // @Failure 500 {object} ApiResponse
    // @Router /v1/api/sub [get]


    app.Get("/v1/api/sub/", func(c *fiber.Ctx) error {
        c.Accepts("application/json")
        valX, errX := apiresponse.ParseBigInt(c.Query("x"))
        if errX != nil {
            resp := apiresponse.CreateApiResponse(false, "Conversion error for X: "+errX.Error(), nil)
            return c.Status(fiber.StatusBadRequest).JSON(resp)
        }
        valY, errY := apiresponse.ParseBigInt(c.Query("y"))
        if errY != nil {
            resp := apiresponse.CreateApiResponse(false, "Conversion error for Y: "+errY.Error(), nil)
            return c.Status(fiber.StatusBadRequest).JSON(resp)
        }
        val := mathlogic.Sub(valX, valY)
        resp := apiresponse.CreateApiResponse(true, "Success", val)
        return c.JSON(resp)
    })
    // GetE godoc
    // @Summary Get the mathematical constant e
    // @Description Get the mathematical constant e to the specified number of decimal places
    // @ID get-e
    // @Tags math
    // @Accept  json
    // @Produce  json
    // @Param decimal_places query int false "Number of decimal places to round e to (default is 2)" minimum(0) maximum(15)
    // @Success 200 {object} apiresponse
    // @Failure 400 {object} apiresponse
    // @Router /v1/api/e [get]

    app.Get("/v1/api/e", func(c *fiber.Ctx) error {
        e := math.E
        resp := apiresponse.CreateApiResponse(true, "Success", fmt.Sprintf("%v", e))
        return c.JSON(resp)
    })

    // GetPi godoc
    // @Summary Get the value of Pi
    // @Description Returns the value of the mathematical constant Pi
    // @ID get-pi
    // @Tags math
    // @Accept

    app.Get("/v1/api/pi", func(c *fiber.Ctx) error {
        pi := math.Pi
        resp := apiresponse.CreateApiResponse(true, "Success", fmt.Sprintf("%v", pi))
        return c.JSON(resp)
    })

	log.Fatal(app.Listen(":3001"))
}
