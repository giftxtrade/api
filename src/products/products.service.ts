import { Injectable } from '@nestjs/common';
import { InjectRepository } from '@nestjs/typeorm';
import { CreateProductDto } from './dto/create-product.dto';
import { UpdateProductDto } from './dto/update-product.dto';
import { Product } from './entities/product.entity';
import { Repository } from 'typeorm';
import { CategoriesService } from '../categories/categories.service';

@Injectable()
export class ProductsService {
  private static readonly selectAll = [
    'products.id AS id',
    'products.title AS title',
    'products.description AS description',
    'products.productKey AS productKey',
    'products.imageUrl AS imageUrl',
    'products.rating AS rating',
    'products.price AS price',
    'products.currency AS currency',
    'products.modified AS modified',
    'categories.name AS category',
    'products.website AS website',
  ];

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

  async findAllWithLimit(limit: number, offset: number, minPrice?: number, maxPrice?: number, search?: string): Promise<Product[]> {
    let where = ''
    let whereValues = {}
    if (minPrice && maxPrice) {
      where = 'price >= :minPrice AND price <= :maxPrice';
      whereValues = { minPrice, maxPrice };

      if (search) {
        where += ' AND (title LIKE :search OR categories.name LIKE :search)';
        whereValues = { minPrice, maxPrice, search: `%${search}%` };
      }
    }

    return await this.productRepository
      .createQueryBuilder('products')
      .select(ProductsService.selectAll)
      .where(where, whereValues)
      .leftJoin('products.category', 'categories')
      .limit(limit)
      .offset(offset)
      .orderBy('products.rating', 'DESC')
      .getRawMany();
  }

  async findOne(id: number): Promise<Product> {
    return await this.productRepository.findOne({ id });
  }

  async findByProductKey(productKey: string): Promise<Product> {
    return await this.productRepository.findOne({ productKey })
  }
}
