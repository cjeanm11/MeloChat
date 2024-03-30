from flask import jsonify, request, Flask, send_from_directory
from audiocraft.models import MusicGen
from audiocraft.data.audio import audio_write
import tempfile
import os, signal

app = Flask(__name__)
tempfile.mkdtemp()
global tmp_file_counter
tmp_file_counter = 0

def setup_model():
    model = MusicGen.get_pretrained('small')
    model.set_generation_params(duration=1) 
    return model

def generate(descriptions=['space']):
    global tmp_file_counter
    model = setup_model()
    wav = model.generate(descriptions)  
    results = []
    for idx, one_wav in enumerate(wav):
        # Will save under {idx}.wav, with loudness normalization at -14 db LUFS.
        fileName = idx + tmp_file_counter
        tempDir = '../.temp'
        audio_write(f'{tempDir}/{fileName}', one_wav.cpu(), model.sample_rate, strategy="loudness", loudness_compressor=True)
        results.append(f'{fileName}.wav')
    tmp_file_counter = tmp_file_counter + 1
    return results

@app.route("/generate_audio", methods = ['POST'])
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
    return jsonify({ "success": True, "message": "Server is shutting down..." })