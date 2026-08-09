package com.orymar.reliablems.controller;


import com.orymar.reliablems.service.KafkaProducerService;
import lombok.AllArgsConstructor;
import org.springframework.http.ResponseEntity;
import org.springframework.web.bind.annotation.GetMapping;
import org.springframework.web.bind.annotation.PathVariable;
import org.springframework.web.bind.annotation.RequestParam;
import org.springframework.web.bind.annotation.RestController;

import java.util.UUID;

@RestController
@AllArgsConstructor
public class BaseController {
    private final KafkaProducerService producerService;

    private static final String KAFKA_TOPIC = "my-topic";

    /// for easy test in browser
    @GetMapping("/send")
    public ResponseEntity send(@RequestParam int countRequest) {
        var num = 0.0;
        for (int i = 0; i < countRequest; i++) {
            num = Math.random() * 10000;
            String message = String.valueOf(num);
            producerService.sendMessage(KAFKA_TOPIC, UUID.randomUUID().toString(),message);
        }
        return ResponseEntity.ok().build();
    }
}
