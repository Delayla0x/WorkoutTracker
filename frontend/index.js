import React from 'react';
import ReactDOM from 'react-dom';
import App from './App';

const API_URL = process.env.REACT_APP_API_URL;

ReactDOM.render(
  <React.StrictMode>
    <App apiUrl={API_URL} />
  </React.StrictMode>,
  document.getElementById('root')
);