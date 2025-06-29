import { useState } from "react";
import { useContext } from 'react'
import {UserTypeContext} from "../App"
import { jwtDecode } from 'jwt-decode';
import { Link, useNavigate } from "react-router-dom";

function Login() {
  const [data, setData] = useState(); 
  const [loading, setLoading] = useState(false);
  const [error, setError] = useState(null);
  const [userType,setuserType] = useContext(UserTypeContext);

  const navigation = useNavigate(); 

  const handleSubmit = async (event) => {
    event.preventDefault();
    setLoading(true);
    setError(null);
    const user = {
      email: document.getElementById("email").value,
      password: document.getElementById("password").value,
    };


    try {
      const response = await fetch('http://localhost:8523/login', {
        method: 'POST',
        headers: {
          'Content-Type': 'text/plain',
        },
        body: JSON.stringify(user), // Replace with your data
      });

      if (!response.ok) {
        throw new Error(`HTTP error! status: ${response.status}`);
      }

      const responseData = await response.json();
      setData(responseData);
      localStorage.setItem('userToken', responseData.token);
      const decoded = jwtDecode(localStorage.getItem('userToken'))
      setuserType(decoded.usertype)
      navigation("/")
    } catch (e) {
      console.log(e)
      alert("Usuario o contraseña incorrecto")
      return
    } finally {
      setLoading(false);

    }
  };



    return (
      <div className="container-login">
        <form className="form">
          <h3>Log in</h3>
            <input placeholder="Email" type="email" id="email"></input>
            <input placeholder="password" type="password" id="password"></input>
            <button onClick={handleSubmit}>Log in</button>
        </form>
      </div>
    )
}

export default Login;

/*
async function LoginRequest() {
  try {

    const response = await fetch('http://localhost:8523/login', {
    method: 'POST' ,
    headers: {
        'Content-Type': 'text/plain'
    },
    body: JSON.stringify(user)
    }).then(res => {
    if (response.ok) {
          console.log(response.json())
          alert("Logged in");
            // REDIRECT USER 
        } else {
          alert("Could not log in"); 
        }
    })
  } catch (error) {
    console.error("Fetch error:", error);
    alert("Error ocurred while loggin in."); 
  }
}*/