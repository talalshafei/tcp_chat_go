# Simple TCP Chat
A lightweight chat application built with Golang that allows multiple users to communicate in real time over a TCP connection. Users can set nicknames, join chat rooms, send messages, and leave conversations easily.

## Commands
- `/nick <name>` – Set a nickname (default is "anonymous").
- `/join <room>` – Join a chat room. If the room doesn’t exist, a new one is created. Users can only be in one room at a time.
- `/rooms` – List all available chat rooms.
- `/msg <message>` – Send a message to everyone in the current room.
- `/quit` – Disconnect from the server.

## Simulation
1. Start the chat server and connect two clients.
   ![Screenshot from 2025-03-10 00-30-03](https://github.com/user-attachments/assets/b40c3518-b013-4aaa-afae-81711c307c24)

2. Each client sets a nickname using /nick <name> for easier identification.
   ![Screenshot from 2025-03-10 00-29-37](https://github.com/user-attachments/assets/167b598c-9db1-4852-ad89-5b828e1bbaa9)

3. Talal joins a room using /join <room>. If the room doesn’t exist, it is created automatically.
   ![image](https://github.com/user-attachments/assets/284896e8-59c3-41e8-84f3-2bba41e5747c)

4. John joins the same room. A notification is broadcasted to all room members, informing them that John has joined. Talal sees this message.
   ![image](https://github.com/user-attachments/assets/1da2f068-eabe-4bab-82fb-3dc5d1217e3f)

5. Talal and John can now chat freely within the room using /msg <message>.
   ![image](https://github.com/user-attachments/assets/bca1d2fa-260e-40bd-97dc-50c4a7aead39)

  
6. Talal decides to join a different room using /join <new-room>. He is automatically removed from his previous room and enters the new one. A message is broadcasted in both rooms to notify members of his departure and arrival.
   ![image](https://github.com/user-attachments/assets/81b24e86-68f7-40b4-9b68-99ce7ab29e3b)

7. Talal can use /rooms at any time to see a list of all available chat rooms.
   ![image](https://github.com/user-attachments/assets/81033a34-9b5f-4371-8dab-8c272fac9e7b)

8. Finally, if Talal decides to leave the chat entirely, he uses /quit, and a message is sent to his current room informing members that he has left.
   ![image](https://github.com/user-attachments/assets/dfb7d2d3-5ef7-4316-905b-3842e2223756)

This setup allows seamless communication between multiple users in different chat rooms while enabling easy room switching and exploration.

### Acknowledgement
This project was inspired by a YouTube [video](https://www.youtube.com/watch?v=Sphme0BqJiY), where I explored TCP in Go and learned about iota.
