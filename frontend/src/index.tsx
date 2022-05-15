import "./index.css";

import App from "./App";
import ReactDOM from "react-dom/client";

declare global {
  interface Window {
    global: {
      api_location: string;
    };
  }
}

const root = ReactDOM.createRoot(
  document.getElementById("root") as HTMLElement
);
root.render(<App />);
