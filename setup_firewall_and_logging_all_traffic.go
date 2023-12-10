package main

import (
    “fmt”
    “os/exec”
)

func main() {
    // Allow inbound HTTP traffic on port 80
    allowHTTPCmd := exec.Command(“iptables”, “-A”, “INPUT”, “-p”, “tcp”, “–dport”, “80”, “-j”, “ACCEPT”)
    if err := allowHTTPCmd.Run(); err != nil {
        fmt.Println(“Error allowing inbound HTTP traffic:”, err)
        return
    }
    fmt.Println(“Inbound HTTP traffic allowed on port 80.”)

    // Log all incoming traffic
    logInputCmd := exec.Command("iptables", "-A", "INPUT", "-j", "LOG", "--log-prefix", "Incoming Traffic: ")
    if err := logInputCmd.Run(); err != nil {
        fmt.Println("Error setting up logging for incoming traffic:", err)
        return
    }
    fmt.Println("Logging for all incoming traffic set up.")

    // Log all outgoing traffic
    logOutputCmd := exec.Command("iptables", "-A", "OUTPUT", "-j", "LOG", "--log-prefix", "Outgoing Traffic: ")
    if err := logOutputCmd.Run(); err != nil {
        fmt.Println("Error setting up logging for outgoing traffic:", err)
        return
    }
    fmt.Println("Logging for all outgoing traffic set up.")
}
