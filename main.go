package main

import (
	"fmt"
	"log"
	"net/http"
	"os/exec"

	"github.com/gin-gonic/gin"

	"github.com/coreos/go-systemd/daemon"
)

func escpRestart(c *gin.Context) {
	cmd := exec.Command("sudo", "systemctl", "restart", "mz_escp")
	out, err := cmd.CombinedOutput()
	if err != nil {
		log.Fatalf("cmd.Run() failed with %s\n", err)
		c.String(http.StatusOK, fmt.Sprintf("error: %s", err))
	}
	fmt.Printf("combined out:\n%s\n", string(out))
	c.HTML(http.StatusOK, "done.html", gin.H{
		"msg": "重啟印表機成功",
	})
	// c.Writer.WriteHeader(200)
	// idx, _ := bindata.Asset("assets/done.html")
	// c.Writer.Write(idx)
	// c.Writer.Header().Add("Accept", "text/html")
	// c.Writer.Flush()
}

func odooRestart(c *gin.Context) {
	cmd := exec.Command("sudo", "systemctl", "restart", "biznavi")
	out, err := cmd.CombinedOutput()
	if err != nil {
		log.Fatalf("cmd.Run() failed with %s\n", err)
		c.String(http.StatusOK, fmt.Sprintf("error: %s", err))
	}
	fmt.Printf("combined out:\n%s\n", string(out))
	c.HTML(http.StatusOK, "done.html", gin.H{
		"msg": "重啟伺服器成功",
	})
	// c.Writer.WriteHeader(200)
	// idx, _ := bindata.Asset("assets/done.html")
	// c.Writer.Write(idx)
	// c.Writer.Header().Add("Accept", "text/html")
	// c.Writer.Flush()
}

func gitPull(c *gin.Context) {
	cmd := exec.Command("git", "-C", "/opt/biznavi", "pull")
	out, err := cmd.CombinedOutput()
	if err != nil {
		log.Fatalf("cmd.Run() failed with %s\n", err)
		c.String(http.StatusOK, fmt.Sprintf("error: %s", err))
	}
	fmt.Printf("combined out:\n%s\n", string(out))
	c.HTML(http.StatusOK, "done.html", gin.H{
		"msg": "拉版成功",
	})
	// c.Writer.WriteHeader(200)
	// idx, _ := bindata.Asset("assets/done.html")
	// c.Writer.Write(idx)
	// c.Writer.Header().Add("Accept", "text/html")
	// c.Writer.Flush()
}

func odoo(c *gin.Context) {
	pwd := c.PostForm("pwd")
	srv_act := c.PostForm("srv_act")
	fmt.Println("pwd:", pwd)
	fmt.Println("srv_act:", srv_act)
	if pwd == "sf17BY23" {
		switch srv_act {
		case "escp":
			escpRestart(c)
		case "odoo":
			odooRestart(c)
		case "git":
			gitPull(c)
		default: //default:當前面條件都沒有滿足時將會執行此處內包含的方法
			fmt.Println("no action")
			// c.String(200, "No Action........")
			c.HTML(http.StatusOK, "error.html", gin.H{
				"msg": "No Action........",
			})
		}
	} else {
		// c.String(200, "XXXXX密碼錯誤XXXXX")
		c.HTML(http.StatusOK, "error.html", gin.H{
			"msg": "XXXXX密碼錯誤XXXXX",
		})
	}

}

func index(c *gin.Context) {
	// c.Writer.WriteHeader(200)
	// idx, _ := bindata.Asset("assets/index.html")
	// c.Writer.Write(idx)
	// c.Writer.Header().Add("Accept", "text/html")
	// c.Writer.Flush()
	c.HTML(http.StatusOK, "index.html", nil)
}

// go-bindata-assetfs.exe  -o bindata/bind.go -pkg=bindata assets/...
// go build
func main() {
	// gin.DisableConsoleColor()
	// gin.SetMode(gin.ReleaseMode)
	// Logging to a file.
	// f, _ := os.Create("log/odoo_cmd.log")
	// gin.DefaultWriter = io.MultiWriter(f)

	r := gin.Default()
	r.LoadHTMLGlob("assets/*.html")
	r.Static("/js", "./assets/js")
	r.Static("/css", "./assets/css")
	// fs := assetfs.AssetFS{Asset: bindata.Asset, AssetDir: bindata.AssetDir, AssetInfo: bindata.AssetInfo, Prefix: "assets", Fallback: "index.html"}
	// r.StaticFS("/css", &fs)
	// r.StaticFS("/js", &fs)

	r.GET("/", index)
	r.POST("/odoo/action", odoo)
	// r.GET("/odoo/restart", odooRestart)
	// r.GET("/odoo/gitpull", gitPull)
	// r.GET("/odoo/escp", escpRestart)

	daemon.SdNotify(false, "REEADY=1")

	r.Run(":3980")
}
