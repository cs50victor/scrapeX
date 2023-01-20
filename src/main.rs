mod response;
mod routes;
mod utils;

use actix_web::{middleware, App, HttpServer};
use routes::index;
use std::env;

#[actix_rt::main]
async fn main() -> std::io::Result<()> {
    let port = 3000;
    println!("\n=== STARTING SERVER ===\nhttp://localhost:{port}\n");

    env::set_var("RUST_LOG", "info");
    pretty_env_logger::init();

    HttpServer::new(|| {
        App::new()
            .wrap(middleware::Logger::new("IP - %a | Time - %D ms"))
            .wrap(middleware::DefaultHeaders::new().add(("Content-Type", "application/json")))
            .service(index())
    })
    .bind(("127.0.0.1", port))?
    .run()
    .await
}
