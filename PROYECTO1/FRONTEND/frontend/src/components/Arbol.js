import React, { useState, useEffect } from 'react';
import { Select, Input } from 'antd';
import axios from 'axios';
import { Network } from 'vis-network';

const { Option } = Select;
const { Search } = Input;

const VisjsTree = () => {
    const [processes, setProcesses] = useState([]);
    const [selectedPid, setSelectedPid] = useState(null);
    const [selectedProcess, setSelectedProcess] = useState(null);
    const [network, setNetwork] = useState(null);
    const [searchTerm, setSearchTerm] = useState('');

    useEffect(() => {
        axios.get('http://192.168.1.42:8080/procesos')
            .then(response => {
                setProcesses(response.data);
            })
            .catch(error => {
                console.error('Error fetching processes:', error);
            });
    }, []);

    const handlePidChange = (value) => {
        setSelectedPid(value);
        const selected = processes.find(process => process.pid === value);
        setSelectedProcess(selected);
    };

    const handleSearch = (value) => {
        setSearchTerm(value);
        const nodeId = processes.find(process => process.name === value)?.pid;
        if (nodeId) {
            network.selectNodes([nodeId]);
        } else {
            network.unselectAll();
        }
    };

    useEffect(() => {
        if (!selectedProcess || !network) return;

        const { nodes, edges } = generateData(selectedProcess);

        network.setData({ nodes, edges });
    }, [selectedProcess, network]);

    useEffect(() => {
        const container = document.getElementById('vis-tree');
        const options = {
            layout: {
                hierarchical: {
                    direction: 'UD',
                    sortMethod: 'directed',
                    nodeSpacing: 100, // Espaciado horizontal entre nodos
                },
            },
        };
        const newNetwork = new Network(container, {}, options);
        setNetwork(newNetwork);
        return () => {
            newNetwork.destroy();
        };
    }, []);

    const generateData = (process) => {
        const nodes = [];
        const edges = [];

        const processNode = { id: process.pid, label: process.name };
        nodes.push(processNode);

        if (process.child) {
            process.child.forEach(child => {
                const childNode = { id: child.pid, label: child.name };
                const edge = { from: process.pid, to: child.pid };
                nodes.push(childNode);
                edges.push(edge);
            });
        }

        return { nodes, edges };
    };

    return (
        <div style={{ display: 'flex', flexDirection: 'column', alignItems: 'center', justifyContent: 'center' }}>
            <Select
                placeholder="Seleccionar PID padre"
                onChange={handlePidChange}
                style={{ width: 200, marginBottom: 16 }}
            >
                {processes.map((process) => (
                    <Option key={process.pid} value={process.pid}>
                        {process.name}
                    </Option>
                ))}
            </Select>
            <Search
                placeholder="Buscar nodo"
                onSearch={handleSearch}
                style={{ width: 200, marginBottom: 16 }}
            />
            <div id="vis-tree" style={{ height: '600px', width: '80%' }}></div>
        </div>
    );
};

export default VisjsTree;