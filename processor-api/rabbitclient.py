import pika
import logging
from services.loader_service import LoaderService
from services.extraction_service import ExtractionService
from services.summarization_service import SummarizationService
import os

logger = logging.getLogger()


# Callback function to process messages
def callback(ch, method, properties, body):
    document_path = body.decode()
    logger.debug(f"Received message: {document_path}")

    # Load file
    loader_service = LoaderService()
    loader_service.load(document_path)

    logger.debug(f"Loaded file: {document_path}")

    #  Extract text
    extraction_service = ExtractionService()

    text = extraction_service.extract_from_pdf(f"temp/{document_path}")

    logger.debug(f"Extracted text for: {document_path}")

    # Summarize text
    summarization_service = SummarizationService()

    summary = summarization_service.summarize(text)

    logger.debug(f"Summarized text for: {document_path}")

    # Write summary to file
    
    with open(f"summaries/{document_path}", mode='w') as file:
        file.write(summary)

    logger.debug(f"Wrote summary to: summaries/{document_path}.")

    # Acknowledge the message as processed
    ch.basic_ack(delivery_tag=method.delivery_tag)


def consume():    
    # Connect to RabbitMQ server
    rabbit_mq_url = os.getenv("RABBIT_MQ_URL")

    # Create connection parameters from the URL
    connection_parameters = pika.URLParameters(rabbit_mq_url)
    
    connection = pika.BlockingConnection(connection_parameters)  # Adjust if RabbitMQ is hosted elsewhere
    
    channel = connection.channel()

    # Declare the queue to listen to
    queue_name = 'my_queue'
    channel.queue_declare(queue=queue_name, durable=True)  # Make sure the queue is durable


    # Set up the consumer to listen to the queue
    channel.basic_consume(queue=queue_name, on_message_callback=callback, auto_ack=False)

    print(f"Listening for messages in queue '{queue_name}'...")

    # Start listening for messages
    channel.start_consuming()
