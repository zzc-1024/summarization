package main

import (
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hello, World!",
		})
	})
	r.POST("/upload", func(c *gin.Context) {
		// 获取名为 "file" 的上传文件
		file, err := c.FormFile("file")
		if err != nil {
			log.Println("Failed to get file:", err)
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}
		
		// 保存上传的文件
		if err := c.SaveUploadedFile(file, file.Filename); err != nil {
			log.Println("Failed to save file:", err)
			c.JSON(500, gin.H{"error": "Failed to save file"})
			return
		}

		c.JSON(200, gin.H{
			"message":  "File uploaded successfully",
			"filename": file.Filename,
		})
	})

	r.Run(":8000")
}
