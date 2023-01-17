package config

type KafkaConf struct {
    Address string `ini:"address"`
}

type EtcdConf struct {
    Address string `ini:"address"`
    Timeout int    `ini:"timeout"`
}

type AppConf struct {
    KafkaConf `ini:"kafka"`
    EtcdConf  `ini:"etcd"`
}

// TaillogConf ---- unused ↓ ----
type TaillogConf struct {
    FileName string `ini:"filename"`
}
