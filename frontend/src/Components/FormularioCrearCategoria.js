import { useState, useEffect } from 'react';
import { useNavigate } from 'react-router-dom'
import './CreateActivity.css' 

function FormularioCrearCategoria({setiscreatingCategory,setiscreating}) {
    const navigation = useNavigate()

    const headers_Create = {
        'Authorization': localStorage.getItem('userToken'),
        'Content-Type': 'text/plain',
    };

    async function PostCrearCategoria() {
        const data = {
            name: document.getElementById("name").value
        }
        console.log("JSON: ",JSON.stringify(data))
        const url = "/activities/category/add"
        await fetch(url,{
            method:"POST",
            headers: headers_Create,
            body: JSON.stringify(data),
        }).then((res)=>{
            if (res.status == 200) {
                alert("Categoria Creada con exito")
            }
            if (res.status == 401) {
                console.log("token is not valid")
                localStorage.removeItem('userToken')
                alert("Debes iniciar sesion nuevamente")
                navigation("/Login")
            }
            setiscreatingCategory(false)
            setiscreating(true)
        })
    }
    return (
        <div className='CreateActivityForm-container'>
            <div className='CreateActivityForm'>
                <h5>Formulario Crear Horario</h5>
                <input id='name' placeholder='Nombre Categoria'></input>
                <button onClick={()=>{PostCrearCategoria()}}>Crear Horario</button>
                <button onClick={() =>{setiscreatingCategory(false);setiscreating(true)} }>Cancelar</button>
            </div>
        </div>
    );
}

export default FormularioCrearCategoria;