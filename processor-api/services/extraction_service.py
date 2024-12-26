from PyPDF2 import PdfReader
import logging
import os

class ExtractionService:
    def __init__(self):
        self.logger = logging.getLogger()

    def extract_from_pdf(self, pdf_path):
        """Extracts text from a PDF using PyPDF2."""
        text = ""
        with open(pdf_path, 'rb') as file:
            reader = PdfReader(file)
            for page in reader.pages:
                text += page.extract_text() + "\n"
        
        self.logger.debug(text)

        os.remove(path=pdf_path)

        return text