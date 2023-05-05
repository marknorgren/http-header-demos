import http.server
import socketserver

class CaseSensitiveRequestHandler(http.server.SimpleHTTPRequestHandler):
    def send_header(self, keyword, value):
        self._headers_buffer.append(("%s: %s\r\n" % (keyword, value)).encode('latin-1', 'strict'))

    def do_GET(self):
        self.send_response(200)
        self.send_header("Content-Type", "text/html; charset=utf-8")
        
        # Echo back request headers
        for keyword, value in self.headers.items():
            self.send_header(keyword, value)

        self.end_headers()

        response = b"<html><body><h1>Case-sensitive Headers Echo Server</h1></body></html>"
        self.wfile.write(response)

PORT = 8000

Handler = CaseSensitiveRequestHandler
httpd = socketserver.TCPServer(("", PORT), Handler)
print(f"Serving on port {PORT}")
httpd.serve_forever()
