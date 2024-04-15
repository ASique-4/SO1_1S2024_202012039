import React from 'react';
import { Link } from 'react-router-dom';

const Navbar = () => {
    return (
        <nav className="navbar">
            <ul className="navbar-list">
                <li className="navbar-item">
                    <Link to="/home" className="navbar-link">Home</Link>
                </li>
                <div className="navbar-divider"></div>
                <li className="navbar-item">
                    <Link to="/MonitoreoEnTiempoReal" className="navbar-link">Monitoreo En Tiempo Real</Link>
                </li>
                <div className="navbar-divider"></div>
                <li className="navbar-item">
                    <Link to="/MonitoreoHistorico" className="navbar-link">Monitoreo Historico</Link>
                </li>
                <div className="navbar-divider"></div>
                <li className="navbar-item">
                    <Link to="/ArbolProcesos" className="navbar-link">Árbol de Procesos</Link>
                </li>
                <div className="navbar-divider"></div>
                <li className="navbar-item">
                    <Link to="/CambioEstados" className="navbar-link">Simulación de Cambios de Estado</Link>
                </li>
            </ul>
        </nav>
    );
};

export default Navbar;
