import { useState } from 'react'
import './App.css'
import { Route, BrowserRouter as Router, Routes } from "react-router-dom";

import ResumeInput from './components/ResumeInput'
import { ToastContainer } from 'react-toastify';

function App() {

  return (
    <Router>
      <ToastContainer />
      <Routes>
        <Route path="/*" element={<ResumeInput />}></Route>
      </Routes>
    </Router>
  )
}

export default App
