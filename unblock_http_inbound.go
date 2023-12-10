package main

import (
    “fmt”
    “os/exec”
)

func main() {
    // Unblock inbound HTTP traffic on port 80
    cmd := exec.Command(“iptables”, “-D”, “INPUT”, “-p”, “tcp”, “–dport”, “80”, “-j”, “DROP”)

    err := cmd.Run()
    if err != nil {
        fmt.Println("Error unblocking inbound HTTP traffic:", err)
    } else {
        fmt.Println("Inbound HTTP traffic unblocked.")
    }
}
