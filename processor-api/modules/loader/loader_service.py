import boto3
from botocore.exceptions import NoCredentialsError, PartialCredentialsError
import logging

class LoaderService:
    def __init__(self):
        self.logger = logging.getLogger()

    def load(self, document_path: str) -> str:
        s3 = boto3.client("s3")
        file_path = f"temp/{document_path}"

        try:
            s3.download_file("web-bucket-8a9e1fc", document_path, file_path)
            self.logger.debug('File downloaded successfully.')
            return file_path
        except (NoCredentialsError, PartialCredentialsError) as e:
            self.logger.debug(f"Credentials error: {e}")
            raise
        except Exception as e:
            self.logger.debug(f"Unexpected error: {e}")
            raise