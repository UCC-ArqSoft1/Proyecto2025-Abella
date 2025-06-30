import { useSearchParams } from 'react-router-dom';
import { Link, useNavigate } from "react-router-dom";
import { useContext,useEffect,useState } from "react";
import {UserTypeContext} from "../App"
import './styles/activities-list.css';
import banner from '../assets/activities-banner.jpg'

export default function Actividades() {
    const [searchParams,setSearchParams] = useSearchParams();
    const [selectedActivity,setselectedActivity] = useState(null)
    const [Actividades, setActividades] = useState([])
    const [userType,setuserType] = useContext(UserTypeContext);
    const [profesores,setprofesores] = useState([])
    const [categorias,setcategorias] = useState([])
    const keyword = searchParams.get('keyword');
    const navigation = useNavigate();
    const [iscreating,setiscreating] = useState(false)

    // Variables
    var url = "/actividades"
    keyword != null ? url = url+'?keyword='+keyword : url = url
    // Final declaracion de variables

    const headers = {
        'Authorization': localStorage.getItem('userToken'),
       'Content-Type': 'text/plain',
    };

    function updateSearch() {
                    try {
                fetch(url,{
                    method: 'GET',
                    headers: headers
                }).then((res)=> {

                    return(res.json());
                }).then((data)=>{
                    console.log(data)
                    setActividades(data);
                })
            } catch (error) {
                console.log("404:Could not Fetch")
                return
            }
    }

    function updatesearchparam() {
        const keyword2 = document.getElementById("search-input").value
        keyword2 != null ? url = "/actividades"+'?keyword='+keyword2 : url = url
        console.log(url)
        updateSearch()
    }


    useEffect(()=> {
        updateSearch()
        }, []);
 function Horarios(props) {

        var close = false;
        if (close == false) {
        return (
            <div onClick={() => {setselectedActivity(null)}} className="Inscripcion-horarios">
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

    function FormularioCrearActividad() {
    const headers_Create = {
        'Authorization': localStorage.getItem('userToken'),
       'Content-Type': 'text/plain',
    };
        try {
            fetch("/users/profesores",{
                method:"GET",
                headers:headers_Create
            }).then((res)=>{
                return(res.json());
            }).then((data)=>{
                console.log(data)
                setprofesores(data)
            })
        } catch (error) {
            alert("Error al realizar fetch de profesores")
        }

        try {
            fetch("/activities/types",{
                method:"GET",
                headers:headers_Create
            }).then((res)=>{
                return(res.json());
            }).then((data)=>{
                console.log(data)
                setcategorias(data)
            })
        } catch (error) {
            alert("Error al realizar fetch de Categorias")
        }
        return(
            <div className='CreateActivityForm-container'>
                <form className='CreateActivityForm'>
                    <h5>Formulario Crear Actividad</h5>
                    <input placeholder='Nombre Actividad'></input>
                    <input placeholder='Descripcion'></input>
                    <p>Profesores:</p>
                    <select id="profesores_select">
                        {profesores?.map((Profesor)=>{
                            return(
                                <option value={Profesor.id}>{Profesor.name}</option>
                            )
                        })}
                    </select>
                    <p>Categoria:</p>
                    <select id="categorias_select">
                        {categorias?.map((Categoria)=>{
                            return(
                                <option value={Categoria.id}>{Categoria.name}</option>
                            )
                        })}
                    </select>
                    <input placeholder='Nombre Actividad'></input>
                    <button>Crear Actividad</button>
                    <button onClick={()=>{setiscreating(false)}}>Cancelar</button>
                    </form>
            </div>
        )
    }



    function Editbtn() {
        return (
            <button>Editar</button>
        )
    }

    function divider() {
        return (
            <hr className='solid'></hr>
        )
    }

    let quantity = 0
      if (Actividades != null) {
          quantity = Actividades.length
      }

    return ( 
        <div className="activities-page">
            <h1>Empieza a entrenar hoy</h1>
            <img width={"90%"} src={banner}></img>
            <div className="activities-content-table">
                <div className='search-side-menu'>
                    <p>{quantity} Resultados</p>
                    <input placeholder='Buscar por palabra clave o categoria' id='search-input'></input>
                    <button onClick={()=> {updatesearchparam()}}>Buscar</button>
                    <button>Quitar Filtros</button>
                    {userType == 2 && <button onClick={()=>{setiscreating(true)}}>Añadir actividad</button>}
                    {iscreating == true && <FormularioCrearActividad/>}
                </div>
                <div className='activities-list'>
                    {Actividades?.map((actividad)=>{
                    return(
                    <div>
                    <div key={actividad.id} className="activity-item">
                        <div className='activity-props'>
                            <h3>{actividad.name}</h3>
                            <p>{actividad.description}</p>
                            <div className='Activity-actors'>
                                <p>Tipo de Actividad: {actividad.activitytype}</p>
                                <p>Profesor: {actividad.coach_name}</p>
                            </div>
                        </div>
                        <button onClick={ ()=> navigation("/Actividad",{state:{id:actividad.id}})}>Inscribirse</button>
                        {userType == 2 ? Editbtn() : ()=>{return ""}}
                    </div>
                    {quantity > 1 ? divider() : null}
                    </div>
                    )})}
                </div>
            </div>
        </div>
    )
}

