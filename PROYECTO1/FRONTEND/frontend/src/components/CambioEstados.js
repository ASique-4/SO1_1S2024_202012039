import React, { useState } from 'react';
import ReactFlow from 'react-flow-renderer';

const CambioEstados = () => {
    const [pid, setPid] = useState('');
    const [elements, setElements] = useState([]);

    const handleNew = async () => {
        const response = await fetch('/procesos/iniciar', { method: 'POST' });
        const data = await response.json();
        setPid(data.pid);
        setElements((els) => [
            ...els,
            { id: 'new', data: { label: `New: ${data.pid}` }, position: { x: 0, y: 0 } },
        ]);
    };

    const handleStop = async () => {
        await fetch('/procesos/detener', { method: 'POST', body: JSON.stringify({ pid }) });
        setElements((els) => [
            ...els,
            { id: 'stop', data: { label: 'Stop' }, position: { x: 100, y: 0 }, sourcePosition: 'right', targetPosition: 'left' },
            { id: 'e1-2', source: 'new', target: 'stop', animated: true },
        ]);
    };

    const handleReady = async () => {
        await fetch('/procesos/reanudar', { method: 'POST', body: JSON.stringify({ pid }) });
        setElements((els) => [
            ...els,
            { id: 'ready', data: { label: 'Ready' }, position: { x: 200, y: 0 }, sourcePosition: 'right', targetPosition: 'left' },
            { id: 'e2-3', source: 'stop', target: 'ready', animated: true },
        ]);
    };

    const handleKill = async () => {
        await fetch('/procesos/matar', { method: 'POST', body: JSON.stringify({ pid }) });
        setElements((els) => [
            ...els,
            { id: 'kill', data: { label: 'Kill' }, position: { x: 300, y: 0 }, sourcePosition: 'right', targetPosition: 'left' },
            { id: 'e3-4', source: 'ready', target: 'kill', animated: true },
        ]);
    };

    return (
        <div className="button-container">
            <button className="new-button" onClick={handleNew}>New</button>
            <button className="stop-button" onClick={handleStop}>Stop</button>
            <button className="ready-button" onClick={handleReady}>Ready</button>
            <button className="kill-button" onClick={handleKill}>Kill</button>
            <p className="pid-text">PID: {pid}</p>
            <ReactFlow elements={elements} />
        </div>
    );
};

export default CambioEstados;