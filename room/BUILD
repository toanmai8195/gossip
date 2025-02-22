load("@rules_kotlin//kotlin:kotlin.bzl", "kt_jvm_library")

kt_jvm_library(
    name = "hello_lib",
    srcs = glob(["*.kt"]),
    deps = [
        "@maven//:com_datastax_oss_java_driver_core",
    ],
)

java_binary(
    name = "hello_world",
    main_class = "com.example.MainKt",
    runtime_deps = [":hello_lib"],
)