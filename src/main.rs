mod response;
mod routes;
mod utils;

use actix_web::{middleware, App, HttpServer};
use routes::index;
use std::env;

#[actix_rt::main]
async fn main() -> std::io::Result<()> {
    let port = env::var("PORT")
        .unwrap_or_else(|_| "8080".to_string())
        .parse::<u16>()
        .expect("PORT must be a number");

    println!("\n=== STARTING SERVER ===\nhttp://localhost:{port}\n");

    env::set_var("RUST_LOG", "info,scrape=debug");
    pretty_env_logger::init();

    HttpServer::new(|| {
        App::new()
            .wrap(middleware::Logger::new("IP - %a | Time - %D ms"))
            .wrap(middleware::DefaultHeaders::new().add(("Content-Type", "application/json")))
            .service(index())
    })
    .bind(("0.0.0.0", port))?
    .run()
    .await
}
