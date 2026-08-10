package com.orymar.reliablems.entity;

import lombok.AllArgsConstructor;
import lombok.Data;

@Data
@AllArgsConstructor
public class KafkaMessagePayload {
    private String key;
    private String value;
}