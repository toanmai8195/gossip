load("@rules_proto//proto:defs.bzl", "proto_library")
load("@rules_go//proto:def.bzl", "go_proto_library")

proto_library(
    name = "chat_proto",
    srcs = ["chat.proto"],
    visibility = ["//visibility:public"],
)

go_proto_library(
    name = "chat_go_proto",
    protos = [":chat_proto"],
    visibility = ["//visibility:public"],
    deps = ["@com_google_protobuf//:go_default_library"]
)