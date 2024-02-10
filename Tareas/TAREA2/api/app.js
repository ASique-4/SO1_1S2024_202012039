const express = require('express');
const cors = require('cors');
const mongoose = require('mongoose');
const app = express();


mongoose.connect(`mongodb://MongoDB:27017/tarea2-db`, { useNewUrlParser: true, useUnifiedTopology: true })
    .then(() => console.log('Conexión a MongoDB exitosa'))
    .catch(err => console.error('Error de conexión a MongoDB:', err));

// Define el modelo de imagen
const Foto = mongoose.model('Foto', new mongoose.Schema({
    image: String,
    date: String
}), 'fotos');

app.use(cors()); // Usa CORS en todas las rutas
app.use(express.json({ limit: '50mb' })); // Aumenta el límite a 50MB

app.post('/upload', (req, res) => {
    const base64Image = req.body.image;

    // Crea una nueva imagen y la guarda en la base de datos
    const foto = new Foto({ image: base64Image });
    foto.save()
        .then(() => res.status(200).json({ message: 'Imagen recibida y guardada correctamente' }))
        .catch(err => res.status(500).json({ message: 'Error al guardar la imagen', error: err }));
});

app.listen(8080, () => {
    console.log(`Servidor escuchando en 8080`);
});