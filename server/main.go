package main

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type ChatMessage struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

type ChatRequest struct {
	Messages []ChatMessage `json:"messages"`
}

type ChatResponse struct {
	ID      string    `json:"id"`
	Object  string    `json:"object"`
	Created int64     `json:"created"`
	Model   string    `json:"model"`
	Choices []Choice  `json:"choices"`
	Usage   Usage     `json:"usage"`
}

type Choice struct {
	Index   int     `json:"index"`
	Message Message `json:"message"`
	Delta   Message `json:"delta"`
}

type Message struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

type Usage struct {
	PromptTokens     int `json:"prompt_tokens"`
	CompletionTokens int `json:"completion_tokens"`
	TotalTokens      int `json:"total_tokens"`
}

func main() {
	e := echo.New()

	// CORS middleware
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"http://localhost:3000"},
		AllowMethods: []string{http.MethodGet, http.MethodPost, http.MethodOptions},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept},
	}))

	// Routes
	e.POST("/api/chat", handleChat)
	e.POST("/api/chat/stream", handleChatStream)

	// Start server
	fmt.Println("Server starting on :8080")
	e.Logger.Fatal(e.Start(":8080"))
}

func handleChat(c echo.Context) error {
	var req ChatRequest
	if err := c.Bind(&req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid request")
	}

	// Mock response
	response := ChatResponse{
		ID:      "chatcmpl-" + generateID(),
		Object:  "chat.completion",
		Created: time.Now().Unix(),
		Model:   "gpt-3.5-turbo",
		Choices: []Choice{
			{
				Index: 0,
				Message: Message{
					Role:    "assistant",
					Content: generateMockResponse(req.Messages),
				},
			},
		},
		Usage: Usage{
			PromptTokens:     100,
			CompletionTokens: 50,
			TotalTokens:      150,
		},
	}

	return c.JSON(http.StatusOK, response)
}

func handleChatStream(c echo.Context) error {
	var req ChatRequest
	if err := c.Bind(&req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid request")
	}

	// Set headers for SSE
	c.Response().Header().Set("Content-Type", "text/event-stream")
	c.Response().Header().Set("Cache-Control", "no-cache")
	c.Response().Header().Set("Connection", "keep-alive")
	c.Response().Header().Set("Access-Control-Allow-Origin", "*")
	c.Response().Header().Set("Access-Control-Allow-Headers", "Cache-Control")

	// Get the response content
	content := generateMockResponse(req.Messages)
	words := splitIntoWords(content)

	// Stream the response
	for i, word := range words {
		choice := Choice{
			Index: 0,
			Delta: Message{
				Role:    "assistant",
				Content: word + " ",
			},
		}

		response := ChatResponse{
			ID:      "chatcmpl-" + generateID(),
			Object:  "chat.completion.chunk",
			Created: time.Now().Unix(),
			Model:   "gpt-3.5-turbo",
			Choices: []Choice{choice},
		}

		data, _ := json.Marshal(response)
		fmt.Fprintf(c.Response().Writer, "data: %s\n\n", data)
		c.Response().Flush()

		// Simulate typing delay
		time.Sleep(50 * time.Millisecond)

		// Send done event at the end
		if i == len(words)-1 {
			doneEvent := map[string]interface{}{
				"id":      "chatcmpl-" + generateID(),
				"object":  "chat.completion.chunk",
				"created": time.Now().Unix(),
				"model":   "gpt-3.5-turbo",
				"choices": []map[string]interface{}{
					{
						"index": 0,
						"delta": map[string]interface{}{},
					},
				},
			}
			doneData, _ := json.Marshal(doneEvent)
			fmt.Fprintf(c.Response().Writer, "data: %s\n\n", doneData)
			c.Response().Flush()
		}
	}

	return nil
}

