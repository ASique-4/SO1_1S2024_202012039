import React from 'react';
import { LineChart, Line, XAxis, YAxis, CartesianGrid, Tooltip, Legend } from 'recharts';

const MonitoreoHistorico = () => {
    const data = [
        { name: 'Enero', value1: 100, value2: 200 },
        { name: 'Febrero', value1: 150, value2: 220 },
        { name: 'Marzo', value1: 200, value2: 250 },
        // Agrega más datos aquí
    ];

    return (
        <div className="chart-container" style={{ display: "flex", justifyContent: "center" }}>
            <div className="line-chart">
            <h1>RAM</h1>
            <LineChart width={600} height={300} data={data}>
                <CartesianGrid strokeDasharray="3 3" />
                <XAxis dataKey="name" />
                <YAxis />
                <Tooltip />
                <Legend />
                <Line type="monotone" dataKey="value1" stroke="#8884d8" />
            </LineChart>
            </div>

            <div className="line-chart">
            <h1>CPU</h1>
            <LineChart width={600} height={300} data={data}>
                <CartesianGrid strokeDasharray="3 3" />
                <XAxis dataKey="name" />
                <YAxis />
                <Tooltip />
                <Legend />
                <Line type="monotone" dataKey="value2" stroke="#82ca9d" />
            </LineChart>
        </div>
        </div>
    );
};

export default MonitoreoHistorico;