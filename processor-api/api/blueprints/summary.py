from flask import Blueprint
import logging
from flask_injector import inject


summary_routes = Blueprint("summary_routes", __name__, url_prefix="/api/v1/processor/summaries")
logger = logging.getLogger()

@inject
@summary_routes.get("/<document>")
def get_summary(document: str):
    
    summary = ""
    with open(f"summaries/{document}", mode='r') as file:
       lines = file.readlines()
       for line in lines:
          summary +=line + "\n" 

    return summary, 200


