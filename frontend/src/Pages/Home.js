import { Link, useNavigate } from "react-router-dom";
import { useEffect,useState } from "react";
import { jwtDecode } from 'jwt-decode';
import './styles/activities.css';
import HomeBg from '../assets/HomeBg.mp4'
import search from '../assets/search.svg'
import pool from '../assets/pool.jpg'
import Horarios from "../Components/Inscription_hours";

function Home() {
    const [selectedActivity,setselectedActivity] = useState(null)
    const [Actividades, setActividades] = useState([])
    const navigation = useNavigate();
    let firstthree;
    useEffect(()=> {
        try {
            const headers = {
               'Content-Type': 'text/plain',
            };
            fetch("/actividades",{
                method: 'GET',
                headers: headers
            }).then((res)=> {
                return(res.json());
            }).then((data)=>{
                setActividades(data);
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



    // Solo queremos mostrar los primeros 3 (si es que hay) en los "destacados"
    var firstThreeItems = []
    if (Actividades.length > 3) {
        firstThreeItems = Actividades.slice(0, 3);
        setActividades(firstThreeItems)
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
            {Actividades.map((actividad)=>{
                return(
               <div key={actividad.id} className="activity-card">
                    <img src={pool}  width={"100%"}></img>
                    <div className="container">
                        <h3>{actividad.name}</h3>
                        <p>Tipo de Actividad: {actividad.activitytype}</p>
                        <button onClick={ ()=> navigation("/Actividad",{state:{id:actividad.id}})}>Realizar Inscripcion</button>
                    </div>
                </div>

            )})}
            </div>
        </div>
    )
};

export default Home;