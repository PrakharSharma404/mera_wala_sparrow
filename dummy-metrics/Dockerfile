FROM python:3.9-slim
WORKDIR /app
RUN pip install prometheus-client
COPY main.py .
CMD ["python", "main.py"]
