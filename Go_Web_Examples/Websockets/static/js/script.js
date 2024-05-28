// script.js

let input = document.getElementById("input");
let output = document.getElementById("output");
let socket = new WebSocket("ws://localhost:8080/echo");

socket.onopen = function() {
  output.innerHTML += "Status: Connected\n\n";
};

socket.onmessage = function(e) {
  output.innerHTML += "Server: " + e.data + "\n";
};

function send() {
  socket.send(input.value);
  input.value = "";
}
