import axios from 'axios';
import { useEffect, useState } from 'react';
import { LineChart, Line, CartesianGrid, XAxis, YAxis, Tooltip, Legend } from 'recharts';

const MonitoreoHistorico = () => {
    const [cpuData, setCpuData] = useState([]);
    const [ramData, setRamData] = useState([]);

    useEffect(() => {
        const fetchData = () => {
            axios.get('http://localhost:8080/ram/all')
                .then(response => {
                    setRamData(response.data.map(item => ({
                        name: item.fechaHora,
                        Uso: item.porcentaje
                    })));
                    console.log('RAM data:', "correcto");
                })
                .catch(error => {
                    console.error('Error fetching RAM data:', error);
                });

            axios.get('http://localhost:8080/cpu/all')
                .then(response => {
                    setCpuData(response.data.map(item => ({
                        name: item.fechaHora,
                        Uso: item.cpu_porcentaje
                    })));

                    console.log('CPU data:', "correcto");
                })
                .catch(error => {
                    console.error('Error fetching CPU data:', error);
                });
        };

        fetchData(); // Fetch data immediately
        const intervalId = setInterval(fetchData, 10000); // Fetch data every 10 seconds

        return () => clearInterval(intervalId); // Clean up on component unmount
    }, []);

    return (
        <div className="chart-container" style={{ display: "flex", justifyContent: "center" }}>
            <div className="line-chart">
                <h1>RAM</h1>
                <LineChart width={600} height={300} data={ramData}>
                    <CartesianGrid strokeDasharray="3 3" />
                    <XAxis dataKey="name" tick={{ fontSize: 15 }} />
                    <YAxis />
                    <Tooltip />
                    <Legend />
                    <Line type="monotone" dataKey="Uso" stroke="#8884d8" />
                </LineChart>
            </div>

            <div className="line-chart">
                <h1>CPU</h1>
                <LineChart width={600} height={300} data={cpuData}>
                    <CartesianGrid strokeDasharray="3 3" />
                    <XAxis dataKey="name" tick={{ fontSize: 15 }} />
                    <YAxis />
                    <Tooltip />
                    <Legend />
                    <Line type="monotone" dataKey="Uso" stroke="#82ca9d" />
                </LineChart>
            </div>
        </div>
    );
};

export default MonitoreoHistorico;