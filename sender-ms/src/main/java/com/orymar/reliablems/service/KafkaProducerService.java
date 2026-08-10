package com.orymar.reliablems.service;

import lombok.AllArgsConstructor;
import org.springframework.kafka.core.KafkaTemplate;
import org.springframework.stereotype.Service;

@Service
@AllArgsConstructor
public class KafkaProducerService {

    private final KafkaTemplate<String, Object> kafkaTemplate;

    public void sendMessage(String topic, String key, Object payload) {
        kafkaTemplate.send(topic, key, payload);
    }
}