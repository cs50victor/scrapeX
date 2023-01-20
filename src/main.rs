mod response;

use actix_web::{post, middleware, web, App, HttpRequest, HttpResponse, HttpServer, Responder, error};
use std::{env, error::Error};
use response::Response;
use serde::{Deserialize, Serialize};
use log::info;

#[derive(Debug, Deserialize, Serialize)]
struct ScrapeInfo {
    url: String,
    tag: String,
    class: String,
    child_tag: Option<String>,
}

async fn index(req: HttpRequest, info: web::Json::<ScrapeInfo>) -> impl Responder {
    if req.method().as_str() != "POST" {
        info!("Method not allowed");
        return HttpResponse::MethodNotAllowed().json(Response::error("Method not allowed".to_string()));
    }

    let url = info.url.clone();
    let selector = format_selector(info.into_inner());
    match scrape(&url, selector).await {
        Ok(res) => HttpResponse::Ok().json(Response::data(res)),
        Err(err) => HttpResponse::NoContent().json(Response::error(err.to_string())),
    }
}

#[actix_rt::main]
async fn main() -> std::io::Result<()> {
    let port = 3000;
    println!("\n=== STARTING SERVER ===\nhttp://localhost:{port}\n");
    env::set_var("RUST_LOG", "info");
    pretty_env_logger::init();
    
    // web::JsonConfig::default().error_handler(| err, req | {
    //     let err = serde_json::from_slice::<serde_json::Value>(&res.body()).unwrap();
    //     let err = err["message"].as_str().unwrap();
    //     error::InternalError::from_response(err, HttpResponse::BadRequest().finish()).into()
    // })

    HttpServer::new(|| {
        App::new()
            .app_data(web::JsonConfig::default().error_handler(| err, req | {
                let res = HttpResponse::BadRequest().json(Response::error(format!("err {err}, req {req:?}")));
                error::InternalError::from_response(err, res).into()
            }))
            .wrap(middleware::Logger::new("IP - %a | Time - %D ms"))
            .wrap(middleware::DefaultHeaders::new().add(("Content-Type", "application/json")))
            .service(web::resource("/").to(index))
    })
    .bind(("127.0.0.1", port))?
    .run()
    .await
}

fn format_selector(el: ScrapeInfo) -> String {
    match el.child_tag {
        Some(child_tag) => format!("{}.{}>{}", el.tag, el.class, child_tag),
        None => format!("{}.{}", el.tag, el.class),
    }
}

async fn scrape(url: &str, selector: String) -> Result<Vec<String>, Box<dyn Error>> {
    let html = reqwest::get(url).await?.text().await?;
    let document = scraper::Html::parse_document(&html);
    let title_selector = scraper::Selector::parse(&selector).unwrap();
    let titles = document.select(&title_selector).map(|x| x.inner_html());

    let data = titles.map(|item| item).collect::<Vec<String>>();
    Ok(data)
}
