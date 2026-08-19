import { BrowserRouter, Route, Routes } from 'react-router-dom'
import './App.css'
import Landing from './pages/Landing'
import Login from './pages/Login'


function App() {
  return (
    <>
      <div>
       

      
      <BrowserRouter>
        <Routes>
          <Route path="/" element={<Landing />}/>
          <Route path="/login" element={<Login />}/>
        </Routes>
      </BrowserRouter>
      </div>
    </>
 )
}

export default App

        