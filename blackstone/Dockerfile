FROM python:3.7

ADD . .


RUN python -m pip install -r requirements.txt

EXPOSE 9999

CMD [ "python", "server.py", "--port", "9999" ]