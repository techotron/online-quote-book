import React from 'react';
import ReactDOM from 'react-dom';
import 'bootstrap/dist/css/bootstrap.min.css';

import Router from './router';
import MainNavBar from './navbar';

ReactDOM.render(
    <React.StrictMode>
        <MainNavBar />
        <Router />
    </React.StrictMode>,
    document.getElementById('root')
);
