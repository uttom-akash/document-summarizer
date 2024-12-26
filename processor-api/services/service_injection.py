from flask_injector import Binder, request
from services.loader_service import LoaderService
from services.extraction_service import ExtractionService
from services.summarization_service import SummarizationService

def inject_services(binder: Binder):
    binder.bind(LoaderService, to=LoaderService, scope=request)
    binder.bind(ExtractionService, to=ExtractionService, scope=request)
    binder.bind(SummarizationService, to=SummarizationService, scope=request)