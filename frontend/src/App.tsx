import { BrowserRouter, Route, Routes } from 'react-router-dom'
import './App.css'
import Landing from './pages/Landing'
import Login from './pages/Login'
import Register from './pages/Registor'
import Home from './pages/Home'


function App() {
  return (
    <>
      <div>
       

      
      <BrowserRouter>
        <Routes>
          <Route path="/" element={<Landing />}/>
          <Route path="/login" element={<Login />}/>
          <Route path="/register" element={<Register />}/>
          <Route path="/home" element={<Home/>}/>
        </Routes>
      </BrowserRouter>
      </div>
    </>
 )
}

export default App

        