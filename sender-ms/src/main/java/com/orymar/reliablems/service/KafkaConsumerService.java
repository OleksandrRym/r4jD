package com.orymar.reliablems.service;

import org.springframework.kafka.annotation.KafkaListener;
import org.springframework.stereotype.Component;

@Component
public class KafkaConsumerService {
    private static final String KAFKA_TOPIC = "my-topic";

    @KafkaListener(topics = KAFKA_TOPIC, groupId = "my-group")
    public void listen(String message) {
        System.out.println("Отримано: " + message);
    }
}