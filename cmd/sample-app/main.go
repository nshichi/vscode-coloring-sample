package main

/* VSCode の色分け修正サンプルコード */

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()

	e.Static("/", "./html")

	// e.File("/", "./html/index.html")

	fmt.Printf("hello\n")
	e.Start(":3000")
}

func GET_hello(c echo.Context) error {
	return c.String(http.StatusOK, "hello")
}

type CNo string
type GKey string

type BillRec struct {
	CNo  CNo
	GKey GKey
	// ...
}

func sample(bills []BillRec) {
	const N = 1_000

	cnoset := make(map[CNo]struct{}, 0)
	clistByGKey := make(map[GKey][]CNo)

	for _, b := range bills {
		if _, ok := cnoset[b.CNo]; ok {

			clistByGKey[b.GKey] = append(clistByGKey[b.GKey], b.CNo)
		}
	}
}
