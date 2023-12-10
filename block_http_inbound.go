package main

import (
    “fmt”
    “os/exec”
)

func main() {
    // Block inbound HTTP traffic on port 80
    cmd := exec.Command(“iptables”, “-A”, “INPUT”, “-p”, “tcp”, “–dport”, “80”, “-j”, “DROP”)

    err := cmd.Run()
    if err != nil {
        fmt.Println("Error blocking inbound HTTP traffic:", err)
    } else {
        fmt.Println("Inbound HTTP traffic blocked.")
    }
}
