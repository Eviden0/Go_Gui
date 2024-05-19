package test

import (
	"bytes"
	"fmt"
	"log"
	"os/exec"
	"testing"
)

func TestCommand(t *testing.T) {
    cmd := exec.Command("sh", "/Users/pixiao/WorkSpace/Tomcat/bin/shutdown.sh")
    var stdout, stderr bytes.Buffer
    cmd.Stdout = &stdout  // 标准输出
    cmd.Stderr = &stderr  // 标准错误
    err := cmd.Run()
    outStr, errStr := stdout.String(), stderr.String()
    fmt.Printf("out:\n%s\nerr:\n%s\n", outStr, errStr)
    if err != nil {
        log.Fatalf("cmd.Run() failed with %s\n", err)
    }
}
