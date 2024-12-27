import logging
from logging.handlers import RotatingFileHandler

def configure_logging(app):
    """Set up logging handlers for the application."""
    filelog_handler = RotatingFileHandler("logs/doc_processing.log", maxBytes=10 * 1024)
    filelog_handler.setLevel(logging.INFO)

    consolelog_handler = logging.StreamHandler()
    consolelog_handler.setLevel(logging.INFO)

    root = logging.getLogger()
    root.setLevel(logging.INFO)
    root.addHandler(filelog_handler)
    root.addHandler(consolelog_handler)
