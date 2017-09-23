package main

import (
	"fmt"
	"os"
	"os/exec"
	"strconv"
	"strings"
)

// Prints the help
func printHelp() {
	fmt.Printf("usage: midiconnect <src> <dest>\n")
	fmt.Printf("   src    The name of the keyboard MIDI interface\n")
	fmt.Printf("   dest   The name of the MIDI interface to be driven by the keyboard\n")
}

// The main function
func main() {

	// Parse arguments
	if len(os.Args) < 3 {
		printHelp()
		return
	}
	src := sanitizeString(os.Args[1])
	dst := sanitizeString(os.Args[2])

	// Perform aconnect -i to list interfaces
	out, err := exec.Command("aconnect", "-i").Output()

	if err != nil {
		fmt.Printf("Could not perform aconnect - make sure it is installed: x%s\n", err)
		return
	}

	// Parse result of aconnect
	var cp ConnectionParser
	clients, err := cp.parse(string(out))

	if err != nil {
		fmt.Printf("Parsing of aconnect output failed: x%s\n", err)
		return
	}

	// Find source and destination interface
	var srcClient, dstClient *ConnectionClient
	for _, client := range clients {
		if client.name == src {
			srcClient = client
		} else if client.name == dst {
			dstClient = client
		}
	}

	if srcClient == nil || len(srcClient.ports) == 0 {
		fmt.Printf("No valid client found for src parameter\n")
		return
	}
	if dstClient == nil || len(dstClient.ports) == 0 {
		fmt.Printf("No valid client found for dest parameter\n")
		return
	}

	// Build aconnect command to link interfaces
	aconnectSrc := []string{strconv.Itoa(srcClient.id), strconv.Itoa(srcClient.ports[0].id)}
	aconnectDst := []string{strconv.Itoa(dstClient.id), strconv.Itoa(dstClient.ports[0].id)}

	aconnectSrcString := strings.Join(aconnectSrc, ":")
	aconnectDstString := strings.Join(aconnectDst, ":")

	fmt.Printf("Executing: %s %s %s\n", "aconnect", aconnectSrcString, aconnectDstString)
	_, err = exec.Command("aconnect", aconnectSrcString, aconnectDstString).Output()

	if err != nil {
		fmt.Printf("Error while connecting: %s\n", err)
		return
	}
}
