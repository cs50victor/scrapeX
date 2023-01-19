mod utils;
use actix_web::{middleware, web, App, HttpResponse, HttpServer, Responder};
use std::error::Error;
use utils::Response;

const PORT: u16 = 3000;

async fn index() -> impl Responder {
    let url = "https://www.imdb.com/search/title/?groups=top_100&sort=user_rating";

    match scrape(url).await {
        Ok(res) => HttpResponse::Ok().json(Response::data(res)),
        Err(err) => HttpResponse::NoContent().json(Response::error(err.to_string())),
    }
}

#[actix_rt::main]
async fn main() -> std::io::Result<()> {
    println!("\n[ starting server at http://localhost:{PORT} ]");
    println!("==================================");

    HttpServer::new(|| {
        App::new()
            .wrap(middleware::Logger::default())
            .service(web::resource("/").to(index))
    })
    .bind(("127.0.0.1", PORT))?
    .run()
    .await
}

fn get_selector(_html: &str) -> String {
    "h3.lister-item-header>a".to_string()
}

async fn scrape(url: &str) -> Result<Vec<String>, Box<dyn Error>> {
    let html = reqwest::get(url).await?.text().await?;
    let document = scraper::Html::parse_document(&html);
    let selector = get_selector(&html);
    let title_selector = scraper::Selector::parse(&selector).unwrap();
    let titles = document.select(&title_selector).map(|x| x.inner_html());

    let data = titles.map(|item| item).collect::<Vec<String>>();
    Ok(data)
}
