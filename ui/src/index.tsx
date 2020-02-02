import React from 'react';
import ReactDOM from 'react-dom';
import './index.css';
import * as serviceWorker from './serviceWorker';

console.log(module.hot)
if (module.hot) {
    module.hot.accept("./App", () => {
      const NewApp = require("./App").default;
      ReactDOM.render(<NewApp />, document.getElementById("root") as HTMLElement);
    });
  }

serviceWorker.unregister();
