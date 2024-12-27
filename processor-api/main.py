from api.create_api import create_api
from prometheus_client import start_http_server
import threading
from common.rabbitclient import rab_listen
from consumers.consumer_dispatcher import dispatch_consumer
from dotenv import load_dotenv

if __name__ == "__main__":
    
    load_dotenv()

    rabbit_client_thread = threading.Thread(target=rab_listen, args=(dispatch_consumer,))
    
    rabbit_client_thread.start()

    create_api()
    
    rabbit_client_thread.join()