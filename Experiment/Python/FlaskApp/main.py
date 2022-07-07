from flask import Flask, request
from Session.PHPSessionReader import ReadSession

app = Flask(__name__)

@app.route('/')
def index():
    SessionId = request.cookies.get("PHPSESSID")
    print(SessionId)
    

if __name__ == "__main__":
    app.run(port=3000)