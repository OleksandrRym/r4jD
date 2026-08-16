package com.orymar.reliablems.controller;

import com.orymar.reliablems.entity.Food;
import com.orymar.reliablems.entity.dto.FoodCreatedRequestDto;
import com.orymar.reliablems.entity.mapper.FoodMapper;
import com.orymar.reliablems.service.FoodService;
import lombok.AllArgsConstructor;
import org.springframework.http.HttpStatus;
import org.springframework.http.ResponseEntity;
import org.springframework.web.bind.annotation.*;

import java.util.List;

@RestController
@RequestMapping("/food")
@AllArgsConstructor
public class FoodController {
    private final FoodService foodService;
    private final FoodMapper foodMapper;

    @PostMapping
    public ResponseEntity<?> add(@RequestBody FoodCreatedRequestDto dto) {
        Food entity = foodMapper.toEntity(dto);
        Food food = foodService.add(entity);
        FoodCreatedRequestDto result = foodMapper.toDto(food);
        return ResponseEntity
                .status(HttpStatus.CREATED)
                .body(result);
    }

    @GetMapping
    public ResponseEntity<List<Food>> get() {
        return ResponseEntity
                .status(HttpStatus.FOUND)
                .body(foodService.list());
    }

}
