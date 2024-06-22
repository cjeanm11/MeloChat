#!/bin/bash
cd ./audiocraft
python -m venv .venv
source ./.venv/bin/activate
pip install -r requirements.txt
flask --app audio_gen run
