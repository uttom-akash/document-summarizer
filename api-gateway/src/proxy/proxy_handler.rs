use actix_web::{web, HttpRequest, HttpResponse};
use reqwest::Client;
use std::env;

pub async fn proxy_requests(
    req: HttpRequest,
    body: web::Bytes,
    client: web::Data<Client>,
) -> HttpResponse {

    let splitted_path = req.path().split('/').collect::<Vec<&str>>();

    if splitted_path.len() < 4 {
        return HttpResponse::NotFound().body("");
    }

    let service = req.path().split('/').collect::<Vec<&str>>()[3];
    println!("service: {}",service);

    let storage_api = env::var("STORAGE_API").expect("STORAGE_API address must be set");
    let processor_api = env::var("PROCESSOR_API").expect("PROCESSOR_API address must be set");
    
    println!("{}", processor_api);
    // let storage_api = env::var("STORAGE_API").expect("STORAGE api address must be set");


    match service {
        "auth" => HttpResponse::NotImplemented().body("The service is not implemented yet."),
        "storage" => {
            let backend_url = format!("{storage_api}{}", req.uri()); // Forward to backend
            
            println!("backend url:{}", backend_url);

            let forward_request = client
                .request(req.method().clone().into(), &backend_url)
                .headers(req.headers().clone().into())
                .body(body);

            match forward_request.send().await {
                Ok(res) => {
                    let status = res.status();
                    let headers = res.headers().clone();
                    let body = res.bytes().await.unwrap_or_default();

                    let mut response = HttpResponse::build(status);
                    for (key, value) in headers.iter() {
                        response.append_header((key.clone(), value.clone()));
                    }
                    response.body(body)
                }
                Err(e) => {
                    println!("{}", e); 
                    HttpResponse::BadGateway().body("Failed to proxy request.")
                }
            }
        },
        "processor" => {
            let backend_url = format!("{processor_api}{}", req.uri()); // Forward to backend

            println!("backend url: {}", backend_url);

            let forward_request = client
                .request(req.method().clone().into(), &backend_url)
                .headers(req.headers().clone().into())
                .body(body);

            match forward_request.send().await {
                Ok(res) => {
                    let status = res.status();
                    let headers = res.headers().clone();
                    let body = res.bytes().await.unwrap_or_default();

                    let mut response = HttpResponse::build(status);
                    for (key, value) in headers.iter() {
                        response.append_header((key.clone(), value.clone()));
                    }
                    response.body(body)
                }
                Err(e) => {
                    println!("{}", e);
                    HttpResponse::BadGateway().body("Failed to proxy request.")
                }
            }
        }
        "dashboard" => HttpResponse::NotImplemented().body("The service is not implemented yet."),
        _ => HttpResponse::NotFound().body("There is no service handling this request."),
    }
}
