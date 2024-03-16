import React, { useState } from 'react';
import { Select } from 'antd';

const { Option } = Select;

const Arbol = () => {
    const [selectedPid, setSelectedPid] = useState(null);

    const handlePidChange = (value) => {
        setSelectedPid(value);
        console.log(`Selected PID: ${selectedPid}`);
    };

    // Datos de ejemplo para el árbol
    const treeData = [
        {
            name: 'Nodo 1',
            children: [
                {
                    name: 'Nodo 1.1',
                    children: [
                        { name: 'Nodo 1.1.1' },
                        { name: 'Nodo 1.1.2' },
                    ],
                },
                {
                    name: 'Nodo 1.2',
                    children: [
                        { name: 'Nodo 1.2.1' },
                        { name: 'Nodo 1.2.2' },
                    ],
                },
            ],
        },
        {
            name: 'Nodo 2',
            children: [
                { name: 'Nodo 2.1' },
                { name: 'Nodo 2.2' },
            ],
        },
    ];

    return (
        <div>
            <Select
                placeholder="Seleccionar PID padre"
                onChange={handlePidChange}
                style={{ width: 200, marginBottom: 16 }}
            >
                {treeData.map((node) => (
                    <Option key={node.name} value={node.name}>
                        {node.name}
                    </Option>
                ))}
            </Select>
        </div>
    );
};

export default Arbol;
