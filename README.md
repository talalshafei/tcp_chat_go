# Simple TCP Chat
Simple chat written with Golang.

## Commands
- `/nick <name>` - get a name, otherwise user will stay anonymous.
- `/join <name>` - join a room, if room doesn't exist, a new room will be created. User can join one room at the same time.
- `/rooms` - show list of available rooms to join.
- `/msg <msg>` - broadcast a message to everyone in a room
- `/quit` - disconnects from the chat server

## Simulation