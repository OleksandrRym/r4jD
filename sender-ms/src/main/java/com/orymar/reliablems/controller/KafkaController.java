package com.orymar.reliablems.controller;

import com.orymar.reliablems.service.KafkaProducerService;
import lombok.AllArgsConstructor;
import org.springframework.http.ResponseEntity;
import org.springframework.web.bind.annotation.GetMapping;
import org.springframework.web.bind.annotation.RequestParam;
import org.springframework.web.bind.annotation.RestController;

import java.util.UUID;
import com.orymar.reliablems.entity.KafkaMessagePayload;

@RestController
@AllArgsConstructor
public class KafkaController {
    private final KafkaProducerService producerService;
    private static final String KAFKA_TOPIC = "my-topic";

    /// for easy test in browser
    @GetMapping("/send")
    public ResponseEntity<Void> send(@RequestParam int countRequest) {
        double num = 0.0;

        for (int i = 0; i < countRequest; i++) {
           num = Math.random() * 10000;
           KafkaMessagePayload payload = new KafkaMessagePayload(
                UUID.randomUUID().toString(), String.valueOf(num)
           );
           producerService.sendMessage(KAFKA_TOPIC, payload.getKey(), payload);
        }

        return ResponseEntity.ok().build();
    }
}