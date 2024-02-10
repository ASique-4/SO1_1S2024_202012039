import React, { useRef, useState } from 'react';

function App() {
  const videoRef = useRef(null);
  const canvasRef = useRef(null);
  const [base64Image, setBase64Image] = useState(null);

  const startCamera = async () => {
    const constraints = { video: true };
    const stream = await navigator.mediaDevices.getUserMedia(constraints);
    videoRef.current.srcObject = stream;
    videoRef.current.play();
  };

  const takePhoto = () => {
    const context = canvasRef.current.getContext('2d');
    context.drawImage(videoRef.current, 0, 0, canvasRef.current.width, canvasRef.current.height);
    const base64String = canvasRef.current.toDataURL('image/png').split(",")[1];
    setBase64Image(base64String);

    
    fetch('http://localhost:8080/upload', {
      method: 'POST',
      headers: {
        'Content-Type': 'application/json'
      },
      body: JSON.stringify({ 
        image: base64Image,
        date: new Date().toISOString()
      })
    })
      .then(response => response.json())
      .then(data => {
        // Handle the response data here
        console.log(data);
      })
      .catch(error => {
        // Handle any errors here
      });
  };

  return (
    <div className="App">
      <center>
        <div className='video-container'>
          <video className='video' ref={videoRef} width="640" height="480" />
          <canvas className='canvas' ref={canvasRef} width="640" height="480" />
        </div>
        <br></br>
        <button className='camara-btn' onClick={startCamera}>Iniciar cámara</button>
        <button className='foto-btn' onClick={takePhoto}>Tomar foto</button>
      </center>
    </div>
  );
}

export default App;