use serde::{Serialize, Deserialize};

#[allow(dead_code)]
pub enum ResponseType <T> {
    Data(T),
    Error(T),
}

#[derive(Debug, Serialize, Deserialize)]
pub struct Response <T> {
    data: Option<T>,
    error: Option<T>,
}

impl <T> Response <T> {
    pub fn new(res:ResponseType<T>) -> Self {
        match res {
            ResponseType::Data(data) => Self { data: Some(data), error: None },
            ResponseType::Error(error) => Self { data: None, error: Some(error) },
        }
    }
}