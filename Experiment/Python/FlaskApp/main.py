from flask import Flask, request
from Session.PHPSessionReader import ReadSession

app = Flask(__name__)

@app.route('/')
def index():
    SessionId = request.cookies.get("PHPSESSID")
    SessionData = ReadSession(SessionId)
    return SessionData[b'libname']
    

if __name__ == "__main__":
    app.run(port=3000)