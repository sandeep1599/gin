package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)
type Book  struct{
	ID     int `json:"id"`
	Title  string `json:"title"`
	Author string `json:"author"`
}

var books []Book

func main (){

	router := gin.Default()

	router.GET("/ping", Pong)
	router.GET("/books",GetBooks)
	router.POST("/books",CreateBooks)

	fmt.Println("Starting the server on the port: 8080")
	err :=router.Run("localhost:8080")

	if err != nil{
		fmt.Printf("cannot start the sever: %v\n", err)
		return
	}
}


func Pong (c *gin.Context){
	c.JSON(http.StatusOK,"pong")
}

func GetBooks (c *gin.Context){
	c.JSON(http.StatusOK, books)
}

func CreateBooks( c *gin.Context){
	
	var newBook Book

	err := c.BindJSON(&newBook)

	if err != nil{
		c.JSON(http.StatusBadRequest,"cannot create Book")
		return
	}

	newBook.ID = len(books)+1

	books = append(books, newBook)

	c.JSON(http.StatusCreated,books)


}