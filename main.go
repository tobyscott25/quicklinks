package main

import (
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
	"gopkg.in/yaml.v3"
)

type Config struct {
	Redirects map[string]string `yaml:"redirects"`
}

func loadConfig(filePath string) (Config, error) {
	var config Config
	file, err := os.ReadFile(filePath)
	if err != nil {
		return config, err
	}
	err = yaml.Unmarshal(file, &config)
	return config, err
}

func main() {

	configFilePath := "redirects.yaml"
	config, err := loadConfig(configFilePath)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error loading config file: %v\n", err)
		os.Exit(1)
	}

	r := gin.Default()

	r.GET("/:path", func(c *gin.Context) {
		requestedPath := c.Param("path")
		normalizedPath := strings.ToLower(requestedPath)

		redirectURL, exists := config.Redirects[normalizedPath]
		if !exists {
			c.JSON(http.StatusNotFound, gin.H{
				"error": "Link not found",
			})
			return
		}

		c.Redirect(http.StatusFound, redirectURL)
	})

	r.Run()
}
