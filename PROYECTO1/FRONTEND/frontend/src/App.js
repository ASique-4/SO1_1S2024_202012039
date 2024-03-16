import React from 'react';
import { BrowserRouter as Router, Route, Routes } from 'react-router-dom';
import Home from './components/Home';
import Navbar from './components/Navbar';
import MonitoreoEnTiempoReal from './components/MonitoreoEnTiempoReal';
import MonitoreoHistorico from './components/MonitoreoHistorico';
import Arbol from './components/Arbol';
import CambioEstados from './components/CambioEstados';

function App() {
  return (
    <Router>
      <Navbar />
      <Routes>
        <Route path='/' element={<Home />} />
        <Route path='/home' element={<Home />} />
        <Route path='/MonitoreoEnTiempoReal' element={<MonitoreoEnTiempoReal />} />
        <Route path='/MonitoreoHistorico' element={<MonitoreoHistorico />} />
        <Route path='/ArbolProcesos' element={<Arbol />} />
        <Route path='/CambioEstados' element={<CambioEstados />} />
      </Routes>
    </Router>
  );
}

export default App;