import logo from './logo.svg';
import './App.css';
import Home from "./Pages/Home"
import {BrowserRouter, Routes, Route } from "react-router-dom";
import { useState, createContext, useContext, useEffect } from "react";
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

  useEffect(()=> { // We update the 
    if (localStorage.getItem('userToken') != "null" && localStorage.getItem('userToken') != "") { // If token exists
      const decoded = jwtDecode(localStorage.getItem('userToken'))
      if (decoded.usertype != null) {
        setuserType(decoded.usertype)
      }
    }
  })

  

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
