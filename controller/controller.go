package controller

import (
	"fmt"
	"math"
	"net/http"

	"github.com/Mubashir01234/echoPostCalculator/database"
	"github.com/labstack/echo/v4"
)

// var db *sql.DB
type Numbers struct {
	Number1 float64 `json:"number1"`
	Number2 float64 `json:"number2"`
}
type Response struct {
	Result float64 `json:"result"`
}

type GetRecordResponse struct{
	Id 			int 	`json:"id"`
	Number1 	float64 `json:"number1"`
	Number2 	float64 `json:"number2"`
	Operation 	string 	`json:"operation"`
	Result 		float64 `json:"result"`
	CreatedAt 	string 	`json:"createdAt"`
}
type GetAllRecordResponse []struct{
	getAllrecord []GetRecordResponse `json:"allRecord"`
}
func (number Numbers) Connect(result float64, operation string) {
	db := database.Conc()
    defer db.Close()
    sql := "INSERT INTO calculate(number1, number2, operation, result) VALUES( ?, ?,?, ?)"
    stmt, err := db.Prepare(sql)
    if err != nil {
        fmt.Print(err.Error())
    }
    defer stmt.Close()
    _, err2 := stmt.Exec(number.Number1, number.Number2, operation, result)
    if err2 != nil {
        panic(err2)
    }
}
func Addition(c echo.Context) error {
	n := new(Numbers)
	if err := c.Bind(n); err != nil {
		return err
	}
	fmt.Println(n)
	add := n.Number1 + n.Number2
	result := Response{
		add,
	}
	n.Connect(add, "+")

	return c.JSON(http.StatusOK, result)
	// return c.String(http.StatusOK, "Number1 + Number2 = "+strconv.Itoa(add))
}
func Subtraction(c echo.Context) error {
	n := new(Numbers)
	if err := c.Bind(n); err != nil {
		return err
	}
	fmt.Println(n)
	sub := n.Number1 - n.Number2
	result := Response{
		sub,
	}
	n.Connect(sub, "-")
	return c.JSON(http.StatusOK, result)
}
func Multiplication(c echo.Context) error {
	n := new(Numbers)
	if err := c.Bind(n); err != nil {
		return err
	}
	fmt.Println(n)
	mul := n.Number1 * n.Number2
	result := Response{
		mul,
	}
	n.Connect(mul, "*")
	return c.JSON(http.StatusOK, result)
}
func Division(c echo.Context) error {
	n := new(Numbers)
	if err := c.Bind(n); err != nil {
		return err
	}
	fmt.Println(n)
	div := n.Number1 / n.Number2
	result := Response{
		div,
	}
	n.Connect(div, "/")
	return c.JSON(http.StatusOK, result)
}
func Modulus(c echo.Context) error {
	n := new(Numbers)
	if err := c.Bind(n); err != nil {
		return err
	}
	fmt.Println(n)
	mod := int(n.Number1) % int(n.Number2)
	result := Response{
		float64(mod),
	}
	n.Connect(float64(mod), "%")
	return c.JSON(http.StatusOK, result)
}
func Power(c echo.Context) error {
	n := new(Numbers)
	if err := c.Bind(n); err != nil {
		return err
	}
	fmt.Println(n)
	power := math.Pow(n.Number1, n.Number2)
	result := Response{
		power,
	}
	n.Connect(power, "^")
	return c.JSON(http.StatusOK, result)
}
func Square(c echo.Context) error {
	n := new(Numbers)
	if err := c.Bind(n); err != nil {
		return err
	}
	fmt.Println(n)
	squ := n.Number1 * n.Number1
	result := Response{
		squ,
	}
	n.Connect(squ, "^2")
	return c.JSON(http.StatusOK, result)
}
func SquareRoot(c echo.Context) error {
	n := new(Numbers)
	if err := c.Bind(n); err != nil {
		return err
	}
	fmt.Println(n)
	sqrt := math.Sqrt(n.Number1)
	result := Response{
		sqrt,
	}
	n.Connect(sqrt, "???")
	return c.JSON(http.StatusOK, result)
}
func GetRecord(c echo.Context) error{
	fmt.Println("HEADER")
	fmt.Println(c.Request().Header["Authorization"])
	db := database.Conc()
    defer db.Close()
	requestedId := c.Param("id")
	fmt.Println(requestedId)
	var id int
	var number1, number2, result float64
	var operation, createdAt string
	err :=db.QueryRow("select * from calculate where ID = ?", requestedId).Scan(&id, &number1, &number2, &operation, &result, &createdAt)
	if err!=nil{
		fmt.Print(err.Error())
	}
	res:=GetRecordResponse{
		Id: id,
		Number1: number1,
		Number2: number2,
		Operation: operation,
		Result: result,
		CreatedAt: createdAt,
		
	}
	return c.JSON(http.StatusOK, res)
}
func GetAllRecord(c echo.Context) error{
	db:= database.Conc()
	defer db.Close()
	response:=make([]GetRecordResponse,0)
	var id int
	var number1, number2, result float64
	var operation, createdAt string
	rows, err :=db.Query("SELECT * FROM calculate")
	if err!=nil{
		fmt.Print(err.Error())
	}
	defer rows.Close()
	for rows.Next(){
		err=rows.Scan(&id, &number1, &number2, &operation, &result, &createdAt)
		if err!=nil{
			panic(err)
		}
		res:=GetRecordResponse{
			Id: id,
			Number1: number1,
			Number2: number2,
			Operation: operation,
			Result: result,
			CreatedAt: createdAt,
		}
		response= append(response,res)
	}
	if err := rows.Err(); err != nil {
		panic(err)
	}
	return c.JSON(http.StatusOK, response)
}
func DeleteRecord(c echo.Context) error{
	db:= database.Conc()
	defer db.Close()
	requestedId := c.Param("id")
	fmt.Println(requestedId)
	_,err :=db.Query("DELETE from calculate where ID = ?", requestedId)
	if err!=nil{
		fmt.Print(err.Error())
	}
	// res:=ResDel{
	// 	Message: "ID information is deleted",	
	// }
	return c.JSON(http.StatusOK, "Success")
}