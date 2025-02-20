load("@rules_go//go:def.bzl", "go_binary","go_library")

go_library(
    name = "connector_lib",
    srcs = [
        "main.go",
    ],
    visibility = ["//visibility:public"],
    deps = [
        "//server",
        "//proto:chat_go_proto",
    ],
)

go_binary(
    name = "connector",
    embed = [":connector_lib"],
    deps = [
    "//server",
     "//proto:chat_go_proto", 
    ],
)