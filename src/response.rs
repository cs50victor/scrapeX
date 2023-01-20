use crate::routes::ScrapeInfo;
use serde::{Deserialize, Serialize};

#[derive(Debug, Serialize, Deserialize)]
pub struct Response<T> {
    data: Option<T>,
    error: Option<T>,
}

impl<T> Response<T> {
    pub fn data(data: T) -> Self {
        Self {
            data: Some(data),
            error: None,
        }
    }
    pub fn error(error: T) -> Self {
        Self {
            data: None,
            error: Some(error),
        }
    }
}

pub fn deserialize_json_error(err: serde_json::Error) -> Response<String> {
    match err.classify() {
        serde_json::error::Category::Io => Response::error(format!("IO error - failure to read or write bytes on an IO stream")),
        serde_json::error::Category::Syntax => Response::error(format!("Syntax error - payload was not syntactically valid JSON")),
        serde_json::error::Category::Data => Response::error(format!("couldn't deserialize JSON. Make sure your payload is valid JSON and matches the expected type {:?}", ScrapeInfo::dummy())),
        serde_json::error::Category::Eof => Response::error("EOF error".to_string()),
    }
}
