
import logging
from injector import inject
from modules.extractor.extraction_service import ExtractionService
from modules.loader.loader_service import LoaderService
from modules.summarizer.summarization_service import SummarizationService


class SummarizerConsumer:
    @inject
    def __init__(self, loader: LoaderService,
                  extractor: ExtractionService,
                    summarizer: SummarizationService):
        self.extractor = extractor
        self.loader = loader
        self.summarizer = summarizer
        self.logger = logging.getLogger()

    def consume(self, ch, method, properties, body):
        document_path = body.decode()
        self.logger.debug(f"Received message: {document_path}")

        # Load file
        self.loader.load(document_path)

        self.logger.debug(f"Loaded file: {document_path}")

        #  Extract text
        text = self.extractor.extract_from_pdf(f"temp/{document_path}")

        self.logger.debug(f"Extracted text for: {document_path}")

        # Summarize text
        summary = self.summarizer.summarize(text)

        self.logger.debug(f"Summarized text for: {document_path}")

        # Write summary to file
        
        with open(f"summaries/{document_path}", mode='w') as file:
            file.write(summary)

        self.logger.debug(f"Wrote summary to: summaries/{document_path}.")

        # Acknowledge the message as processed
        ch.basic_ack(delivery_tag=method.delivery_tag)