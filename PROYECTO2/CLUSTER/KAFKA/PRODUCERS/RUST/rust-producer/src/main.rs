use rocket::response::status::BadRequest;
use rocket::serde::json::Json;
use rocket::config::SecretKey;
use rocket_cors::{AllowedOrigins, CorsOptions};
use rdkafka::producer::{FutureProducer, FutureRecord};
use rdkafka::config::ClientConfig;
use serde_json::Value;
use std::env;

#[derive(serde::Deserialize)]
struct KafkaMessage {
    topic: String,
    message: Value,
}

#[rocket::post("/sendMessage", data = "<data>")]
async fn send_message(data: Json<KafkaMessage>) -> Result<String, BadRequest<String>> {
    let request_data = data.into_inner();

    let kafka_brokers = match env::var("KAFKA_BROKERS") {
        Ok(brokers) => brokers.split(',').map(|s| s.to_owned()).collect(),
        Err(_) => vec!["localhost:9092".to_owned()], 
    };

    let producer: FutureProducer = ClientConfig::new()
        .set("bootstrap.servers", &kafka_brokers.join(","))
        .create()
        .expect("Error creating Kafka producer");

    let message_string = request_data.message.to_string();

    let result = producer.send(
        FutureRecord::to(&request_data.topic)
            .payload(&message_string)
            .key("key"),
        std::time::Duration::from_secs(0),
    ).await;

    match result {
        Ok(_) => {
            Ok("Mensaje enviado correctamente a Kafka. :D".to_string())
        }
        Err(e) => {
            Err(BadRequest(format!("Failed to send message to Kafka: {:?}", e)))
        }
    }
}

#[rocket::main]
async fn main() {
    let secret_key = SecretKey::generate();

    let cors = CorsOptions::default()
        .allowed_origins(AllowedOrigins::all())
        .to_cors()
        .expect("failed to create CORS fairing");

    let config = rocket::Config {
        address: "0.0.0.0".parse().unwrap(),
        port: 5004,
        secret_key: secret_key.unwrap(),
        ..rocket::Config::default()
    };

    rocket::custom(config)
        .attach(cors)
        .mount("/", rocket::routes![send_message])
        .launch()
        .await
        .unwrap();
}
