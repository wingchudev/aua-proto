PROTOC = protoc
PLUGINS = --plugin=protoc-gen-go=./protoc-gen-go --plugin=protoc-gen-micro=./protoc-gen-micro

PROTO_MICRO_BUILD = $(PROTOC) $(PLUGINS) --go_out=paths=source_relative:. --micro_out=paths=source_relative:.
PROTO_GRPC_BUILD = $(PROTOC) $(PLUGINS) -I. --go_out=paths=source_relative,plugins=grpc:.

DOC_PLUGINS = --plugin=protoc-gen-doc=./protoc-gen-doc
PROTO_DOC_BUILD = $(PROTOC) $(DOC_PLUGINS) --doc_out=./doc --doc_opt=html,index.html

target ?= aua

ALL:
	$(PROTO_MICRO_BUILD) common/*.proto
	$(PROTO_MICRO_BUILD) event/*.proto
	$(PROTO_MICRO_BUILD) rights/rights.proto
	$(PROTO_MICRO_BUILD) user/*.proto
	$(PROTO_MICRO_BUILD) cms/*.proto
	$(PROTO_MICRO_BUILD) aua/*.proto
	$(PROTO_MICRO_BUILD) wallet/*.proto
	$(PROTO_DOC_BUILD) */*.proto

micro-proto:
	$(PROTO_MICRO_BUILD) $(target)/*.proto

grpc-proto:
	$(PROTO_GRPC_BUILD) $(target)/*.proto

proto-doc:
	$(PROTO_DOC_BUILD) */*.proto
