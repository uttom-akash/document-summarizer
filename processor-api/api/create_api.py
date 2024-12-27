from flask import Flask
from common.logging_config import configure_logging
from common.metrics_config import setup_metrics
from api.blueprints.summary import summary_routes
from flask_injector import FlaskInjector
from modules.dependency_injection import inject_modules
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

    FlaskInjector(app=app, modules=[inject_modules])

    app.run(host="0.0.0.0", port=5000)
    