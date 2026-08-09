package com.orymar.reliablems.service;

import lombok.AllArgsConstructor;
import org.springframework.kafka.core.KafkaTemplate;
import org.springframework.stereotype.Service;

@Service
@AllArgsConstructor
public class KafkaProducerService {

    private final KafkaTemplate<String, String> kafkaTemplate;

    public void sendMessage(String topic, String key, String value) {
        kafkaTemplate.send(topic, key, value)
                .whenComplete((result, ex) -> {
                    if (ex == null) {
                        System.out.println("→ partition: " + result.getRecordMetadata().partition());
                    }
                });
    }
}