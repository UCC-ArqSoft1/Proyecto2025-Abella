import { useState, useEffect } from 'react';
import { useNavigate } from 'react-router-dom'
import './CreateActivity.css' 

function FormularioEditarHorario({ setiscreating,idactividad,idhorario }) {
    const navigation = useNavigate()
    const headers_Create = {
        'Authorization': localStorage.getItem('userToken'),
        'Content-Type': 'text/plain',
    };

    async function PostEditHorario() {
        
        const data = {
            id: parseInt(idhorario),
            idactividad: parseInt(idactividad),
            day: document.getElementById("day_edit").value,
            starting_hour: parseInt(document.getElementById("starting_hour_input_edit").value),
            finish_hour: parseInt(document.getElementById("finish_hour_input_edit").value)
        }
        console.log("JSON: ",JSON.stringify(data))
        const url = "/activities/hours/edit"
        await fetch(url,{
            method:"POST",
            headers: headers_Create,
            body: JSON.stringify(data),
        }).then((res)=>{
            if (res.status == 200) {
                alert("Horario editado con exito")
            }
            if (res.status == 401) {
                console.log("token is not valid")
                localStorage.removeItem('userToken')
                alert("Debes iniciar sesion nuevamente para realizar esta accion.")
                navigation("/Login")
            }
        })
    }
    return (
        <div className='CreateActivityForm-container'>
            <div className='CreateActivityForm'>
                <h5>Formulario Editar Horario</h5>
                <input id='day_edit' placeholder='Dia'></input>
                <input id='starting_hour_input_edit' placeholder='Hora Inicio'></input>
                <input id='finish_hour_input_edit' placeholder='Hora finaliza'></input>
                <button onClick={()=>{PostEditHorario()}}>Confirmar</button>
                <button onClick={() => setiscreating(false)}>Cancelar</button>
            </div>
        </div>
    );
}

export default FormularioEditarHorario;