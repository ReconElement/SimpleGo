package main

import (
	"os"
	"fmt"
	"net/http"
	"github.com/gin-gonic/gin"
)

type Note struct{
	Title string `json: "title" binding: "required"`
	Content string `json: "content" binding: "required"`
}

func main() {
	r := gin.Default()
	r.POST("/post", writeIntoFile)
	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "hello world"})
	})
	r.Run(":8000")
}

func writeIntoFile(c *gin.Context) {
	var input Note
	if err:= c.ShouldBindJSON(&input);
	err!=nil{
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error":err.Error()});
	}
	fmt.Println(input.Title, input.Content);
	thing := fmt.Sprintf("\nTitle: %s, Content: %s",input.Title, input.Content);
	fileName := "test.txt"
	if doesFileExist(fileName){
		file, err:= os.OpenFile(fileName, os.O_WRONLY, 0644);
		if err!=nil{
			panic(err)
		}
		// file.WriteString(thing);
		createOrAppend(fileName, thing);
		defer file.Close();
	}else{
		file, err:= os.Create(fileName);
		if err!=nil{
			panic(err)
		}
		// file.WriteString(thing);
		createOrAppend(fileName, thing);
		defer file.Close();
	}
}

func doesFileExist(fileName string) bool{
	_,err:=os.Stat(fileName)
	return !os.IsNotExist(err);
}

func createOrAppend(fileName string, data string){
	file, err := os.OpenFile(fileName, os.O_WRONLY | os.O_APPEND, 0644);
	if err!=nil{
		panic(err)
	}else{
		file.WriteString(data)
	}
	defer file.Close();
}