package com.orymar.reliablems.entity;

import java.util.UUID;
import jakarta.persistence.*;
import lombok.*;

@Entity
@Table(name = "foods")
@Data
public class Food {

    @Id
    @GeneratedValue(strategy = GenerationType.UUID)
    @Column(name = "uuid")
    private UUID uuid;

    @Column(nullable = false)
    private String name;

    private String descriptions;

    private double protein;
    private double fat;
    private double carbohydrates;
    private double calories;
}