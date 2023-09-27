package handlers

import (
	"math/rand"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type HelloWorld struct {
	Msg string `json:"message"`
}

type Foo struct {
	UUID string `json:"uuid"`
	Bar  int    `json:"bar"`
}

type BarsJSONResponse struct {
	Bars []Foo `json:"bars"`
}

type BarJSONResponse struct {
	Bar int `json:"bar"`
}

var bars []Foo

func Greetings(c echo.Context) error {
	loadTestBars()
	return c.JSON(http.StatusOK, HelloWorld{
		Msg: "Hello World!!",
	})
}

func loadTestBars() {
	for i := 1; i < 6; i++ {
		tb := Foo{UUID: strconv.Itoa(i), Bar: rand.Intn(100)}
		bars = append(bars, tb)
	}
}

func GetBars(c echo.Context) error {
	if len(bars) < 1 {
		return c.JSON(http.StatusNotFound, "No bars found")
	} else {
		op := BarsJSONResponse{
			Bars: bars,
		}
		res := make(map[string]any)
		res["status"] = http.StatusOK
		res["data"] = op

		return c.JSON(http.StatusOK, res)
	}
}

func contains(requestedUuid string) (bool, int) {
	for index, bar := range bars {
		if bar.UUID == requestedUuid {
			return true, index
		}
	}
	return false, -1
}

func GetBarsId(c echo.Context) error {
	if len(bars) < 1 {
		return c.JSON(http.StatusNotFound, "No bars found")
	} else {
		id := c.Param("uuid")
		var barExist, barIndex = contains(string(id))

		if barExist {
			op := BarJSONResponse{
				Bar: bars[barIndex].Bar,
			}

			res := make(map[string]any)
			res["status"] = http.StatusOK
			res["data"] = op
			return c.JSON(http.StatusOK, res)
		} else {
			res := make(map[string]any)
			res["status"] = http.StatusNotFound
			res["data"] = "Bar not found"
			res["check"] = id
			return c.JSON(http.StatusOK, res)
		}

	}
}

func GetBarsSum(c echo.Context) error {
	sum := 0
	if len(bars) < 1 {
		return c.JSON(http.StatusNotFound, "No bars found")
	} else {
		for i := 0; i < len(bars); i++ {
			sum += bars[i].Bar
		}

		res := make(map[string]any)
		res["status"] = http.StatusOK
		res["sum"] = sum

		return c.JSON(http.StatusOK, res)
	}
}

func DeleteBarsId(c echo.Context) error {
	id := c.Param("uuid")
	var barexist, barindex = contains(string(id))

	if barexist {
		bars = append(bars[:barindex], bars[barindex+1:]...)
		res := make(map[string]any)
		res["status"] = http.StatusOK
		res["data"] = "bar found and deleted"
		res["check"] = id

		return c.JSON(http.StatusOK, res)

	} else {
		res := make(map[string]any)
		res["status"] = http.StatusNotFound
		res["data"] = "BAR not found"
		res["check"] = id
		return c.JSON(http.StatusOK, res)

	}
}

func CreateNewBar(c echo.Context) error {
	u := &Foo{
		UUID: strconv.Itoa(len(bars) + 1),
	}
	err := c.Bind(u)

	if err != nil {
		return c.JSON(http.StatusBadRequest, "The json body is empty,try again!!")
	}

	bars = append(bars, *u)
	return c.JSON(http.StatusOK, u)
}
