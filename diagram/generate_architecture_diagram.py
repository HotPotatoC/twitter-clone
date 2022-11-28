from diagrams import Cluster, Diagram, Edge
from diagrams.aws.storage import S3
from diagrams.onprem.analytics import Spark
from diagrams.onprem.client import User
from diagrams.onprem.compute import Server
from diagrams.onprem.database import PostgreSQL
from diagrams.onprem.inmemory import Redis
from diagrams.onprem.network import Nginx
from diagrams.onprem.queue import Kafka
from diagrams.saas.cdn import Cloudflare

with Diagram("High-Level Architecture", show=False):
    client = User("Client")

    with Cluster("Media"):
        cdn = Cloudflare("CDN")
        (
            client
            >> Edge(color="black")
            >> cdn
            >> Edge(color="darkorange")
            << S3("media store")
        )

    with Cluster(""):
        lb = Nginx("Load Balancer")

        client >> lb
        with Cluster("App Servers"):
            svcs = [
                Server("server1"),
                Server("server2"),
                Server("server3"),
            ]
            lb >> Edge(color="darkgreen") << svcs

        with Cluster("Feed Task Queue"):
            pubsub = Kafka("Pub/Sub")
            worker = Spark("Cluster/Worker")
            (
                svcs
                >> Edge(color="forestgreen")
                >> pubsub
                >> Edge(color="forestgreen")
                >> worker
            )

        with Cluster("Caching Layer"):
            cache = Redis("Cache")
            feed_cache = Redis("Feed cache")
            svcs >> Edge(color="red") >> cache
            svcs >> Edge(color="red") >> feed_cache
            worker >> Edge(color="forestgreen") >> feed_cache

        with Cluster("DB Layer"):
            master = PostgreSQL("master")
            db_read_replicas = PostgreSQL("read-replicas")
            master - Edge(style="dotted") - db_read_replicas
            svcs >> Edge(color="blue") >> master
            worker >> Edge(color="forestgreen") >> master
