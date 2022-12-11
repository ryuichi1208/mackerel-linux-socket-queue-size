import socket
import time

from datetime import datetime

host = '192.168.1.182'
port = 9000
bind_address = (host, port)

backlog_size = 1

with socket.socket(socket.AF_INET, socket.SOCK_STREAM) as server_socket:
    server_socket.setsockopt(socket.SOL_SOCKET, socket.SO_REUSEADDR, 1)
    server_socket.bind(bind_address)
    server_socket.listen(backlog_size)

    print('[{}] Server startup'.format(datetime.now().strftime('%Y-%m-%d %H:%M:%S')))

    try:
        client_socket, addr = server_socket.accept()

        remote_addr = client_socket.getpeername()
    
        print('[{}] - handle connection, start - {}'.format(datetime.now().strftime('%Y-%m-%d %H:%M:%S'), remote_addr))
    
        with client_socket:
            while True:
                time.sleep(1)

    except KeyboardInterrupt:
        print('[{}] Server stop'.format(datetime.now().strftime('%Y-%m-%d %H:%M:%S')))
