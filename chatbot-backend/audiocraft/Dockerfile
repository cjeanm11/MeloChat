FROM python:3.9

RUN apt-get update && apt-get install -y libsndfile1
ENV LD_LIBRARY_PATH="/usr/lib/x86_64-linux-gnu/:$LD_LIBRARY_PATH"

WORKDIR /app

COPY . /app

RUN pip install --no-cache-dir -r requirements.txt

EXPOSE 5000

ENV FLASK_APP=audio_gen

CMD ["python", "-m", "flask", "run", "--host=0.0.0.0"]
