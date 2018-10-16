import React from "react";
import ReactDOM from "react-dom";

import App from "/components";

window.addEventListener("DOMContentLoaded", () => {
  const el = document.querySelector("#app");
  ReactDOM.render(<App />, el);
});
