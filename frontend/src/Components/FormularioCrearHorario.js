import { useState, useEffect } from 'react';
import { useNavigate } from 'react-router-dom'
import './CreateActivity.css' 

function FormularioCrearHorario({ setiscreating,idactividad }) {
    const navigation = useNavigate()

    const headers_Create = {
        'Authorization': localStorage.getItem('userToken'),
        'Content-Type': 'text/plain',
    };

    async function PostCrearHorario() {
        
        const data = {
            id: parseInt(idactividad),
            day: document.getElementById("day").value,
            starting_hour: parseInt(document.getElementById("starting_hour_input").value),
            finish_hour: parseInt(document.getElementById("finish_hour_input").value)
        }
        console.log("JSON: ",JSON.stringify(data))
        const url = "/activities/addhour"
        await fetch(url,{
            method:"POST",
            headers: headers_Create,
            body: JSON.stringify(data),
        }).then((res)=>{
            if (res.status == 202) {
                alert("Actividad Creada con exito")
            }
            if (res.status == 401) {
                console.log("token is not valid")
                localStorage.removeItem('userToken')
                alert("Debes iniciar sesion nuevamente")
                navigation("/Login")
            }
        })
    }
    return (
        <div className='CreateActivityForm-container'>
            <div className='CreateActivityForm'>
                <h5>Formulario Crear Horario</h5>
                <input id='day' placeholder='Dia'></input>
                <input id='starting_hour_input' placeholder='Hora Inicio'></input>
                <input id='finish_hour_input' placeholder='Hora finaliza'></input>
                <button onClick={()=>{PostCrearHorario()}}>Crear Horario</button>
                <button onClick={() => setiscreating(false)}>Cancelar</button>
            </div>
        </div>
    );
}

export default FormularioCrearHorario;