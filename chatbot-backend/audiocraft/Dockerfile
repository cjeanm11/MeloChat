# Use an official Python runtime as a parent image
FROM python:3.9

RUN apt-get update && apt-get install -y libsndfile1
ENV LD_LIBRARY_PATH="/usr/lib/x86_64-linux-gnu/:$LD_LIBRARY_PATH"

# Set the working directory in the container
WORKDIR /app

# Copy the current directory contents into the container at /app
COPY . /app

# Install dependencies directly in the container
RUN pip install --no-cache-dir -r requirements.txt

# Make port 5000 available to the world outside this container
EXPOSE 5000

# Define environment variable
ENV FLASK_APP=audio_gen

# Run the Flask app when the container launches
CMD ["python", "-m", "flask", "run", "--host=0.0.0.0"]
