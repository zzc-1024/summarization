package main

import (
	"log"
	"os"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/parakeet-nest/parakeet/completion"
	"github.com/parakeet-nest/parakeet/embeddings"
	"github.com/parakeet-nest/parakeet/enums/option"
	"github.com/parakeet-nest/parakeet/llm"
)

func main() {
	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hello, World!",
		})
	})

	r.GET("/task", func(c *gin.Context) {
		// 获取uploads文件夹下文件夹的数量
		entries, err := os.ReadDir("uploads")
		if err != nil {
			log.Println("Failed to get tast count:", err)
			return
		}
		// 返回所有文件夹名称
		var folderNames []string
		for _, entry := range entries {
			folderNames = append(folderNames, entry.Name())
		}
		c.JSON(200, gin.H{
			"folderNames": folderNames,
		})
	})

	r.GET("/task/:id", func(c *gin.Context) {
		// 获取id参数
		id := c.Param("id")

		// 获取该id下的文件名称
		files, err := os.ReadDir("uploads/" + id)
		if err != nil {
			log.Println("Failed to get files:", err)
			c.JSON(500, gin.H{"error": "Failed to get files"})
			return
		}

		// 获取文件名称
		var filenames []string
		for _, file := range files {
			filenames = append(filenames, file.Name())
		}

		// 返回文件名称
		c.JSON(200, gin.H{
			"filenames": filenames,
		})
	})

	r.POST("/task", func(c *gin.Context) {
		// 获取上传id
		uploadId := c.PostForm("uploadId")

		// 创建uploads文件夹
		err := os.Mkdir("uploads/"+uploadId, 0755)
		if err != nil {
			log.Println("Failed to create uploads folder:", err)
			c.JSON(500, gin.H{"error": "Failed to create uploads folder"})
			return
		}
	})

	r.GET("/task/count", func(c *gin.Context) {
		// 获取uploads文件夹下文件夹的数量
		entries, err := os.ReadDir("uploads")
		if err != nil {
			log.Println("Failed to get tast count:", err)
			return
		}
		c.JSON(200, gin.H{"count": len(entries)})
	})

	r.POST("/task/file", func(c *gin.Context) {
		// 获取名为 "file" 的上传文件
		file, err := c.FormFile("file")
		if err != nil {
			log.Println("Failed to get file:", err)
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}

		// 获取上传id
		uploadId := c.PostForm("taskId")

		// 保存上传的文件
		if err := c.SaveUploadedFile(file, "uploads/"+uploadId+"/"+file.Filename); err != nil {
			log.Println("Failed to save file:", err)
			c.JSON(500, gin.H{"error": "Failed to save file"})
			return
		}

		c.JSON(200, gin.H{
			"message":  "File uploaded successfully",
			"filename": file.Filename,
		})
	})

	r.POST("/task/:id/abstract", func(c *gin.Context) {
		// 获取任务id
		id := c.Param("id")
		// 读取文件夹下所有的文件内容
		files, err := os.ReadDir("uploads/" + id)
		if err != nil {
			log.Println("Failed to get files:", err)
			c.JSON(500, gin.H{"error": "Failed to get files"})
			return
		}

		// 读取文件内容
		var fileContents []string
		for _, file := range files {
			content, err := os.ReadFile("uploads/" + id + "/" + file.Name())
			if err != nil {
				log.Println("Failed to read file:", err)
				c.JSON(500, gin.H{"error": "Failed to read file"})
				return
			}
			fileContents = append(fileContents, string(content))
		}
		ollamaUrl := "http://localhost:11434"
		embeddingsModel := "bge-m3:latest"
		smallChatModel := "qwen2.5:latest"
		store := embeddings.MemoryVectorStore{
			Records: make(map[string]llm.VectorRecord),
		}
		// Create embeddings from documents and save them in the store
		for idx, doc := range fileContents {
			log.Println("Creating embedding from document ", idx)
			embedding, err := embeddings.CreateEmbedding(
				ollamaUrl,
				llm.Query4Embedding{
					Model:  embeddingsModel,
					Prompt: doc,
				},
				strconv.Itoa(idx),
			)
			if err != nil {
				log.Println("😡:", err)
			} else {
				store.Save(embedding)
			}
		}
		userContent := `这篇文章讲了什么？`

		systemContent := `你是一个摘要机器人，你需要按照用户的要求，从给定的文本中提取出摘要。`
		// Create an embedding from the question
		embeddingFromQuestion, err := embeddings.CreateEmbedding(
			ollamaUrl,
			llm.Query4Embedding{
				Model:  embeddingsModel,
				Prompt: userContent,
			},
			"question",
		)
		if err != nil {
			log.Fatalln("😡:", err)
		}

		//🔎 searching for similarity...
		similarity, _ := store.SearchMaxSimilarity(embeddingFromQuestion)

		documentsContent := `<context><doc>` + similarity.Prompt + `</doc></context>`

		query := llm.Query{
			Model: smallChatModel,
			Messages: []llm.Message{
				{Role: "system", Content: systemContent},
				{Role: "system", Content: documentsContent},
				{Role: "user", Content: userContent},
			},
			Options: llm.SetOptions(map[string]interface{}{
				option.Temperature: 0.0,
			}),
		}

		log.Println(query)
		log.Println("🤖 answer:")

		// Answer the question
		answer, err := completion.Chat(ollamaUrl, query)
		if err != nil {
			log.Fatalln("😡:", err)
		}
		c.JSON(200, gin.H{
			"answer": answer,
		})
	})

	r.Run(":8000")
}
