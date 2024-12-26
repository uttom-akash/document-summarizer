use actix_web::{web, App, HttpServer};
use proxy::proxy_handler::proxy_requests;
use reqwest::Client;
use dotenv::dotenv;
pub mod proxy;

#[actix_web::main]
async fn main() -> std::io::Result<()> {
    dotenv().ok();
    
    println!("Starting the server");

    let client = Client::new(); 

    HttpServer::new(move || {
        App::new()
            .app_data(web::Data::new(client.clone())) 
            .default_service(web::route().to(proxy_requests)) 
    })
    .bind(("0.0.0.0", 3000))?
    .run()
    .await
}