package kafka

import (
    "github.com/IBM/sarama"
    "crypto/tls"
    "crypto/x509"
    "fmt"
    "log"
    "os"
)

func InitKafka(kafkaConfig *KafkaConfig) (sarama.SyncProducer, error) {    
    config := sarama.NewConfig()
    version, err := sarama.ParseKafkaVersion(kafkaConfig.Version)
    if err != nil {
        return nil, err
    }

    config.Version = version
    
    config.ClientID = "snorlax"
    config.Metadata.Full = true
    config.Net.SASL.Enable = true
    config.Net.SASL.User = kafkaConfig.Authentcation.Username
    config.Net.SASL.Password = kafkaConfig.Authentcation.Password
    config.Net.SASL.Handshake = true
    if kafkaConfig.Authentcation.Algorithm == "sha512" {
        config.Net.SASL.SCRAMClientGeneratorFunc = func() sarama.SCRAMClient { return &XDGSCRAMClient{HashGeneratorFcn: SHA512} }
        config.Net.SASL.Mechanism = sarama.SASLTypeSCRAMSHA512
    } else if kafkaConfig.Authentcation.Algorithm == "sha256" {
        config.Net.SASL.SCRAMClientGeneratorFunc = func() sarama.SCRAMClient { return &XDGSCRAMClient{HashGeneratorFcn: SHA256} }
        config.Net.SASL.Mechanism = sarama.SASLTypeSCRAMSHA256
    } 

    if kafkaConfig.UseTLS {
        caCert, err := os.ReadFile(kafkaConfig.CAFile)
        if err != nil {
            log.Fatal(err)
        }

        caCertPool := x509.NewCertPool()
        caCertPool.AppendCertsFromPEM(caCert)

        config.Net.TLS.Enable = true
        config.Net.TLS.Config = &tls.Config{
            RootCAs:            caCertPool,
            InsecureSkipVerify: true,
        }
    }


    producer, err := newProducer(config, kafkaConfig)
    if err != nil {
        fmt.Println("Could not create producer: ", err)
        return nil, err
    }

    return producer, nil
}


func newProducer(config *sarama.Config, kafkaConfig *KafkaConfig) (sarama.SyncProducer, error) {
    config.Producer.Partitioner = sarama.NewRandomPartitioner
    config.Producer.RequiredAcks = sarama.WaitForAll
    config.Producer.Return.Successes = true
    producer, err := sarama.NewSyncProducer(kafkaConfig.Brokers, config)

    return producer, err
}

func SendMessage(producer sarama.SyncProducer, topic string, message string) (error) {
    msg := &sarama.ProducerMessage{
        Topic:     topic,
        Partition: -1,
        Value:     sarama.StringEncoder(message),
    }

    _, _, err := producer.SendMessage(msg)
    if err != nil {
        fmt.Println(err)
        return err
    }

    return nil
}
