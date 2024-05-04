use rocket::{serde::json::Json, response::status::BadRequest};
use rocket_cors::{AllowedOrigins, CorsOptions};
use reqwest::Error;
use serde::{Serialize, Deserialize};
use rocket::config::SecretKey;
use std::env;

#[derive(Serialize, Deserialize, Debug)]
struct Data {
    name: String,
    album: String,
    year: String,
    rank: String
}

async fn enviar_json(data: &Data) -> Result<(), Error> {
    
    let target_url = env::var("ENV_RUTA_SERVER_GO").expect("ENV_RUTA_SERVER_GO environment variable not set");

    let message = serde_json::json!({
        "topic": "topic-sopes1",
        "message": data
    });

    let client = reqwest::Client::new();
    client.post(&target_url)
        .header("Content-Type", "application/json")
        .json(&message)
        .send()
        .await?;

    Ok(())
}


#[rocket::post("/insert", data = "<data>")]
async fn receive_data(data: Json<Data>) -> Result<String, BadRequest<String>> {

    // Print data received
    println!("Received data: Name: {}, Album: {}, Year: {}, Rank: {}",
                        data.name, data.album, data.year, data.rank);

    let received_data = data.into_inner();

    match enviar_json(&received_data).await {
        Ok(_) => Ok("Data received successfully!".to_string()),
        Err(_) => Err(BadRequest::from(rocket::response::status::BadRequest("Failed to send data".to_string()))),
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
        port: 8080,
        secret_key: secret_key.unwrap(), // Desempaqueta la clave secreta generada
        ..rocket::Config::default()
    };

    // Montar la aplicación Rocket con el middleware CORS
    rocket::custom(config)
        .attach(cors)
        .mount("/rust", rocket::routes![receive_data]) // Cambiado a "/rust"
        .launch()
        .await
        .unwrap();
}
