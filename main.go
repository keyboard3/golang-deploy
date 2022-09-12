package main

import (
  "net/http"
  "github.com/gin-gonic/gin"
  "log"
  "os"
  "os/exec"
)
var (
  fileInfo *os.FileInfo
  err      error
)

func main() {
  r := gin.Default()
  r.GET("/deploy/keyboard3/:project", func(c *gin.Context) {
    name := c.Param("project")

    _, err := os.Stat("../"+name)
    if err != nil {
        if os.IsNotExist(err) {
            log.Println("project "+name+" not exist")
            cmd := exec.Command("/bin/sh", "-c", "cd .. && git clone https://github.com/keyboard3/"+name +" && ./deploy.sh")
            err :=cmd.Run()
            if err != nil {
              log.Printf("Command finished with error: %v", err)
              c.String(http.StatusOK, "Command finished with error: %v", err)
            }
        } else {
          c.String(http.StatusOK, "error: %v", err)
        }
    } else {
      log.Println("project "+name+" exist")
      cmd := exec.Command("/bin/sh", "-c", "cd ../"+name+" && git pull && kubectl delete deployment "+name+" && ./deploy.sh")
      err :=cmd.Run()
      if err != nil {
        log.Printf("Command finished with error: %v", err)
        c.String(http.StatusOK, "Command finished with error: %v", err)
      } else {
        c.String(http.StatusOK,"success");
      }
    }
  })
  r.Run(":3000") // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}