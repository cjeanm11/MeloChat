import { createEffect, createSignal } from 'solid-js';
import styles from './App.module.css';
import Chat from './componants/Chat';
import Prompt from './componants/Prompt';

function App() {
  const[getChat, setChat] = createSignal([{"text":"What is your name ?", "user": true}, {"text":"My name is AI", "user": false}]);
  const [Loading, setLoading] = createSignal(false);
  return (
    <div class={styles.App}>
      <div class='on flex h-screen flex-col overscroll-none' >
        <div class='text-gray-200 bg-gradient-to-t from-slate-600 to-slate-800 pt-4 pb-3 sticky top-0 flex-1 border-4 border-[#243c5a] border-t-0 border-l-0 border-r-0' style='-webkit-app-region: drag;' >
          <div class="flex-1 flex-line justify-left">
            <h1 class="text-grey-900 text-lg blur-none">
             <svg style='display: inline;' width="20" height="20" xmlns="http://www.w3.org/2000/svg" viewBox="0 0 166 155.3"><defs><linearGradient id="a" gradientUnits="userSpaceOnUse" x1="27.5" y1="3" x2="152" y2="63.5"><stop offset=".1" stop-color="#76b3e1"/><stop offset=".3" stop-color="#dcf2fd"/><stop offset="1" stop-color="#76b3e1"/></linearGradient><linearGradient id="b" gradientUnits="userSpaceOnUse" x1="95.8" y1="32.6" x2="74" y2="105.2"><stop offset="0" stop-color="#76b3e1"/><stop offset=".5" stop-color="#4377bb"/><stop offset="1" stop-color="#1f3b77"/></linearGradient><linearGradient id="c" gradientUnits="userSpaceOnUse" x1="18.4" y1="64.2" x2="144.3" y2="149.8"><stop offset="0" stop-color="#315aa9"/><stop offset=".5" stop-color="#518ac8"/><stop offset="1" stop-color="#315aa9"/></linearGradient><linearGradient id="d" gradientUnits="userSpaceOnUse" x1="75.2" y1="74.5" x2="24.4" y2="260.8"><stop offset="0" stop-color="#4377bb"/><stop offset=".5" stop-color="#1a336b"/><stop offset="1" stop-color="#1a336b"/></linearGradient></defs><path d="M163 35S110-4 69 5l-3 1c-6 2-11 5-14 9l-2 3-15 26 26 5c11 7 25 10 38 7l46 9 18-30z" fill="#76b3e1"/><path d="M163 35S110-4 69 5l-3 1c-6 2-11 5-14 9l-2 3-15 26 26 5c11 7 25 10 38 7l46 9 18-30z" opacity=".3" fill="url(#a)"/><path d="M52 35l-4 1c-17 5-22 21-13 35 10 13 31 20 48 15l62-21S92 26 52 35z" fill="#518ac8"/><path d="M52 35l-4 1c-17 5-22 21-13 35 10 13 31 20 48 15l62-21S92 26 52 35z" opacity=".3" fill="url(#b)"/><path d="M134 80a45 45 0 00-48-15L24 85 4 120l112 19 20-36c4-7 3-15-2-23z" fill="url(#c)"/><path d="M114 115a45 45 0 00-48-15L4 120s53 40 94 30l3-1c17-5 23-21 13-34z" fill="url(#d)"/></svg>
             &nbsp; MeloChat 
            </h1>
          </div>
      </div>

      <div class='flex-1'
      >
        <Chat getChat={getChat} Loading={Loading}></Chat>
      </div>
      <div class='flex-1'>
      <Prompt setChat={setChat} getChat={getChat} setLoading={setLoading}></Prompt>
      </div>

      </div>
  </div>
  );
}

export default App;
