import { jwtDecode } from 'jwt-decode'
import { useEffect, useState } from 'react'
import { useNavigate } from 'react-router-dom'
import './styles/profile.css'


export default function Profile() {
    const navigation = useNavigate()
    const [MyActivities,setMyActivities] = useState([])


    const token = localStorage.getItem('userToken')
    useEffect(() => {
        if (token == null) {
            navigation("/Login")
            return
        }
        // Si el token no es null hacemos la request aca
        var decoded

        try {
            decoded = jwtDecode(token)
        } catch (error) {
            localStorage.removeItem('userToken')
            alert("Debes iniciar sesion nuevamente")
            navigation("/Login")
            return
        }
        let userid = decoded.userid
        let url = "/user/" + userid + "/activities"

        const headers = {
            'Authorization': localStorage.getItem('userToken'),
            'Content-Type': 'text/plain',
        };


        try {
            fetch(url, {
                method: 'GET',
                headers: headers
            }).then((res) => {
                if (res.status == 401) {
                    console.log("token is not valid")
                    localStorage.removeItem('userToken')
                    alert("Debes iniciar sesion nuevamente")
                    navigation("/Login")
                }
                return (res.json());
            }).then((data) => {
                console.log(data)
                setMyActivities(data)
            })
        } catch (error) {
            console.log("404:Could not Fetch")
            return
        }

    },[])

    function handleDelete(id) {
        const param = id
        const url = "/users/inscription/"+id+"/delete"
        console.log(url)
        try {
            fetch(url,{
                headers:{
                    'Authorization':token,
                    'Content-Type': 'text/plain',  
                },
                method:'DELETE'
            }).then((res)=>{
                if (!res.ok) {
                    if (res.status == 401) {
                        console.log("token is not valid")
                        localStorage.removeItem('userToken')
                        alert("Debes iniciar sesion nuevamente")
                        navigation("/Login")
                    }
                } else {
                    alert("Se ha salido de la actividad")
                }
            })
        } catch (error) {
            alert("Ha ocurrido un error al salir de la actividad")
        }
    }



    return (
        <section className='myactivities-section'>
            <h1>Mis Actividades</h1>
            <div className='myactivities-container'>
            {MyActivities?.map((myactividad)=>{
                return (
                    <div key={myactividad.id}>
                        <div className='myactivity'>
                        <h4>Actividad: {myactividad.name}</h4>
                        <div className='myactivities-details'>
                            <p>Dia: {myactividad.day}</p>
                            <p>Hora Inicio: {myactividad.hour_start}{myactividad.hour_start >= 1200 && myactividad.hour_start <= 2300 ? "PM" : "AM"}</p>
                            <p>Hora Inicio: {myactividad.hour_finish}{myactividad.hour_finish >= 1200 && myactividad.hour_finish <= 2300 ? "PM" : "AM"}</p>
                        </div>
                    </div>
                    <button className='right' onClick={()=>{handleDelete(myactividad.id)}}>Abandonar</button>
                    </div>
                )
            })}
            </div>
        </section>
    )
}