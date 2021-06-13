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
    const productFound = await this.findByProductKey(createProductDto.productKey)
    if (productFound) {
      // If product already exists then update rating and price, then return
      productFound.price = createProductDto.price;
      productFound.rating = createProductDto.rating;
      return await productFound.save();
    }

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
    return await this.productRepository.findOne({ id });
  }

  async findByProductKey(productKey: string): Promise<Product> {
    return await this.productRepository.findOne({ productKey })
  }
}
