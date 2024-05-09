package kafka 

type KafkaConfig struct {
    Brokers []string
    UseTLS	bool
    CAFile  string
    Version string

    Authentcation Authentcation
    Topic string
}

type Authentcation struct {
    Username string
    Password string
    Algorithm string
}
