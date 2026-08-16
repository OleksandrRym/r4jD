package com.orymar.reliablems.service;


import com.orymar.reliablems.entity.Food;
import com.orymar.reliablems.repo.FoodRepository;
import lombok.RequiredArgsConstructor;
import org.springframework.stereotype.Service;

import java.util.List;

@Service
@RequiredArgsConstructor
public class FoodService {
    private final FoodRepository foodRepository;

    public Food add(Food food){
       return foodRepository.save(food);
    }

    public List<Food> list(){
        return foodRepository.findAll();
    }
}
