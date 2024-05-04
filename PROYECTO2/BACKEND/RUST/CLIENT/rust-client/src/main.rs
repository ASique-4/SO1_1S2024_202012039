use rocket::{post, serde::json::Json};
use reqwest::Client;
use serde::{Deserialize, Serialize};
use rocket::config::SecretKey;
use rocket_cors::{AllowedOrigins, CorsOptions};

#[derive(Debug, Serialize, Deserialize)]
struct Data {
    name: String,
    album: String,
    year: String,
    rank: String,
}

#[post("/insert", data = "<data>")] // Cambiado a "/insert"
async fn send_data(data: Json<Data>) -> String {
    let client = Client::new();
    // Cambiado el URL del servidor a la ruta correcta
    let server_url = "http://localhost:8080/rust/insert"; // Cambiado a la ruta correcta
    let response = client.post(server_url).json(&data.into_inner()).send().await;

    match response {
        Ok(_) => "Data sent successfully!".to_string(),
        Err(e) => format!("Failed to send data: {}", e),
    }
}

#[rocket::main]
async fn main() {
    let secret_key = SecretKey::generate(); // Genera una nueva clave secreta

    // Configuración de opciones CORS
    let cors = CorsOptions::default()
        .allowed_origins(AllowedOrigins::all())
        .to_cors()
        .expect("failed to create CORS fairing");

    let config = rocket::Config {
        address: "0.0.0.0".parse().unwrap(),
        port: 8000,
        secret_key: secret_key.unwrap(), // Desempaqueta la clave secreta generada
        ..rocket::Config::default()
    };

    rocket::custom(config)
        .attach(cors)
        .mount("/rust", rocket::routes![send_data]) // Cambiado a "/rust/insert"
        .launch()
        .await
        .unwrap();
}
