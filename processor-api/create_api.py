from flask import Flask
from configs.logging_config import configure_logging
from configs.metrics_config import setup_metrics
from routes.summary_routes import summary_routes
from flask_injector import FlaskInjector
from services.service_injection import inject_services
from flask import jsonify
from dotenv import load_dotenv

def create_api():
    
    app = Flask(__name__)
    
    configure_logging(app)
    
    setup_metrics(app)
    
    app.register_blueprint(summary_routes) 

    @app.errorhandler(500)
    def handle_internal_server_error(error: Exception):
        return error, 500

    FlaskInjector(app=app, modules=[inject_services])

    app.run(host="0.0.0.0", port=5000)
    