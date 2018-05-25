from django.apps import AppConfig


class PipelinesConfig(AppConfig):
    name = 'pipelines'
    verbose_name = 'Pipelines'

    def ready(self):
        import pipelines.signals  # noqa
