package config

type KafkaConf struct {
    Address     string `ini:"address"`
    ChanMaxSize int    `ini:"chan_max_size"`
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
