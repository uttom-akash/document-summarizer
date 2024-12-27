import logging
from huggingface_hub import InferenceClient
import os

class SummarizationService:
    def __init__(self):
        self.logger = logging.getLogger()

    def summarize(self, text: str):
        """Summarizes the text using facebook/bart-large-cnn model."""

        client = InferenceClient(
            "facebook/bart-large-cnn",
            token=os.getenv('HUGGING_FACE_TOKEN'),
        )

        summary = client.summarization(text[0:1000])

        self.logger.debug(summary.summary_text)

        return summary.summary_text
    