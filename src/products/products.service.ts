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
      let changed = false

      if (createProductDto.price !== productFound.price) {
        productFound.price = createProductDto.price
        changed ||= true
      }

      if (createProductDto.rating !== productFound.rating) {
        productFound.rating = createProductDto.rating
        changed ||= true
      }

      if (changed) {
        productFound.modified = new Date(Date.now());
        return await productFound.save();
      }
      return productFound;
    }

    const product = new Product();
    product.title = createProductDto.title.trim();
    product.description = createProductDto.description.trim();
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

  async findAllWithLimit(limit: number, offset: number): Promise<Product[]> {
    return await this.productRepository
      .createQueryBuilder('products')
      .limit(limit)
      .offset(offset)
      .getRawMany();
  }

  async findOne(id: number): Promise<Product> {
    return await this.productRepository.findOne({ id });
  }

  async findByProductKey(productKey: string): Promise<Product> {
    return await this.productRepository.findOne({ productKey })
  }
}
