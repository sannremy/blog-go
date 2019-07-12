# Build a real-time application using HTML5 WebSockets

In this very first post on my new blog, I show you how I've carried out a simple push system using WebSocket specified by [the RFC 6455 document](https://tools.ietf.org/html/rfc6455). Unfortunately, I can't bring you a live demo because of some hosting service restrictions that don't allow me to create sockets. So, I worked on localhost with my own computer using Ubuntu, Apache and PHP-CLI. Nevertheless, you can download scripts on [my github account](https://github.com/srchea/PHP-Push-WebSocket/).

You'll find one of several basic implementations of a communication between clients and a server by WebSocket in PHP. I deliberately didn't complicate the code (by adding other features or improvements) to make it easy to understand.

## How does a WebSocket basically work?

WebSocket communications are managed by a server. Each client have to introduce itself by sending a handshake request via the WebSocket protocol (ws). The server accepts or not to open a socket connection by sending a handshake response. I think that schemas below are telling more than a long text, so let's have a look on these pictures:

<p style="text-align: center;">
  <img src="/static/assets/posts/websocket-handshake.png" />
</p>

Only after the handshake acceptance, both sides can communicate. Many clients can be connected to the server. For some performance reasons, the limit of clients can be set before creating a new socket.

<p style="text-align: center;">
  <img src="/static/assets/posts/websocket-clients-server.png" />
</p>

## The server side

Like I said earlier in this post, I used Apache 2 and PHP-CLI 5.3 under Ubuntu 10.10. In this section, we're going to see some tricky parts of the server side especially the handshake method, encoding and unmasking data methods. And to finish, I'm going to explain the part of pushing data to clients. The server works like a daemon and is executed by the following command:

```bash
php -q phppushwebsocket.php
```

### Handshake

First of all, the handshake has to be done to let both, client and server, communicate. The browser introduces itself by sending HTTP headers, something like:

```http
GET / HTTP/1.1
Upgrade: websocket
Connection: Upgrade
Host: 127.0.0.1:5001
Origin: http://localhost
Sec-WebSocket-Key: k2towQT28s50DtKptTjZbg==
Sec-WebSocket-Version: 13
```

The `Sec-WebSocket-Version` determines the protocol version of the connection. The HTTP Headers could differ regarding the version. That's why we make sure that the information that we want to extract are correct. For instance, Origin was `Sec-WebSocket-Origin` in the version 8. Moreover, server responses are not the same regarding versions. In the following code, it deals with the version 13. At the moment I'm writing, the unique web-browser which supports this implementation is Google Chrome 16, others would be rejected.

```php
/**
 * Do the handshaking between client and server
 * @param $client
 * @param $headers
 */
private function handshake($client, $headers) {
  $this->console("Getting client WebSocket version...");
  if(preg_match("/Sec-WebSocket-Version: (.*)\r\n/", $headers, $match))
  $version = $match[1];
  else {
  $this->console("The client doesn't support WebSocket");
  return false;
  }

  $this->console("Client WebSocket version is {$version}, (required: 13)");
  if($version == 13) {
  // Extract header variables
  $this->console("Getting headers...");
  if(preg_match("/GET (.*) HTTP/", $headers, $match))
    $root = $match[1];
  if(preg_match("/Host: (.*)\r\n/", $headers, $match))
    $host = $match[1];
  if(preg_match("/Origin: (.*)\r\n/", $headers, $match))
    $origin = $match[1];
  if(preg_match("/Sec-WebSocket-Key: (.*)\r\n/", $headers, $match))
    $key = $match[1];

  $this->console("Client headers are:");
  $this->console("\t- Root: ".$root);
  $this->console("\t- Host: ".$host);
  $this->console("\t- Origin: ".$origin);
  $this->console("\t- Sec-WebSocket-Key: ".$key);

  $this->console("Generating Sec-WebSocket-Accept key...");
  $acceptKey = $key.'258EAFA5-E914-47DA-95CA-C5AB0DC85B11';
  $acceptKey = base64_encode(sha1($acceptKey, true));

  $upgrade = "HTTP/1.1 101 Switching Protocols\r\n".
       "Upgrade: websocket\r\n".
       "Connection: Upgrade\r\n".
       "Sec-WebSocket-Accept: $acceptKey".
       "\r\n\r\n";

  $this->console(
    "Sending this response to the client #{$client->getId()}:"
    ."\r\n".$upgrade
  );
  socket_write($client->getSocket(), $upgrade);
  $client->setHandshake(true);
  $this->console("Handshake is successfully done!");
  return true;
  }
  else {
  $this->console(
    "WebSocket version 13 required"
    ."(the client supports version {$version})"
  );
  return false;
  }
}
```

### Unmasking/Encoding data frames

