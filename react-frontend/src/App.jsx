import { useState } from 'react'
import './App.css'
import { Route, BrowserRouter as Router, Routes } from "react-router-dom";

import ResumeInput from './components/ResumeInput'

function App() {

  return (
    <Router>
      <Routes>
        <Route path="/*" element={<ResumeInput />}></Route>
      </Routes>
    </Router>
  )
}

export default App
