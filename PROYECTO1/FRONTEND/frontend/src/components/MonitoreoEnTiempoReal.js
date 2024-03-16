import React from "react";
import { PieChart, Pie, Cell } from "recharts";

const data = [
    { name: "A", value: 400 },
    { name: "B", value: 300 },
    { name: "C", value: 300 },
    { name: "D", value: 200 },
];

const data2 = [
    { name: "A", value: 100 },
    { name: "B", value: 100 },
    { name: "C", value: 100 },
    { name: "D", value: 100 },
];

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
    return (
        <div className="chart-container" style={{ display: "flex", justifyContent: "center" }}>
            <div className="pie-chart" >
                <h1>RAM</h1>
                <PieChart width={400} height={400}>
                    <Pie
                        data={data}
                        cx={200}
                        cy={200}
                        outerRadius={80}
                        fill="#8884d8"
                        dataKey="value"
                    >
                        {data.map((entry, index) => (
                            <Cell
                                key={`cell-${index}`}
                                fill={COLORS[index % COLORS.length]}
                            />
                        ))}
                    </Pie>
                </PieChart>
            </div>
            <div className="pie-chart">
                <h1>CPU</h1>
                <PieChart width={400} height={400}>
                    <Pie
                        data={data2}
                        cx={200}
                        cy={200}
                        outerRadius={80}
                        fill="#8884d8"
                        dataKey="value"
                    >
                        {data2.map((entry, index) => (
                            <Cell
                                key={`cell-${index}`}
                                fill={COLORS[index % COLORS.length]}
                            />
                        ))}
                    </Pie>
                </PieChart>
            </div>
        </div>
    );
};

export default MonitoreoEnTiempoReal;
