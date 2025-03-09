package main

import (
	"bufio"
	"fmt"
	"net"
	"strings"
)

type client struct {
	conn     net.Conn
	nick     string
	room     *room
	commands chan<- command
}

func (c *client) readInput() {
	for {
		msg, err := bufio.NewReader(c.conn).ReadString('\n')
		if err != nil {
			return
		}

		msg = strings.Trim(msg, "\r\n")
		args := strings.Split(msg, " ")
		cmd := strings.TrimSpace(args[0])

		initCommand := func(id commandID) command {
			return command{
				id:     id,
				client: c,
				args:   args,
			}
		}

		switch cmd {
		case "/nick":
			c.commands <- initCommand(CMD_NICK)

		case "/join":
			c.commands <- initCommand(CMD_JOIN)

		case "/rooms":
			c.commands <- initCommand(CMD_ROOMS)

		case "/msg":
			c.commands <- initCommand(CMD_MSG)

		case "/quit":
			c.commands <- initCommand(CMD_QUIT)

		default:
			c.err(fmt.Errorf("Unknown command: %s", cmd))
		}
	}
}

func (c *client) err(err error) {
	c.conn.Write([]byte("Err: " + err.Error() + "\n"))
}

func (c *client) msg(message string) {
	c.conn.Write([]byte("> " + message + "\n"))
}
