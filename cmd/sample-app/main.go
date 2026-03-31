package main

/* VSCode の色分けサンプルコード */
import (
	"net/http"

	"github.com/labstack/echo/v4"
)

const N = 1_000

func main() {
	e := echo.New()
	e.Static("/", "./html") // e.File("/", "./html/index.html")
	e.Start(":3000")
}

func GET_hello(c echo.Context) error {
	return c.String(http.StatusOK, "hello\n")
}

type CNo string
type GKey string

type BillRec struct {
	CNo  CNo
	GKey GKey
	// ...
}

func sample(bills []BillRec) {
	clistByGKey := make(map[GKey][]CNo)
	for _, b := range bills {
		clistByGKey[b.GKey] = append(clistByGKey[b.GKey], b.CNo)
	}
}
