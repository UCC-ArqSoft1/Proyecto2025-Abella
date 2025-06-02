import { useState,useEffect } from "react";
import './styles/Login.css';


async function CreateUser() {
  try {
    const user = {
      email: document.getElementById("email-field").value,
      password: document.getElementById("password-field").value,
      name: document.getElementById("name-field").value,
      lastname: document.getElementById("lastname-field").value,
      documentation: parseInt(document.getElementById("documentation-field").value, 10),
    };

    const response = await fetch('http://localhost:8523/register', {
    method: 'POST',
    headers: {
        'Content-Type': 'text/plain'
    },
    body: JSON.stringify(user)
    });

    if (response.ok) {
      alert("Usuario creado con exito");
      // Optionally, you could redirect the user or update the UI here
    } else {
      const errorData = await response.json(); // Try to get more specific error info from the server
      console.error("Error creating user:", errorData);
      alert("No se ha podido crear el usuario");
      // Optionally, display a more informative error message to the user
    }
  } catch (error) {
    console.error("Fetch error:", error);
    alert("Ocurrió un error al intentar crear el usuario."); // Generic error message for network issues
  }
}

function Register() {
    return (
        <div className="container-login">
            <form className="form">
                <h3>Crear Cuenta</h3>
                <input id="email-field" placeholder="Email" type="email"></input>
                <input id="password-field" placeholder="Contraseña" type="password"></input>
                <input id="name-field" placeholder="Nombre"></input>
                <input id="lastname-field" placeholder="Apellido"></input>
                <input id="documentation-field" placeholder="DNI"></input>
                <button type="button" onClick={CreateUser}>Crear Cuenta</button>
                <span>
                  <p>¿Ya tienes una cuenta? <a href="/login">Inicia sesion</a> </p>
                </span>
            </form>
        </div>
    )
}

export default Register;