import { Link, useNavigate } from "react-router-dom";
import { useEffect,useState } from "react";
import { jwtDecode } from 'jwt-decode';
import './styles/activities.css';
import HomeBg from '../assets/HomeBg.mp4'
import search from '../assets/search.svg'
import pool from '../assets/pool.jpg'

function Home() {
    const [selectedActivity,setselectedActivity] = useState(null)
    const [Actividades, setActividades] = useState([])
    const [isJWTvalid,setisJWTvalid] = useState()
    const navigation = useNavigate();
    let firstthree;
    useEffect(()=> {
        try {
            const token = localStorage.getItem('userToken') 
            const headers = {
                'Authorization': token,
               'Content-Type': 'text/plain',
            };
            fetch("/actividades",{
                method: 'GET',
                headers: headers
            }).then((res)=> {
                setisJWTvalid(res.headers.get("istokenvalid"))
                if (isJWTvalid == false) {
                    console.log("token invalidated")
                    localStorage.setItem('userToken',null)
                }
                console.log(res.headers.get("istokenvalid"))
                return(res.json());
            }).then((data)=>{
                setActividades(data);
                firstthree = Actividades.slice(0,3)
            })
        } catch (error) {
            console.log("404:Could not Fetch")
            return
        }
    }, []);

    function SearchByParam(event) {
        event.preventDefault();
        const param = document.getElementById("search-activity").value
        const url = 'actividades?keyword='+param;
        navigation(url)
    }

    function Horarios(props) {

        var close = false;
        if (close == false) {
        return (
            <div onClick={()=> {setselectedActivity(null)}} className="Inscripcion-horarios">
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
            return false;
        }
    };


    // Solo queremos mostrar los primeros 3 (si es que hay) en los destacados
    var firstThreeItems = []
    if (Actividades.length > 3) {
        firstThreeItems = Actividades.slice(0, 3);
    }


    return ( 
        <div className="Home-page">
            <section className="wrapper">
            <video src={HomeBg} autoPlay loop muted className="background-video" />
            <div className="titles">
                    <h1 className="title">GYMNAME</h1>
                    <h3 className="subtitle">Empieza a entrenar hoy</h3>
            </div>
            </section>
            <form onSubmit={SearchByParam} className="search-form">
                <div className="search-activity">  
                <button className="form-button" type="submit"><img src={search} width={"18px"} height={"18px"}></img></button>    
                <input placeholder="Buscar por palabra clave o categoria" type="text" id="search-activity" className="search-input"></input>
                </div>
                <p>Busca tu proxima actividad deportiva.</p>
            </form>
            <h1 className="search-form featuredh1">Actividades Destacadas</h1>
            <div className="Activities-container">
            {firstThreeItems.map((actividad)=>{
                return(
               <div key={actividad.id} className="activity-card">
                    <img src={pool}  width={"100%"}></img>
                    <div className="container">
                        <h3>{actividad.name}</h3>
                        <p>Tipo de Actividad: {actividad.activitytype}</p>
                        <button onClick={ ()=> setselectedActivity(actividad.id)}>Realizar Inscripcion</button>
                        {selectedActivity == actividad.id && <Horarios actividad={actividad}/>}
                    </div>
                </div>

            )})}
            </div>
        </div>
    )
};

export default Home;