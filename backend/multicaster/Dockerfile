FROM python:3.10.8-buster

WORKDIR /app
COPY requirements.txt /app/
RUN pip install -r requirements.txt

COPY ./one_to_multiple_cast_skyway.py /app/

CMD ["python3", "/app/one_to_multiple_cast_skyway.py"]
