#/bin/bash

## /opt/bitnami/kafka/bin/kafka-topics.sh --create --topic $TEST_TOPIC_NAME --bootstrap-server kafka:9092
/opt/bitnami/kafka/bin/kafka-topics.sh --create --topic transactions --bootstrap-server kafka:29092
echo "Topico transactions criado com sucesso!"
/opt/bitnami/kafka/bin/kafka-topics.sh --create --topic balances --bootstrap-server kafka:29092
echo "Topico balances criado com sucesso!"
