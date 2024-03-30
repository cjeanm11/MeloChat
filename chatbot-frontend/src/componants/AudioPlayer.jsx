import { createSignal } from "solid-js";
const url = "http://localhost:8080/temp/"

function AudioPlayer(props) {
  const [result] = createSignal(props.result);

  return (
  <audio controls class="bg-gray-200 rounded-lg p-2">
    <source src={ url + result()} type="audio/wav" />
    Your browser does not support the audio element.
  </audio>
  );
}

export default AudioPlayer;
