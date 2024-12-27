
from injector import Injector

from consumers.summarizer_consumer import SummarizerConsumer
from modules.dependency_injection import inject_modules


def dispatch_consumer(ch, method, properties, body):

    injector = Injector(modules=[inject_modules])

    summarizer_consumer = injector.get(SummarizerConsumer)

    summarizer_consumer.consume(ch, method, properties, body)


