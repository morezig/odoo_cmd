package main

import (
	"fmt"
	"log"
	"net/http"
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

	r.GET("/odoo/restart", odooRestart)

	r.Run()
}
