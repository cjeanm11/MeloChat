# Changelog

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.0.0/).

## [0.0.2] - 2023-08-01

Improved demo, fixed top p (thanks @jnordberg).

Compressor tanh on output to avoid clipping with some style (especially piano).
Now repeating the conditioning periodically if it is too short.

More options when launching Gradio app locally (thanks @ashleykleynhans).

Testing out PyTorch 2.0 memory efficient attention.

Added extended generation (infinite length) by slowly moving the windows.
Note that other implementations exist: https://github.com/camenduru/MusicGen-colab.

## [0.0.1] - 2023-06-09

Initial release, with model evaluation only.
