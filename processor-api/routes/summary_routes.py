from flask import Blueprint,render_template
import logging
from flask_injector import inject
from services.loader_service import LoaderService
from services.extraction_service import ExtractionService
from services.summarization_service import SummarizationService


summary_routes = Blueprint("summary_routes", __name__, url_prefix="/api/v1/processor/summaries")
logger = logging.getLogger()

@inject
@summary_routes.get("/<document>")
def get_summary(document: str,
                 loader_service: LoaderService,
                   extraction_service: ExtractionService,
                     summarization_service: SummarizationService):
    
    summary = ""
    with open(f"summaries/{document}", mode='r') as file:
       lines = file.readlines()
       for line in lines:
          summary +=line + "\n" 

    return summary, 200


