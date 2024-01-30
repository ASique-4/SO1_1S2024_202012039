import React, { useState } from 'react';
import "../App.css";

const UserData = () => {
  const [usuario, setUsuario] = useState([]);

  function getData() {
    fetch("http://localhost:8080/data")
      .then((response) => response.json())
      .then((data) => {
        setUsuario([data]);   
        console.log(data);     
      })
      .catch((error) => console.log(error));
  }

  return (
    <div>
      <center>
        <button className='obtenerDatos' onClick={getData}>Get Data</button>
        <div className="mis-datos">
          <table>
            <thead>
              <tr>
                <th>Nombre</th>
                <th>Apellido</th>
                <th>Edad</th>
                <th>Carnet</th>
              </tr>
            </thead>
            <tbody>
              {usuario.map((user) => (
                <tr key={user.carnet}>
                  <td>{user.nombres}</td>
                  <td>{user.apellidos}</td>
                  <td>{user.edad}</td>
                  <td>{user.carnet}</td>
                </tr>
              ))}
            </tbody>
          </table>
        </div>
      </center>
    </div>
  );
};

export default UserData;
