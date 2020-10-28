package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"os/exec"

	"github.com/gin-gonic/gin"
)

func odooRestart(c *gin.Context) {
	cmd := exec.Command("systemctl", "restart", "biznavi")
	out, err := cmd.CombinedOutput()
	if err != nil {
		log.Fatalf("cmd.Run() failed with %s\n", err)
		c.String(http.StatusOK, fmt.Sprintf("error: %s", err))
	}
	fmt.Printf("combined out:\n%s\n", string(out))
	c.String(http.StatusOK, string(out))
}

func main() {
	r := gin.Default()

	gin.DisableConsoleColor()
	gin.SetMode(gin.ReleaseMode)

	// Logging to a file.
	f, _ := os.Create("log/odoo_cmd.log")
	gin.DefaultWriter = io.MultiWriter(f)

	r.GET("/odoo/restart", odooRestart)

	r.Run(":8888")
}
