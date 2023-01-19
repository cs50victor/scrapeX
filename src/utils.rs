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
