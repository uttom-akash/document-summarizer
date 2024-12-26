from create_api import create_api
from prometheus_client import start_http_server
import threading
from rabbitclient import consume
from dotenv import load_dotenv

if __name__ == "__main__":
    
    load_dotenv()

    rabbit_client_thread = threading.Thread(target=consume)
    
    rabbit_client_thread.start()

    create_api()
    
    rabbit_client_thread.join()