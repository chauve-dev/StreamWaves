from http.server import BaseHTTPRequestHandler, HTTPServer
import threading

list_of_spawn = []

hostName = "localhost"
serverPort = 51525

class MyServer(BaseHTTPRequestHandler):
    def do_GET(self):
        if self.path.startswith("/add"):
            self.add()
        elif self.path == "/fetch":
            self.fetch()
        else:
            self.send_response(200)
            self.send_header("Content-type", "text/html")
            self.end_headers()
            self.wfile.write(bytes("failed", "utf-8"))


    def add(self):
        self.send_response(200)
        self.send_header("Content-type", "text/html")
        self.end_headers()
        monster = self.path.replace('/add/', '')
        list_of_spawn.append(monster)

    def fetch(self):
        self.send_response(200)
        self.send_header("Content-type", "text/html")
        self.end_headers()
        self.wfile.write(bytes("|".join(list_of_spawn), "utf-8"))
        list_of_spawn.clear()

webServer = HTTPServer((hostName, serverPort), MyServer)
print("Server started http://%s:%s" % (hostName, serverPort))
threading.Thread(target=webServer.serve_forever).start()
