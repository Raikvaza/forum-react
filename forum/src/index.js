import React from 'react';
import ReactDOM from 'react-dom/client';
import './index.css';
import App from './routes/App';
import SignUpPage from './routes/Signup';

import { BrowserRouter as Router, Routes, Route, useLoaderData } from "react-router-dom";
import {
  createBrowserRouter,
  RouterProvider,
  
} from "react-router-dom";





const router = createBrowserRouter([
  {
    path: "/",
    element: <App/>,
  },
  {
    path:"/signup",
    element: <SignUpPage/>,
    
    // loader: ({params}) => {
    //   return params.para;
    // },
  },
]);


const root = ReactDOM.createRoot(document.getElementById('root'));
root.render(
  <React.StrictMode>
    <RouterProvider router={router} />
  </React.StrictMode>
);


