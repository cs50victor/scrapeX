use actix_web::{web, HttpRequest, HttpResponse, Resource, Responder};
use serde::{Deserialize, Serialize};

use crate::response::{deserialize_json_error, Response};
use crate::utils::scrape;

// ------------------------------ '/' ------------------------------
#[derive(Debug, Deserialize, Serialize)]
pub struct ScrapeInfo {
    url: String,
    tag: String,
    class: String,
    child_tag: Option<String>,
}

impl ScrapeInfo {
    pub fn dummy() -> Self {
        Self {
            url: String::from("https://www.google.com"),
            tag: String::from("div"),
            class: String::from("text-xl font-bold"),
            child_tag: Some(String::from("a")),
        }
    }
}

async fn handler(req: HttpRequest, body: web::Bytes) -> impl Responder {
    if req.method().as_str() != "POST" {
        return HttpResponse::MethodNotAllowed()
            .json(Response::error("Method not allowed".to_string()));
    }

    let info = match serde_json::from_slice::<ScrapeInfo>(&body) {
        Ok(info) => info,
        Err(e) => return HttpResponse::BadRequest().json(deserialize_json_error(e)),
    };

    let (url, selector) = (info.url.clone(), format_selector(info));

    match scrape(&url, selector).await {
        Ok(res) => HttpResponse::Ok().json(Response::data(res)),
        Err(err) => HttpResponse::NoContent().json(Response::error(err.to_string())),
    }
}

fn format_selector(el: ScrapeInfo) -> String {
    match el.child_tag {
        Some(child_tag) => format!("{}.{}>{}", el.tag, el.class, child_tag),
        None => format!("{}.{}", el.tag, el.class),
    }
}

pub fn index() -> Resource {
    web::resource("/").to(handler)
}
