
import { useContext, useState } from 'react'
import './Navbar.css'
import {UserTypeContext} from "../App"
import gymlogo from "../assets/gymlogo.png"


export default function Navbar(props) {

    const [userType,setuserType] = useContext(UserTypeContext);
    const [visibility,seyvisibility] = useState(true)
    function notloggedelements() { // not logged in
        return (
            <>
                <a className="navbar-element" href='/login'>Log in</a>
                <a className="navbar-element" href='/register'>Register</a>
            </>
        )
    }
    
    
    function loggedelements() {  //  logged in
        return (
            <>
                <a className="navbar-element" href='/Profile'>Profile</a>
            </>
        )
    }

    window.addEventListener("scroll", (event) => {
    let scroll = window.pageYOffset
    if (scroll > 10) {seyvisibility(false)} else {seyvisibility(true)}
    });

    return (
        <div className={`navbar ${visibility ? "visible":"hidden"}`}>
        <nav className="navbar-container">
            <div className="navbar-items"><a href="/"><img src={gymlogo} width={"100px"}></img></a></div>
            <div className="navbar-items">
                <li>
                    <a className="navbar-element" href='/actividades'>Actividades</a>
                    <a className="navbar-element" href='/actividades'>Nosotros</a>
                    <a className="navbar-element" href='/actividades'>Contacto</a>
                    <a className="navbar-element"></a>
                </li>
            </div>
            <div className="navbar-items">
                {userType != null ? loggedelements() : notloggedelements()}
            </div>
        </nav>
        </div>
    )
}