import aiohttp
from flask import Flask


app = Flask(__name__)

@app.route("/queue")
def compute():
    pass

if __name__ == '__main__':
    app.run(host='0.0.0.0', port=8081)