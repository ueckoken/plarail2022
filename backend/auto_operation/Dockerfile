FROM python:3.9.6-slim

WORKDIR /app
COPY ./requirements.txt ./
RUN pip install -r requirements.txt

COPY ./ ./

CMD [ "python3", "main.py" ]
