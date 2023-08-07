import React from "react";
import axios from "axios";
import ReactDOM from "react-dom/client";
import App from "@/App.tsx";

axios.defaults.baseURL = import.meta.env.VITE_BASE_API;

ReactDOM.createRoot(document.getElementById("root")!).render(
  <React.StrictMode>
    <App />
  </React.StrictMode>
);
