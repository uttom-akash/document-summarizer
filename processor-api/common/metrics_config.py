import time
from flask import g, request, Response
from prometheus_client import Counter, Gauge, CollectorRegistry, generate_latest

custom_registry = CollectorRegistry()
REQUESTS = Counter('http_requests_total', 'Total number of HTTP requests', ['method', 'endpoint'], registry=custom_registry)
RESPONSE_TIME = Gauge('http_response_duration_seconds', 'Duration of HTTP responses in seconds', ['method', 'endpoint'], registry=custom_registry)

def setup_metrics(app):
    """Attach request timing metrics to the app."""
    
    @app.before_request
    def before_request():
        g.request_start_time = time.time()

    @app.after_request
    def after_request(response):
        duration = time.time() - g.request_start_time
        RESPONSE_TIME.labels(method=request.method, endpoint=request.path).set(duration)
        REQUESTS.labels(method=request.method, endpoint=request.path).inc()
        return response
    
    @app.route('/metrics')
    def metrics():
        return Response(generate_latest(registry=custom_registry), mimetype='text/plain')
