import React, { useEffect, useRef, useState } from 'react';
import { DataSet } from 'vis-data';
import { Network } from 'vis-network';
import axios from 'axios';

const CambioEstados = () => {
    const networkRef = useRef(null);
    const [nodes, setNodes] = useState(new DataSet([]));
    const [edges, setEdges] = useState(new DataSet([]));
    const [pid, setPid] = useState(null);

    const handleNew = async () => {
        const response = await axios.post('http://192.168.1.42:8080/procesos/iniciar');
        setPid(response.data.pid);
    
        nodes.add({ id: 'new', label: 'New', color: '#f44336' }); 
        nodes.add({ id: 'ready', label: 'Ready', color: '#4caf50' }); 
        nodes.add({ id: 'running', label: 'Running', color: '#2196f3' }); 

        edges.add({ id: 'new-ready', from: 'new', to: 'ready' });
        edges.add({ id: 'ready-running', from: 'ready', to: 'running' });
    };

    const handleStop = async () => {
        await axios.post('http://192.168.1.42:8080/procesos/detener?pid=' + pid);

        edges.add({ id: 'running-ready', from: 'running', to: 'ready' });
    };

    const handleResume = async () => {
        await axios.post('http://192.168.1.42:8080/procesos/reanudar?pid=' + pid);

        const edgeId = edges.getIds().find(id => id.startsWith('running-ready'));
        edges.update({ id: edgeId, color: { color: 'red' } });
    };

    const handleKill = async () => {
        await axios.post('http://192.168.1.42:8080/procesos/matar?pid=' + pid);

        nodes.add({ id: 'terminated', label: 'Terminated', color: '#9e9e9e' }); 
        edges.add({ id: 'running-terminated', from: 'running', to: 'terminated' });
    };

    const handleClean = () => {
        setNodes(new DataSet([]));
        setEdges(new DataSet([]));
        setPid(null);
    }

    useEffect(() => {
        const container = networkRef.current;

        const data = {
            nodes,
            edges,
        };

        const options = {
            // specify your network options here
        };

        new Network(container, data, options);
    }, [nodes, edges]); // Agrega nodes y edges como dependencias del useEffect


    return (
        <div className="button-container">
            <button className="new-button" onClick={handleNew}>New</button>
            <button className="stop-button" onClick={handleStop}>Stop</button>
            <button className="ready-button" onClick={handleResume}>Resume</button>
            <button className="kill-button" onClick={handleKill}>Kill</button>
            <button className="clean-button" onClick={handleClean}>Clean</button>
            <p className="pid-text">PID: {pid}</p>
            <div ref={networkRef} style={{ height: '400px' }}></div>
        </div>
    );
};

export default CambioEstados;