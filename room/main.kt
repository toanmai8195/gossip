package com.example

import com.datastax.oss.driver.api.core.CqlSession
import java.net.InetSocketAddress

fun main() {
    println("Connecting to Cassandra...")

    val session = CqlSession.builder()
        .addContactPoint(InetSocketAddress("127.0.0.1", 9042)) // Đổi IP nếu cần
        .withLocalDatacenter("datacenter1") // Đổi datacenter nếu cần
        .build()

    println("Connected to Cassandra: ${session.name}")

    session.close()
}