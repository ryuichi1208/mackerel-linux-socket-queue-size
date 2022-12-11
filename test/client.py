import socket

IPADDR = "127.0.0.1"
PORT = 9000

sock = socket.socket(socket.AF_INET)
sock.connect((IPADDR, PORT))
sock.close()
