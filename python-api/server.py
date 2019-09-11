
from flask import Flask, request,jsonify
import requests
import json
import logging

#logger config
logging.basicConfig(
                    level=logging.DEBUG,
                    format="%(asctime)s [%(threadName)-12.12s] [%(levelname)-5.5s]  %(message)s",
                    handlers=[
                        logging.FileHandler("python.log",delay=False),
                        logging.StreamHandler()
                        ]
                    )

app = Flask(__name__)

COMPUTE_URL = "http://localhost:8080/compute"

@app.route("/queue", methods=['POST'])
def queue():
    task_array = []
    content = request.json
    for item in content:
        r = requests.post(COMPUTE_URL,json={"exp":item})
        if r.status_code == 200:
            jsonResponse = r.json()
            logging.debug("compute request success: "+item)
            logging.info("result: "+jsonResponse['result']+" time: "+jsonResponse["time"])
            task_array.append({"time":jsonResponse["time"],"result":jsonResponse['result']})
        else:
            jsonResponse = r.json()
            logging.debug("compute request failed: "+item)
            logging.info(jsonResponse['result'])
            #append results to array
            task_array.append({"time":"","result":"error"})
    return jsonify(task_array)

if __name__ == '__main__':
    app.run(host='0.0.0.0', port=8081)