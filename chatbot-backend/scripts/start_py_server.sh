#!/bin/bash
cd ./audiocraft && . ./.venv/bin/activate
flask --app audio_gen run
