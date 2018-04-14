**WebSockets**

Message Room Arch:

```
{
    "roomId": { 
        "userId": userChannel
    }
}
```


Server:

* Starts a goroutine for every new client that is connected.
* Handles the Messages Accordingly.
* Writes the message to all the other users except for the one that has sent that message.

Client:

* Starts two goroutines.
* One for the reading from the stdin.
* Another for the reading from the Socket Connection.

How to use:

* Start the server
```
$ go run server.go
```

* Start multiple clients
```
$ go run client.go

# Terminal 1
userid:username //registers the users
start_room:roomname //Creates a room with that name #TO BE CHANGED

# Terminal 2
userid:username2
join_room:roomname //Joins the room
msg:message //Sends the message to other users if any in the same room
```