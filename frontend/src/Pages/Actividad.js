
import { useEffect, useState } from 'react';
import { useParams } from 'react-router-dom';
import { useLocation } from 'react-router-dom';
import './styles/actividad.css';
import { jwtDecode } from 'jwt-decode';

function Actividad() {
    const [actividad,setactividad] = useState()
    const location = useLocation();
    const props = location.state;
    const idactividad = props.id
    const [selectedActivity,setselectedActivity] = useState()
    var usertoken = localStorage.getItem('userToken')
    var userid = null
    if (usertoken != null) {
        const decoded = jwtDecode(usertoken)
        userid = decoded.userid
    }



    useEffect((async) => {
        // Get Activity by id
        console.log("Start fetch")
        
        try {
            fetch("/actividades/"+idactividad,{
            method:"GET"}).then((res)=>{
            return res.json();
        }).then((data) => {
            setactividad(data)
            console.log(actividad)
        })
        } catch (error) {
            console.log(error)
        }
        
    },[])

    if (actividad == null) { // esperamos hasta que la request termine
        return <></>
    }

    function RealizarInscripcion(_userid,_activityid,_day,_hourstart,_hourfinish) {
        const data = {
            userid: _userid,
            activityid: _activityid,
            day: _day,
            hour_start: _hourstart,
            hour_finish: _hourfinish,
        }
        console.log("JSON: ",JSON.stringify(data))
        fetch("/users/inscription",{
        method:"POST",
        headers: {
          'Content-Type': 'text/plain',
        },
        body: JSON.stringify(data), // Replace with your data
      });
    }


function Horarios() {

        var close = false;
        if (close == false) {
        return (
            <div onClick={() => {setselectedActivity(null)}} className="Inscripcion-horarios">
            <div className="inscription-container">
            <h4>Horarios</h4>
            <div className="Horarios">
                {actividad.activity_hours != null ? actividad.activity_hours.map((hour) => {
                return (
                    <div key={hour.id} className="hour">
                        <p>{hour.day}</p>
                        <p>{hour.hour_start}{hour.hour_start > 1100 && hour.hour_start <= 2300 ? "PM" : "AM"}</p>
                        <p>{hour.hour_finish}{hour.hour_finish > 1100 && hour.hour_finish <= 2300 ? "PM" : "AM"}</p>
                        <button onClick={()=>{RealizarInscripcion(userid,actividad.id,hour.day,hour.hour_start,hour.hour_finish)}}>Inscribirse</button>
                    </div>
                );
            }) : <p>No hay horarios disponibles</p>}
            </div>
            </div>
        </div>
        )  
        } else {
            return false;
        }
    };
    return (
    <div class="Actividad-container">
        <h3>{actividad.name}</h3>
        <p>{actividad.description}</p>
        <p>{actividad.activitytype}</p>
        <p>{actividad.coach_name}</p>
        <p>{actividad.duration} horas</p>
        <button onClick={()=> setselectedActivity(actividad.id)}>Elegir Horario</button>
        {selectedActivity == actividad.id && <Horarios actividad={actividad}/>}
    </div>
    )
}



export default Actividad;