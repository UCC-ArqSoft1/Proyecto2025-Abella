import { useState } from "react";


export default function Horarios(props) {
        const [closed,setclosed] = useState(false)
        if (closed == false) {
        return (
            <div onClick={()=> {setclosed(true)}} className="Inscripcion-horarios">
                <div className="inscription-container">
                <h4>Horarios</h4>
                    <div className="Horarios">
                        {props.actividad.activity_hours != null ? props.actividad.activity_hours.map((hour) => {
                        return (
                            <div key={hour.id} className="hour">
                                <p>{hour.day}</p>
                                <p>{hour.hour_start}{hour.hour_start > 1100 && hour.hour_start <= 2300 ? "PM" : "AM"}</p>
                                <p>{hour.hour_finish}{hour.hour_finish > 1100 && hour.hour_finish <= 2300 ? "PM" : "AM"}</p>
                                <button>Inscribirse</button>
                            </div>
                        );
                    }) : <p>No hay horarios disponibles</p>}
                    </div>
                </div>
            </div>
        )  
        } else {
            console.log("closed")
            return false;
        }
};