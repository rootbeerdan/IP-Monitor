package main

import (
    "log"
    "net/http"
    "time"

    "github.com/gen2brain/beeep"
    "github.com/getlantern/systray"
)

var currentStatus string

func main() {
    // Initialize the system tray icon
    systray.Run(onReady, onExit)
}

func onReady() {
    // Set the initial status to "unknown" and icon to gray
    currentStatus = "unknown"
    systray.SetIcon(getIcon("gray"))

    // Set up a ticker to check connectivity every 30 seconds
    ticker := time.NewTicker(30 * time.Second)
    go func() {
        for {
            select {
            case <-ticker.C:
                // Check connectivity to http://ipv6.whatismyip.akamai.com
                ipv6Resp, ipv6Err := http.Get("http://ipv6.whatismyip.akamai.com")
                ipv6Available := ipv6Err == nil && ipv6Resp.StatusCode == http.StatusOK
                if ipv6Err != nil {
                    log.Printf("Error checking connectivity to ipv6 endpoint: %s", ipv6Err)
                }

                // Check connectivity to http://whatismyip.akamai.com
                ipv4Resp, ipv4Err := http.Get("http://whatismyip.akamai.com")
                ipv4Available := ipv4Err == nil && ipv4Resp.StatusCode == http.StatusOK
                if ipv4Err != nil {
                    log.Printf("Error checking connectivity to ipv4 endpoint: %s", ipv4Err)
                }

                // Set the system tray icon and send a notification if the connectivity status has changed
                var newStatus string
                if ipv6Available && ipv4Available {
                    newStatus = "connected"
                    if currentStatus != "connected" {
                        systray.SetIcon(getIcon("green"))
                        sendNotification("Connected", "Both endpoints are accessible.")
                    }
                } else if ipv6Available || ipv4Available {
                    newStatus = "partially connected"
                    if currentStatus != "partially connected" {
                        systray.SetIcon(getIcon("yellow"))
                        sendNotification("Partially Connected", "One endpoint is accessible.")
                    }
                } else {
                    newStatus = "disconnected"
                    if currentStatus != "disconnected" {
                        systray.SetIcon(getIcon("red"))
                        sendNotification("Disconnected", "Neither endpoint is accessible.")
                    }
                }
                currentStatus = newStatus
            }
        }
    }()
}

func onExit() {
    // Clean up resources on exit
}

func getIcon(color string) []byte {
    // Return the icon bytes for the specified color
    // In a real application, you would have multiple icons of different sizes and resolutions for different platforms
    // This example just returns a dummy icon for demonstration purposes
    switch color {
    case "gray":
        return []byte{0x7F, 0x7F, 0x7F}
    case "green":
        return []byte{0x00, 0xFF, 0x00}
    case "yellow":
        return []byte{0xFF, 0xFF, 0x00}
    case "red":
        return []byte{0xFF, 0x00, 0x00}
    default:
        return []byte{}
    }
}

func sendNotification(title string, message string) {
    // Send a desktop notification using the "beeep" package
    err := beeep.Notify(title, message, "")
    if err !=
