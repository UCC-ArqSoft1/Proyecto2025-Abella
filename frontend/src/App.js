import './App.css';
import Home from "./Pages/Home"
import {BrowserRouter, Routes, Route } from "react-router-dom";
import { useState, useEffect } from "react";
import Register from './Pages/Register';
import Login from './Pages/Login';
import Actividad from './Pages/Actividad';
import Actividades from './Pages/Actividades';
import Navbar from './Components/Navbar';
import react from 'react';
import { jwtDecode } from 'jwt-decode';
import Profile from './Pages/Profile';

export const UserTypeContext = react.createContext()

function App() {
  
  const [userType,setuserType] = useState(null)

  useEffect(()=> { // La funcion es un poco redundante pero de esa manera se asegura de que el token no sea invalido y que todos los errores esten manejados.
    const token = localStorage.getItem('userToken')
        console.log(token)
        if (token == null) {
            setuserType(null)
        } else {
          let decoded = null
          try {
              decoded = jwtDecode(token)
          } catch (error) {
            setuserType(null)
            localStorage.setItem('userToken',null)
            return;
          }
          setuserType(decoded.usertype)
        }
        return
  },[])

  return (
      <UserTypeContext.Provider value={[userType,setuserType]}> 
      <BrowserRouter>
      <Navbar></Navbar>
      <Routes>
          <Route path='/' element={<Home/>}/>
          <Route path="Register" element={<Register />} />
          <Route path="Login" element={<Login />} />
          <Route path="Actividad" element={<Actividad />} />
          <Route path="actividades" element={<Actividades />} />
          <Route path='Profile' element={<Profile/>}></Route>
      </Routes>
    </BrowserRouter>
    </UserTypeContext.Provider>
  );
}

export default App;
