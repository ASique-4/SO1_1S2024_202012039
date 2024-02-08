const express = require('express');
const cors = require('cors');
const app = express();

app.use(cors()); // Usa CORS en todas las rutas
app.use(express.json({ limit: '50mb' })); // Aumenta el límite a 50MB

app.post('/upload', (req, res) => {
    const base64Image = req.body.image;
    // Aquí puedes procesar la imagen en base64 como desees
    // Por ejemplo, puedes guardarla en el sistema de archivos o enviarla a un servicio externo
    
    // Envía una respuesta de éxito
    res.status(200).json({ message: 'Imagen recibida correctamente' });
});

app.listen(8080, () => {
    console.log('Servidor escuchando en el puerto 8080');
});