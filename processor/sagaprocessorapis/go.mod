module github.com/awe76/saga/processor/sagaprocessorapis

go 1.17

require (
	github.com/awe76/saga/api/sagabrokerapis v0.0.0-20220118103239-d76224848f32
	github.com/awe76/saga/api/sagatransactionapis v0.0.0-20220117062005-4913b5035ccb
	github.com/awe76/saga/state/sagastateapis v0.0.0-20220114230303-dcd8eb398529
	github.com/awe76/saga/tracer/sagatracerapis v0.0.0-20220117121131-595bce182008
	github.com/nats-io/nats.go v1.13.1-0.20211122170419-d7c1d78a50fc
	google.golang.org/grpc v1.43.0
	google.golang.org/protobuf v1.27.1
)

require (
	github.com/golang/protobuf v1.5.2 // indirect
	github.com/nats-io/nats-server/v2 v2.7.0 // indirect
	github.com/nats-io/nkeys v0.3.0 // indirect
	github.com/nats-io/nuid v1.0.1 // indirect
	golang.org/x/crypto v0.0.0-20220112180741-5e0467b6c7ce // indirect
	golang.org/x/net v0.0.0-20211112202133-69e39bad7dc2 // indirect
	golang.org/x/sys v0.0.0-20220111092808-5a964db01320 // indirect
	golang.org/x/text v0.3.6 // indirect
	google.golang.org/genproto v0.0.0-20200526211855-cb27e3aa2013 // indirect
)
