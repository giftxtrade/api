import { Injectable } from '@nestjs/common';
import { InjectRepository } from '@nestjs/typeorm';
import { CreateProductDto } from './dto/create-product.dto';
import { UpdateProductDto } from './dto/update-product.dto';
import { Product } from './entities/product.entity';
import { Repository } from 'typeorm';
import { CategoriesService } from '../categories/categories.service';

@Injectable()
export class ProductsService {
  constructor(
    @InjectRepository(Product)
    private readonly productRepository: Repository<Product>,
    private readonly categoryServices: CategoriesService,
  ) { }

  async create(createProductDto: CreateProductDto): Promise<Product> {
    const product = new Product();
    product.title = createProductDto.title;
    product.description = createProductDto.description;
    product.productKey = createProductDto.productKey;
    product.imageUrl = createProductDto.imageUrl;
    product.rating = createProductDto.rating;
    product.price = createProductDto.price;
    product.currency = createProductDto.currency;
    product.category = await this.categoryServices.findLikeOrCreate(createProductDto.category);
    product.website = createProductDto.website;

    return await product.save();
  }

  async findAll(): Promise<Product[]> {
    return await this.productRepository.find();
  }

  async findOne(id: number): Promise<Product> {
    return await this.productRepository.findOne(id);
  }

  /* update(id: number, updateProductDto: UpdateProductDto) {
    return `This action updates a #${id} product`;
  }

  remove(id: number) {
    return `This action removes a #${id} product`;
  } */
}