Well, the client can now send messages to the server and vice-versa. Sent messages have to be formatted in a special way. In fact, if you try to display them, they will be encrypted, so they have to be unmasked to be human readable. The RFC 6455 describes [the way to unmask data](https://tools.ietf.org/html/rfc6455#section-5.2).

An implementation of that description in PHP would be this:

```php
/**
 * Unmask a received payload
 * @param $payload
 */
private function unmask($payload) {
  $length = ord($payload[1]) & 127;

  if($length == 126) {
    $masks = substr($payload, 4, 4);
    $data = substr($payload, 8);
  }
  elseif($length == 127) {
    $masks = substr($payload, 10, 4);
    $data = substr($payload, 14);
  }
  else {
    $masks = substr($payload, 2, 4);
    $data = substr($payload, 6);
  }

  $text = '';
  for ($i = 0; $i < strlen($data); ++$i) {
    $text .= $data[$i] ^ $masks[$i%4];
  }
  return $text;
}
```

We have to encode the text before sending to the client. If a sent text is wrongly encoded, the client might close the connection or not correctly receive it. This is a basic implementation without encryption:

```php
/**
 * Encode a text for sending to clients via ws://
 * @param $text
 */
private function encode($text) {
  // 0x1 text frame (FIN + opcode)
  $b1 = 0x80 | (0x1 & 0x0f);
  $length = strlen($text);

  if($length  125 && $length < 65536)
    $header = pack('CCS', $b1, 126, $length);
  elseif($length >= 65536)
    $header = pack('CCN', $b1, 127, $length);

  return $header.$text;
}
```

### Pushing data

Now that we can send/receive messages and each side can display them properly, we're going to see the pushing part. First of all, the script creates one process for the server + _n_ processes for _n_ connected clients. We're going to fork the server process to make it still accepts new clients and children processes will send data to clients. This method below have to be called after the handshake.

```php
/**
 * Start a child process for pushing data
 * @param unknown_type $client
 */
private function startProcess($client) {
  $this->console("Start a child process");
  $pid = pcntl_fork();
  if($pid == -1) {
    die('could not fork');
  }
  elseif($pid) { // process
    $client->setPid($pid);
  }
  else {
    // we are the child
    while(true) {
      // push something to the client
      $seconds = rand(2, 5);
      $this->send($client, "I am waiting {$seconds} seconds");
      sleep($seconds);
    }
  }
}
```

Note that we set the PID of the Client object in the server process. When a client leaves the application, it will send a 'quit' command to the server (more information in the client side section). The server will treat the message and will send a kill request to the process. However, we can make it in another way: instead of forking the server process, we could execute another program and kill it properly.

## The client side

This part is really easy. Nothing special to figure out other than the WebSocket initialization and events which handle WebSockets: `onopen`, `onclose` and `onerror`. We will see the `onmessage` event in the receiving data section. This is a basic client side implementation:

```javascript
var ws = new WebSocket('ws://127.0.0.1:5001');
ws.onopen = function(msg) {
  write('Connection successfully opened (readyState ' + this.readyState+')');
};
ws.onclose = function(msg) {
  if(this.readyState == 2)
    write(
      'Closing... The connection is going throught'
      + 'the closing handshake (readyState '+this.readyState+')'
    );
  else if(this.readyState == 3)
    write(
      'Connection closed... The connection has been closed'
      + 'or could not be opened (readyState '+this.readyState+')'
    );
  else
    write('Connection closed... (unhandled readyState '+this.readyState+')');
};
ws.onerror = function(event) {
  terminal.innerHTML = ''+event.data+''
  + terminal.innerHTML;
};
```

### Receiving data

The client doesn't have to make a request to the server to know if there are new data. When there is something new, the server will send automatically to the concerned client. That new data is handled by the `onmessage` event.

```javascript
ws.onmessage = function(msg) {
  write('Server says: '+msg.data);
};
```

### Disconnect the client

When a client leaves, a 'quit' message is sent on the before unload window event (`window.onbeforeunload`). This message is directly handled by the server.

```javascript
window.onbeforeunload = function() {
  ws.send('quit');
};
```

Subsequently, the server will execute something like:

```php
if($action == "quit") {
  $this->console("Killing a child process");
  posix_kill($client->getPid(), SIGTERM);
  $this->console("Process {$client->getPid()} is killed!");
}
```

## More information

This is some useful links for more information and details about WebSockets:

 - [The WebSocket Protocol, RFC 6455 on the IETF Website](https://tools.ietf.org/html/rfc6455)
 - [The WebSocket API on the W3C Website](https://dev.w3.org/html5/websockets/)

Feel free to copy and adapt [the source code of PHP Push WebSocket](https://github.com/srchea/PHP-Push-WebSocket/).

In few years, we might say good bye to the pull technology like AJAX and optimize network performance by the way. 🙂
