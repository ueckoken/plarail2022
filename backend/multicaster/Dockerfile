FROM python:3.11.3-buster

WORKDIR /app
COPY requirements.txt /app/
RUN pip install -r requirements.txt

COPY ./one_to_multiple_cast_skyway.py /app/

CMD ["python3", "/app/one_to_multiple_cast_skyway.py"]
