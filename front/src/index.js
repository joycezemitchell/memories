import React from 'react';
import ReactDOM from 'react-dom';
import './index.css';
import App from './App';
import * as serviceWorker from './serviceWorker';
import { BrowserRouter } from 'react-router-dom';
import JavascriptTimeAgo from 'javascript-time-ago'
 
import en from 'javascript-time-ago/locale/en'
import ru from 'javascript-time-ago/locale/ru'

JavascriptTimeAgo.addLocale(en)
JavascriptTimeAgo.addLocale(ru)

ReactDOM.render(
  <React.StrictMode>
    <BrowserRouter>
      <App />
    </BrowserRouter>
  </React.StrictMode>,
  document.getElementById('root')
);

serviceWorker.unregister();
