load("@rules_go//go:def.bzl", "go_library")

go_library(
    name = "handler",
    srcs = glob(["*.go"]),
    visibility = ["//visibility:public"],
    deps = [
        "@com_github_gorilla_websocket//:go_default_library",
        "//proto:chat_go_proto",
         "@org_golang_google_protobuf//proto:go_default_library",
         "//server"
    ],
     importpath = "ws-server/handler",
)