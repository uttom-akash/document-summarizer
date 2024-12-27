import pika
import logging
from modules.loader.loader_service import LoaderService
from modules.extractor.extraction_service import ExtractionService
from modules.summarizer.summarization_service import SummarizationService
import os

logger = logging.getLogger()


def rab_listen(callback):    
    # Connect to RabbitMQ server
    rabbit_mq_url =  os.getenv("RABBIT_MQ_URL")
    print(f"Connecting to RabbitMQ server at {rabbit_mq_url}...")
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
