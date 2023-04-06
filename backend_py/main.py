from flask import Flask, request
import buildings.v1
from logging import Logger
from uuid import uuid4
from typing import Dict

app = Flask(__name__)
app.register_blueprint(buildings.v1.endpoint)
log = Logger("backend_app")

login_tokens: Dict[str, bool] = {}


@app.route("/", methods=["GET"])
def index():
    log.info("Index hit")
    return "Welcome to the ESC Reference OAuth2 Impl.", 200


@app.route("/login", methods=["GET"])
def login():
    new_token = uuid4()
    login_tokens[str(new_token)] = True
    return (
        f"""
        <meta http-equiv="refresh" content="0; URL='http://localhost:1339/api/login/callback?token={new_token}'"/>
        """,
        200
    )

@app.route("/validate", methods=["GET"])
def validate():
    token: str = request.args.get("token")
    if token in login_tokens:
        return {token: str(uuid4())}, 200
    else:
        return "we not good", 401

if __name__ == "__main__":
    app.run(debug=True, host="localhost", port=1338)
