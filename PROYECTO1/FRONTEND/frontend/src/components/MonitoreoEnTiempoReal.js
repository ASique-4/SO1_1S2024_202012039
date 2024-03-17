import React, { useEffect, useState } from "react";
import { PieChart, Pie, Cell, Tooltip } from "recharts";
import axios from 'axios';

const COLORS = [
    "#C7B299",
    "#a9957d",
    "#675640",
    "#A68B5B",
    "#483507",
    "#4D4D4D",
    "#797979",
    "#F5EFE7",
    "#ebe5dd",
    "#c2bdb5",
];

const MonitoreoEnTiempoReal = () => {
    const [ramData, setRamData] = useState([]);
    const [cpuData, setCpuData] = useState([]);

    useEffect(() => {
        const intervalRam = setInterval(() => {
            axios.get('http://192.168.1.42:8080/ram')
                .then(response => {
                    const data = response.data;
                    setRamData([
                        { name: 'En Uso', value: data.memoriaEnUso },
                        { name: 'Libre', value: data.libre },
                    ]);
                    console.log('RAM data:', data);
                })
                .catch(error => {
                    console.error('Error fetching RAM data:', error);
                });
        }, 1000); // Actualiza cada 3 segundos

        const intervalCpu = setInterval(() => {
            axios.get('http://192.168.1.42:8080/cpu')
                .then(response => {
                    const data = response.data;
                    setCpuData([
                        { name: 'En Uso', value: data.cpu_uso },
                        { name: 'Libre', value: data.cpu_libre },
                    ]);
                    console.log('CPU data:', data);
                })
                .catch(error => {
                    console.error('Error fetching CPU data:', error);
                });
        }, 1000); // Actualiza cada 3 segundos

        // Limpia los intervalos cuando el componente se desmonta
        return () => {
            clearInterval(intervalRam);
            clearInterval(intervalCpu);
        };
    }, []);

    return (
        <div className="chart-container" style={{ display: "flex", justifyContent: "center" }}>
            <div className="pie-chart">
                <h1>RAM</h1>
                <PieChart width={400} height={400}>
                    <Pie
                        data={ramData}
                        cx={200}
                        cy={200}
                        outerRadius={80}
                        fill="#8884d8"
                        dataKey="value"
                    >
                        {ramData.map((entry, index) => (
                            <Cell key={`cell-${index}`} fill={COLORS[index % COLORS.length]} />
                        ))}
                    </Pie>
                    <Tooltip />
                </PieChart>
            </div>
            <div className="pie-chart">
                <h1>CPU</h1>
                <PieChart width={400} height={400}>
                    <Pie
                        data={cpuData}
                        cx={200}
                        cy={200}
                        outerRadius={80}
                        fill="#8884d8"
                        dataKey="value"
                    >
                        {cpuData.map((entry, index) => (
                            <Cell key={`cell-${index}`} fill={COLORS[index % COLORS.length]} />
                        ))}
                    </Pie>
                    <Tooltip />
                </PieChart>
            </div>
        </div>
    );
};
export default MonitoreoEnTiempoReal;