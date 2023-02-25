from flask import Flask
import buildings.v1
from logging import Logger

app = Flask(__name__)
app.register_blueprint(buildings.v1.endpoint)
log = Logger("backend_app")


@app.route("/", methods=["GET"])
def index():
    log.info("Index hit")
    return "Welcome to Elon Launchpad", 200


if __name__ == "__main__":
    app.run(debug=True, host="localhost", port=5000)
