import { useSearchParams } from 'react-router-dom';
import { Link, useNavigate } from "react-router-dom";
import { useContext,useEffect,useState } from "react";
import {UserTypeContext} from "../App"
import './styles/activities-list.css';
import banner from '../assets/activities-banner.jpg'
import FormularioCrearActividad from '../Components/FormularioCrearActividad';
import FormularioEditarActividad from '../Components/FormularioEditarActividad';

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
    const [isediting,setisediting] = useState(false)

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

    function DeleteActivity(actId) {
        try {
            fetch("/actividades/"+actId+"/delete",{
                method:'POST',
                headers: headers
            }).then((res)=>{
            if (res.status == 200) {
                alert("Actividad borrada con exito")
            }
            if (res.status == 401) {
                console.log("token is not valid")
                localStorage.removeItem('userToken')
                alert("Debes iniciar sesion nuevamente")
                navigation("/Login")
            }
            })
        } catch (error) {
            console.log(error)
        }
    }

    useEffect(()=> {
        updateSearch()
    }, []);

    function Editbtn() {
        return (
            <button onClick={()=>{setisediting(true)}}>Editar</button>
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
                    {iscreating == true && <FormularioCrearActividad setiscreating={setiscreating}/>}
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
                        <button onClick={ ()=> navigation("/Actividad",{state:{id:actividad.id}})}>Detalles</button>
                        {userType == 2 ? Editbtn() : ()=>{return""}}
                        {userType == 2 && <button onClick={()=>{DeleteActivity(actividad.id)}}>Borrar</button>}
                        {isediting == true && <FormularioEditarActividad setiscreating={setisediting} idactividad={actividad.id} name={actividad.name} description={actividad.description} duration={actividad.duration} capacity={actividad.capacity}/>}
                    </div>
                    {quantity > 1 ? divider() : null}
                    </div>
                    )})}
                </div>
            </div>
        </div>
    )
}


