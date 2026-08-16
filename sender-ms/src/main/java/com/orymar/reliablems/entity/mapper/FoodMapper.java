package com.orymar.reliablems.entity.mapper;

import com.orymar.reliablems.entity.Food;
import com.orymar.reliablems.entity.dto.FoodCreatedRequestDto;
import org.mapstruct.Mapper;

@Mapper(componentModel = "spring")
public interface FoodMapper {
    Food toEntity(FoodCreatedRequestDto foodCreatedRequestDto);

    FoodCreatedRequestDto toDto(Food entity);
}
