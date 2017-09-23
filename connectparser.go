package main

import (
	"bufio"
	"errors"
	"regexp"
	"strconv"
	"strings"
)

type ConnectionParser struct{}

// Returns a ConnectionClient object based on strings that were extracted from the aconnect output for a new client
func (cp ConnectionParser) createClient(idString string, nameString string) (ConnectionClient, error) {
	var client ConnectionClient

	idString = sanitizeString(idString)
	nameString = sanitizeString(nameString)

	id, err := strconv.Atoi(idString)
	if err != nil {
		return client, err
	}

	client = ConnectionClient{id: id, name: nameString, ports: make([]ConnectionPort, 0)}
	return client, nil
}

// Returns a ConnectionPort object based on strings that were extracted from the aconnect output for a new port
func (cp ConnectionParser) createPort(idString string, nameString string) (ConnectionPort, error) {
	var port ConnectionPort

	idString = sanitizeString(idString)
	nameString = sanitizeString(nameString)

	id, err := strconv.Atoi(idString)
	if err != nil {
		return port, err
	}

	port = ConnectionPort{id: id, name: nameString}
	return port, nil
}

// Perform the actual parsing, using all output from the aconnect command
func (cp ConnectionParser) parse(clientsString string) ([]*ConnectionClient, error) {
	clients := make([]*ConnectionClient, 0)

	reClient := regexp.MustCompile("^client ([0-9]+): \\'(.+)\\'")
	rePort := regexp.MustCompile("^\\s+ ([0-9]+) \\'(.+)\\'")

	scanner := bufio.NewScanner(strings.NewReader(clientsString))

	for scanner.Scan() {
		scanLine := scanner.Text()

		clientMatches := reClient.FindStringSubmatch(scanLine)
		portMatches := rePort.FindStringSubmatch(scanLine)

		if len(clientMatches) > 0 { // is it a client?
			client, err := cp.createClient(clientMatches[1], clientMatches[2])
			if err != nil {
				return clients, err
			}

			clients = append(clients, &client)

		} else if len(portMatches) > 0 { // is it a port?

			if len(clients) < 1 { // we can only append ports when we already have a client
				return clients, errors.New("No clients to add port to")
			}
			client := clients[len(clients)-1]

			port, err := cp.createPort(portMatches[1], portMatches[2])
			if err != nil {
				return clients, err
			}

			client.ports = append(client.ports, port)

		} else { // invalid line
			return clients, errors.New("Unknown line")
		}

	}

	return clients, nil
}
