import { createSignal } from "solid-js";
import AudioPlayer from './AudioPlayer';
const disabled_class =
  "py-4 pr-0 w-44 px-8 me-2 text-sm font-medium text-gray-900 bg-white rounded-full border border-gray-200 hover:bg-gray-100 hover:text-blue-700 focus:z-10 focus:ring-2 focus:ring-blue-700 focus:text-blue-700 dark:bg-gray-800 dark:text-gray-400 dark:border-gray-600 dark:hover:text-white dark:hover:bg-gray-700 inline-flex items-center";
const button_class =
  "py-4 w-35 px-9 me-2 text-sm font-medium text-gray-900 bg-white rounded-full border border-gray-200 hover:bg-gray-100 hover:text-blue-700 focus:z-10 focus:ring-2 focus:ring-blue-700 focus:text-blue-700 dark:bg-gray-800 dark:text-gray-400 dark:border-gray-600 dark:hover:text-white dark:hover:bg-gray-700 inline-flex items-center";
const url = "http://localhost:8080/api/chat";

function Prompt(props) {
  const [getPrompt, setPrompt] = createSignal("");
  const [getDisabled, setDisabled] = createSignal(false);
  const [loadingText, setLoadingText] = createSignal("Loading");
  const dots = "...";

  const handle_submit = () => {
    props.setChat((c) => {
      c.push({
        user: true,
        text: getPrompt(),
      });
      setPrompt("");
      props.setLoading(true);
      return [...c];
    });
    setDisabled(true);
    const lastText = props.getChat()[props.getChat().length - 1].text 
    const message = JSON.stringify({ description: lastText });
    let timerId = setInterval(() => {
      const currentDots = loadingText();
      setLoadingText(currentDots === dots.repeat(3) ? "Loading" : currentDots + ".");
      if (loadingText() === "Loading....") {
        setLoadingText("Loading");
      }
    }, 300);

    fetch(url, {
      method: "POST",
      headers: {
        "Content-Type": "application/json",
        "origin": 'http://localhost:8000'
      },
      body: message,
    })
      .then((res) => res.json())
      .then((data) => {
        props.setLoading(false);
        setDisabled(false);
        props.setChat((c) => {
          c.push({
            user: false,
            text: <AudioPlayer result={data.result} />
          });

          return [...c];
        });
        clearInterval(timerId)
      });
  };

  return (
    <div  class="h-24 w-full rounded-full fixed bottom-0 flex justify-center items-center p-1 bg-grey-800 opacity-80 backdrop-blur-sm">
      <form
        class="w-full flex justify-center items-center gap-4"
        onsubmit={(e) => {
          e.preventDefault();
          if (getPrompt() == "") {
            return;
          }
          handle_submit();
        }}
      >
        <input
          class="w-2/3 p-2.5 border bg-gradient-to-r from-[#2c4f7c] to-slate-600 text-white placeholder-white border-slate-500 rounded-full abg-shade-10"
          style="resize: none;"
          type="text"
          placeholder="Enter your prompt"
          value={getPrompt()}
          onChange={(event) => setPrompt(event.target.value)}
        />
        <button
          class={getDisabled() ? disabled_class : button_class}
          type="submit"
          disabled={getDisabled()}>
           { getDisabled() ? (<div>
              <svg aria-hidden="true" role="status" class="inline w-5 h-4 me-3 text-gray-200 animate-spin dark:text-gray-600" viewBox="0 0 100 101" fill="none" xmlns="http://www.w3.org/2000/svg">
                <path d="M100 50.5908C100 78.2051 77.6142 100.591 50 100.591C22.3858 100.591 0 78.2051 0 50.5908C0 22.9766 22.3858 0.59082 50 0.59082C77.6142 0.59082 100 22.9766 100 50.5908ZM9.08144 50.5908C9.08144 73.1895 27.4013 91.5094 50 91.5094C72.5987 91.5094 90.9186 73.1895 90.9186 50.5908C90.9186 27.9921 72.5987 9.67226 50 9.67226C27.4013 9.67226 9.08144 27.9921 9.08144 50.5908Z" fill="currentColor"/>
                <path d="M93.9676 39.0409C96.393 38.4038 97.8624 35.9116 97.0079 33.5539C95.2932 28.8227 92.871 24.3692 89.8167 20.348C85.8452 15.1192 80.8826 10.7238 75.2124 7.41289C69.5422 4.10194 63.2754 1.94025 56.7698 1.05124C51.7666 0.367541 46.6976 0.446843 41.7345 1.27873C39.2613 1.69328 37.813 4.19778 38.4501 6.62326C39.0873 9.04874 41.5694 10.4717 44.0505 10.1071C47.8511 9.54855 51.7191 9.52689 55.5402 10.0491C60.8642 10.7766 65.9928 12.5457 70.6331 15.2552C75.2735 17.9648 79.3347 21.5619 82.5849 25.841C84.9175 28.9121 86.7997 32.2913 88.1811 35.8758C89.083 38.2158 91.5421 39.6781 93.9676 39.0409Z" fill="#1C64F2"/>
                </svg>
                <span>{loadingText()}</span>
            </div>) : (<p>Submit</p>)}
          
        </button>
      </form>
    </div>
  );
}

export default Prompt;
