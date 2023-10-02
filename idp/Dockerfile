# syntax=docker/dockerfile:1.2
FROM python:bookworm as builder

WORKDIR /app

COPY requirements.txt requirements.txt


RUN apt-get update && apt-get install -y --no-install-recommends \
    git \
    build-essential \
    python3-dev \
    python3-pip \
    libsasl2-dev \
    libssl-dev \
    libffi-dev \
    xmlsec1

RUN pip install -r requirements.txt

COPY . .

EXPOSE 8088

CMD ["./idp.py", "idp_conf"]
