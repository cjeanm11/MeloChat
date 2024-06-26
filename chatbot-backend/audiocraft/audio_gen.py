from flask import jsonify, request, Flask, send_from_directory
from audiocraft.models import MusicGen
from audiocraft.data.audio import audio_write
import tempfile
import os, signal
import uuid  # Import the uuid module

app = Flask(__name__)
tempfile.mkdtemp()

def setup_model():
    model = MusicGen.get_pretrained('small')
    model.set_generation_params(duration=1)
    return model

def generate(descriptions=['space']):
    model = setup_model()
    wav = model.generate(descriptions)
    results = []
    for one_wav in wav:
        # Generate a unique file name using UUID
        fileName = str(uuid.uuid4())
        tempDir = '../.temp'
        audio_write(f'{tempDir}/{fileName}', one_wav.cpu(), model.sample_rate, strategy="loudness", loudness_compressor=True)
        results.append(f'{fileName}.wav')
    return results

@app.route("/generate_audio", methods=['POST'])
def generate_audio():
    data = request.get_json()
    description = data.get("description")
    if description:
        results = generate([description])
        if results:
            return jsonify({"result": results[-1]})
    return jsonify({"error": "No descriptions provided"})

@app.route('/stopServer', methods=['GET'])
def stopServer():
    os.kill(os.getpid(), signal.SIGINT)
    return jsonify({"success": True, "message": "Server is shutting down..."})

if __name__ == "__main__":
    app.run()
