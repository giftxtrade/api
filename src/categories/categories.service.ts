import { Injectable } from '@nestjs/common';
import { InjectRepository } from '@nestjs/typeorm';
import { Repository } from 'typeorm';
import { Category } from './entities/category.entity';

@Injectable()
export class CategoriesService {
  constructor(
    @InjectRepository(Category)
    private readonly categoryRepository: Repository<Category>,
  ) { }

  async insert(name: string, description?: string, categoryUrl?: string): Promise<Category> {
    const category: Category = new Category();
    category.name = name;
    category.description = description;
    category.categoryUrl = categoryUrl;
    return await category.save();
  }

  async findOne(name: string): Promise<Category> {
    return await this.categoryRepository
      .findOne({ where: { name: name } });
  }

  async findOneLike(name: string): Promise<Category> {
    name = name.toLowerCase().trim();
    return await this.categoryRepository
      .createQueryBuilder('categories')
      .where(`
        name like :name or
        description like :name
      `, { name: `%${name}%` })
      .getOne();
  }

  async findOrCreate(name: string, description?: string, categoryUrl?: string): Promise<Category> {
    const category = await this.findOne(name);
    if (!category)
      return await this.insert(name);
    return category;
  }

  async findLikeOrCreate(name: string, description?: string, categoryUrl?: string): Promise<Category> {
    const category = await this.findLikeOrCreate(name);
    if (!category)
      return await this.insert(name);
    return category;
  }
}
