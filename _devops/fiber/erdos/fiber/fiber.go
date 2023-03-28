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

	app.Use(keyauth.New(keyauth.Config{
		KeyLookup: "cookie:token",
		Validator: validateAPIKey,
	}))

	app.Get("/v1/api/healthy", func(c *fiber.Ctx) error {
		resp := apiresponse.CreateApiResponse(true, "healthy", "fiber node")
		return c.JSON(resp)
	})

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
    app.Get("/v1/api/e", func(c *fiber.Ctx) error {
        e := math.E
        resp := apiresponse.CreateApiResponse(true, "Success", fmt.Sprintf("%v", e))
        return c.JSON(resp)
    })

    app.Get("/v1/api/pi", func(c *fiber.Ctx) error {
        pi := math.Pi
        resp := apiresponse.CreateApiResponse(true, "Success", fmt.Sprintf("%v", pi))
        return c.JSON(resp)
    })

	log.Fatal(app.Listen(":3001"))
}
