import React from 'react';

const CambioEstados = () => {
    const [pid, setPid] = React.useState('');

    const handleNew = () => {
        // Add your logic for the "New" button here
    };

    const handleStop = () => {
        // Add your logic for the "Stop" button here
    };

    const handleReady = () => {
        // Add your logic for the "Ready" button here
    };

    const handleKill = () => {
        // Add your logic for the "Kill" button here
    };

    return (
        <div className="button-container">
            <button className="new-button" onClick={handleNew}>New</button>
            <button className="stop-button" onClick={handleStop}>Stop</button>
            <button className="ready-button" onClick={handleReady}>Ready</button>
            <button className="kill-button" onClick={handleKill}>Kill</button>
            <p className="pid-text">PID: {pid}</p>
        </div>
    );
};

export default CambioEstados;