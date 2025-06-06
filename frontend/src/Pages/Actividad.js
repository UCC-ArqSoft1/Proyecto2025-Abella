
import { act, useEffect, useState } from 'react';

import React from 'react';
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
            <div className="Inscripcion-horarios">
            <div className="inscription-container">
            <h4>Horarios</h4>
            <div className="Horarios">
                <button class="close-btn" onClick={() => {setselectedActivity(null)}}>X</button>
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
    <div className='activity-wrapper'>
 <div className="card">
        <div className="header">
            <div className="icon-wrapper">
                <svg className="icon" viewBox="0 0 24 24">
                    <path strokeLinecap="round" strokeLinejoin="round" d="M13 10V3L4 14h7v7l9-11h-7z" />
                </svg>
            </div>
            <h1 className="title" id="activity-name">{actividad.name}</h1>
            <div className="activity-type" id="activity-type">{actividad.activitytype}</div>
        </div>

        <div className="description">
            <p id="activity-description">
            {actividad.description}
            </p>
        </div>

        <div className="details">
            <div className="detail-card coach-card">
                <div className="detail-header">
                    <svg className="detail-icon coach-icon" viewBox="0 0 24 24">
                        <path strokeLinecap="round" strokeLinejoin="round" d="M16 7a4 4 0 11-8 0 4 4 0 018 0zM12 14a7 7 0 00-7 7h14a7 7 0 00-7-7z" />
                    </svg>
                    <h3 className="detail-label">Profesor</h3>
                </div>
                <p className="detail-value" id="coach-name">{actividad.coach_name}</p>
            </div>

            <div className="detail-card duration-card">
                <div className="detail-header">
                    <svg className="detail-icon duration-icon" viewBox="0 0 24 24">
                        <path strokeLinecap="round" strokeLinejoin="round" d="M12 8v4l3 3m6-3a9 9 0 11-18 0 9 9 0 0118 0z" />
                    </svg>
                    <h3 className="detail-label">Duracion</h3>
                </div>
                <p className="detail-value" id="activity-duration">{actividad.duration} Horas</p>
            </div>
        </div>

        <div className="action-section">
            <button className="start-button" onClick={()=>{setselectedActivity(actividad.id)}}>Inscribirse ahora</button>
        </div>
    </div>
        {selectedActivity == actividad.id && <Horarios/>}
    </div>
   
    )
}



export default Actividad;