func generateMockResponse(messages []ChatMessage) string {
	responses := []string{
		"こんにちは！何かお手伝いできることはありますか？",
		"興味深い質問ですね。詳しく説明させていただきます。興味深い質問ですね。詳しく説明させていただきます。興味深い質問ですね。詳しく説明させていただきます。興味深い質問ですね。詳しく説明させていただきます。興味深い質問ですね。詳しく説明させていただきます。興味深い質問ですね。詳しく説明させていただきます。興味深い質問ですね。詳しく説明させていただきます。",
		"それは良いアイデアだと思います。さらに詳しく話してみましょう。それは良いアイデアだと思います。さらに詳しく話してみましょう。それは良いアイデアだと思います。さらに詳しく話してみましょう。それは良いアイデアだと思います。さらに詳しく話してみましょう。それは良いアイデアだと思います。さらに詳しく話してみましょう。それは良いアイデアだと思います。さらに詳しく話してみましょう。",
		"確かに、その通りですね。他にも考えられる方法があります。確かに、その通りですね。他にも考えられる方法があります。確かに、その通りですね。他にも考えられる方法があります。確かに、その通りですね。他にも考えられる方法があります。確かに、その通りですね。他にも考えられる方法があります。確かに、その通りですね。他にも考えられる方法があります。",
		"素晴らしい質問です！これについて詳しく説明します。素晴らしい質問です！これについて詳しく説明します。素晴らしい質問です！これについて詳しく説明します。素晴らしい質問です！これについて詳しく説明します。素晴らしい質問です！これについて詳しく説明します。素晴らしい質問です！これについて詳しく説明します。",
		"なるほど、理解しました。それについて詳しくお話ししましょう。なるほど、理解しました。それについて詳しくお話ししましょう。なるほど、理解しました。それについて詳しくお話ししましょう。なるほど、理解しました。それについて詳しくお話ししましょう。なるほど、理解しました。それについて詳しくお話ししましょう。なるほど、理解しました。それについて詳しくお話ししましょう。",
		"とても良いポイントですね。これについて詳しく説明させていただきます。とても良いポイントですね。これについて詳しく説明させていただきます。とても良いポイントですね。これについて詳しく説明させていただきます。とても良いポイントですね。これについて詳しく説明させていただきます。とても良いポイントですね。これについて詳しく説明させていただきます。とても良いポイントですね。これについて詳しく説明させていただきます。",
		"確かに、その観点は重要です。さらに詳しく検討してみましょう。確かに、その観点は重要です。さらに詳しく検討してみましょう。確かに、その観点は重要です。さらに詳しく検討してみましょう。確かに、その観点は重要です。さらに詳しく検討してみましょう。確かに、その観点は重要です。さらに詳しく検討してみましょう。確かに、その観点は重要です。さらに詳しく検討してみましょう。",
		"興味深い視点ですね。これについて詳しく説明します。興味深い視点ですね。これについて詳しく説明します。興味深い視点ですね。これについて詳しく説明します。興味深い視点ですね。これについて詳しく説明します。興味深い視点ですね。これについて詳しく説明します。興味深い視点ですね。これについて詳しく説明します。興味深い視点ですね。これについて詳しく説明します。",
		"素晴らしいアイデアです！これについて詳しく話してみましょう。素晴らしいアイデアです！これについて詳しく話してみましょう。素晴らしいアイデアです！これについて詳しく話してみましょう。素晴らしいアイデアです！これについて詳しく話してみましょう。素晴らしいアイデアです！これについて詳しく話してみましょう。素晴らしいアイデアです！これについて詳しく話してみましょう。",
	}

	return responses[rand.Intn(len(responses))]
}

func generateID() string {
	const charset = "abcdefghijklmnopqrstuvwxyz0123456789"
	b := make([]byte, 29)
	for i := range b {
		b[i] = charset[rand.Intn(len(charset))]
	}
	return string(b)
}

func splitIntoWords(text string) []string {
	// Simple word splitting for Japanese text
	var words []string
	var currentWord string
	
	for _, char := range text {
		if char == ' ' || char == '。' || char == '、' || char == '！' || char == '？' {
			if currentWord != "" {
				words = append(words, currentWord)
				currentWord = ""
			}
			words = append(words, string(char))
		} else {
			currentWord += string(char)
		}
	}
	
	if currentWord != "" {
		words = append(words, currentWord)
	}
	
	return words
} 