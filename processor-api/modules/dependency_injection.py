from flask_injector import Binder, request
from modules.loader.loader_service import LoaderService
from modules.extractor.extraction_service import ExtractionService
from modules.summarizer.summarization_service import SummarizationService

def inject_modules(binder: Binder):
    binder.bind(LoaderService, to=LoaderService, scope=request)
    binder.bind(ExtractionService, to=ExtractionService, scope=request)
    binder.bind(SummarizationService, to=SummarizationService, scope=request)